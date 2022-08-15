package main

import (
	"fmt"
	"github.com/google/uuid"
	"notepad-slack/application"
	"notepad-slack/configuration"
	"notepad-slack/domain"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const FileNamePrefix = "ncs_file_request_2022"

var counter CounterStore

type CounterStore struct {
	Counter int32
	sync.Mutex
}

// ApiConsumerSubscriber api test consumer
type ApiConsumerSubscriber struct {
	Events chan domain.FileRequest
}

// Listen listener
//
//	Parameter:
//		subscriber:
//			api consumer
func (subscriber *ApiConsumerSubscriber) Listen(store *CounterStore) {
	for {
		// track total incoming requests
		store.Lock()
		atomic.AddInt32(&store.Counter, 1)
		store.Unlock()

		select {
		case event := <-subscriber.Events:
			fmt.Println(strconv.Itoa(int(store.Counter)) + ". received file: '" + event.Name + "' with id: " + event.Id)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout...")
		}
	}
}

func main() {
	app := application.App{}
	app.StartConfiguration()
	counter = CounterStore{}

	events := make(chan domain.FileRequest, 1)
	defer close(events)

	fileRequest := &domain.FileRequest{
		Payload: map[string]string{
			"tax_id":        "txid_12345",
			"tax_item_code": "PSP12001",
			"amount":        "900",
			"fee":           "5",
		},
	}

	fileRequestManager := domain.FileRequestManager{
		Events: events,
	}
	apiConsumer := ApiConsumerSubscriber{
		Events: events,
	}

	go func() {
		for {
			now := time.Now()
			fileRequest.Id = uuid.New().String()
			fileRequest.Name = fmt.Sprintf("%s_%d_%d_%d_%d_%d", FileNamePrefix, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

			_ = fileRequestManager.Send(fileRequest)
			time.Sleep(time.Duration(configuration.GetAppDefaultTimer()))
		}
	}()

	apiConsumer.Listen(&counter)
}

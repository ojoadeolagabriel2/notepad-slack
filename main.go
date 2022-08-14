package main

import (
	"fmt"
	"github.com/google/uuid"
	"notepad-slack/application"
	"notepad-slack/configuration"
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

type FileRequest struct {
	Id      string
	Name    string
	Payload map[string]string
}

type FileRequestManager struct {
	Events chan FileRequest
}

// Send send file requests
//	Parameters:
//		request:
//			file request [main.FileRequest]
//		fileRequestManager:
//			file request manager [main.FileRequestManager]
//	Returns:
//		error
func (fileRequestManager *FileRequestManager) Send(request *FileRequest) error {
	fmt.Println("sending ..." + request.Name)
	fileRequestManager.Events <- *request
	fmt.Println("sent ..." + request.Name)
	return nil
}

// ApiConsumerSubscriber api test consumer
type ApiConsumerSubscriber struct {
	Events chan FileRequest
}

// Listen listener
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

	events := make(chan FileRequest, 1)
	defer close(events)

	fileRequest := &FileRequest{
		Payload: map[string]string{
			"tax_id":        "txid_12345",
			"tax_item_code": "PSP12001",
			"amount":        "900",
			"fee":           "5",
		},
	}

	fileRequestManager := FileRequestManager{
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

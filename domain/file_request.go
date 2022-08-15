package domain

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"notepad-slack/dto"
	"notepad-slack/utils"
)

type FileRequest struct {
	Id      string
	Name    string
	Payload map[string]string
}

type FileRequestManager struct {
	Events chan FileRequest
}

// Send send file requests
//
//	Parameters:
//		request:
//			file request [main.FileRequest]
//		fileRequestManager:
//			file request manager [main.FileRequestManager]
//	Returns:
//		error
func (fileRequestManager *FileRequestManager) Send(request *FileRequest) error {
	_ = toFileRequestDto(request)

	fmt.Println("sending ..." + request.Name)
	fileRequestManager.Events <- *request
	fmt.Println("sent ..." + request.Name)
	return nil
}

func validate(request *FileRequest) (bool, error) {
	return true, nil
}

func toFileRequestDto(request *FileRequest) *dto.FileRequestDto {
	currency := utils.Currency("GBP")
	if !currency.IsValid() {
		return nil
	}

	if success, err := validate(request); err == nil && success {
		fileRequestDto := &dto.FileRequestDto{
			TraceId:           uuid.New().String(),
			MerchantRequestId: request.Id,
			PaymentItemCode:   request.Name,
			AmountPayable: utils.Amount{
				Value:    utils.ToInt64(request.Payload["amount"]),
				Currency: currency,
			},
			Fee: utils.Amount{
				Value:    utils.ToInt64(request.Payload["fee"]),
				Currency: currency,
			},
		}
		if result, err := json.Marshal(fileRequestDto); err == nil {
			fmt.Println(string(result))
		}
		return fileRequestDto
	}
	return nil
}

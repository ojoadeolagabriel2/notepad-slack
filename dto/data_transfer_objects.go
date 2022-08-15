package dto

import "notepad-slack/utils"

type FileRequestDto struct {
	TraceId           string       `json:"trace_id"`
	MerchantRequestId string       `json:"merchant_request_id"`
	PaymentItemCode   string       `json:"payment_item_code"`
	AmountPayable     utils.Amount `json:"amount_payable"`
	Fee               utils.Amount `json:"fee"`
}

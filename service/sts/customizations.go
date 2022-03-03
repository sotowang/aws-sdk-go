package sts

import "github.com/sotowang/aws-sdk-go/aws/request"

func init() {
	initRequest = customizeRequest
}

func customizeRequest(r *request.Request) {
	r.RetryErrorCodes = append(r.RetryErrorCodes, ErrCodeIDPCommunicationErrorException)
}

package gosns

import (
	"net/http"

	"github.com/Admiral-Piett/goaws/app"
)

var (
	ErrNonExistentTopic = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "Not Found",
		Code:      "AWS.SimpleNotificationService.NonExistentTopic",
		Message:   "The specified topic does not exist for this wsdl version.",
	}
	ErrNonExistentSubscription = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "Not Found",
		Code:      "AWS.SimpleNotificationService.NonExistentSubscription",
		Message:   "The specified subscription does not exist for this wsdl version.",
	}
	ErrTopicAlreadyExists = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "Duplicate",
		Code:      "AWS.SimpleNotificationService.TopicAlreadyExists",
		Message:   "The specified topic already exists.",
	}
	ErrInvalidParameterValue = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "InvalidParameter",
		Code:      "AWS.SimpleNotificationService.ValidationError",
		Message:   "The input fails to satisfy the constraints specified by an AWS service.",
	}
)

package models

import (
	"net/http"

	"github.com/Admiral-Piett/goaws/app"
)

var (
	SnsErrNonExistentTopic = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "NonExistentTopic",
		Code:      "AWS.SimpleNotificationService.NonExistentTopic",
		Message:   "The specified topic does not exist for this wsdl version.",
	}
	SnsErrNonExistentSubscription = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "NonExistentSubscription",
		Code:      "AWS.SimpleNotificationService.NonExistentSubscription",
		Message:   "The specified subscription does not exist for this wsdl version.",
	}
	SnsErrTopicAlreadyExists = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "TopicAlreadyExists",
		Code:      "AWS.SimpleNotificationService.TopicAlreadyExists",
		Message:   "The specified topic already exists.",
	}
	SnsErrInvalidParameterValue = &app.SnsErrorType{
		HttpError: http.StatusBadRequest,
		Type:      "ValidationError",
		Code:      "AWS.SimpleNotificationService.ValidationError",
		Message:   "The input fails to satisfy the constraints specified by an AWS service.",
	}
)

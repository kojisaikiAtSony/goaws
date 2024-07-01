package models

import (
	"net/url"

	"github.com/Admiral-Piett/goaws/app"
)

var BASE_XMLNS = "http://queue.amazonaws.com/doc/2012-11-05/"
var BASE_RESPONSE_METADATA = app.ResponseMetadata{RequestId: "00000000-0000-0000-0000-000000000000"}

var AVAILABLE_QUEUE_ATTRIBUTES = map[string]bool{
	"DelaySeconds":                          true,
	"MaximumMessageSize":                    true,
	"MessageRetentionPeriod":                true,
	"Policy":                                true,
	"ReceiveMessageWaitTimeSeconds":         true,
	"VisibilityTimeout":                     true,
	"RedrivePolicy":                         true,
	"RedriveAllowPolicy":                    true,
	"ApproximateNumberOfMessages":           true,
	"ApproximateNumberOfMessagesDelayed":    true,
	"ApproximateNumberOfMessagesNotVisible": true,
	"CreatedTimestamp":                      true,
	"LastModifiedTimestamp":                 true,
	"QueueArn":                              true,
}

// Ref: https://docs.aws.amazon.com/sns/latest/api/API_CreateTopic.html
type CreateTopicRequest struct {
	Name string `json:"Name" schema:"Name"`

	// Goaws unsupports below properties currently.
	DataProtectionPolicy string            `json:"DataProtectionPolicy" schema:"DataProtectionPolicy"`
	Attributes           map[string]string `json:"Attributes" schema:"Attributes"`
	Tags                 map[string]string `json:"Tags" schema:"Tags"`
}

func (r *CreateTopicRequest) SetAttributesFromForm(values url.Values) {}

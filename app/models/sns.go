package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/Admiral-Piett/goaws/app"

	log "github.com/sirupsen/logrus"
)

func NewCreateTopicRequest() *CreateTopicRequest {
	return &CreateTopicRequest{}
}

type CreateTopicRequest struct {
	Name string `json:"Name" schema:"Name"`

	// Goaws unsupports below properties currently.
	DataProtectionPolicy string            `json:"DataProtectionPolicy" schema:"DataProtectionPolicy"`
	Attributes           map[string]string `json:"Attributes" schema:"Attributes"`
	Tags                 map[string]string `json:"Tags" schema:"Tags"`
}

func (r *CreateTopicRequest) SetAttributesFromForm(values url.Values) {}

func NewSubscribeRequest() *SubscribeRequest {
	return &SubscribeRequest{}
}

type SubscribeRequest struct {
	TopicArn   string                 `json:"TopicArn" schema:"TopicArn"`
	Endpoint   string                 `json:"Endpoint" schema:"Endpoint"`
	Protocol   string                 `json:"Protocol" schema:"Protocol"`
	Attributes SubscriptionAttributes `json:"Attributes"`
}

func (r *SubscribeRequest) SetAttributesFromForm(values url.Values) {
	for i := 1; true; i++ {
		nameKey := fmt.Sprintf("Attributes.entry.%d.key", i)
		attrName := values.Get(nameKey)
		if attrName == "" {
			break
		}

		valueKey := fmt.Sprintf("Attributes.entry.%d.value", i)
		attrValue := values.Get(valueKey)
		if attrValue == "" {
			continue
		}
		switch attrName {
		case "RawMessageDelivery":
			tmp, err := strconv.ParseBool(attrValue)
			if err != nil {
				log.Debugf("Failed to parse form attribute - %s: %s", attrName, attrValue)
				continue
			}
			r.Attributes.RawMessageDelivery = tmp
		case "FilterPolicy":
			var tmp map[string][]string
			err := json.Unmarshal([]byte(attrValue), &tmp)
			if err != nil {
				log.Debugf("Failed to parse form attribute - %s: %s", attrName, attrValue)
				continue
			}
			r.Attributes.FilterPolicy = tmp
		}
	}
	return
}

type SubscriptionAttributes struct {
	FilterPolicy       app.FilterPolicy `json:"FilterPolicy" schema:"FilterPolicy"`
	RawMessageDelivery bool             `json:"RawMessageDelivery" schema:"RawMessageDelivery"`
	//DeliveryPolicy      map[string]interface{} `json:"DeliveryPolicy" schema:"DeliveryPolicy"`
	//FilterPolicyScope   string                 `json:"FilterPolicyScope" schema:"FilterPolicyScope"`
	//RedrivePolicy       RedrivePolicy          `json:"RedrivePolicy" schema:"RawMessageDelivery"`
	//SubscriptionRoleArn string                 `json:"SubscriptionRoleArn" schema:"SubscriptionRoleArn"`
	//ReplayPolicy        string                 `json:"ReplayPolicy" schema:"ReplayPolicy"`
	//ReplayStatus        string                 `json:"ReplayStatus" schema:"ReplayStatus"`
}

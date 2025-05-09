package gosqs

import (
	"net/http"
	"strings"
	"time"

	"github.com/Admiral-Piett/goaws/app/interfaces"
	"github.com/Admiral-Piett/goaws/app/models"
	"github.com/Admiral-Piett/goaws/app/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func ChangeMessageVisibilityV1(req *http.Request) (int, interfaces.AbstractResponseBody) {
	requestBody := models.NewChangeMessageVisibilityRequest()
	ok := utils.REQUEST_TRANSFORMER(requestBody, req, false)
	if !ok {
		log.Error("Invalid Request - ChangeMessageVisibilityV1")
		return utils.CreateErrorResponseV1("InvalidParameterValue", true)
	}

	vars := mux.Vars(req)

	queueUrl := requestBody.QueueUrl
	queueName := ""
	if queueUrl == "" {
		queueName = vars["queueName"]
	} else {
		uriSegments := strings.Split(queueUrl, "/")
		queueName = uriSegments[len(uriSegments)-1]
	}

	receiptHandle := requestBody.ReceiptHandle

	visibilityTimeout := requestBody.VisibilityTimeout
	if visibilityTimeout > 43200 {
		return utils.CreateErrorResponseV1("ValidationError", true)
	}

	if _, ok := models.SyncQueues.Queues[queueName]; !ok {
		return utils.CreateErrorResponseV1("QueueNotFound", true)
	}

	models.SyncQueues.Lock()
	messageFound := false
	for i := 0; i < len(models.SyncQueues.Queues[queueName].Messages); i++ {
		queue := models.SyncQueues.Queues[queueName]
		msgs := queue.Messages
		if msgs[i].ReceiptHandle == receiptHandle {
			timeout := models.SyncQueues.Queues[queueName].VisibilityTimeout
			if visibilityTimeout == 0 {
				msgs[i].ReceiptTime = time.Now().UTC()
				msgs[i].ReceiptHandle = ""
				msgs[i].VisibilityTimeout = time.Now().Add(time.Duration(timeout) * time.Second)
				msgs[i].Retry++
				if queue.MaxReceiveCount > 0 &&
					queue.DeadLetterQueue != nil &&
					msgs[i].Retry >= queue.MaxReceiveCount {
					queue.DeadLetterQueue.Messages = append(queue.DeadLetterQueue.Messages, msgs[i])
					queue.Messages = append(queue.Messages[:i], queue.Messages[i+1:]...)
				}
			} else {
				msgs[i].VisibilityTimeout = time.Now().Add(time.Duration(visibilityTimeout) * time.Second)
			}
			messageFound = true
			break
		}
	}
	models.SyncQueues.Unlock()
	if !messageFound {
		return utils.CreateErrorResponseV1("MessageNotInFlight", true)
	}

	respStruct := models.ChangeMessageVisibilityResult{
		Xmlns:    models.BaseXmlns,
		Metadata: models.BaseResponseMetadata,
	}

	return http.StatusOK, &respStruct
}

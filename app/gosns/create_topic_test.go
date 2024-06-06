package gosns

import (
	"net/http"
	"testing"

	"github.com/Admiral-Piett/goaws/app"
	"github.com/Admiral-Piett/goaws/app/fixtures"
	"github.com/Admiral-Piett/goaws/app/interfaces"
	"github.com/Admiral-Piett/goaws/app/models"
	"github.com/Admiral-Piett/goaws/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateTopicV1_success(t *testing.T) {
	app.CurrentEnvironment = fixtures.LOCAL_ENVIRONMENT
	defer func() {
		utils.ResetApp()
		utils.REQUEST_TRANSFORMER = utils.TransformRequest
	}()

	request_success := models.CreateTopicRequest{
		Name: "new-topic-1",
	}
	utils.REQUEST_TRANSFORMER = func(resultingStruct interfaces.AbstractRequestBody, req *http.Request, emptyRequestValid bool) (success bool) {
		v := resultingStruct.(*models.CreateTopicRequest)
		*v = request_success
		return true
	}

	// No topic yet
	assert.Equal(t, len(app.SyncTopics.Topics), 0)

	// Request
	_, r := utils.GenerateRequestInfo("POST", "/", nil, true)
	status, response := CreateTopicV1(r)

	// Result
	assert.Equal(t, http.StatusOK, status)
	createTopicResponse, ok := response.(models.CreateTopicResponse)
	assert.True(t, ok)
	assert.Contains(t, createTopicResponse.Result.TopicArn, "arn:aws:sns:")
	assert.Contains(t, createTopicResponse.Result.TopicArn, "new-topic-1")
	// 1 topic there
	assert.Equal(t, len(app.SyncTopics.Topics), 1)
}
func TestCreateTopicV1_existant_topic(t *testing.T) {
	app.CurrentEnvironment = fixtures.LOCAL_ENVIRONMENT
	defer func() {
		utils.ResetApp()
		utils.REQUEST_TRANSFORMER = utils.TransformRequest
	}()

	// Same topic name with existant topic
	request_success := models.CreateTopicRequest{
		Name: "new-topic-1",
	}
	utils.REQUEST_TRANSFORMER = func(resultingStruct interfaces.AbstractRequestBody, req *http.Request, emptyRequestValid bool) (success bool) {
		v := resultingStruct.(*models.CreateTopicRequest)
		*v = request_success
		return true
	}

	// Prepare existant topic
	topic := &app.Topic{
		Name: "new-topic-1",
		Arn:  "arn:aws:sns:us-east-1:123456789012:new-topic-1",
	}
	app.SyncTopics.Topics["new-topic-1"] = topic
	assert.Equal(t, len(app.SyncTopics.Topics), 1)

	// Reques
	_, r := utils.GenerateRequestInfo("POST", "/", nil, true)
	status, response := CreateTopicV1(r)

	// Result
	assert.Equal(t, http.StatusOK, status)
	createTopicResponse, ok := response.(models.CreateTopicResponse)
	assert.True(t, ok)
	assert.Equal(t, createTopicResponse.Result.TopicArn, "arn:aws:sns:us-east-1:123456789012:new-topic-1") // Same with existant topic
	// No additional topic
	assert.Equal(t, len(app.SyncTopics.Topics), 1)
}

func TestCreateTopicV1_request_transformer_error(t *testing.T) {
	app.CurrentEnvironment = fixtures.LOCAL_ENVIRONMENT
	defer func() {
		utils.ResetApp()
		utils.REQUEST_TRANSFORMER = utils.TransformRequest
	}()

	utils.REQUEST_TRANSFORMER = func(resultingStruct interfaces.AbstractRequestBody, req *http.Request, emptyRequestValid bool) (success bool) {
		return false
	}

	_, r := utils.GenerateRequestInfo("POST", "/", nil, true)
	code, _ := CreateTopicV1(r)

	assert.Equal(t, http.StatusBadRequest, code)
}

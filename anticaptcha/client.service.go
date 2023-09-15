package anticaptcha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha/model"
)

const (
	BaseAddress = "https://api.anti-captcha.com"
)

type Client struct {
	Key         string
	SoftId      int
	Language    string
	CallbackUrl string
	Verbose     bool
	client      *http.Client
	baseAddress string
}

type Tasker interface {
}

func NewClient(key string) *Client {
	return &Client{
		client:      &http.Client{},
		baseAddress: BaseAddress,
		Key:         key,
		SoftId:      0,
		Language:    "en",
		CallbackUrl: "",
		Verbose:     false,
	}
}

func (c *Client) EnableVerbose() {
	c.Verbose = true
}

func (c *Client) SetSoftId(value int) {
	c.SoftId = value
}

func (c *Client) SetLanguage(value string) {
	c.Language = value
}

func (c *Client) SetCallbackUrl(value string) {
	c.CallbackUrl = value
}

func (c *Client) SetBaseAddress(value string) {
	c.baseAddress = value
}

// Retrieve an account balance
// Please don't call this method more than once every 30 seconds, and use a memory/disk cached value instead.
func (c *Client) GetBalance() (*model.Balance, error) {
	endpoint := fmt.Sprintf("%s/getBalance", c.baseAddress)

	data := map[string]string{
		"clientKey": c.Key,
	}

	r, err := c.httpPostRequest(endpoint, data)

	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{})
	err = json.Unmarshal(r, &response)

	if err != nil {
		return nil, err
	}

	switch response["balance"].(type) {
	case float64:
		var balance model.Balance
		err = json.Unmarshal(r, &balance)

		if err != nil {
			return nil, err
		}

		return &balance, nil
	default:
		return nil, ErrorMapping[response["errorCode"].(string)]
	}
}

// Obtain queue load statistics
// This method makes it possible to define a suitable time for uploading a new task. Output is cached for 10 seconds.
func (c *Client) GetQueueStats(queueId model.QueueId) (*model.QueueStats, error) {
	endpoint := fmt.Sprintf("%s/getQueueStats", c.baseAddress)

	data := map[string]interface{}{
		"clientKey": c.Key,
		"queueId":   queueId,
	}

	r, err := c.httpPostRequest(endpoint, data)

	if err != nil {
		return nil, err
	}

	var queueStats model.QueueStats
	err = json.Unmarshal(r, &queueStats)

	if err != nil {
		return nil, err
	}

	return &queueStats, nil
}

func (c *Client) Report(taskId float64, reportTaskType model.ReportTaskType) (*model.ReportResponse, error) {
	endpoint := fmt.Sprintf("%s/%s", c.baseAddress, reportTaskType)

	data := map[string]interface{}{
		"clientKey": c.Key,
		"taskId":    taskId,
	}

	r, err := c.httpPostRequest(endpoint, data)

	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{})
	err = json.Unmarshal(r, &response)

	if err != nil {
		return nil, err
	}

	switch response["status"].(type) {
	case float64:
		var report model.ReportResponse
		err = json.Unmarshal(r, &report)

		if err != nil {
			return nil, err
		}

		return &report, nil
	default:
		return nil, ErrorMapping[response["errorCode"].(string)]
	}
}

// Pull a captcha task status
func (c *Client) GetResult(task Tasker) (*model.TaskResultResponse, error) {
	endpoint := fmt.Sprintf("%s/getTaskResult", c.baseAddress)

	res, err := c.createTask(task)

	if err != nil {
		return nil, err
	}

	for {
		data := model.TaskResultRequest{
			ClientKey: c.Key,
			TaskId:    res.TaskId,
		}

		r, err := c.httpPostRequest(endpoint, data)

		if err != nil {
			return nil, err
		}

		response := make(map[string]interface{})
		err = json.Unmarshal(r, &response)

		if err != nil {
			return nil, err
		}

		switch response["status"].(type) {
		case string:
			var taskResult model.TaskResultResponse
			err = json.Unmarshal(r, &taskResult)

			if err != nil {
				return nil, err
			}

			if taskResult.Status == "ready" {
				return &taskResult, nil
			}

			time.Sleep(2 * time.Second)
		default:
			return nil, ErrorMapping[response["errorCode"].(string)]
		}
	}
}

func (c *Client) createTask(task Tasker) (*model.TaskResponse, error) {
	endpoint := fmt.Sprintf("%s/createTask", c.baseAddress)

	data := model.TaskRequest{
		ClientKey:    c.Key,
		Task:         task,
		SoftId:       c.SoftId,
		LanguagePool: c.Language,
		CallbackUrl:  c.CallbackUrl,
	}

	r, err := c.httpPostRequest(endpoint, data)

	if err != nil {
		return nil, err
	}

	response := make(map[string]interface{})
	err = json.Unmarshal(r, &response)

	if err != nil {
		return nil, err
	}

	switch response["taskId"].(type) {
	case float64:
		var taskResponse model.TaskResponse
		err = json.Unmarshal(r, &taskResponse)

		if err != nil {
			return nil, err
		}

		return &taskResponse, nil
	default:
		return nil, ErrorMapping[response["errorCode"].(string)]
	}
}

func (c *Client) httpPostRequest(endpoint string, data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	resp, err := c.client.Post(endpoint, "application/json", bytes.NewBuffer(b))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if c.Verbose {
		log.Println(string(r))
	}

	return r, nil
}

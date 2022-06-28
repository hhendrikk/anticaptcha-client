package model

type TaskResultRequest struct {
	ClientKey string  `json:"clientKey"`
	TaskId    float64 `json:"taskId"`
}

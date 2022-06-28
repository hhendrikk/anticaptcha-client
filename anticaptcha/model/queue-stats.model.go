package model

type QueueId int8

const (
	ImageToTextEnglish                QueueId = 1
	ImageToTextRussian                QueueId = 2
	RecaptchaV2WithProxy              QueueId = 5
	RecaptchaV2WithoutProxy           QueueId = 6
	FuncaptchaWithProxy               QueueId = 7
	FuncaptchaWithoutProxy            QueueId = 10
	GeeTestWithProxy                  QueueId = 12
	GeeTestWithoutProxy               QueueId = 13
	RecaptchaV3s03                    QueueId = 18
	RecaptchaV3s07                    QueueId = 19
	RecaptchaV3s09                    QueueId = 20
	HCaptchaWithProxy                 QueueId = 21
	HCaptchaWithoutProxy              QueueId = 22
	RecaptchaEnterpriseV2WithProxy    QueueId = 23
	RecaptchaEnterpriseV2WithoutProxy QueueId = 24
	AntiGateTask                      QueueId = 25
)

type QueueStats struct {
	// Number of idle workers online, waiting for a task
	Waiting int `json:"waiting"`
	// Queue load as a percentage
	Load float64 `json:"load"`
	// Average task solution cost in USD
	Bid float64 `json:"bid"`
	// Average task solution speed in seconds
	Speed float64 `json:"speed"`
	// Total number of workers
	Total int `json:"total"`
}

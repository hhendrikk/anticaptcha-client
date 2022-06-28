package model

type TaskResultSolutionResponse struct {
	// ImageToTextTask
	Text string `json:"text"`

	// ImageToTextTask
	Url string `json:"url"`

	// Recaptcha
	GRecaptchaResponse string `json:"gRecaptchaResponse"`

	// Funcaptcha
	Token string `json:"token"`

	// GeeTest V3
	Challenge string `json:"challenge"`
	// GeeTest V3
	Validate string `json:"validate"`
	// GeeTest V3
	Seccode string `json:"seccode"`

	// GeeTest V4
	CaptchaId string `json:"captcha_id"`
	// GeeTest V4
	LotNumber string `json:"lot_number"`
	// GeeTest V4
	PassToken string `json:"pass_token"`
	// GeeTest V4
	GenTime string `json:"gen_time"`
	// GeeTest V4
	CaptchaOutput string `json:"captcha_output"`
}

type TaskResultResponse struct {
	ErrorId    int                        `json:"errorId"`
	Status     string                     `json:"status"`
	Solution   TaskResultSolutionResponse `json:"solution"`
	Cost       string                     `json:"cost"`
	Ip         string                     `json:"ip"`
	CreateTime float64                    `json:"createTime"`
	EndTime    float64                    `json:"endTime"`
	SolveCount int                        `json:"solveCount"`
}

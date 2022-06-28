package model

type TaskResultSolutionResponse struct {
	Text               string `json:"text"`
	Url                string `json:"url"`
	GRecaptchaResponse string `json:"gRecaptchaResponse"`
	Token              string `json:"token"`
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

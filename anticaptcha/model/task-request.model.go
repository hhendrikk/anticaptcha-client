package model

type TaskRequest struct {
	ClientKey    string      `json:"clientKey"`
	Task         interface{} `json:"task"`
	SoftId       int         `json:"softId"`
	LanguagePool string      `json:"languagePool,omitempty"`
	CallbackUrl  string      `json:"callbackUrl,omitempty"`
}

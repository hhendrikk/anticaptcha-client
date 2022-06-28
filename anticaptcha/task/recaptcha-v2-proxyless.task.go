package task

type RecaptchaV2ProxylessTask struct {
	// Previous name of task type: NoCaptchaTaskProxyless.
	// It will be supported forever, no need to update your code.
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// Recaptcha website key.
	WebsiteKey string `json:"websiteKey"`
	// Value of 'data-s' parameter. Applies only to Recaptchas on Google web sites.
	RecaptchaDataSValue string `json:"recaptchaDataSValue"`
	// Specify whether or not Recaptcha is invisible.
	// This will render an appropriate widget for our workers.
	IsInvisible bool `json:"isInvisible"`
}

func NewRecaptchaV2ProxylessTask(websiteURL string, websiteKey string) *RecaptchaV2ProxylessTask {
	return &RecaptchaV2ProxylessTask{
		Type:       "RecaptchaV2TaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

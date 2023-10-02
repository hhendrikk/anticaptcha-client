package task

type HCaptchaProxylessTask struct {
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// hCaptcha sitekey
	WebsiteKey string `json:"websiteKey"`
	// Specify whether or not Hcaptcha is invisible.
	// This will render an appropriate widget for our workers.
	IsInvisible bool `json:"isInvisible"`
	// 	Additional parameters which we'll use to render Hcaptcha widget for Enterprise version.
	// Property		Type	Required
	// rqdata		String	No
	// sentry		Boolean	No
	// apiEndpoint	String	No
	// endpoint		String	No
	// reportapi	String	No
	// assethost	String	No
	// imghost		String	No
	EnterprisePayload map[string]interface{} `json:"enterprisePayload"`
}

// hCaptcha devs call their captcha "a drop-in replacement for Recaptcha".
// We tried to create the same thing in our API, so task properties are absolutely the same as in RecaptchaV2TaskProxyless except for the "type" property.
func NewHCaptchaProxylessTask(websiteURL string, websiteKey string, userAgent string) *HCaptchaProxylessTask {
	return &HCaptchaProxylessTask{
		Type:       "HCaptchaTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

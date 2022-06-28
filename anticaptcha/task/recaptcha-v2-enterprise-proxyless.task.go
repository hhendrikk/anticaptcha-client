package task

type RecaptchaV2EnterpriseProxylessTask struct {
	// Previous name of task type: NoCaptchaTaskProxyless.
	// It will be supported forever, no need to update your code.
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// Recaptcha website key.
	WebsiteKey string `json:"websiteKey"`
	// Additional parameters which should be passed to "grecaptcha.enterprise.render" method along with sitekey
	EnterprisePayload map[string]string `json:"enterprisePayload"`
	// Use this parameter to send the domain name from which the Recaptcha script should be served.
	// Can have only one of two values: "www.google.com" or "www.recaptcha.net".
	// Do not use this parameter unless you understand what you are doing.
	ApiDomain string `json:"apiDomain"`
}

// This type of task is for solving Google Recaptcha Enterprise V2 from the worker's IP address.
// It is mostly similar to RecaptchaV2TaskProxyless, except tasks are solved using an Enterprise API and assigned to workers with the best Recaptcha V3 score.
func NewRecaptchaV2EnterpriseProxylessTask(websiteURL string, websiteKey string) *RecaptchaV2EnterpriseProxylessTask {
	return &RecaptchaV2EnterpriseProxylessTask{
		Type:       "RecaptchaV2EnterpriseTaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
	}
}

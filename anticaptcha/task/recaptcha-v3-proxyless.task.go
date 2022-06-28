package task

type RecaptchaV3TaskProxyless struct {
	// RecaptchaV3TaskProxyless
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// Recaptcha website key
	WebsiteKey string `json:"websiteKey"`
	// Filters workers with a particular score. It can have one of the following values:
	// 0.3
	// 0.7
	// 0.9
	MinScore float64 `json:"minScore"`
	// Recaptcha's "action" value. Website owners use this parameter to define what users are doing on the page.
	// Example:

	// grecaptcha.execute('site_key', {action:'login_test'})
	PageAction string `json:"pageAction"`
	// Set this flag to "true" if you need this V3 solved with Enterprise API.
	// Default value is "false" and Recaptcha is solved with non-enterprise API.
	// Can be determined by a javascript call like in the following example:

	// grecaptcha.enterprise.execute('site_key', {..})
	IsEnterprise bool `json:"isEnterprise"`
	// Use this parameter to send the domain name from which the Recaptcha script should be served.
	// Can have only one of two values: "www.google.com" or "www.recaptcha.net".
	// Do not use this parameter unless you understand what you are doing.
	ApiDomain string `json:"apiDomain"`
}

// As V3 Enterprise is virtually the same as V3 non-Enterprise, we decided to roll out it’s support within the usual V3 tasks.
// Differences between V3 Enterprise and V3 non-Enterprise:
// widget code is loaded via enterprise.js (vs api.js)
// user score retrieval is made with grecaptcha.enterprise.execute call (vs grecaptcha.execute)
// Thus to mark your Enterprise V3 task, you simply need to add a flag "isEnterprise": true to your non-Enterprise V3 payload
func NewRecaptchaV3TaskProxyless(websiteURL string, websiteKey string, minScore float64, pageAction string) *RecaptchaV3TaskProxyless {
	return &RecaptchaV3TaskProxyless{
		Type:       "RecaptchaV3TaskProxyless",
		WebsiteURL: websiteURL,
		WebsiteKey: websiteKey,
		MinScore:   minScore,
		PageAction: pageAction,
	}
}

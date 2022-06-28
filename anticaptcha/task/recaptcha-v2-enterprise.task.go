package task

type RecaptchaV2EnterpriseTask struct {
	//Previous task name: 'NoCaptchaTask'.
	// We will continue supporting old task name forever, no need to update your code.
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
	// Type of proxy
	// http - usual http/https proxy
	// socks4 - socks4 proxy
	// socks5 - socks5 proxy
	ProxyType string `json:"proxyType"`
	// Proxy IP address ipv4/ipv6.
	// No host names or IP addresses from local networks.
	ProxyAddress string `json:"proxyAddress"`
	// Proxy port
	ProxyPort int `json:"proxyPort"`
	// Login for proxy which requires authorization (basic)
	ProxyLogin string `json:"proxyLogin"`
	// Proxy password
	ProxyPassword string `json:"proxyPassword"`
	// Browser's User-Agent used in emulation.
	// You must use a modern-browser signature; otherwise, Google will ask you to "update your browser".
	UserAgent string `json:"userAgent"`
	// Additional cookies that we should use in Google domains.
	Cookies string `json:"cookies"`
}

// This type of task is for solving Google Recaptcha Enterprise V2 using the provided proxy.
// It is mostly similar to RecaptchaV2Task, except tasks are solved using an Enterprise API and assigned to workers with the best Recaptcha V3 score.
func NewRecaptchaV2EnterpriseTask(websiteURL string, websiteKey string, proxyType string, proxyAddress string, proxyPort int, userAgent string) *RecaptchaV2EnterpriseTask {
	return &RecaptchaV2EnterpriseTask{
		Type:         "RecaptchaV2EnterpriseTask",
		WebsiteURL:   websiteURL,
		WebsiteKey:   websiteKey,
		ProxyType:    proxyType,
		ProxyAddress: proxyAddress,
		ProxyPort:    proxyPort,
		UserAgent:    userAgent,
	}
}

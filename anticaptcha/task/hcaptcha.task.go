package task

type HCaptchaTask struct {
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// hCaptcha sitekey
	WebsiteKey string `json:"websiteKey"`
	// Type of proxy
	// http - usual http/https proxy
	// socks4 - socks4 proxy
	// socks5 - socks5 proxy
	ProxyType string `json:"proxyType"`
	// Proxy IP address ipv4/ipv6. No host names or IP addresses from local networks.
	ProxyAddress string `json:"proxyAddress"`
	// Proxy port
	ProxyPort int `json:"proxyPort"`
	// Login for proxy which requires authorization (basic)
	ProxyLogin string `json:"proxyLogin"`
	// Proxy password
	ProxyPassword string `json:"proxyPassword"`
	// Specify whether or not Hcaptcha is invisible.
	// This will render an appropriate widget for our workers.
	IsInvisible bool `json:"isInvisible"`
	// 	Additional parameters which we'll use to render Hcaptcha widget for Enterprise version.
	EnterprisePayload map[string]interface{} `json:"enterprisePayload"`
}

// hCaptcha devs call their captcha "a drop-in replacement for Recaptcha".
// We tried to create the same thing in our API, so task properties are absolutely the same as in RecaptchaV2Task except for the "type" property.
// IMPORTANT: hCaptcha seems to have a limit on solved tasks from one IP: about 3 items per 12 hours.
// Take this into account when you build the solving process through your proxy.
func NewHCaptchaTask(websiteURL string, websiteKey string, proxyType string, proxyAddress string, proxyPort int, userAgent string) *HCaptchaTask {
	return &HCaptchaTask{
		Type:         "HCaptchaTask",
		WebsiteURL:   websiteURL,
		WebsiteKey:   websiteKey,
		ProxyType:    proxyType,
		ProxyAddress: proxyAddress,
		ProxyPort:    proxyPort,
	}
}

package task

type RecaptchaV2Task struct {
	//Previous task name: 'NoCaptchaTask'.
	// We will continue supporting old task name forever, no need to update your code.
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// Recaptcha website key.
	WebsiteKey string `json:"websiteKey"`
	// Value of 'data-s' parameter. Applies only to Recaptchas on Google web sites.
	RecaptchaDataSValue string `json:"recaptchaDataSValue"`
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
	// Specify whether or not Recaptcha is invisible.
	// This will render an appropriate widget for our workers.
	IsInvisible bool `json:"isInvisible"`
}

func NewRecaptchaV2Task(websiteURL string, websiteKey string, proxyType string, proxyAddress string, proxyPort int, userAgent string) *RecaptchaV2Task {
	return &RecaptchaV2Task{
		Type:         "RecaptchaV2Task",
		WebsiteURL:   websiteURL,
		WebsiteKey:   websiteKey,
		ProxyType:    proxyType,
		ProxyAddress: proxyAddress,
		ProxyPort:    proxyPort,
		UserAgent:    userAgent,
	}
}

package task

type FunCaptchaTask struct {
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// Arkose Labs public key
	WebsitePublicKey string `json:"websitePublicKey"`
	// Custom Arkose Labs subdomain from which the Javascript widget is loaded.
	// Required for some cases, but most Arkose Labs integrations run without it.
	FuncaptchaApiJSSubdomain string `json:"funcaptchaApiJSSubdomain"`
	// An additional parameter that may be required by Arkose Labs implementation.
	// Use this property to send "blob" value as an object converted to a string.
	// See an example of what it might look like. {"\blob\":\"HERE_COMES_THE_blob_VALUE\"}
	Data string `json:"data"`
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
}

// This type of task solves arkoselabs.com puzzles in our workers' browsers.
// Your app submits the website address and public key and receives a token after task completion.
// Use this token to submit the form with the Arkose Labs captcha.
func NewFunCaptchaTask(websiteURL string, WebsitePublicKey string, proxyType string, proxyAddress string, proxyPort int, userAgent string) *FunCaptchaTask {
	return &FunCaptchaTask{
		Type:             "FunCaptchaTask",
		WebsiteURL:       websiteURL,
		WebsitePublicKey: WebsitePublicKey,
		ProxyType:        proxyType,
		ProxyAddress:     proxyAddress,
		ProxyPort:        proxyPort,
		UserAgent:        userAgent,
	}
}

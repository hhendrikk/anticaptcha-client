package task

type GeeTestTask struct {
	Type string `json:"type"`
	// Address of a target web page. Can be located anywhere on the web site, even in a member area.
	// Our workers don't navigate there but simulate the visit instead.
	WebsiteURL string `json:"websiteURL"`
	// The domain public key, rarely updated.
	Gt string `json:"gt"`
	// Changing token key.
	// Make sure you grab a fresh one for each captcha; otherwise, you'll be charged for an error task.
	Challenge string `json:"challenge"`
	// Optional API subdomain. May be required for some implementations.
	GeetestApiServerSubdomain string `json:"geetestApiServerSubdomain"`
	// Required for some implementations.
	// Send the JSON encoded into a string. The value can be traced in browser developer tools.
	// Put a breakpoint before calling the "initGeetest" function.
	GeetestGetLib string `json:"geetestGetLib"`
	// Version number.
	// Default version is 3.
	// Supported versions: 3 and 4.
	Version int `json:"version"`
	// Additional initialization parameters for version 4
	InitParameters map[string]interface{} `json:"initParameters"`
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

// This type of task solves GeeTest captchas in our workers' browsers.
// Your app submits the website address, gt key, challenge key and after task completion receives a solution consisting of 3 tokens.
// For version GeeTest version 4 output consists of 5 values and challenge key is not required.
func NewGeeTestTask(websiteURL string, gt string, challenge string, proxyType string, proxyAddress string, proxyPort int, userAgent string) *GeeTestTask {
	return &GeeTestTask{
		Type:         "GeeTestTask",
		WebsiteURL:   websiteURL,
		Gt:           gt,
		Challenge:    challenge,
		ProxyType:    proxyType,
		ProxyAddress: proxyAddress,
		ProxyPort:    proxyPort,
		UserAgent:    userAgent,
	}
}

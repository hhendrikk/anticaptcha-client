package task

type FunCaptchaProxylessTask struct {
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
}

// This type of task solves Arkose Labs captcha (or Funcaptcha) without proxy.
// Task will be executed using our own proxy servers and/or workers' IP addresses.
// Arkose Labs API provides information to the website owner about the solver's IP address.
// However it's worth trying first to bypass captcha without proxy, and if it doesn't work - switch to FuncaptchaTask with proxy.
func NewFunCaptchaProxylessTask(websiteURL string, WebsitePublicKey string) *FunCaptchaProxylessTask {
	return &FunCaptchaProxylessTask{
		Type:             "FunCaptchaTaskProxyless",
		WebsiteURL:       websiteURL,
		WebsitePublicKey: WebsitePublicKey,
	}
}

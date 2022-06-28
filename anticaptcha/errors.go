package anticaptcha

import (
	"fmt"
)

var ErrorMapping = map[string]error{
	"ERROR_KEY_DOES_NOT_EXIST":              ErrKeyDoesNotExist,
	"ERROR_NO_SLOT_AVAILABLE":               ErrNoSlotAvailable,
	"ERROR_ZERO_CAPTCHA_FILESIZE":           ErrZeroCaptchaFilesize,
	"ERROR_TOO_BIG_CAPTCHA_FILESIZE":        ErrTooBigCaptchaFilesize,
	"ERROR_ZERO_BALANCE":                    ErrZeroBalance,
	"ERROR_IP_NOT_ALLOWED":                  ErrIpNotAllowed,
	"ERROR_CAPTCHA_UNSOLVABLE":              ErrCaptchaUnsolvable,
	"ERROR_BAD_DUPLICATES":                  ErrBadDuplicates,
	"ERROR_NO_SUCH_METHOD":                  ErrNoSuchMethod,
	"ERROR_IMAGE_TYPE_NOT_SUPPORTED":        ErrImageTypeNotSupported,
	"ERROR_NO_SUCH_CAPCHA_ID":               ErrNoSuchCapchaId,
	"ERROR_IP_BLOCKED":                      ErrIpBlocked,
	"ERROR_TASK_ABSENT":                     ErrTaskAbsent,
	"ERROR_TASK_NOT_SUPPORTED":              ErrTaskNotSupported,
	"ERROR_INCORRECT_SESSION_DATA":          ErrIncorrectSessionData,
	"ERROR_PROXY_CONNECT_REFUSED":           ErrProxyConnectRefused,
	"ERROR_PROXY_CONNECT_TIMEOUT":           ErrProxyConnectTimeout,
	"ERROR_PROXY_READ_TIMEOUT":              ErrProxyReadTimeout,
	"ERROR_PROXY_BANNED":                    ErrProxyBanned,
	"ERROR_PROXY_TRANSPARENT":               ErrProxyTransparent,
	"ERROR_RECAPTCHA_TIMEOUT":               ErrRecaptchaTimeout,
	"ERROR_RECAPTCHA_INVALID_SITEKEY":       ErrRecaptchaInvalidSitekey,
	"ERROR_RECAPTCHA_INVALID_DOMAIN":        ErrRecaptchaInvalidDomain,
	"ERROR_RECAPTCHA_OLD_BROWSER":           ErrRecaptchaOldBrowser,
	"ERROR_TOKEN_EXPIRED":                   ErrTokenExpired,
	"ERROR_PROXY_HAS_NO_IMAGE_SUPPORT":      ErrProxyHasNoImageSupport,
	"ERROR_PROXY_INCOMPATIBLE_HTTP_VERSION": ErrProxyIncompatibleHttpVersion,
	"ERROR_PROXY_NOT_AUTHORISED":            ErrProxyNotAuthorised,
	"ERROR_INVISIBLE_RECAPTCHA":             ErrInvisibleRecaptcha,
	"ERROR_FAILED_LOADING_WIDGET":           ErrFailedLoadingWidget,
	"ERROR_VISIBLE_RECAPTCHA":               ErrVisibleRecaptcha,
	"ERROR_ALL_WORKERS_FILTERED":            ErrAllWorkersFiltered,
	"ERROR_ACCOUNT_SUSPENDED":               ErrAccountSuspended,
	"ERROR_TEMPLATE_NOT_FOUND":              ErrTemplateNotFound,
	"ERROR_TASK_CANCELED":                   ErrTaskCanceled,
}

var (
	ErrKeyDoesNotExist              = fmt.Errorf("id: %d, code: %s, description: %s", 1, "ERROR_KEY_DOES_NOT_EXIST", "Authorization key not found in the system. Make sure you copied it correctly without spaces or tabulation signs.")
	ErrNoSlotAvailable              = fmt.Errorf("id: %d, code: %s, description: %s", 2, "ERROR_NO_SLOT_AVAILABLE", "No idle captcha workers are available at the moment; customers should increase their maximum bid in the API settings in the client area or choose less busy hours to create tasks. You'll find more information about how bids work in FAQ.")
	ErrZeroCaptchaFilesize          = fmt.Errorf("id: %d, code: %s, description: %s", 3, "ERROR_ZERO_CAPTCHA_FILESIZE", "The size of the captcha you are uploading is less than 100 bytes.")
	ErrTooBigCaptchaFilesize        = fmt.Errorf("id: %d, code: %s, description: %s", 4, "ERROR_TOO_BIG_CAPTCHA_FILESIZE", "The size of the captcha you are uploading is more than 500,000 bytes.")
	ErrZeroBalance                  = fmt.Errorf("id: %d, code: %s, description: %s", 10, "ERROR_ZERO_BALANCE", " has zero or negative balance.")
	ErrIpNotAllowed                 = fmt.Errorf("id: %d, code: %s, description: %s", 11, "ERROR_IP_NOT_ALLOWED", "You are not allowed to make requests from your IP with your current account key. Please refer to IP list section located in security settings inside client area.")
	ErrCaptchaUnsolvable            = fmt.Errorf("id: %d, code: %s, description: %s", 12, "ERROR_CAPTCHA_UNSOLVABLE", "5 different workers could not solve the Captcha. Customers are charged for these tasks because our workers spend time performing them.")
	ErrBadDuplicates                = fmt.Errorf("id: %d, code: %s, description: %s", 12, "ERROR_BAD_DUPLICATES", "100% recognition feature did not work because of insufficient number of guesses.")
	ErrNoSuchMethod                 = fmt.Errorf("id: %d, code: %s, description: %s", 14, "ERROR_NO_SUCH_METHOD", "Request made to API with a method that does not exist. This usually happens when programmers mistype method names.")
	ErrImageTypeNotSupported        = fmt.Errorf("id: %d, code: %s, description: %s", 15, "ERROR_IMAGE_TYPE_NOT_SUPPORTED", "Could not determine captcha file type from its exif header or the image type is not supported. The only allowed formats are JPG, GIF, PNG. Images must contain an EXIF header containing information about the image type.")
	ErrNoSuchCapchaId               = fmt.Errorf("id: %d, code: %s, description: %s", 16, "ERROR_NO_SUCH_CAPCHA_ID", "The captcha you are requesting does not exist in your active captchas list or has expired. Captchas are removed from the API 60 seconds after the worker completes the task. During this period, your app should send all task result polls and correct/incorrect reporting requests.")
	ErrIpBlocked                    = fmt.Errorf("id: %d, code: %s, description: %s", 21, "ERROR_IP_BLOCKED", "Your IP is blocked due to improper use of API. Check the reason here.")
	ErrTaskAbsent                   = fmt.Errorf("id: %d, code: %s, description: %s", 22, "ERROR_TASK_ABSENT", "\"task\" property is empty or not set in the createTask method.")
	ErrTaskNotSupported             = fmt.Errorf("id: %d, code: %s, description: %s", 23, "ERROR_TASK_NOT_SUPPORTED", "Task type is not supported or typed incorrectly. Please check \"type\" property in the task object.")
	ErrIncorrectSessionData         = fmt.Errorf("id: %d, code: %s, description: %s", 24, "ERROR_INCORRECT_SESSION_DATA", "Some of the required values for successive user emulation are missing. API output contains more details about what is missing.")
	ErrProxyConnectRefused          = fmt.Errorf("id: %d, code: %s, description: %s", 25, "ERROR_PROXY_CONNECT_REFUSED", "Could not connect to task proxy, connection refused. Learn more about working with proxy in this FAQ section.")
	ErrProxyConnectTimeout          = fmt.Errorf("id: %d, code: %s, description: %s", 26, "ERROR_PROXY_CONNECT_TIMEOUT", "Could not connect to task proxy, connection timed out. Learn more about working with proxy in this FAQ section.")
	ErrProxyReadTimeout             = fmt.Errorf("id: %d, code: %s, description: %s", 27, "ERROR_PROXY_READ_TIMEOUT", "Reading timeout of task's proxy, Learn more about working with proxy in this FAQ section.")
	ErrProxyBanned                  = fmt.Errorf("id: %d, code: %s, description: %s", 28, "ERROR_PROXY_BANNED", "Proxy IP banned by target service.")
	ErrProxyTransparent             = fmt.Errorf("id: %d, code: %s, description: %s", 29, "ERROR_PROXY_TRANSPARENT", "Task denied at proxy checking state. Proxy must be non-transparent to hide our server IP. Use our proxy checker tool to debug your proxy.")
	ErrRecaptchaTimeout             = fmt.Errorf("id: %d, code: %s, description: %s", 30, "ERROR_RECAPTCHA_TIMEOUT", "Recaptcha task timeout, probably due to slow proxy server or Google server")
	ErrRecaptchaInvalidSitekey      = fmt.Errorf("id: %d, code: %s, description: %s", 31, "ERROR_RECAPTCHA_INVALID_SITEKEY", "Captcha provider reported that the site key is invalid.")
	ErrRecaptchaInvalidDomain       = fmt.Errorf("id: %d, code: %s, description: %s", 32, "ERROR_RECAPTCHA_INVALID_DOMAIN", "Captcha provider reported that the domain for this site key is invalid.")
	ErrRecaptchaOldBrowser          = fmt.Errorf("id: %d, code: %s, description: %s", 33, "ERROR_RECAPTCHA_OLD_BROWSER", "Captcha provider reported that the browser user-agent is not compatible with their javascript")
	ErrTokenExpired                 = fmt.Errorf("id: %d, code: %s, description: %s", 34, "ERROR_TOKEN_EXPIRED", "Captcha provider server reported that the additional variable token has expired. Please try again with a new token.")
	ErrProxyHasNoImageSupport       = fmt.Errorf("id: %d, code: %s, description: %s", 35, "ERROR_PROXY_HAS_NO_IMAGE_SUPPORT", "Proxy does not support transfer of image data from Google servers. Use our proxy checker tool to debug your proxy.")
	ErrProxyIncompatibleHttpVersion = fmt.Errorf("id: %d, code: %s, description: %s", 36, "ERROR_PROXY_INCOMPATIBLE_HTTP_VERSION", "Proxy does not support long GET requests with length about 2000 bytes and does not support SSL connections. Use our proxy checker tool to debug your proxy.")
	ErrProxyNotAuthorised           = fmt.Errorf("id: %d, code: %s, description: %s", 49, "ERROR_PROXY_NOT_AUTHORISED", "Proxy login and password are incorrect. Use our proxy checker tool to debug your proxy.")
	ErrInvisibleRecaptcha           = fmt.Errorf("id: %d, code: %s, description: %s", 51, "ERROR_INVISIBLE_RECAPTCHA", "An attempt was made to solve an Invisible Recaptcha as if it was a regular one.")
	ErrFailedLoadingWidget          = fmt.Errorf("id: %d, code: %s, description: %s", 52, "ERROR_FAILED_LOADING_WIDGET", "Could not load captcha provider widget in worker browser. Please try sending a new task.")
	ErrVisibleRecaptcha             = fmt.Errorf("id: %d, code: %s, description: %s", 53, "ERROR_VISIBLE_RECAPTCHA", "Attempted solution of usual Recaptcha V2 as Recaptcha V2 invisible. Remove flag 'isInvisible' from the API payload.")
	ErrAllWorkersFiltered           = fmt.Errorf("id: %d, code: %s, description: %s", 54, "ERROR_ALL_WORKERS_FILTERED", "No workers left that were not filtered by reportIncorrectRecaptcha method.")
	ErrAccountSuspended             = fmt.Errorf("id: %d, code: %s, description: %s", 55, "ERROR_ACCOUNT_SUSPENDED", "The system has suspended your account for a significant reason. Contact support for details.")
	ErrTemplateNotFound             = fmt.Errorf("id: %d, code: %s, description: %s", 56, "ERROR_TEMPLATE_NOT_FOUND", "AntiGate template not found by its name during task creation.")
	ErrTaskCanceled                 = fmt.Errorf("id: %d, code: %s, description: %s", 57, "ERROR_TASK_CANCELED", "AntiGate task was canceled by worker. See \"errorDescription\" field for the cancellation reason.")
)


# AntiCaptcha Client for Golang

Golang library to usage anti-captcha.com.

## API Documentation

[Documentation](https://anti-captcha.com/apidoc)


## Installation

```bash
  go get -u github.com/hhendrikk/anticaptcha-client@v1.0.1
```

## Get balance
```golang
package main

import (
	"fmt"
	"log"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha"
)

func main() {
	client := anticaptcha.NewClient("YOUR_API_KEY")
	result, err := client.GetBalance()

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Balance: %0.2f\n", result.Balance)

}
```

## ImageToText

- From Url
```golang
package main

import (
	"fmt"
	"log"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha"
	"github.com/hhendrikk/anticaptcha-client/anticaptcha/task"
)

func main() {
	client := anticaptcha.NewClient("YOUR_API_KEY")

	imageToTextTask, err := task.NewImageToTextTaskFromUrl("https://www.anti-captcha.com/images/test.jpg")

	if err != nil {
		log.Fatalln(err.Error())
	}

	result, err := client.GetResult(imageToTextTask)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Text: %s, Url: %s\n", result.Solution.Text, result.Solution.Url)

}

```

- Image base64

```golang
package main

import (
	"fmt"
	"log"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha"
	"github.com/hhendrikk/anticaptcha-client/anticaptcha/task"
)

func main() {
	client := anticaptcha.NewClient("YOUR_API_KEY")

	imageToTextTask := task.NewImageToTextTask("base64")

	result, err := client.GetResult(imageToTextTask)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Text: %s, Url: %s\n", result.Solution.Text, result.Solution.Url)

}
```

## Recaptcha V2

- Proxyless

```golang
package main

import (
	"fmt"
	"log"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha"
	"github.com/hhendrikk/anticaptcha-client/anticaptcha/task"
)

func main() {
	client := anticaptcha.NewClient("YOUR_API_KEY")

	recaptchaV2 := task.NewRecaptchaV2ProxylessTask("websiterURL", "websiteKey")

	result, err := client.GetResult(recaptchaV2)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Recaptcha Token: %s\n", result.Solution.GRecaptchaResponse)

}
```

## Recaptcha V3

```golang
package main

import (
	"fmt"
	"log"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha"
	"github.com/hhendrikk/anticaptcha-client/anticaptcha/task"
)

func main() {
	client := anticaptcha.NewClient("YOUR_API_KEY")

	recaptchaV3 := task.NewRecaptchaV3TaskProxyless("websiterURL", "websiteKey", 0.3, "pageAction")

	result, err := client.GetResult(recaptchaV3)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Recaptcha Token: %s\n", result.Solution.GRecaptchaResponse)

}
```

## FunCaptcha

- Proxyless

```golang
ppackage main

import (
	"fmt"
	"log"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha"
	"github.com/hhendrikk/anticaptcha-client/anticaptcha/task"
)

func main() {
	client := anticaptcha.NewClient("YOUR_API_KEY")

	funCaptcha := task.NewFunCaptchaProxylessTask("websiterURL", "websiteKey")

	result, err := client.GetResult(funCaptcha)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("Token: %s\n", result.Solution.Token)

}
```

## Support

- GeeTestTask
- GeeTestTaskProxyless
- HCaptchaTask
- HCaptchaTaskProxyless
- ReportIncorrectImageCaptcha
- ReportIncorrectRecaptcha
- ReportCorrectRecaptcha
- ReportIncorrectHcaptcha



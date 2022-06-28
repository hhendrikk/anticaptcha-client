package main

import (
	"fmt"
	"os"

	"github.com/hhendrikk/anticaptcha-client/anticaptcha/service"
	"github.com/hhendrikk/anticaptcha-client/anticaptcha/task"
)

func main() {
	client := service.NewClient("")
	task := task.NewRecaptchaV2ProxylessTask("https://www.google.com/recaptcha/api2/demo", "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-")

	res, err := client.GetResult(task)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("gResponseToken: %s\n", res.Solution.GRecaptchaResponse)
}

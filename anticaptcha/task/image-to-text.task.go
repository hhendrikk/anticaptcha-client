package task

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

type ImageToTextTask struct {
	// Required
	// Defines the type of task.
	Type string `json:"type"`
	// Required
	// File body encoded in base64. Make sure to send it without line breaks. Do not include 'data:image/png,' or similar tags, only clean base64!
	Body string `json:"body"`
	// Optional
	// false - no requirements
	// true - requires workers to enter an answer with at least one "space". If there are no spaces, they will skip the task, so use it with caution.
	Phrase bool `json:"phrase"`
	// Optional
	// false - no requirements
	// true - workers see a special mark indicating that the answer must be entered with case sensitivity.
	Case bool `json:"case"`
	// Optional
	// 0 - no requirements
	// 1 - only numbers are allowed
	// 2 - any letters are allowed except numbers
	Numeric int `json:"numeric"`
	// Optional
	// false - no requirements
	// true - workers see a special mark telling them the answer must be calculated
	Math bool `json:"math"`
	// Optional
	// 0 - no requirements
	// >1 - defines minimum length of the answer
	MinLength int `json:"minLength"`
	// Optional
	// 0 - no requirements
	// >1 - defines maximum length of the answer
	MaxLength int `json:"maxLength"`
	// Optional
	// Additional comments for workers like "enter red text".
	// The result is not guaranteed and is totally up to the worker.
	Comment string `json:"comment"`
	// Optional
	// Parameter to distinguish source of image captchas in spending statistics.
	WebsiteURL string `json:"websiteURL"`
}

// Post an image body and receive text from it. Text can only contain digits, letters, special characters and a space.
// GIF animations are supported, up to 500kb.
// Custom captchas like "find a cat in this series of images and enter its number" are not supported.
// body File body encoded in base64.
// Make sure to send it without line breaks. Do not include 'data:image/png,' or similar tags, only clean base64!
func NewImageToTextTask(body string) *ImageToTextTask {
	return &ImageToTextTask{
		Type: "ImageToTextTask",
		Body: body,
	}
}

// Post an image body and receive text from it. Text can only contain digits, letters, special characters and a space.
// GIF animations are supported, up to 500kb.
// Custom captchas like "find a cat in this series of images and enter its number" are not supported.
func NewImageToTextTaskFromUrl(url string) (*ImageToTextTask, error) {
	client := http.Client{}
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	b64 := base64.StdEncoding.EncodeToString(body)

	r := &ImageToTextTask{
		Type:       "ImageToTextTask",
		Body:       b64,
		Phrase:     false,
		Case:       true,
		Numeric:    0,
		Math:       false,
		MinLength:  0,
		MaxLength:  0,
		Comment:    "",
		WebsiteURL: "",
	}

	return r, nil
}

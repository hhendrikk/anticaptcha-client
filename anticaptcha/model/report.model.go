package model

type ReportResponse struct {
	Status string `json:"status"`
}

type ReportTaskType string

const (
	// Send complaint on an image captcha
	// Complaints are accepted for image captchas only. Your complaint will be checked by 3 moderators,
	// 2 of them must confirm it. Only then will you get a full refund.
	// If you have fewer than a 50% mistake confirmation ratio, your reports will be ignored.
	// Reports must be sent within 60 seconds of task completion.
	// If you send a report later, the API will return ERROR_NO_SUCH_CAPCHA_ID error.
	// Only one report can be sent per task.
	// Why 50%? Your reports must be very precise and you must be 100% certain that they are correct.
	// You can't just report every other captcha; you must code your software with strict testing standards and be sure that the target website does not detect you some other way.
	// Please note that it's not possible to reset confirmation statistics.
	// Please note that we accept complaints only for captchas in English language.
	// Complaints for Russian language are not accepted and API will return ERROR_NO_SUCH_CAPCHA_ID error code.
	ReportIncorrectImageCaptcha ReportTaskType = "reportIncorrectImageCaptcha"

	// Send complaint on a Recaptcha token
	// Complaints are accepted for V2 and V3 Recaptchas only, including Enterprise Recaptcha.
	// It is important to read the following description; otherwise, our system may ban your reports.
	// Because we can't check your report as we do with image captchas, your complaint goes first to our backend for statistics analysis and is only accepted if the results are positive.
	// We compare your report statistics with the statistics of other trusted customers. If your report rate differs significantly from other reports, then it will be ignored for several days.
	// Thus, to ensure the best results, you should always monitor your automation processes and send reports only if you are 100% certain the Recaptcha was wrong.
	// Not all reports are accepted. To calculate your average fail rate as accurately as possible, a minimum 100 recaptcha tasks per account must be sent for recognition per 24 hours.
	// Reports must be sent within 60 seconds of task completion.
	// If you send a report later, the API will return ERROR_NO_SUCH_CAPTCHA_ID error.
	// Only one report can be sent per task.
	ReportIncorrectRecaptcha ReportTaskType = "reportIncorrectRecaptcha"

	// Report correctly solved Recaptcha tokens
	// Use this method along with reportIncorrectRecaptcha for your Recaptcha V3 and Recaptcha V2 Enterprise tasks.
	// Reports for Recaptcha V2 are currently accepted but eventually ignored. Our quality for V2 is about 99%, and you don't need to build a whitelist of successful workers.
	// Behind the scenes, we will put the worker you've reported on a whitelist, and at the next round of assigning workers to your captchas, our system will match this whitelist with your captchas.
	// If any of the reported workers are online and idle, they will be put at the front of the line for assignment to your new task.
	// The system will keep this record for the next hour and may remove it if you send a reportIncorrectRecaptcha request for a task, which is executed by the same whitelisted worker.
	// In short: reportCorrectRecaptcha adds workers to your whitelist, reportIncorrectRecaptcha removes them from it.
	// Reporting correctly solved tokens works on noticeable volumes, starting from 10 tasks per minute.
	// This is due to the heavy migration of our workers between different task queues and their current busy/idle status. With volumes like 1 task per minute, you won't notice any improvements in quality.
	// Reports must be sent within 60 seconds of task completion.
	// If you send a report later, the API will return ERROR_NO_SUCH_CAPCHA_ID error.
	// Only one report can be sent per task.
	ReportCorrectRecaptcha ReportTaskType = "reportCorrectRecaptcha"

	// Send complaint on a Hcaptcha token
	// Use this method to send us information about tokens which did not pass on target service.
	// Refunds are not guaranteed, however sending reports helps system to filter out workers which were banned in Hcaptcha network.
	// Complaints are accepted for Hcaptchas only, including Enterprise Hcaptcha.
	ReportIncorrectHcaptcha ReportTaskType = "reportIncorrectHcaptcha"
)

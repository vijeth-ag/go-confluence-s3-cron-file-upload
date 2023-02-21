package main

// 1
import (
	"clonfluence/report"
	"fmt"

	"gopkg.in/robfig/cron.v2"
)

func runCronJobs() {
	s := cron.New()

	s.AddFunc("@every 5s", func() {
		report.UploadReportToConfluence()
		report.UploadToS3()
	})
	s.Start()
}

func main() {
	runCronJobs()
	fmt.Scanln()
}

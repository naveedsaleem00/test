package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {

	abc()

	/*Create a new cron scheduler*/
	go CronJob()

}

func CronJob() {
	/*Create a new cron scheduler*/
	c := cron.New()

	/*Add a cron job (in this case, it prints a message every minute)*/
	_, err := c.AddFunc("@every 5s", func() {
		fmt.Println(time.Now())
		abc()
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	/*Start the cron scheduler*/
	c.Start()

	fmt.Println("Cron job removed. Exiting...")

}

func abc() {
	fmt.Println("Cron Job done ... ")
}

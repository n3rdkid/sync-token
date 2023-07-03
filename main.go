package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

func refreshToken(id int) {
	uuid := uuid.NewString()
	message := fmt.Sprintf("%v : %v", id, uuid)
	fmt.Println(message)
}

func initCron() {
	runner := cron.New()
	runner.AddFunc("@every 3s", func() {
		refreshToken(5)
	})
	runner.Start()
}

func main() {
	initCron()
	http.ListenAndServe(":8000", nil)
}

package main

import (
	"context"
	"fmt"
	"time"
	//	"github.com/zhashkevych/scheduler"
)

func main() {
	ctx := context.Background()

	//t := time.Now()
	//fmt.Println(t)

	/*	worker := scheduler.NewScheduler()
		worker.Add(ctx, parseSubcriptionData, time.Second*5)
		worker.Add(ctx, sendStatistics, time.Second*10)

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Interrupt)

		<-quit
		worker.Stop()*/
}

func parseSubcriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("subcription parsed succesfuly at &s\n", time.Now().String())
}

func sendStatistics(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("statistics send at %s\n", time.Now().String())
}

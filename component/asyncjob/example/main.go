package main

import (
	"context"
	"log"
	"rest-api/component/asyncjob"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job1")

		//return errors.New("something went wrong with job 1")
		return nil
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 1)
		log.Println("I am job 2")

		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 1)
		log.Println("I am job 3")

		return nil
	})

	group := asyncjob.NewGroup(false, job1, job2, job3)
	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}

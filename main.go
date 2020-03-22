package main

import (
	"fmt"
	"github.com/CHYinYiHang/gocron/bootstartp"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	bootstartp.InitApplication()
	c := cron.New(cron.WithSeconds())
	c.Start()

	fmt.Println("11111")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	id, err := c.AddFunc("*/2 * * * * *", func() {
		//fmt.Println("111")
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})

	fmt.Println(id, "\n", err)

	fmt.Println("222")

	fmt.Println("2222")


	select {
	}
}

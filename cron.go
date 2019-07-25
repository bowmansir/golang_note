package main

import (
	"blog/models"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting ...")

	// 会根据本地时间创建一个新（空白）的 Cron job runner
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}

package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

type DfCronJob struct {
	JobID string
}

func (cj *DfCronJob) Run() {

}

func TestXxx(t *testing.T) {
	c := cron.New(
		cron.WithParser(
			cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor),
		),
	)

	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	parser.Parse("spec string") // 检查cron表达式正确性

	// 启动定时任务
	c.Start()

	// eID, err := c.AddFunc("*/5 * * * * ?", func() { fmt.Println("每5秒执行定时任务") })
	// if err != nil {
	// 	fmt.Println("添加定时任务失败：", err)
	// 	return
	// }

	eID2, _ := c.AddFunc("@every 1s", func() { fmt.Println("每秒执行定时任务") })
	c.AddFunc("@every 1s", func() {
		fmt.Println("每秒执行定时任务,但睡眠")
		time.Sleep(time.Second * 3)
		fmt.Println("睡眠")
	})

	// e := c.Entry(eID)
	// dfj := e.Job.(*DfCronJob)
	// jid := dfj.JobID
	// fmt.Println(jid)

	c.Entries()

	fmt.Println(1)
	time.Sleep(5 * time.Second)
	fmt.Println(2)

	c.Remove(eID2)

	time.Sleep(6 * time.Second)
	c.Stop()
}

func TestCron2(t *testing.T) {
	c := cron.New()
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	c.Stop()
}

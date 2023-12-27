// package job

// import (
// 	"fmt"
// 	"strconv"
// 	"sync"
// 	"testing"
// 	"time"
// )

// /*
// Job 和 task 原型设计
// */


// func Test_dJob1(t *testing.T) {
// 	// job req
// 	jobId := "hello1"
// 	jtype := "hello"

// 	// new job
// 	var subJob dIJob

// 	if jtype == "hello" {
// 		subJob = &dHelloJob{
// 			name: "hzw",
// 		}
// 	}

// 	// 创建job
// 	job1 := &dJob{
// 		JobID:   jobId,
// 		JobType: jtype,
// 		dIJob:   subJob,
// 	}

// 	// job 生成作业
// 	job1.GenTasks()

// 	for i := 0; i < 5; i++ {
// 		// job 启动
// 		job1.Start()
// 		time.Sleep(time.Second * 2)
// 	}

// }

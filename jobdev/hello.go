package jobdev

import (
	"fmt"
	"time"
)

// 一次性的
// 定时job

// 大数据量的(全量或增量) - poc模型
// 准实时的 - jdbc-connector-source的实现逻辑

type dHelloJob struct {
	name string
}

type dHelloTask struct {
	name string
	time time.Time
}

// == hello Job ==
func (j *dHelloJob) Type() string {
	return "hello"

}
func (j *dHelloJob) ResetTask(task dITask) error {
	// 判断传入的task类型
	t, ok := task.(*dHelloTask)
	if !ok {
		return fmt.Errorf("helloJob rest task is not hellotask %T", task)
	}
	// 初始化作业需要使用的参数
	t.time = time.Now()
	return nil
}

func (j *dHelloJob) Retryable() bool {
	return true
}

func (j *dHelloJob) SplitTasks() []dITask {
	return []dITask{&dHelloTask{name: j.name + "_task"}}
}

func (d *dHelloTask) Execute(dt *dTask) error {
	fmt.Printf("[%s] hello task, time:%v\n", dt.TaskID, d.time)
	return nil
}

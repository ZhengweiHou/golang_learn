package jobdev

import (
	"strconv"
	"sync"
	"time"
)

type dJob struct {
	dIJob
	JobID                  string
	JobType                string
	ScheduleType           string
	ScheduleConf           string
	ExecutorFailRetryCount int
	ExecutorBlockStrategy  string
	Status                 string
	tasks                  []*dTask
	retryTime              int
	mu                     sync.Mutex
}

type dIJob interface {
	// 返回当前Job实现的类型名
	Type() string
	// 当前job生成作业实例，如分片作业
	SplitTasks() []dITask
	// 是否可重试
	Retryable() bool
	// 重置task，如增量作业的增量条件等
	ResetTask(task dITask) error
}

func (djob *dJob) GenTasks() {
	itasks := djob.SplitTasks()

	djob.tasks = make([]*dTask, 0)
	for _, itask := range itasks {
		dT := &dTask{
			TaskID: strconv.FormatInt(time.Now().UnixNano(), 10), // TODO 生成taskID，需全局唯一
			job:    djob,                                         // TODO 是否循环引用
			JobID:  djob.JobID,
			Status: TaskStatCreating,
			dITask: itask,
		}
		djob.tasks = append(djob.tasks, dT)
		//  TODO 存 tasks 进 task containner
	}
}

// Start 启动job
func (djob *dJob) Start() {
	djob.Status = JobStatReady

	for _, t := range djob.tasks {

		djob.ResetTask(t.dITask)
		go t.Run() // TODO 运行前如果要向task传入某些参数怎么办，如准实时同步的增量条件
	}

	djob.Status = JobStatRuning

	// TODO job除了task回调外，应该还有一个兜底的检查，由jobmanager轮询筛选检查？
}

func (djob *dJob) Success() {

}

func (djob *dJob) Faild() {
	// TODO retry次数控制
	needretry := djob.Retryable()
	if needretry {
		djob.Start() // TODO 其他方式
	}
}

// 作业状态更新，触发job的状态收集
func (djob *dJob) TaskCall(task *dTask) {
	djob.mu.Lock()
	defer djob.mu.Unlock()

	if djob.Status == JobStatFailed || djob.Status == JobStatSucceed {
		return
	}

	// job 层面的锁控制
	// TODO 并发处理，待进一步设计
	// 1. 判断task状态
	tStat := task.Status
	jStat := djob.Status
	if tStat == TaskStatSucceed {
		// 检查其他 task 状态
		i := 0
		for _, t := range djob.tasks {
			if t.Status == TaskStatFailed {
				jStat = JobStatFailed
				break
			}
			if t.Status == JobStatSucceed {
				i++
			}
		}
		if i == len(djob.tasks) {
			jStat = JobStatSucceed
		}
	} else {
		jStat = JobStatFailed
	}

	switch jStat {
	case JobStatSucceed:
		djob.Success()
	case JobStatFailed:
		for _, t := range djob.tasks {
			t.Stop()
		}
		djob.Faild()
	}
}

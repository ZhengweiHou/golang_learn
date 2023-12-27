package jobdev

type dTask struct {
	TaskID string
	dITask
	job    *dJob // 循环引用 golang GC采用 标记-清除法 能够处理
	JobID  string
	Status string
}

type dITask interface {
	Execute(*dTask) error
	Destroy() error
}

func (dt *dTask) Run() {
	dt.Status = TaskStatRuning
	// TODO task的状态怎么同步,
	err := dt.dITask.Execute(dt)
	if err != nil {
		dt.Status = TaskStatFailed
	} else {
		dt.Status = TaskStatSucceed
	}
	dt.End()
}

func (dt *dTask) Stop() {
	// task stop
}

func (dt *dTask) End() {
	// 通知job,子task完成
	// 获取task所属job
	job := dt.job
	job.TaskCall(dt) // task结束时回调job
}

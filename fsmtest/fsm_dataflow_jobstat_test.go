package fsmtest

import (
	"context"
	"testing"

	"github.com/looplab/fsm"
	"github.com/sirupsen/logrus"
)

const (
	JobStatCreating  = "creating"
	JobStatInited    = "inited"
	JobStatScheduled = "scheduled"
	JobStatReady     = "ready"
	JobStatRuning    = "running"
	JobStatFailed    = "failed"
	JobStatSucceed   = "succeed"
)

type Job struct {
	ID  string
	xxx string
	FSM *fsm.FSM
	ctx context.Context
}

func NewJob(id string) *Job {
	job := &Job{
		ID:  id,
		ctx: context.Background(), // 初始化job ctx
	}

	job.FSM = fsm.NewFSM(
		JobStatCreating,
		fsm.Events{
			{Name: JobStatInited, Src: []string{JobStatCreating}, Dst: JobStatInited},
			{Name: JobStatScheduled, Src: []string{JobStatInited, JobStatSucceed}, Dst: JobStatScheduled},
			{Name: JobStatReady, Src: []string{JobStatInited, JobStatSucceed, JobStatFailed}, Dst: JobStatReady},
			{Name: JobStatRuning, Src: []string{JobStatReady}, Dst: JobStatRuning},
			{Name: JobStatSucceed, Src: []string{JobStatRuning}, Dst: JobStatSucceed},
			{Name: JobStatFailed, Src: []string{JobStatRuning}, Dst: JobStatFailed},
		},
		fsm.Callbacks{
			"before_event": func(ctx context.Context, e *fsm.Event) {
				logrus.Infof("before_event job[%s]当前状态[%s],状态转变:[%s] -> [%s]", job.ID, e.FSM.Current(), e.Src, e.Dst)
			},
			"enter_state": func(ctx context.Context, e *fsm.Event) {
				logrus.Infof("enter_state  job[%s]当前状态[%s],状态转变:[%s] -> [%s]", job.ID, e.FSM.Current(), e.Src, e.Dst)
			},
			"leave_state": func(ctx context.Context, e *fsm.Event) {
				logrus.Infof("leave_state  job[%s]当前状态[%s],状态转变:[%s] -> [%s]", job.ID, e.FSM.Current(), e.Src, e.Dst)
			},
			"after_event": func(ctx context.Context, e *fsm.Event) {
				logrus.Infof("after_event  job[%s]当前状态[%s],状态转变:[%s] -> [%s]", job.ID, e.FSM.Current(), e.Src, e.Dst)
			},
			JobStatInited: func(ctx context.Context, e *fsm.Event) {
				logrus.Infof("inited  job[%s]当前状态[%s],状态转变:[%s] -> [%s]", job.ID, e.FSM.Current(), e.Src, e.Dst)
			},
			"before_" + JobStatSucceed: func(ctx context.Context, e *fsm.Event) {
				// 检查job内作业的状态是否终态
			},
		},
	)

	return job
}

func TestFsmJob1(t *testing.T) {
	job := NewJob("hhh")
	err := job.FSM.Event(context.Background(), JobStatInited)
	te(err)

	err = job.FSM.Event(context.Background(), JobStatReady)
	te(err)

	err = job.FSM.Event(context.Background(), JobStatRuning)
	te(err)

	if job.FSM.Can(JobStatSucceed) {
		logrus.Infof("job 当前状态%s,可以转换为状态%s", job.FSM.Current(), JobStatSucceed)
	}

	err = job.FSM.Event(context.Background(), JobStatRuning)
	te(err)

}

func te(err error) {
	if err != nil {
		logrus.Error(err.Error())
	}
}

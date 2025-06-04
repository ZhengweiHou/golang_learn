package hzwrefl

import (
	"fmt"
	"reflect"
	"testing"
)

type IJob interface {
	DoJob()
}

type Job1 struct{}

func (j Job1) DoJob() {
	fmt.Println("Job1 is doing the job")
}

type Job2 struct{}

func (j Job2) DoJob() {
	fmt.Println("Job2 is doing the job")
}

func instantiateJobs() []IJob {
	jobTypes := []reflect.Type{
		reflect.TypeOf(Job1{}),
		reflect.TypeOf(Job2{}),
		// Add more types as needed
	}

	jobs := make([]IJob, len(jobTypes))

	for i, jobType := range jobTypes {
		job := reflect.New(jobType).Interface()
		jobs[i] = job.(IJob)
	}

	return jobs
}

func TestRef2(t *testing.T) {

	jobs := instantiateJobs()

	for _, job := range jobs {
		job.DoJob()
	}
}

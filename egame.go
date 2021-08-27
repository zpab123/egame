package egame

import (
	"time"

	"github.com/zpab123/egame/job"
	"github.com/zpab123/egame/worker"
)

var (
	eapp *Application = NewApplication() // 默认app
)

// /////////////////////////////////////////////////////////////////////////////
// public

// 添加工人
func AddWorker(num int) {
	for i := 0; i < num; i++ {
		worker.NewWorker()
	}
}

// 添加工作
func AddJob(j job.IJob) {
	worker.AddJob(j)
}

// 启动
func Run() {
	worker.Run()

	runing := true
	for runing {
		runing = worker.Working()
		time.Sleep(1 * time.Second)
	}
}

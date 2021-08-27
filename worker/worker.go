// /////////////////////////////////////////////////////////////////////////////
// 工人

package worker

import (
	"sync"
	"time"

	"github.com/zpab123/egame/job"
	"github.com/zpab123/egame/syncs"
)

var (
	workers map[uint32]*Worker = make(map[uint32]*Worker, 0)
	ids     uint32             = 0 // id
	last    uint32             = 0 // 用于 job 分配
)

// /////////////////////////////////////////////////////////////////////////////
// public

// 添加工作
func AddJob(j job.IJob) {
	// 如何均分分布 job ?
	if len(workers) == 0 || j == nil {
		return
	}

	now := last % ids
	w, ok := workers[now]
	if ok {
		w.addJob(j)
		last++
	}
}

// worker 是否正在工作
func Working() bool {
	for _, w := range workers {
		if w.working.Load() {
			return true
		}
	}

	return false
}

// 启动 worker
func Run() {
	for _, w := range workers {
		w.Run()
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Worker

// 工人
type Worker struct {
	id       uint32           // 工人id
	jobs     []job.IJob       // 工作列表
	tmpjob   []job.IJob       // 缓存
	jobMutex sync.Mutex       // jobs 数据锁
	working  syncs.AtomicBool // 是否正在工作
}

// 创建一个工人
func NewWorker() *Worker {
	w := Worker{
		id:     ids,
		jobs:   make([]job.IJob, 0),
		tmpjob: make([]job.IJob, 0),
	}
	workers[ids] = &w
	ids++

	return &w
}

// -----------------------------------------------------------------------------
// public

// 开始工作
func (w *Worker) Run() {
	w.working.Store(true)

	go w.start()
}

// 是否正在工作
func (w *Worker) Working() bool {
	return w.working.Load()
}

// -----------------------------------------------------------------------------
// private

// 开始工作
func (w *Worker) start() {
	for w.working.Load() {
		w.update()
		time.Sleep(1 * time.Second)
	}
}

// 帧函数
func (w *Worker) update() {
	w.jobMutex.Lock()
	ln := len(w.jobs)
	if ln == 0 {
		w.jobMutex.Unlock()
		return
	}

	w.tmpjob = make([]job.IJob, ln)
	copy(w.tmpjob, w.jobs)
	w.jobMutex.Unlock()

	for _, j := range w.tmpjob {
		j.Update()
	}
}

// 添加工作
func (w *Worker) addJob(j job.IJob) {
	if j != nil {
		w.jobMutex.Lock()
		w.jobs = append(w.jobs, j)
		w.jobMutex.Unlock()
	}
}

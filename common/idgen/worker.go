package idgen

import (
	"log"
	"sync"
	"time"
)

//转自链接：https://learnku.com/articles/26821

const (
	workerIdBits     int64 = 5
	datacenterIdBits int64 = 5
	sequenceBits     int64 = 12

	maxWorkerId     int64 = -1 ^ (-1 << uint64(workerIdBits))
	maxDatacenterId int64 = -1 ^ (-1 << uint64(datacenterIdBits))
	maxSequence     int64 = -1 ^ (-1 << uint64(sequenceBits))

	timeLeft uint8 = 22
	dataLeft uint8 = 17
	workLeft uint8 = 12

	twepoch int64 = 1525705533000
)

type worker struct {
	mu           sync.Mutex
	lastStamp    int64
	workerId     int64
	dataCenterId int64
	sequence     int64
}

func NewWorker(workerId int64, dataCenterId int64, sequence int64) *worker {
	return &worker{lastStamp: -1, workerId: workerId, dataCenterId: dataCenterId, sequence: sequence}
}

func (w *worker) getCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}

//var i int = 1
func (w *worker) nextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	timestamp := w.getCurrentTime()
	if timestamp < w.lastStamp {
		log.Fatal("can not generate id")
	}
	if w.lastStamp == timestamp {
		// 这其实和 <==>
		// w.sequence++
		// if w.sequence++ > maxSequence  等价
		w.sequence = (w.sequence + 1) & maxSequence
		if w.sequence == 0 {
			// 之前使用 if, 只是没想到 GO 可以在一毫秒以内能生成到最大的 Sequence, 那样就会导致很多重复的
			// 这个地方使用 for 来等待下一毫秒
			for timestamp <= w.lastStamp {
				//i++
				//fmt.Println(i)
				timestamp = w.getCurrentTime()
			}
		}
	} else {
		w.sequence = 0
	}
	w.lastStamp = timestamp

	return ((timestamp - twepoch) << timeLeft) | (w.dataCenterId << dataLeft) | (w.workerId << workLeft) | w.sequence
}
func (w *worker) tilNextMillis() int64 {
	timestamp := w.getCurrentTime()
	if timestamp <= w.lastStamp {
		timestamp = w.getCurrentTime()
	}
	return timestamp
}

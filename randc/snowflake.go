package randc

// 雪花算法 生成唯一的值
import (
	"fmt"
	"github.com/Cc360428/HelpPackage/other"
	"sync"
	"time"
)

type IdWorker struct {
	startTime             int64
	workerIdBits          uint
	datacenterIdBits      uint
	maxWorkerId           int64
	maxDatacenterId       int64
	sequenceBits          uint
	workerIdLeftShift     uint
	datacenterIdLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerId              int64
	datacenterId          int64
	sequence              int64
	lastTimestamp         int64
	signMask              int64
	idLock                *sync.Mutex
}

func (w *IdWorker) InitIdWorker(workerId, datacenterId int64) error {
	var baseValue int64 = -1
	w.startTime = 1463834116272
	w.workerIdBits = 5
	w.datacenterIdBits = 5
	w.maxWorkerId = baseValue ^ (baseValue << w.workerIdBits)
	w.maxDatacenterId = baseValue ^ (baseValue << w.datacenterIdBits)
	w.sequenceBits = 12
	w.workerIdLeftShift = w.sequenceBits
	w.datacenterIdLeftShift = w.workerIdBits + w.workerIdLeftShift
	w.timestampLeftShift = w.datacenterIdBits + w.datacenterIdLeftShift
	w.sequenceMask = baseValue ^ (baseValue << w.sequenceBits)
	w.sequence = 0
	w.lastTimestamp = -1
	w.signMask = ^baseValue + 1

	w.idLock = &sync.Mutex{}

	if w.workerId < 0 || w.workerId > w.maxWorkerId {
		return fmt.Errorf("workerId[%v] is less than 0 or greater than maxWorkerId[%v]", workerId, datacenterId)
	}
	if w.datacenterId < 0 || w.datacenterId > w.maxDatacenterId {
		return fmt.Errorf("datacenterId[%d] is less than 0 or greater than maxDatacenterId[%d]", workerId, datacenterId)
	}
	w.workerId = workerId
	w.datacenterId = datacenterId
	return nil
}

func (w *IdWorker) NextId() (int64, error) {
	w.idLock.Lock()
	timestamp := time.Now().UnixNano()
	if timestamp < w.lastTimestamp {
		return -1, fmt.Errorf("clock moved backwards.  Refusing to generate id for %v milliseconds", w.lastTimestamp-timestamp)
	}

	if timestamp == w.lastTimestamp {
		w.sequence = (w.sequence + 1) & w.sequenceMask
		if w.sequence == 0 {
			timestamp = w.tilNextMillis()
			w.sequence = 0
		}
	} else {
		w.sequence = 0
	}

	w.lastTimestamp = timestamp

	w.idLock.Unlock()

	id := ((timestamp - w.startTime) << w.timestampLeftShift) |
		(w.datacenterId << w.datacenterIdLeftShift) |
		(w.workerId << w.workerIdLeftShift) |
		w.sequence

	if id < 0 {
		id = -id
	}

	return id, nil
}

func (w *IdWorker) tilNextMillis() int64 {
	timestamp := time.Now().UnixNano()
	if timestamp <= w.lastTimestamp {
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return timestamp
}

func GetStringOrder() (order string) {
	currWoker := IdWorker{}
	_ = currWoker.InitIdWorker(1000, 1)
	t, _ := currWoker.NextId()
	order = other.Int64TurnString(t)
	return order
}

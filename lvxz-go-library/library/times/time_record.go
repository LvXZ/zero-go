package times

import (
	"time"
)

type TimeRecord struct {
	start time.Time
}

func initTimeRecord() *TimeRecord {
	var timeRecord = new(TimeRecord)
	timeRecord.start = time.Now()
	return timeRecord
}

func (tr *TimeRecord) reset() {
	tr.start = time.Now()
}

func (tr *TimeRecord) getMilliSecond() int64 {
	timeUsed := time.Since(tr.start)
	return timeUsed.Milliseconds()
}

func (tr *TimeRecord) getSecond() float64 {
	timeUsed := time.Since(tr.start)
	return timeUsed.Seconds()
}

func (tr *TimeRecord) getMicroSecond() int64 {
	timeUsed := time.Since(tr.start)
	return timeUsed.Microseconds()
}

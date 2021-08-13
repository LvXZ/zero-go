package times

import (
	"testing"
	"time"
)

func TestTimeRecord_Reset(t *testing.T) {
	timeRecord := initTimeRecord()
	if timeRecord == nil {
		t.Fatal("初始化失败")
	}
	startTime := timeRecord.start
	timeRecord.reset()
	if startTime == timeRecord.start {
		t.Fatal("reset失败")
	}
}

/**
 * 测试用例: 获取毫秒
 */
func TestTimeRecord_GetMilliSecond(t *testing.T) {
	timeRecord := initTimeRecord()
	if timeRecord == nil {
		t.Fatal("初始化失败")
	}
	time.Sleep(time.Millisecond * 100)
	usedTime := timeRecord.getMilliSecond()
	if usedTime < 100 {
		t.Fatal("获取毫秒时间失败")
	}
}

/**
 * 测试用例: 获取秒
 */
func TestTimeRecord_GetSecond(t *testing.T) {
	timeRecord := initTimeRecord()
	if timeRecord == nil {
		t.Fatal("初始化失败")
	}
	time.Sleep(time.Second * 1)
	usedTime := timeRecord.getSecond()
	if usedTime < 1 {
		t.Fatal("获取秒时间失败")
	}
}

/**
 * 测试用例: 获取纳秒
 */
func TestTimeRecord_GetMicroSecond(t *testing.T) {
	timeRecord := initTimeRecord()
	if timeRecord == nil {
		t.Fatal("初始化失败")
	}
	time.Sleep(time.Microsecond * 20)
	usedTime := timeRecord.getMicroSecond()
	if usedTime < 20 {
		t.Fatal("获取纳秒时间失败")
	}
}

func TestTimeRecord_AA(t *testing.T) {
	var a map[string]string
	for k, v := range a {
		print(k)
		print(v)
	}
}

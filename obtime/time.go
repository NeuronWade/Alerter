package obtime

import "time"

const (
	DAY_TIME  = 86400
	WEEK_TIME = 86400 * 7
)

type Offset struct {
	Time   time.Time
	Offset time.Duration
}

var timerecords = make(map[string]int64)
var offset = Offset{
	Time: time.Now(),
}

func Now() time.Time {
	return time.Now().Add(offset.Offset)
}

func SetUnixSec(value int64) time.Time {
	return time.Unix(value, 0)
}

// offset : 当天的偏移量, 按天计算
func GetDay(offset int) time.Time {
	t := Now().AddDate(0, 0, offset)
	year, month, day := t.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return today
}

func SetLocal(localTime time.Time) {
	offset.Time = localTime
	offset.Offset = localTime.Sub(time.Now())
}

func ResetLocal() {
	offset.Time = time.Now()
	offset.Offset = 0
}

func SetUnixTimeSign(key string, value int64) {
	timerecords[key] = value
}

func GetUnixTimeSign(key string) int64 {
	return timerecords[key]
}

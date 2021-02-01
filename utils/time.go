package utils

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

var timeBeijin = time.FixedZone("CST", 3600*8)

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func FormatTime(t time.Time) string {
	t = t.In(timeBeijin)
	return t.Format(taobao.DATE_TIME_FORMAT)
}

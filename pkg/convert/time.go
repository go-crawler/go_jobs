package convert

import (
	"time"

	"github.com/jinzhu/now"
)

func MustDateToUnix(date string) int64 {
	if date == `` || date == `0000-00-00` {
		return 0
	}

	loc, _ := time.LoadLocation("Asia/Chongqing")
	t, err := now.ParseInLocation(loc, date)
	if err != nil {
		return 0
	}

	return t.Unix()
}

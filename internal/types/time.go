package types

import (
	"fmt"
	"time"
)

var (
	formatTime = "2006-01-02 15:04:05"
)

// Time 自定义时间类型
type Time struct {
	time.Time
}

// MarshalJSON实现json.Marshaler 接口
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
	str := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(str), nil
}

// UnmarshalJSON实现json.Unmarshaler接口
func (t *Time) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == `""` {
		t.Time = time.Time{}
		return nil
	}
	parsed, err := time.ParseInLocation(`"`+formatTime+`"`, str, time.Local)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

// GormDataType确保GORM使用datetime类型
func (Time) GormDataType() string {
	return "datetime"
}

package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

const DateLayout = "2006-01-02"
const DateTimeLayout = "2006-01-02 15:04:05"

var TimeZone, _ = time.LoadLocation("Asia/Shanghai")

type Date time.Time

func (d *Date) MarshalJSON() ([]byte, error) {
	time := time.Time(*d)
	return []byte(fmt.Sprintf("\"%s\"", time.Format(DateLayout))), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(DateLayout, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	_d := Date(t)
	*d = _d
	return nil
}

func (d *Date) String() string {
	if d == nil {
		return ""
	}
	return time.Time(*d).Format(DateLayout)
}

func (t Date) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(DateLayout), nil
}

func (t *Date) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = Date(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

type DateTime time.Time

func (d *DateTime) MarshalJSON() ([]byte, error) {
	time := time.Time(*d)
	return []byte(fmt.Sprintf("\"%s\"", time.Format(DateTimeLayout))), nil
}

func (d *DateTime) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(DateTimeLayout, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	_d := DateTime(t)
	*d = _d
	return nil
}

func (d *DateTime) String() string {
	if d == nil {
		return ""
	}
	return time.Time(*d).Format(DateTimeLayout)
}

func (t DateTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(DateTimeLayout), nil
}

func (t *DateTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = DateTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

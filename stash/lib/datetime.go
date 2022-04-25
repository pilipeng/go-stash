//Go-json解析时间格式

package lib

import "time"

/*import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	// Ignore null, like in the main JSON package.
	if string(b) == "null" || string(b) == `""` {
		return nil
	}

	var err error
	// 指定时区
	d.Time, err = time.ParseInLocation(`"2006-01-02"`, string(b),time.Local)
	if err != nil {
		return err
	}
	return nil
}

type Datetime struct {
	time.Time
}

func (d *Datetime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", d.Format(`"2006-01-02 15:04:05"`))), nil
}

func (d *Datetime) UnmarshalJSON(b []byte) error {
	var err error
	d.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	return nil
}*/
type Datetime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Datetime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Datetime(now)
	return
}

func (t Datetime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Datetime) String() string {
	return time.Time(t).Format(timeFormart)
}

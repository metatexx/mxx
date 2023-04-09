package dbtime

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "15:04:05"

type DBTime struct {
	time.Time
}

func NewTime(t time.Time) DBTime {
	t = t.UTC()
	return DBTime{
		Time: time.Date(0, 1, 1, t.Hour(), t.Minute(), t.Second(), 0, time.UTC),
	}
}

func Now() DBTime {
	return NewTime(time.Now())
}

func (dt *DBTime) AddDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), dt.Hour()+t.Hour(), dt.Minute()+t.Minute(),
		dt.Second()+t.Second(), t.Nanosecond(), t.Location())
}

var _ driver.Valuer = (*DBTime)(nil)

// Value returns the time as string using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (dt *DBTime) Value() (driver.Value, error) {
	return dt.UTC().Format(timeFormat), nil
}

var _ sql.Scanner = (*DBTime)(nil)

// Scan scans the time parsing it if necessary using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (dt *DBTime) Scan(src interface{}) (err error) {
	switch src := src.(type) {
	case time.Time:
		*dt = NewTime(src)
		return nil
	case string:
		dt.Time, err = time.ParseInLocation(timeFormat, src, time.UTC)
		return err
	case []byte:
		dt.Time, err = time.ParseInLocation(timeFormat, string(src), time.UTC)
		return err
	case nil:
		dt.Time = time.Time{}
		return nil
	default:
		return fmt.Errorf("unsupported data type: %T", src)
	}
}

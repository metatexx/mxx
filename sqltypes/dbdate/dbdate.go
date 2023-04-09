package dbdate

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02"
const dbDateZeroValue = "0000-00-00"

type DBDate struct {
	time.Time
}

func NewTime(t time.Time) DBDate {
	t = t.UTC()
	return DBDate{
		Time: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC),
	}
}

var _ driver.Valuer = (*DBDate)(nil)

// Value returns the time as string using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (tm DBDate) Value() (driver.Value, error) {
	return tm.UTC().Format(timeFormat), nil
}

var _ sql.Scanner = (*DBDate)(nil)

// Scan scans the time parsing it if necessary using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (tm *DBDate) Scan(src interface{}) (err error) {
	switch src := src.(type) {
	case time.Time:
		*tm = NewTime(src)
		return nil
	case string:
		if src == dbDateZeroValue {
			tm.Time = time.Time{}.UTC()
			return nil
		}
		tm.Time, err = time.ParseInLocation(timeFormat, src, time.UTC)
		return err
	case []byte:
		if string(src) == dbDateZeroValue {
			tm.Time = time.Time{}.UTC()
			return nil
		}
		tm.Time, err = time.ParseInLocation(timeFormat, string(src), time.UTC)
		return err
	case nil:
		tm.Time = time.Time{}.UTC()
		return nil
	default:
		return fmt.Errorf("unsupported data type: %T", src)
	}
}

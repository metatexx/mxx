package dbts

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"
const dbDateZeroValue = "0000-00-00 00:00:00"

type DBTS struct {
	time.Time
}

var _ driver.Valuer = (*DBTS)(nil)

// Value returns the time as string using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (tm DBTS) Value() (driver.Value, error) {
	return tm.UTC().Format(timeFormat), nil
}

var _ sql.Scanner = (*DBTS)(nil)

// Scan scans the time parsing it if necessary using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (tm *DBTS) Scan(src interface{}) (err error) {
	switch src := src.(type) {
	case time.Time:
		*tm = DBTS{src}
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

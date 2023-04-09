package ynenum

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

// We need a special date format to support the date '0000-00-00' as we use it in welo

const EnumYes = "Y"
const EnumNo = "N"

type YNEnum bool

var _ driver.Valuer = (*YNEnum)(nil)

//goland:noinspection GoMixedReceiverTypes
func (dbd YNEnum) String() string {
	if dbd {
		return EnumYes
	}
	return EnumNo
}

// Value returns the time as string.
//
//goland:noinspection GoMixedReceiverTypes
func (dbd YNEnum) Value() (driver.Value, error) {
	//log.Print("valuing")
	return dbd.String(), nil
}

//goland:noinspection GoMixedReceiverTypes
func (dbd YNEnum) Yes() bool {
	return bool(dbd)
}

var _ sql.Scanner = (*YNEnum)(nil)

// Scan scans the time parsing it if necessary using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (dbd *YNEnum) Scan(src interface{}) (err error) {
	//log.Print("scanning")
	*dbd = false
	switch src := src.(type) {
	case string:
		//log.Printf("scanning: string %q", src)
		*dbd = src == EnumYes
		return nil
	case []byte:
		//log.Printf("scanning: []byte %q", src)
		*dbd = string(src) == EnumYes
		return nil
	case nil:
		//log.Print("scanning: nil")
		return nil
	default:
		return fmt.Errorf("unsupported data type: %T", src)
	}
}

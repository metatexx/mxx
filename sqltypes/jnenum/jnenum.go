package jnenum

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

// We need a special date format to support the date '0000-00-00' as we use it in welo

const EnumJa = "J"
const EnumNein = "N"

type JNEnum bool

var _ driver.Valuer = (*JNEnum)(nil)

//goland:noinspection GoMixedReceiverTypes
func (dbd JNEnum) String() string {
	if dbd {
		return EnumJa
	}
	return EnumNein
}

// Value returns the time as string.
//
//goland:noinspection GoMixedReceiverTypes
func (dbd JNEnum) Value() (driver.Value, error) {
	//log.Print("valuing")
	return dbd.String(), nil
}

//goland:noinspection GoMixedReceiverTypes
func (dbd JNEnum) Ja() bool {
	return bool(dbd)
}

var _ sql.Scanner = (*JNEnum)(nil)

// Scan scans the time parsing it if necessary using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (dbd *JNEnum) Scan(src interface{}) (err error) {
	//log.Print("scanning")
	*dbd = false
	switch src := src.(type) {
	case string:
		//log.Printf("scanning: string %q", src)
		*dbd = src == EnumJa
		return nil
	case []byte:
		//log.Printf("scanning: []byte %q", src)
		*dbd = string(src) == EnumJa
		return nil
	case nil:
		//log.Print("scanning: nil")
		return nil
	default:
		return fmt.Errorf("unsupported data type: %T", src)
	}
}

package mfenum

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

// We need a special date format to support the date '0000-00-00' as we use it in welo

const enumMale = "M"
const enumFemale = "F"

type MFEnum bool

var _ driver.Valuer = (*MFEnum)(nil)

//goland:noinspection GoMixedReceiverTypes
func (dbd MFEnum) String() string {
	if dbd {
		return enumMale
	}
	return enumFemale
}

// Value returns the time as string.
//
//goland:noinspection GoMixedReceiverTypes
func (dbd MFEnum) Value() (driver.Value, error) {
	//log.Print("valuing")
	return dbd.String(), nil
}

//goland:noinspection GoMixedReceiverTypes
func (dbd MFEnum) Male() bool {
	return bool(dbd)
}

var _ sql.Scanner = (*MFEnum)(nil)

// Scan scans the time parsing it if necessary using timeFormat.
//
//goland:noinspection GoMixedReceiverTypes
func (dbd *MFEnum) Scan(src interface{}) (err error) {
	//log.Print("scanning")
	*dbd = false
	switch src := src.(type) {
	case string:
		//log.Printf("scanning: string %q", src)
		*dbd = src == enumMale
		return nil
	case []byte:
		//log.Printf("scanning: []byte %q", src)
		*dbd = string(src) == enumMale
		return nil
	case nil:
		//log.Print("scanning: nil")
		return nil
	default:
		return fmt.Errorf("unsupported data type: %T", src)
	}
}

package bunx

import (
	"database/sql"
	"errors"

	"github.com/uptrace/bun"
)

func NoRows(err error) bool {
	if errors.Is(err, sql.ErrNoRows) {
		return true
	}
	if err != nil {
		panic(err)
	}
	return false
}

func AnyRows(err error) {
	if err == nil || errors.Is(err, sql.ErrNoRows) {
		return
	}
	panic(err)
}

func ToString(db *bun.DB, q bun.Query) string {
	buf, err := q.AppendQuery(db.Formatter(), nil)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

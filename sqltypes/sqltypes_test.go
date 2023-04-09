package main

import (
	"context"
	"database/sql"
	"testing"

	"git.jetbrains.space/metatexx/mxc/jetindex/pkg/sqltypes/dbdate"
	"git.jetbrains.space/metatexx/mxc/jetindex/pkg/sqltypes/mfenum"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Personal struct {
	bun.BaseModel `bun:"table:personal,alias:pe"`

	ID           int           `bun:"pe_id,pk,autoincrement"`
	Name         string        `bun:"pe_name,notnull"`
	Vorname      string        `bun:"pe_vorname,notnull"`
	Gender       mfenum.MFEnum `bun:"pe_geschlecht,notnull"` // ENUM("M","W")
	GeburtsDatum dbdate.DBDate `bun:"pe_geburtstag,notnull"`
	GeburtsOrt   string        `bun:"pe_geburtsort,notnull"`
	GeburtsName  string        `bun:"pe_geburtsname,notnull"`
	Strasse      string        `bun:"pe_strasse,notnull"`
	PLZ          string        `bun:"pe_plz,notnull"`
	Ort          string        `bun:"pe_ort,notnull"`
	Eintritt     dbdate.DBDate `bun:"pe_eintritt,notnull"`
	Austritt     dbdate.DBDate `bun:"pe_austritt,notnull"`
	Befristet    dbdate.DBDate `bun:"pe_befristet,notnull"`
	EMail        string        `bun:"pe_ko_email,notnull"`
	Mobil        string        `bun:"pe_ko_mobil,notnull"`
	Privat       string        `bun:"pe_ko_privat,notnull"`
}

func TestDBDate1andMF(t *testing.T) {
	// Notice: This uses df_welobn with user myadmin for our testing
	sqldb, err := sql.Open("mysql", "myadmin:@/df_welobn")
	if err != nil {
		panic(err)
	}

	err = sqldb.Ping()
	if err != nil {
		t.Fatal("No DB")
	}
	db := bun.NewDB(sqldb, mysqldialect.New())
	// Debug output for everything!
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(false)))

	ctx := context.Background()
	var personal []Personal
	err = db.NewSelect().
		Model(&personal).
		Where("pe_status='GEPRUEFT'").
		Where("pe_ko_email!=''").
		Order("pe_ko_email").
		GroupExpr("pe_geburtstag='0000-00-00'").
		Scan(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range personal {
		if row.ID == 750876 {
			if !row.GeburtsDatum.IsZero() {
				t.Fatal("not detected as zero")
			}
			if row.GeburtsDatum.Format("2006-01-02") != "0001-01-01" {
				t.Fatalf("GeburtsDatum not detected as '0000-01-01' but %q", row.GeburtsDatum.Format("2006-01-02"))
			}
			if row.Gender.String() != "M" {
				t.Fatalf("Gender detected as 'M' but %q", row.Gender)
			}
		} else {
			if row.GeburtsDatum.IsZero() {
				t.Fatal("detected as zero")
			}
			if row.GeburtsDatum.Format("2006-01-02") != "1944-08-01" {
				t.Fatalf("GeburtsDatum not detected as '1944-08-01' but %q", row.GeburtsDatum)
			}
			if row.Gender.String() != "M" {
				t.Fatalf("Gender detected as 'M' but %q", row.Gender)
			}
		}
	}
}

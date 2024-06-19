package mssql

import (
	"database/sql"
	"fmt"
	"github.com/comsma/knead/app"
	_ "github.com/microsoft/go-mssqldb"
	"log"
	"net/url"
	"sync"
)

type Database struct {
	db     *sql.DB
	dbOnce sync.Once
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Connect(c *app.Config) {

	urlq := url.Values{}
	urlq.Add("database", c.Database.DB)

	u := &url.URL{Scheme: "sqlserver",
		User: url.UserPassword(c.Database.Username, c.Database.Password), Host: fmt.Sprintf("%s:%d", c.Database.Host, c.Database.Port), RawQuery: urlq.Encode()}

	d.dbOnce.Do(
		func() {
			db, err := sql.Open(
				"mssql",
				u.String(),
			)
			if err != nil {
				log.Fatalln(err)
			}

			d.db = db
		},
	)

}

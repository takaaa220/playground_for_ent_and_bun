package bunutil

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func Connect() *bun.DB {
	c := mysql.Config{
		DBName:    "sample_db",
		User:      "root",
		Passwd:    "rootpassword",
		Addr:      "localhost:33006",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_general_ci",
	}

	sqldb, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		panic(err)
	}

	return bun.NewDB(sqldb, mysqldialect.New())
}

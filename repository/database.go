// previous to run
// go get -u github.com/lib/pq

package repository

import (
  "database/sql"
  "fmt"
  
  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = ' '
  dbname   = "movies"
)

func QueryDB(query string) *sql.Rows {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil{
      panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil{
      panic(err)
    }

    rows, err := db.Query(query)

    if err != nil{
      //handle this error
      panic(err)
    }

    return rows
}

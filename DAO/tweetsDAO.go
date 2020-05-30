package dao

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type Tweet struct {
  Id int
  User_id int
  Body string
  Created_at string
  // deleated_at string
}

var empty interface{}

func GetTweets() (tweets []Tweet){
  dbms := "mysql"
  // dbuser := "editer"
  dbuser := "root"
  // dbpass := "Tr_kp_514021"
  dbpass := ""
  dbname := "twitter_clone"

  db, err := sql.Open(dbms, dbuser+":"+dbpass+"@/"+dbname)
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()


  rows, err := db.Query("select * from tweets")
  if err != nil {
    panic(err.Error())
  }
  for rows.Next() {
    tweet := Tweet{}
    err = rows.Scan(&tweet.Id, &tweet.User_id, &tweet.Body, &empty, &empty)
    if err != nil {
      panic(err.Error())
    }
    tweets = append(tweets, tweet)
  }
  return
}

func PostTweets(account_id int, body string) {
  dbms := "mysql"
  // dbuser := "editer"
  dbuser := "root"
  // dbpass := "Tr_kp_514021"
  dbpass := ""
  dbname := "twitter_clone"

  db, err := sql.Open(dbms, dbuser+":"+dbpass+"@/"+dbname)
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  stmtInsert, err := db.Prepare("insert into tweets(user_id, body) values(?,?) ")
  if err != nil {
    panic(err.Error())
  }
  defer stmtInsert.Close()
  _, err = stmtInsert.Exec(account_id, body)
  if err != nil {
    panic(err.Error())
  }
}

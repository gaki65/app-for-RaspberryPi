package main

import (
  "net/http"
  "html/template"
  "fmt"
  "twitter_clone/tmp/develop/DAO"
)

type Tweet struct {
  Id int
  User_id int
  Body string
  Created_at string
  // deleated_at string
}

func home(w http.ResponseWriter, r *http.Request) {

  r.ParseForm()
  fmt.Println(r.Form)  //フォームの本文受け取り
  if len(r.Form) != 0 {
    dao.PostTweets(1, r.FormValue("tweet"))
  }

  var t *template.Template
  t, _ = template.ParseFiles("./view/frame.html", "./view/tweets.html", "./view/tweet.html")
  t.ExecuteTemplate(w, "frame", dao.GetTweets())

}

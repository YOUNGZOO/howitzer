package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "gopkg.in/mgo.v2/bson"
  "fmt"
)

func getJournal(c *gin.Context){

  var journals []Journal
  acnDate := c.Param("acn_date")
  acnModel := c.Param("acn_model")

  fmt.Println(acnDate)
  fmt.Println(acnModel)

  err := mongoSession.DB("test").C("journals").Find(
  bson.M{
      "acn_date": acnDate,
      "acn_model": acnModel,
    }).Sort("timestamp").All(&journals)

  if err!= nil{
    panic(err)
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "입력한 날짜, 거래유형에 해당하는 전표가 없습니다"})
  }

  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": journals})
}

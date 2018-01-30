package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "gopkg.in/mgo.v2/bson"
)
func getAcnModel(c *gin.Context){
  var models []Model
  acnModel := c.Param("acn_model")
  err := mongoSession.DB("test").C("models").Find(
    bson.M{
      "acn_model": acnModel,
    }).Sort("timestamp").All(&models)

  if err!= nil{
    panic(err)
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "해당하는 거래유형이 없습니다"})
  }

  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": models})
}

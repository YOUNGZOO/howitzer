package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "time"
  "fmt"
  "strconv"
)

func newJournal(c *gin.Context){
  var journal Journal
  c.BindJSON(&journal)
  fmt.Println(journal.Acn_date)
  valid := journal.checkNil() && journal.checkLength() && validating(journal.Detail)

  if valid == true {
    newOrStill(len(journal.Detail))
    journal.Timestamp = time.Now()
    pool[poolPointer].DB("test").C("journals").Insert(journalConversion(journal))
    c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "전표 등록 성공"})
  }else{
    msg := "전표 입력값 확인"
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusBadRequest, "msg": msg})
  }
}


func validating(jd []JournalDetail) bool{
  var valid2 bool
  if len(jd) ==0 {
    fmt.Println("전표상세 없음")
    return false
  }

  for i:=0; i<len(jd); i++ {
    valid2 = jd[i].checkNil() && jd[i].checkLength()
    fmt.Println("valid2")
    fmt.Println(valid2)
    if !valid2 {
      return valid2
    }
  }

  return true

}


func journalConversion(j Journal) Journal2 {
  var j2 Journal2
  var amt float64
  amt, _ = strconv.ParseFloat(j.Acn_amount,64)
  j2 = Journal2{ Acn_date:j.Acn_date, Acn_model:j.Acn_model, Acn_amount:amt, Detail: journalDetailConversion(j.Detail) , Timestamp: j.Timestamp }
  return j2
}

func journalDetailConversion(jd []JournalDetail) []JournalDetail2 {
  var jd2 []JournalDetail2
  var amt float64
  for i:=0; i<len(jd); i++ {
    amt, _ = strconv.ParseFloat(jd[i].Detail_amount,64)
    jd2 = append(jd2, JournalDetail2{ Acn_code: jd[i].Acn_code, Credit_debit: jd[i].Credit_debit, Detail_amount: amt})
  }
  return jd2
}

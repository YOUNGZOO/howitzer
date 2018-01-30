package main

//입력에 공통적으로 사용할 validator
//true : 통과, false : 실패
import "fmt"

type Checkable interface{
  checkNil() bool
  checkLength() bool
}

func (j Journal) checkNil() bool{
  if j.Acn_date == "" || j.Acn_model =="" {
    fmt.Println(j.Acn_date)
    fmt.Println(j.Acn_model)
    fmt.Println("journal checknil false")
    return false
  }else{
    fmt.Println("journal checknil ok")
    return true
  }
}

func (j Journal) checkLength() bool{
  if len(j.Acn_date)!=8 || len(j.Acn_model) != 5 {
    fmt.Println(j.Acn_date)
    fmt.Println(j.Acn_model)
    fmt.Println("journal checklength false")
    return false
  }else{
    fmt.Println("journal checklength ok")
    return true
  }
}


func (jd JournalDetail) checkNil() bool{
  if jd.Acn_code == "" {
    return false
  }
  return true
}

func (jd JournalDetail) checkLength() bool{
  if len(jd.Acn_code) != 10 {
    return false
  }
  return true
}

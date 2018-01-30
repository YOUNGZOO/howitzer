package main

import "time"

// 회계전표 struct
// 금액필드 : context에 BindJSON시에 0 초가화 https://github.com/gin-gonic/gin/issues/491

type Journal struct {
  Acn_date      string    // 회계일자
  Acn_model     string    // 회계모형
  Acn_amount    string    // 금액
  Detail        []JournalDetail
  Timestamp     time.Time // timestamp
}

// 회계전표상세 struct
type JournalDetail struct {
  Acn_code      string    // 계정과목
  Credit_debit  bool      // T-차변 F-대변
  Detail_amount string    // 상세거래금액
}

// 회계유형 struct
type Model struct {
  Acn_model     string    // 회계모형
  Model_desc    string    // 모형설명
  Timestamp     time.Time // timestamp
}

type Journal2 struct {
  Acn_date      string    // 회계일자
  Acn_model     string    // 회계모형
  Acn_amount    float64    // 금액
  Detail        []JournalDetail2
  Timestamp     time.Time // timestamp
}

// 회계전표상세 struct
type JournalDetail2 struct {
  Acn_code      string    // 계정과목
  Credit_debit  bool      // T-차변 F-대변
  Detail_amount float64    // 상세거래금액
}

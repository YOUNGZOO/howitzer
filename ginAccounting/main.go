package main

import (
  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2"
  "fmt"
  "os"
)

//전역변수
//pool문제 : http://www.popit.kr/golang-mongodb-client-connection-pool-%EA%B4%80%EB%A0%A8/
var mongoSession *mgo.Session //read용
var pool [2]*mgo.Session      //write용 2개
var poolPointer int           //어떤 pool을 사용할 것인가
var err error

func main() {
  //init - router 이후 세션 초기화로 panic 발생
  if os.Getenv("CONTAINERIZED") == "true" {
    mongoSession, err = mgo.Dial(os.Getenv("MONGODB_ADDR"))
  }else{
    mongoSession, err = mgo.Dial("127.0.0.1")
  }

  fmt.Println("containerized :", os.Getenv("CONTAINERIZED"))
  fmt.Println("mongodb_addr :",os.Getenv("MONGODB_ADDR"))

  if err != nil {
    fmt.Println("mongodb 연결 실패")
    panic(err)
  }
  //pool pointer 초기화, write용 connection pooling을 위한 pre-copy
  poolPointer = 0
  for i := 0; i < len(pool) ; i++ {
    pool[i] = mongoSession.Copy()
	}

  //HTTP request routing
  router := gin.Default()
  group1 := router.Group("/")
  {
    group1.POST("/journal", newJournal)                     //회계전표 발생
    group1.GET("/journal/:acn_date/:acn_model", getJournal) //일자, 모형 조건 전표조회
    group1.GET("/model/:acn_model", getAcnModel)            //회계모형 조회
  }
  router.Run()
}


/**

(1) Session 구조체

type Session struct {
	m                sync.RWMutex
	cluster_         *mongoCluster
	slaveSocket      *mongoSocket
	masterSocket     *mongoSocket
	slaveOk          bool
	consistency      Mode
	queryConfig      query
	safeOp           *queryOp
	syncTimeout      time.Duration
	sockTimeout      time.Duration
	defaultdb        string
	sourcedb         string
	dialCred         *Credential
	creds            []Credential
	poolLimit        int
	bypassValidation bool
}

(2) Session 콘솔출력

mongoSession
&{{{0 0} 0 0 0 0} 0xc420218000 <nil> <nil> false 2 {{ <nil> 0 0 <nil> 0 <nil> 0 {<nil> <nil> <nil> false false [] 0 0 } false []} 0.25 0} 0xc42021c000 60000000000 60000000000 test admin <nil> [] 4096 false}
pool 0
&{{{0 0} 0 0 0 0} 0xc420218000 <nil> <nil> false 2 {{ <nil> 0 0 <nil> 0 <nil> 0 {<nil> <nil> <nil> false false [] 0 0 } false []} 0.25 0} 0xc42021c000 60000000000 60000000000 test admin <nil> [] 4096 false}
pool 1
&{{{0 0} 0 0 0 0} 0xc420218000 <nil> <nil> false 2 {{ <nil> 0 0 <nil> 0 <nil> 0 {<nil> <nil> <nil> false false [] 0 0 } false []} 0.25 0} 0xc42021c000 60000000000 60000000000 test admin <nil> [] 4096 false}

**/

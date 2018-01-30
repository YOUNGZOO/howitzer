package main
import "fmt"

//회게 상세가 2보다 긴경우는 무조건 pool을 바꿔가면서
func newOrStill(detail_len int) bool {
  if detail_len > 2 {
    change()
    fmt.Println("쓰기용 pool 바꿈")
    return true
  } else {
    fmt.Println("쓰기용 pool 그대로")
    return false
  }
}

//pool 바꾸기
func change() {
  poolPointer = (poolPointer+1)%len(pool)
}

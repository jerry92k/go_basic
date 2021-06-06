package main
import "fmt"

func main(){
  //맵(Map)
  // 맵 : 해쉬테이블, 딕셔너리(파이썬), key-value로 자료 저장
  // 레퍼런스 타입
  // 비교연산자 사용 불가능(참조 타입이므로)
  // 특징 : 참조타입(key)로 사용 불가능, 값(value)로 모든 타입 사용 가능
  // make 함수 및 축약(리터럴)로 초기화 가능
  // 순서 없음

  // 예제1
  var map1 map[string] int = make(map[string]int) // key:string value:int
  var map2 = make(map[string]int) // 자료형 생략
  map3:=make(map[string]int) // 리터럴

  fmt.Println("ex1 :"map1)
  fmt.Println("ex1 :"map2)
  fmt.Println("ex1 :"map3)

  //예제2
  map4:=map[string]int{} // json형태
  map4["apple"] =25
  map4["banana"]=40
  map4["orange"]=33

  map5:=map[string]int{
    "apple":15,
    "banana":40,
    "orange":23,
  }

  map6=make(map[stirng]int, 10)
  map6["apple"] =25
  map6["banana"]=40
  map6["orange"]=33


}

package testpkg

import (
	"fmt"
)

/*
패키지 import 시 최초 실행되는 메소드
*/
func init() {
	fmt.Println("=====")
	fmt.Println("Package is imported!")
	fmt.Println("=====")
}

// TInterface 인터페이스는 메소드의 집합
type TInterface interface {
	Method()
}

type TStruct struct {
	name string
	age  int
}

/*
NewTStruct 생성자를 이용한 구조체 초기화
*/
func NewTStruct() *TStruct {
	// 구조체 초기화 방법
	return &TStruct{
		name: "Kevin",
		age:  17,
	}
}

/*
Method public method 예시 (대문자는 public 메소드)
메서드 앞의 () 는 receiver 라고 한다.
*/
func (t *TStruct) Method() {
	// Receiver 를 포인터로 선언하면 멤버 변수 값을 수정할 수 있다.
	t.pMethod()
	fmt.Println(t.name, t.age)
}

/*
Private method 는 맨 앞 글자 소문자로 작성
*/
func (t *TStruct) pMethod() {
	t.age += 1
}

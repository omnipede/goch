package testpkg

import (
	"fmt"
	"os"
)

/*
MyError Custom 에러 정의 방법
*/
type MyError struct {
	msg string
}

func (m *MyError) Error() string {
	return m.msg
}

/*
ReadFile 에러 헨들링 및 defer & panic 관련
*/
func ReadFile() (*os.File, error) {
	defer func() {
		// Panic 을 복구한다.
		if r := recover(); r != nil {
			fmt.Println("Read file error", r)
		}
	}()

	f, err := os.Open("NOT_EXIST")
	if err != nil {
		// 프로그램 실행을 멈추고 defer 함수를 실행한다.
		panic(&MyError{msg: "custom error"})
	}

	return f, nil
}

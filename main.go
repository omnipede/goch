package main

import (
	"fmt"
	"go-cheatsheet/testpkg"
)

func main() {
	// Variable decl
	var a [10]int         // array
	b := make([]int, 0)   // slice
	const c = 7           // const
	d := map[int]string{} // map

	// Control (loop and condition)
	for i := 0; i < 10; i += 1 {
		a[i] = c * i
		b = append(b, i)
		if i%2 == 0 {
			d[i] = "Hello"
		} else {
			d[i] = "World"
		}
	}

	// Slice
	// Concat slice
	s := append(b, b...)
	// Delete one
	s = append(s[0:1], s[1:]...)
	fmt.Println(s)

	// Map
	// Key 존재하는지 확인 방법
	_, exists := d[101]
	if !exists {
		fmt.Println("Not exists.")
	}
	// Map iteration
	for key, val := range d {
		fmt.Println(key, val)
	}

	// Function call
	i1, i2 := multiReturn()
	fst, snd := namedReturn()
	// 익명 함수
	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(i1, fst), add(i2, snd))

	// Call external pkg
	var t testpkg.TInterface = testpkg.NewTStruct()
	t.Method()

	// Error 헨들링 예제
	_, err := testpkg.ReadFile()
	if err != nil {
		return
	}

	// Go 루틴
	testpkg.TestGoRoutine()

	// 예제
	testpkg.TestChannel()
}

func multiReturn() (int, int) {
	return 1, 2
}

func namedReturn() (first int, second int) {
	first += 1
	second += 2
	return
}

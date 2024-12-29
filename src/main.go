package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int
	fmt.Printf("%T\n", i)
	reflectedType := reflect.TypeOf(i) // 型情報の取得
	fmt.Printf("%T\n", reflectedType)

	reflectedValue := reflect.ValueOf(i) // 値情報の取得 ( ポインタの場合, 値を変更できる)
	fmt.Printf("%T\n", reflectedValue)

	iInterface := reflectedValue.Interface()
	fmt.Printf("%T\n", iInterface)

	// 値型でアドレス情報を持たないため、値を設定できない
	if reflectedValue.CanSet() != false {
		panic("expected CanSet is false, but I got true.")
	}

	// アドレスを指定すれば値をセットできる
	reflectedValue = reflect.ValueOf(&i)
	if ok := reflectedValue.Elem().CanSet(); ok == false {
		panic(fmt.Sprintf("expected CanSet is true, but I got %t.", ok))
	}

	reflectedValue.Elem().SetInt(10)
	// reflectedValue.Elem().SetString("hoge") // panicが起こる

	type IntEx int

	reflectedType = reflect.TypeFor[IntEx]()
	zeroPointer := reflect.New(reflectedType)
	fmt.Printf("%v\n", zeroPointer)
  fmt.Printf("%v\n", zeroPointer.Elem())

}

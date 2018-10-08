package main

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

func main() {
	// 创建
	u1,_:= uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s\n", u2)

	res:=strings.Split("890f9112-7951-4e27-9f5b-b1fa4a49a059lockTest10000000002","lockTest1")
	for i,val:=range res{
		fmt.Println(i)
		fmt.Println(val)
	}

}
package main

import (
	"fmt"
	"github.com/liujiage/restapi/service"
)

func main() {
	reply := service.Say("jiage liu")
	fmt.Println(reply)

}
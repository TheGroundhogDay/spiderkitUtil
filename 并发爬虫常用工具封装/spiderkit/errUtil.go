package spiderkit

import (
	"fmt"
	"os"
)

func HandleError(err error, when string) {
	if err != nil{
		fmt.Println(when,err)
		os.Exit(1)
	}
}

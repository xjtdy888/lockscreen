package main 

import (
	"github.com/xjtdy888/screenlock"
	"fmt"
)

func main() {
	err := screenlock.Lock()
	fmt.Println(err)
}



package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Timer start")
	t := time.NewTimer(10 * time.Second)
	<-t.C
	fmt.Println("Timer expired")
}

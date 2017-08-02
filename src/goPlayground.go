package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t, time.Unix(0, 0))
}

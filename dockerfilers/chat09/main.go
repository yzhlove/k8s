package main 

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	fmt.Printf("[%s,%s] golang.<hello world> [%s] \n",runtime.GOOS,runtime.GOARCH,time.Now().Format(time.RFC3339))
}
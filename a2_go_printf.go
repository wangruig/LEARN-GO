package main

import (
	"fmt"
)

func main() {
	var stockcode = 30
	var enddate = "20220318"
	var url = "code=%denddate=%s"
	fmt.Printf(url, stockcode, enddate)
}

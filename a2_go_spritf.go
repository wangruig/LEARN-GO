package main

import (
	"fmt"
)

func main() {
	var stockcode = 20
	var enddate = "20220318"
	var url = "Code=%denddate=%s"
	var targeturl = fmt.Sprintf(url, stockcode, enddate)

	fmt.Println(targeturl)
}

package main

import (
	"fmt"
	"time"
)

func GetNow() time.Time {
	return time.Now()
}

func DayOfMonth() int {
	return GetNow().Day()
}

func main() {
	fmt.Println(DayOfMonth())
}

package main

import (
	"fmt"
	"time"
)

type Clock struct {
	name string
}

func (clock *Clock) GetNow() time.Time {
	return time.Now()
}

func Hello() string {
	return "Hello!"
}

func main() {
	clock := &Clock{
		name: "my-clock",
	}
	fmt.Println(clock.GetNow())
	fmt.Println(Hello())
}

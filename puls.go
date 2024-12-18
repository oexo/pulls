package main

import (
	"fmt"
	"log"
	"time"
)

type Pushups struct {
	count map[time.Month][]int
}

type Pullups struct {
	count map[time.Month][]int
}

type Puls struct {
	Pushups
	Pullups
}

func main() {
	fmt.Println("vim-go")
	pushups := Pushups{count: make(map[time.Month][]int)}
	fmt.Println(pushups)
	year, month, day := time.Now().Date()
	log.Printf("%v-%v-%v", year, int(month), day)

	pushups.count[month] = append(pushups.count[month], 1)
	pushups.count[month] = append(pushups.count[month], 1)
	fmt.Println(pushups)
	pulls := Puls{Pushups: pushups, Pullups: Pullups{count: make(map[time.Month][]int)}}
	fmt.Println(pulls)
}

package main

import "fmt"

type mysqlresult struct {
	time  float32
	count int
	total float32
}

func percentile(histogram []int, percentile int) int {
	fmt.Println("true")
}

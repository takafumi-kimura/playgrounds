package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Seedを入れないと毎回同じランタイムで同じランダムなデータが生成される
	//そのため、Seedに毎回変わる値を入れることによって再現性のないランダムな値を取り出せるようになる
	//https://pkg.go.dev/math/rand#Intn
	fmt.Println("毎回同じランダムな値が生成される")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println("本当にランダムな値が生成される")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}

package main

import "fmt"

func main() {
	fmt.Println(sum(5, 2))
	fmt.Println(sum2(5, 2))
	fmt.Println(swap(5, 2))
	fmt.Println(namedReturn(5, 2))

}

//型はスペースで区切る
//返り値は引数の後ろに書く
func sum(x int, y int) int {
	return x + y
}

//引数の型が同じ場合は以下のように書くこともできる
func sum2(x, y int) int {
	return x + y
}

//関数の返値を複数指定することもできる
func swap(x, y int) (int, int) {
	return y, x
}

//関数の返値をに名前をつける
func namedReturn(x, y int) (r1 int, r2 int) {
	r1 = y
	r2 = x
	// returnに何も書かなくて良いメリットがあるけれど、可読性は下がる
	return
}

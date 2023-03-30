package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return => 短い関数でのみ利用すべき
}

func main() {
	fmt.Println(split(17))
}

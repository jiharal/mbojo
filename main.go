package main

import (
	"fmt"

	array "github.com/jiharal/mbojo/array"
	channel "github.com/jiharal/mbojo/channel"
	math "github.com/jiharal/mbojo/math"
)

func main() {
	// Array
	valOfDay3 := []int{1, 2, 1, 3, 4, 5, 6, 7, 33, 22, 32}
	fmt.Println(array.SumWithReturnOnValue(valOfDay3))
	fmt.Println(array.SecondLargeOfArray(valOfDay3))
	fmt.Println(math.Rectangle(4, 5))

	// Channel
	channel.ChannelBasic()
	channel.ChannelAsATypeAndParam()

}
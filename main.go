package main

import (
	"fmt"

	Stack "stack.com/m/v2/custommodules.com/CustomStack"
	mapfunc "stack.com/m/v2/custommodules.com/MapFunc"
	singleton "stack.com/m/v2/custommodules.com/Singleton"
)

func AddTwo(num interface{}) interface{} {
	return num.(int) + 2
}

func main() {
	stack := Stack.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Ende")
	fmt.Println(singleton.Singleton)

	arr := []int{1, 2, 3, 4, 5}
	mappedArr := mapfunc.Map(arr, AddTwo)
	fmt.Println(mappedArr)
}

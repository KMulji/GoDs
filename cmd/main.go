package main

import "fmt"

func main() {
	// var a *ds.DynamicArray[int] = ds.DynamicArrayInit[int](2)
	// a.Append(1)
	// a.Append(2)
	// a.Append(3)
	// a.Insert(2, 50)
	// a.Pop()
	// a.Pop()
	// a.Pop()
	// a.Pop()
	// a.Pop()
	// a.Append(1)
	// a.Append(2)
	// a.Append(3)
	// a.Append(4)
	// a.Append(5)
	// a.Insert(4, 200)
	// a.Insert(3, 100)
	// a.Insert(7, 400)
	// a.Insert(6, 300)
	// a.Delete(3)

	// fmt.Println(a)
	// a.BinSearch(1000)
	// a := []int{4, 4, 1, 1, 4, 8, 7, 7, 7, 8, 4}
	// fmt.Println(ds.FindOdd([]int(a), len(a)))

	//10100
	a := 20
	fmt.Printf("%d", (a & (a - 1)))

}

package ds

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

type DynamicArray[T constraints.Ordered] struct {
	a      []T
	size   int
	length int
}

func DynamicArrayInit[T constraints.Ordered](size int) *DynamicArray[T] {
	return &DynamicArray[T]{make([]T, size), size, 0}
}

func (da *DynamicArray[T]) Append(val T) {
	if da.size == da.length {
		da.grow()
	}
	da.a[da.length] = val
	da.length++
}
func (da *DynamicArray[T]) Pop() (T, error) {
	var zeroValue T
	if da.length == 0 {
		return zeroValue, errors.New("array is empty")
	}
	if da.length == da.size/2 {
		da.shrink()
	}
	temp := da.a[da.length-1]
	da.a[da.length-1] = zeroValue
	da.length--
	return temp, nil

}

func (da *DynamicArray[T]) Insert(index int, val T) {
	if da.size == da.length {
		da.grow()
	}
	i := da.length

	for i != index {
		da.a[i] = da.a[i-1]
		i--
	}
	da.a[index] = val
	da.length++

}

func (da *DynamicArray[T]) Delete(index int) (T, error) {
	var zeroValue T
	if da.length == 0 {
		return zeroValue, errors.New("array is empty")
	}
	if da.length == da.size/2 {
		da.shrink()
	}
	i := index

	for i < da.length-1 {
		da.a[i] = da.a[i+1]
		i++
	}
	temp := da.a[index]
	da.length--
	da.a[da.length] = zeroValue
	return temp, nil

}
func (da *DynamicArray[T]) BinSearch(key T) {
	lo := 0
	hi := da.length - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if da.a[mid] == key {
			fmt.Printf("Found at %d\n", mid)
			return
		} else if key > da.a[mid] {
			lo = mid + 1
		} else if key < da.a[mid] {
			hi = mid - 1
		}

	}
	fmt.Println("Key not found")

}

func (da *DynamicArray[T]) grow() {
	temp := make([]T, da.size*2)

	copy(temp, da.a)
	da.size = da.size * 2
	da.a = temp

}
func (da *DynamicArray[T]) shrink() {
	temp := make([]T, da.size/2)

	copy(temp, da.a)
	da.size = da.size / 2
	da.a = temp
}

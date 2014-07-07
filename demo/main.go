package main

import (
	"fmt"
	"math/rand"
)

// +stream
type T struct {
	A int
	B string
}

func main() {
	fmt.Println("Slice")
	TStreamFromSlice(
		T{1, "a"},
		T{2, "b"},
		T{3, "c"},
		T{4, "d"},
		T{5, "e"},
		T{6, "f"},
		T{7, "g"},
		T{8, "h"},
		T{9, "i"},
	).Filter(func(t T) bool {
		return t.A%2 == 0
	}).Modify(func(t T) T {
		return T{
			A: t.A * 4,
			B: t.B + "FOO",
		}
	}).Each(func(t T) {
		fmt.Println(t)
	}).Drain()

	fmt.Println("Generate")
	count := 0
	TStreamFromGenerator(func() (T, bool) {
		count++
		return T{
			A: rand.Intn(100),
			B: "bar",
		}, count < 10
	}).Filter(func(t T) bool {
		return t.A%2 == 0
	}).Modify(func(t T) T {
		return T{
			A: t.A * 4,
			B: t.B + "FOO",
		}
	}).Each(func(t T) {
		fmt.Println(t)
	}).Drain()

	fmt.Println("Merge")
	TStreamMerge(
		TStreamFromSlice(T{}, T{}),
		TStreamFromSlice(T{12, "x"}, T{54, "boo"}),
	).Filter(func(t T) bool {
		return t.A%2 == 0
	}).Modify(func(t T) T {
		return T{
			A: t.A * 4,
			B: t.B + "FOO",
		}
	}).Each(func(t T) {
		fmt.Println(t)
	}).Drain()

	list := TStreamMerge(
		TStreamFromSlice(T{5, "f"}, T{7, "g"}),
		TStreamFromSlice(T{12, "x"}, T{54, "boo"}),
	).Filter(func(t T) bool {
		return t.A%2 == 0
	}).Modify(func(t T) T {
		return T{
			A: t.A * 4,
			B: t.B + "FOO",
		}
	})

	fmt.Println("Take 2")

	for i := 0; i < 2; i++ {
		fmt.Println(i, list.Next())
	}
	list.Close()

	fmt.Println(list.Drain())
}

package main

import (
	"errors"
	"fmt"
)

// Stack LIFO (Last in First Out)
type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

//-------------------------------------------------

//?? Generic Functions Abstract Algorithms

// Map turn a []T1 to a [T]2 using a mapping function
// This function has two parameter types T1 and T2
// This works with slices of any type
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Filter values from a slice using a filter function
// It returns only a slice with the elements of s
// for which f returned true
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reduce reduces a []T1 to all single values using a reduction function
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

//------------------------------------------------------

//?? Type terms to specify operators

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func divAndRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

// lists all types that support the ==, !=, <, >, <=, and >= operators
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

//lists all types that support the ==, !=, <, >, <=, and >= operators
type PrintableInt interface {
	~int
	Sting() string
}

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	// intStack.Push("nope") compiler will catch this error
	intStack.Push(20)
	intStack.Push(30)
	v, ok := intStack.Pop()
	fmt.Println(v, ok)
	fmt.Println(intStack.Contains(10))

	words := []string{"One", "Potato", "Two", "Potato"}
	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)

	lengths := Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)

	sum := Reduce(lengths, 0, func(acc int, val int) int {
		return acc + val
	})
	fmt.Println(sum)

	//-------------------------------------------------
	var a uint = 18_446_744_073_709_551_615
	var b uint = 9_223_372_036_854_775_808
	fmt.Println(divAndRemainder(a, b))
}

package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

//--------------------------------------

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increase() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

//----rules for passing values to function still exits------

// Function will update but fail to retain update after function life cycle
func doUpdateWrong(c Counter) {
	c.Increase()
	fmt.Println("Update wrong", c.String())
}

func updateCorrectly(c *Counter) {
	c.Increase()
	fmt.Println("Update correctly", c.String())
}

//----methods for nil instances---------------------------

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func main() {
	p := Person{
		FirstName: "Steven",
		LastName:  "Universe",
		Age:       13,
	}

	output := p.String()
	fmt.Println(output)

	//-------------------------------------

	var c Counter
	fmt.Println(c.String())
	c.Increase()
	fmt.Println(c.String())

	//If you call a value receiver on a pointer variable, go automatically dereferences
	o := &Counter{}
	fmt.Println(o)
	o.Increase()
	fmt.Println(o) //the call o.String() is silently converted to (*o).String().

	//-------------------------------------

	var u Counter
	doUpdateWrong(u)
	fmt.Println("in main", u.String())
	updateCorrectly(&u)
	fmt.Println("in main", u.String())

	//_____________________________________

	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
	fmt.Println(*it.left) // false
}

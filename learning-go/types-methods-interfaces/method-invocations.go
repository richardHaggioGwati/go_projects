package main

import (
	"fmt"
	"time"
)

// --------iota----------------------------
type MailCategory int

const (
	Uncategorized MailCategory = iota //(0) is assigned to the first Constant, 1 is assigned to Personal and so on
	Personal
	Spam
	Social
	Advertisement
)

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4        //Assigned to 20 because it has no type | value hence it takes from the previous one
	Field5 = iota //Get assigned 4 because it's the fifth line and iota starts counting from 0

)

// ---------composition and promotion-------
type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee // followed by business logic

// -------embedding functions---------------
type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

//------------interfaces are comparable----------
type Doubler interface {
	Double()
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

// -----------------------------------------
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
	//-----iota--------------------------
	fmt.Println("IOTA", Field1, Field2, Field3, Field4, Field5)

	p := Person{
		FirstName: "Steven",
		LastName:  "Universe",
		Age:       13,
	}

	output := p.String()
	fmt.Println(output)

	//---------embedding-------------------
	outer := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(outer.Double())

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
	fmt.Println(*it.left)        // false

	//------------comparing interfaces---------
	var di1 DoubleInt = 10
	var di2 DoubleInt = 10
	var dis1 = DoubleIntSlice{1, 2,3}
	// var dis2 = DoubleIntSlice{1,2,3}

	DoublerCompare(&di1, &di2) // we are comparing pointers here and not values hence we get false because they do not have the same instance
	DoublerCompare(&di1, dis1) // types do not match
	// DoublerCompare(dis1, dis2) //?? This code will trigger a panic
}

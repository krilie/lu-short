package ndb

import "testing"

type I interface {
	Hello()
}

type O struct{}

func (o *O) Hello() {
	println("o empty")
}

type A struct {
	I
	nameA string
}

type IB interface {
	I
	HelloB()
}

type B struct {
	I
	nameB string
}

func (b *B) HelloB() {
	println(b.nameB + " on IB")
}

func (b *B) Hello() {
	println(b.nameB)
}

type C struct {
	I
	nameC string
}

func TestStruct(t *testing.T) {
	var o = &O{}
	var a = &A{
		I:     o,
		nameA: "name a",
	}
	o.Hello()
	a.Hello()
	o2, ok := interface{}(a).(*O)
	println(ok)
	if ok {
		o2.Hello()
	}
	var b = &B{
		I:     a,
		nameB: "name b",
	}
	b.Hello()
	b.HelloB()
	var c = &C{
		I:     b,
		nameC: "name c",
	}
	c.Hello()
	b2, ok := interface{}(c).(*B)
	if !ok {
		println(false)
	} else {
		b2.HelloB()
	}
}

type one struct {
	name string
}

func (o *one) Hello2() {
	println("hello one")
}

type two struct {
	name2 string
}

func (o *two) Hello2() {
	println("hello two")
}

type outInterface interface {
	Hello2()
}

func Test2(t *testing.T) {
	var a outInterface = &one{}
	a.Hello2()
	var b outInterface = &two{}
	b.Hello2()
	_, ok := a.(*one)
	println(ok)
}

package main

import (
	"container/list"
	"fmt"
)

func spinlockRing(steps int, moves int, value int) {
	l := list.New()
	e := l.PushBack(0)
	pos := 0
	for i := 1; i <= moves; i++ {
		newpos := (pos + steps) % i
		for ; pos < newpos; pos++ {
			e = e.Next()
		}
		for ; newpos < pos; pos-- {
			e = e.Prev()
		}
		e = l.InsertAfter(i, e)
		pos = (pos + 1) % (i + 1)
	}
	e = e.Next()
	pos++
	fmt.Println("pos=", pos, "value=", e.Value)
}

func spinlock(steps int, moves int) {
	value := 0
	pos := 0
	for i := 1; i <= moves; i++ {
		pos = (pos+steps)%i + 1
		if pos == 1 {
			value = i
		}
	}
	fmt.Println("pos=", 1, "value=", value)
}

func main() {
	// Part 1
	spinlockRing(370, 2017, 2017)

	// Part 2
	spinlock(370, 50000000)
}

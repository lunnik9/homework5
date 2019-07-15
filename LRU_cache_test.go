package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	l := LRUCache{
		list.New(),
		make(map[string]*list.Element),
		4,
	}
	test1 := "test1"
	test2 := "test2"
	test3 := "test228"
	l.Add(test1)
	l.Add(test2)
	l.Add(test3)
	l.Add(test1)
	fmt.Println(l)
}

func TestCheck(t *testing.T) {
	l := LRUCache{
		list.New(),
		make(map[string]*list.Element),
		4,
	}
	fmt.Println(l.Check("privet"))
	l.Add("privet")
	fmt.Println(l.Check("privet"))
}

func TestRemove(t *testing.T) {
	l := LRUCache{
		list.New(),
		make(map[string]*list.Element),
		4,
	}
	l.Add("privet")
	fmt.Println(l.Check("privet"))
	l.Remove("privet")
	fmt.Println(l.Check("privet"))

}

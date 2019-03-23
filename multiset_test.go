package multiset

import (
	"fmt"
	"sort"
)

func ExampleNew() {
	m := New("a", "a", "b")
	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a a b]
}

func ExampleMultiset_Add() {
	m := New("a", "a", "a")
	m.Add("b", "b", "c", "d")
	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a a a b b c d]
}

func ExampleMultiset_AddCount() {
	m := New()
	m.AddCount("d", 1)
	m.AddCount("a", 3)
	m.AddCount("c", 1)
	m.AddCount("b", 2)
	var list []string
	for elem := range m.Iter() {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a a a b b c d]
}

func ExampleMultiset_Contains() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.Contains("b", "c", "d"))
	// Output:
	// true
}

func ExampleMultiset_Equals() {
	m1 := New("b", "b", "c", "d")
	m2 := New("c", "b", "d", "b")
	fmt.Println(m1.Equals(m2))
	// Output:
	// true
}

func ExampleMultiset_Count() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.Count("b"))
	// Output:
	// 2
}

func ExampleMultiset_Clear() {
	m := New("b", "b", "c", "d")
	m.Clear()
	fmt.Println(m)
	// Output:
	// []
}

func ExampleMultiset_IsEmpty() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.IsEmpty())
	// Output:
	// false
}

func ExampleMultiset_Size() {
	m := New("b", "b", "c", "d")
	fmt.Println(m.Size())
	// Output:
	// 4
}

func ExampleMultiset_Remove() {
	m := New("a", "a", "a", "b", "b")
	m.Remove("a", "a")
	fmt.Println(m)
	// Output:
	// [a b b]
}

func ExampleMultiset_RemoveCount() {
	m := New("a", "a", "a", "b", "b")
	m.RemoveCount("a", 3)
	fmt.Println(m)
	// Output:
	// [b b]
}

func ExampleMultiset_Iter() {
	m := New("a", "a", "a")
	for elem := range m.Iter() {
		fmt.Println(elem)
	}
	// Output:
	// a
	// a
	// a
}

func ExampleMultiset_DistinctElements() {
	m := New()
	m.AddCount("a", 3)
	m.AddCount("b", 2)
	var list []string
	distinct := m.DistinctElements()
	for _, elem := range distinct {
		list = append(list, elem.(string))
	}
	sort.Strings(list)
	fmt.Println(list)
	// Output:
	// [a b]
}

# Multiset
[![CircleCI branch](https://img.shields.io/circleci/project/github/trivigy/multiset/master.svg?label=master&logo=circleci)](https://circleci.com/gh/trivigy/workflows/multiset)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE.md)
[![](https://godoc.org/github.com/trivigy/multiset?status.svg&style=flat)](http://godoc.org/github.com/trivigy/multiset)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/trivigy/multiset.svg?style=flat&color=e36397&label=release)](https://github.com/trivigy/multiset/releases/latest)

Multiset is a threadsafe abstract data structure library for representing 
collection of distinct values, without any particular order. Unlike a set, 
multiset allows multiple instances for each of its elements.

### Example
```go
package main

import (
    "fmt"
    
    "github.com/trivigy/multiset"
)

func main() {
    m := multiset.New("b", "b", "c", "d")
    fmt.Println(m.Contains("b", "c", "d"))
    
    m1 := multiset.New()
    m1.AddCount("a", 3)
    m1.AddCount("b", 2)
    fmt.Println(m1.DistinctElements())
}
```

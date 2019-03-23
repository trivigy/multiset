// Package multiset implements a map-based multiset data structure.
package multiset

import (
	"fmt"
	"sync"
)

// Multiset represents the multiset data structure.
type Multiset struct {
	sync.RWMutex
	elems map[interface{}]int
}

// New creates and returns a reference to an empty set.  Operations
// on the resulting set are thread-safe.
func New(l ...interface{}) *Multiset {
	m := Multiset{elems: make(map[interface{}]int)}
	for _, e := range l {
		m.addCount(e, 1)
	}
	return &m
}

// String returns a string representation of the multiset.
func (m *Multiset) String() string {
	return fmt.Sprint(m.ToSlice())
}

// Add adds a single occurrence of each of the specified elements to this
// multiset.
func (m *Multiset) Add(s ...interface{}) {
	m.Lock()
	defer m.Unlock()
	for _, e := range s {
		m.addCount(e, 1)
	}
}

// AddCount adds the stated number of occurrences to the specified element of
// this multiset. Returns the number of occurrence of the element before the
// operation; possibly zero. The number of occurrences provided may be zero or
// negative value, in which case no change will be made.
func (m *Multiset) AddCount(e interface{}, c int) int {
	m.Lock()
	defer m.Unlock()
	return m.addCount(e, c)
}

func (m *Multiset) addCount(e interface{}, c int) int {
	count := m.count(e)
	if c <= 0 {
		return count
	}

	m.elems[e] += c
	return count
}

// Contains determines whether this multiset contains at least one occurrence
// of each of the specified elements. If at least one occurrence present
// returns true; otherwise false.
func (m *Multiset) Contains(s ...interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	return m.contains(s...)
}

func (m *Multiset) contains(s ...interface{}) bool {
	for _, e := range s {
		if m.count(e) <= 0 {
			return false
		}
	}
	return true
}

// Equals compares the specified multiset with this multiset for equality.
// Returns true if this multiset contains equal elements with equal counts;
// false otherwise.
func (m *Multiset) Equals(o *Multiset) bool {
	m.RLock()
	o.RLock()
	defer m.RUnlock()
	defer o.RUnlock()
	return m.equals(o)
}

func (m *Multiset) equals(o *Multiset) bool {
	if len(m.elems) != len(o.elems) {
		return false
	}

	for elem, count := range m.elems {
		if !o.contains(elem) || o.count(elem) != count {
			return false
		}
	}
	return true
}

// Count returns the number of occurrences of the specified element if present
// in this multiset; zero otherwise.
func (m *Multiset) Count(e interface{}) int {
	m.RLock()
	defer m.RUnlock()
	return m.count(e)
}

func (m *Multiset) count(s ...interface{}) int {
	var total int
	for _, e := range s {
		if count, ok := m.elems[e]; ok {
			total += count
		}
	}
	return total
}

// Clear removes all of the elements from this multiset. The collection will
// be empty after this method returns.
func (m *Multiset) Clear() {
	m.Lock()
	defer m.Unlock()
	m.elems = make(map[interface{}]int)
}

// IsEmpty returns true if this multiset contains no elements; false otherwise.
func (m *Multiset) IsEmpty() bool {
	m.RLock()
	defer m.RUnlock()
	return m.size() == 0
}

// Size returns the total number of elements in this multiset.
func (m *Multiset) Size() int {
	m.RLock()
	defer m.RUnlock()
	return m.size()
}

func (m *Multiset) size() int {
	var total int
	for _, count := range m.elems {
		total += count
	}
	return total
}

// Remove removes a single occurrence of each of the specified elements from
// this multiset, if present. Returns true if this multiset changed as a result
// of the call
func (m *Multiset) Remove(s ...interface{}) bool {
	m.Lock()
	defer m.Unlock()
	var changed bool
	for _, e := range s {
		if count := m.removeCount(e, 1); count > 1 {
			changed = true
		}
	}
	return changed
}

// RemoveCount removes a number of occurrences of the specified element from
// this multiset. Returns the count of occurrence of the element before the
// operation; possibly zero. The number of occurrences provided may be zero or
// negative value, in which case no change will be made.
func (m *Multiset) RemoveCount(e interface{}, c int) int {
	m.Lock()
	defer m.Unlock()
	return m.removeCount(e, c)
}

func (m *Multiset) removeCount(e interface{}, c int) int {
	count := m.count(e)
	if c <= 0 {
		return count
	}

	if _, ok := m.elems[e]; ok {
		m.elems[e] -= c
		if m.elems[e] <= 0 {
			delete(m.elems, e)
		}
	}
	return count
}

// Iter returns a channel of elements that can be ranged over.
func (m *Multiset) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		m.RLock()
		defer m.RUnlock()
		for elem, count := range m.elems {
			for i := 0; i < count; i++ {
				ch <- elem
			}
		}
		close(ch)
	}()
	return ch
}

// DistinctElements returns a slice containing only distinct elements
// of this multiset.
func (m *Multiset) DistinctElements() []interface{} {
	slice := make([]interface{}, 0, len(m.elems))
	m.RLock()
	defer m.RUnlock()
	for elem := range m.elems {
		slice = append(slice, elem)
	}
	return slice
}

// ToSlice returns a slice containing all elements of this multiset.
func (m *Multiset) ToSlice() []interface{} {
	slice := make([]interface{}, 0, m.Size())
	m.RLock()
	defer m.RUnlock()
	for elem, count := range m.elems {
		for i := 0; i < count; i++ {
			slice = append(slice, elem)
		}
	}
	return slice
}

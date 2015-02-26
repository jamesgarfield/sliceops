/*
The MIT License (MIT)

Copyright (c) 2015 James Garfield

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package sliceops

import (
	"sort"
)

//Impl Types: Requires a slice of any type (interface{}) in source file
//e.g. type Strings []string
type I interface{}
type Slice []I

//Related Type: Enables sorting as a slice operation
//e.g. Expands to 
//type StringsSorter {
//	Strings
//	LessFunc func(string, string) bool}
type _Sorter struct {
	Slice
	LessFunc func(I, I) bool
}

func (s _Sorter) Less(i, j int) bool {
	return s.LessFunc(s.Slice[i], s.Slice[j])
}

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//Sorts the slice in-place using the related type sorter
func (s Slice) Sort(less func(I, I) bool) {
	sort.Sort(_Sorter{s, less})
}

//Returns true if all elements in the slice return true for the provided function
func (s Slice) All(fn func(I) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

//Returns true if any elements in the slice return true for the provided function
func (s Slice) Any(fn func(I) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

//Returns the number of elements in the slice that return true for the provided function
func (s Slice) Count(fn func(I) bool) int {
	count := 0
	for _, v := range s {
		if fn(v) {
			count += 1
		}
	}
	return count
}

//Run the provided function on each element of the slice
func (s Slice) Each(fn func(I)) {
	for _, v := range s {
		fn(v)
	}
}

//Return the first element in the slice to return true for the provided function
func (s Slice) First(fn func(I) bool) (match I, found bool) {
	for _, v := range s {
		if fn(v) {
			match = v
			found = true
			break
		}
	}
	return
}

//Return a new slice of elements that return true for the provided function
func (s Slice) Where(fn func(I) bool) (result Slice) {
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

//Return a new slice of elements that have been removed from this slice
func (s *Slice) Extract(fn func(I) bool) (removed Slice) {
	pos := 0
	kept := *s
	for i := 0; i < kept.Len(); i++ {
		if fn(kept[i]) {
			removed = append(removed, kept[i])
		} else {
			kept[pos] = kept[i]
			pos++
		}
	}

	kept = kept[:pos:pos]
	*s = kept
	return removed
}

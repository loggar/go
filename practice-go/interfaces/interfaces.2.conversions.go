package main

// Sequence ..
import (
	"fmt"
	"sort"
)

// Sequence2 ..
type Sequence2 []int

/*
The String method of Sequence is recreating the work that Sprint already does for slices. We can share the effort if we convert the Sequence to a plain []int before calling Sprint.
*/

/*
func (s Sequence) String() string {
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
*/

/*
This method is another example of the conversion technique for calling Sprintf safely from a String method. Because the two types (Sequence and []int) are the same if we ignore the type name, it's legal to convert between them. The conversion doesn't create a new value, it just temporarily acts as though the existing value has a new type. (There are other legal conversions, such as from integer to floating point, that do create a new value.)
*/

/*
It's an idiom in Go programs to convert the type of an expression to access a different set of methods. As an example, we could use the existing type sort.IntSlice to reduce the entire example to this:
*/

// Method for printing - sorts the elements before printing
func (s Sequence2) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

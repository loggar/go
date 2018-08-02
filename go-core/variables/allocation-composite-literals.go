package main

// Dirinfo is ...
type Dirinfo struct {
	info string
}

// File is a struct ...
type File struct {
	fd      int
	name    string
	dirinfo *Dirinfo
	nepipe  int
}

// NewFile does ...
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}

/*
There's a lot of boiler plate in there. We can simplify it using a composite literal, which is an expression that creates a new instance each time it is evaluated.
*/

// NewFile2 does ...
func NewFile2(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, name, nil, 0} // composite literal
	return &f
}

/*
Note that, unlike in C, it's perfectly OK to return the address of a local variable; the storage associated with the variable survives after the function returns. In fact, taking the address of a composite literal allocates a fresh instance each time it is evaluated, so we can combine these last two lines.
*/

// NewFile3 does ...
func NewFile3(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name, nil, 0}
}

/*
The fields of a composite literal are laid out in order and must all be present. However, by labeling the elements explicitly as field:value pairs, the initializers can appear in any order, with the missing ones left as their respective zero values. Thus we could say
*/

// NewFile4 does ...
func NewFile4(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd: fd, name: name}
}

/*
As a limiting case, if a composite literal contains no fields at all, it creates a zero value for the type. The expressions new(File) and &File{} are equivalent.
*/

/*
Composite literals can also be created for arrays, slices, and maps, with the field labels being indices or map keys as appropriate. In these examples, the initializations work regardless of the values of Enone, Eio, and Einval, as long as they are distinct.
*/
/*
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
*/

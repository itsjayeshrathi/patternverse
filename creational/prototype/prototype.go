package prototype

import "fmt"

// protoype inteface
type Prototype interface {
	Clone() Prototype
	GetDetails() string
}

type File struct {
	Name string
	Type string
	Size int
}

func (f *File) Clone() Prototype {
	copy := *f
	return &copy
}

func (f *File) GetDetails() string {
	return fmt.Sprintf("File: %s, Type: %s, Size: %dKB", f.Name, f.Type, f.Size)
}

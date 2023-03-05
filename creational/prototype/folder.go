package prototype

import "fmt"

type Folder struct {
	name     string
	children []Inode
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name)
	for _, inode := range f.children {
		inode.print(indent + indent)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	tempChildren := []Inode{}

	for _, inode := range f.children {
		copy := inode.clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.children = tempChildren
	return cloneFolder
}

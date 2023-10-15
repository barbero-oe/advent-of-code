package day07

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	tree := buildTree(lines)
	total := 0
	for _, dir := range tree.directories {
		if size := dir.Size(); size <= 100000 {
			total += size
		}
	}
	return total
}

func PartTwo(lines []string) int {
	tree := buildTree(lines)
	diskSize := 70000000
	freeSpace := diskSize - tree.root.Size()
	requiredSpace := 30000000
	minimumSize := requiredSpace - freeSpace
	smallests := make([]*Directory, 0)
	for _, dir := range tree.directories {
		if dir.Size() >= minimumSize {
			smallests = append(smallests, dir)
		}
	}
	slices.SortFunc(smallests, func(a, b *Directory) int { return a.Size() - b.Size() })
	return smallests[0].Size()
}

func buildTree(lines []string) DirTree {
	tree := DirTree{}
	commands := make([]string, 0)
	for _, line := range lines {
		if strings.ContainsRune(line, '$') {
			tree.Execute(commands)
			commands = nil
		}
		commands = append(commands, line)
	}
	tree.Execute(commands)
	return tree
}

type Item interface {
	Size() int
}

type Directory struct {
	Parent      *Directory
	Name        string
	Files       []File
	Directories []*Directory
	cachedSize  int
}

func (d *Directory) Size() int {
	if d.cachedSize >= 0 {
		return d.cachedSize
	}
	total := 0
	for i := range d.Files {
		total += d.Files[i].Size()
	}
	for i := range d.Directories {
		total += d.Directories[i].Size()
	}
	d.cachedSize = total
	return total
}

type File struct {
	Name string
	size int
}

func (f *File) Size() int {
	return f.size
}

type DirTree struct {
	root        *Directory
	current     *Directory
	directories []*Directory
}

func (t *DirTree) Execute(instructions []string) {
	if len(instructions) == 0 {
		return
	}
	command := instructions[0]
	if command == "$ cd /" {
		if t.root == nil {
			t.root = &Directory{Name: "/", cachedSize: -1}
			t.directories = append(t.directories, t.root)
		}
		t.current = t.root
	} else if command == "$ ls" {
		if len(instructions) == 1 {
			return
		}
		for _, line := range instructions[1:] {
			parts := strings.Split(line, " ")
			if parts[0] == "dir" {
				directory := Directory{
					Parent:     t.current,
					Name:       parts[1],
					cachedSize: -1,
				}
				t.current.Directories = append(t.current.Directories, &directory)
				t.directories = append(t.directories, &directory)
			} else {
				size, _ := strconv.Atoi(parts[0])
				name := parts[1]
				t.current.Files = append(t.current.Files, File{Name: name, size: size})
			}
		}
	} else if dirName, found := strings.CutPrefix(command, "$ cd "); found {
		if dirName == ".." {
			t.current = t.current.Parent
		} else {
			for i := range t.current.Directories {
				directory := t.current.Directories[i]
				if directory.Name == dirName {
					t.current = directory
					return
				}
			}
			panic("Directory not found")
		}
	}
}

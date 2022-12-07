package challengeSeven

import (
	"fmt"
	"strconv"
	"strings"
)

type FileLocation struct {
	name        string
	parent      *FileLocation
	isDirectory bool
	size        int
	directories []*FileLocation
	files       []*FileLocation
}

func containsItemWith(list []*FileLocation, name string) bool {
	for _, fileLocation := range list {
		if fileLocation.name == name {
			return true
		}
	}
	return false
}

func (fl *FileLocation) goToDirectory(directory string) *FileLocation {
	if directory == ".." {
		return fl.parent
	}
	for _, dir := range fl.directories {
		if dir.name == directory {
			return dir
		}
	}
	return fl
}

func (fl *FileLocation) updateContent(content string) {
	if content[:3] == "dir" {
		if !containsItemWith(fl.directories, content[4:]) {
			fl.directories = append(fl.directories, NewDirectory(content[4:], fl))
		}
	} else {
		file := CreateFile(content)
		if !containsItemWith(fl.files, file.name) {
			fl.files = append(fl.files, file)
		}
	}
}

func (fl *FileLocation) printFilesAndDirectories(depth int) {
	fmt.Printf("%s (%d)\n", fl.name, fl.size)
	padding := ""
	for i := 0; i < depth; i++ {
		padding += " "
	}
	for _, dir := range fl.directories {
		fmt.Printf("%s- ", padding)
		dir.printFilesAndDirectories(depth + 1)
	}
	for _, file := range fl.files {
		fmt.Printf("%s* ", padding)
		fmt.Println(file.name)
	}
}

func CreateFile(content string) *FileLocation {
	fileAsString := strings.Split(content, " ")
	size, _ := strconv.Atoi(fileAsString[0])
	return NewFile(fileAsString[1], size)
}

func NewDirectory(name string, parent *FileLocation) *FileLocation {
	return &FileLocation{
		name:        name,
		parent:      parent,
		isDirectory: true,
		size:        0,
		directories: []*FileLocation{},
		files:       []*FileLocation{},
	}
}

func NewFile(name string, size int) *FileLocation {
	return &FileLocation{
		name:        name,
		isDirectory: false,
		size:        size,
		directories: nil,
		files:       nil,
	}
}

const COMMAND_OUTPUT = 0
const CHANGE_DIRECTORY = 1
const LIST_DIRECTORY = 2

func challengeSeven(input []string) int {
	root := NewDirectory("/", nil)
	currentDir := root

	for i := 1; i < len(input); i++ {
		command := parseCommand(input[i])
		if command == COMMAND_OUTPUT {
			currentDir.updateContent(input[i])
		}
		if command == CHANGE_DIRECTORY {
			currentDir = currentDir.goToDirectory(input[i][5:])
		}
	}

	calculateFileSizes(root)
	sizes := selectRemovableDirectories(root, 0)
	root.printFilesAndDirectories(0)
	return sizes
}

func selectRemovableDirectories(fileLocation *FileLocation, size int) int {
	if fileLocation.isDirectory {
		for _, dir := range fileLocation.directories {
			if dir.size <= 100000 {
				size += dir.size
			}
			size = selectRemovableDirectories(dir, size)
		}
	}
	return size
}

func calculateFileSizes(fileLocation *FileLocation) {
	if fileLocation.isDirectory {
		for _, dir := range fileLocation.directories {
			calculateFileSizes(dir)
			fileLocation.size += dir.size
		}
		for _, file := range fileLocation.files {
			fileLocation.size += file.size
		}
	}
}

func parseCommand(line string) int {
	if line[2:4] == "cd" {
		return CHANGE_DIRECTORY
	} else if line[2:4] == "ls" {
		return LIST_DIRECTORY
	}
	return COMMAND_OUTPUT
}

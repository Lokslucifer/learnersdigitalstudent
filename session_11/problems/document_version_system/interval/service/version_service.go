package service

import (
	"errors"
	"fmt"
)

type Version struct {
	Ver     int
	Content string
	Prev    *Version
	Next    *Version
}

type VersionManager struct {
	Current    *Version
	UndoStack  []int
	RedoStack  []int
	MaxVersion int
	Num        int
}

func NewVersionManager(maxver int) *VersionManager {
	return &VersionManager{Current: nil, UndoStack: make([]int, 0), RedoStack: make([]int, 0), MaxVersion: maxver, Num: 0}
}
func Top(ver *Version) *Version {

	if ver != nil {

		for ver.Next != nil {
			ver = ver.Next
		}
	}
	return ver
}
func getVersion(cur *Version, ver int) (*Version, error) {
	oldcur := cur
	for cur.Ver < ver {
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	for cur.Ver > ver {
		if cur.Prev == nil {
			break
		}
		cur = cur.Prev
	}
	if cur.Ver != ver {
		return oldcur, errors.New("Version not found")
	} else {
		return cur, nil
	}

}
func (vm *VersionManager) AddVersion(content string) {
	if vm.Num == vm.MaxVersion {
		fmt.Println("Max versions are added")
		return
	}
	oldver := vm.Current

	top := Top(oldver)

	newver := &Version{Content: content, Prev: top, Next: nil, Ver: vm.Num + 1}
	if top != nil {
		top.Next = newver
	}
	if oldver != nil {
		vm.UndoStack = append(vm.UndoStack, oldver.Ver)
		vm.RedoStack = make([]int, 0)
	}

	vm.Current = newver

	vm.Num += 1

}
func (vm *VersionManager) GetCurrentVersion() string {
	cur := vm.Current
	if cur == nil {
		fmt.Println("no version")
		return ""
	}
	fmt.Println("Version No:", cur.Ver)
	fmt.Println(cur.Content)
	return cur.Content
}
func (vm *VersionManager) Undo() {

	
	cur := vm.Current
	
	if cur != nil {
		undover := vm.UndoStack[len(vm.UndoStack)-1]

		newver, err := getVersion(cur, undover)
		if err != nil {
			fmt.Println(err)
			return
		}
		vm.RedoStack = append(vm.RedoStack, cur.Ver)
		vm.UndoStack=vm.UndoStack[:len(vm.UndoStack)-1]
		vm.Current = newver

	}
}

func (vm *VersionManager) Redo() {
	cur := vm.Current
	if cur != nil {
		redover := vm.RedoStack[len(vm.RedoStack)-1]
		newver, err := getVersion(cur, redover)
		if err != nil {
			fmt.Println(err)
			return
		}
		vm.UndoStack = append(vm.UndoStack, redover)
		vm.RedoStack=vm.RedoStack[:len(vm.RedoStack)-1]

		vm.Current = newver

	}
}

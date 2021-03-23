package parser

import (
	"fmt"
	"sort"
)

// CodePos represents a position information in the file.
type CodePos struct {
	Filename string // filename, if any
	Offset   int    // offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (byte count)
}

// IsValid returns true if the position is valid.
func (p CodePos) IsValid() bool {
	return p.Line > 0
}

// String returns a string in one of several forms:
//
//	file:line:column    valid position with file name
//	file:line           valid position with file name but no column (column == 0)
//	line:column         valid position without file name
//	line                valid position without file name and no column (column == 0)
//	file                invalid position with file name
//	-                   invalid position without file name
//
func (p CodePos) String() string {
	s := p.Filename
	if p.IsValid() {
		if s != "" {
			s += ":"
		}
		s += fmt.Sprintf("%d", p.Line)
		if p.Column != 0 {
			s += fmt.Sprintf(":%d", p.Column)
		}
	}
	if s == "" {
		s = "-"
	}
	return s
}

// CodeSet represents a set of source files.
type CodeSet struct {
	Base     int           // base offset for the next file
	Files    []*Code // list of files in the order added to the set
	LastFile *Code   // cache of last file looked up
}

// NewFileSet creates a new file set.
func NewFileSet() *CodeSet {
	return &CodeSet{
		Base: 1, // 0 == NoPos
	}
}

// AddFile adds a new file in the file set.
func (s *CodeSet) AddFile(filename string, base, size int) *Code {
	if base < 0 {
		base = s.Base
	}
	if base < s.Base || size < 0 {
		panic("illegal base or size")
	}
	f := &Code{
		set:   s,
		Name:  filename,
		Base:  base,
		Size:  size,
		Lines: []int{0},
	}
	base += size + 1 // +1 because EOF also has a position
	if base < 0 {
		panic("offset overflow (> 2G of source code in file set)")
	}

	// add the file to the file set
	s.Base = base
	s.Files = append(s.Files, f)
	s.LastFile = f
	return f
}

// File returns the file that contains the position p. If no such file is
// found (for instance for p == NoPos), the result is nil.
func (s *CodeSet) File(p Pos) (f *Code) {
	if p != NoPos {
		f = s.file(p)
	}
	return
}

// Position converts a SourcePos p in the fileset into a CodePos value.
func (s *CodeSet) Position(p Pos) (pos CodePos) {
	if p != NoPos {
		if f := s.file(p); f != nil {
			return f.position(p)
		}
	}
	return
}

func (s *CodeSet) file(p Pos) *Code {
	// common case: p is in last file
	f := s.LastFile
	if f != nil && f.Base <= int(p) && int(p) <= f.Base+f.Size {
		return f
	}

	// p is not in last file - search all files
	if i := searchFiles(s.Files, int(p)); i >= 0 {
		f := s.Files[i]

		// f.base <= int(p) by definition of searchFiles
		if int(p) <= f.Base+f.Size {
			s.LastFile = f // race is ok - s.last is only a cache
			return f
		}
	}
	return nil
}

func searchFiles(a []*Code, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i].Base > x }) - 1
}

// Code represents a source file.
type Code struct {
	// Code set for the file
	set *CodeSet
	// Code name as provided to AddFile
	Name string
	// SourcePos value range for this file is [base...base+size]
	Base int
	// Code size as provided to AddFile
	Size int
	// Lines contains the offset of the first character for each line
	// (the first entry is always 0)
	Lines []int
}

// Set returns CodeSet.
func (f *Code) Set() *CodeSet {
	return f.set
}

// LineCount returns the current number of lines.
func (f *Code) LineCount() int {
	return len(f.Lines)
}

// AddLine adds a new line.
func (f *Code) AddLine(offset int) {
	i := len(f.Lines)
	if (i == 0 || f.Lines[i-1] < offset) && offset < f.Size {
		f.Lines = append(f.Lines, offset)
	}
}

// LineStart returns the position of the first character in the line.
func (f *Code) LineStart(line int) Pos {
	if line < 1 {
		panic("illegal line number (line numbering starts at 1)")
	}
	if line > len(f.Lines) {
		panic("illegal line number")
	}
	return Pos(f.Base + f.Lines[line-1])
}

// FileSetPos returns the position in the file set.
func (f *Code) FileSetPos(offset int) Pos {
	if offset > f.Size {
		panic("illegal file offset")
	}
	return Pos(f.Base + offset)
}

// Offset translates the file set position into the file offset.
func (f *Code) Offset(p Pos) int {
	if int(p) < f.Base || int(p) > f.Base+f.Size {
		panic("illegal SourcePos value")
	}
	return int(p) - f.Base
}

// Position translates the file set position into the file position.
func (f *Code) Position(p Pos) (pos CodePos) {
	if p != NoPos {
		if int(p) < f.Base || int(p) > f.Base+f.Size {
			panic("illegal SourcePos value")
		}
		pos = f.position(p)
	}
	return
}

func (f *Code) position(p Pos) (pos CodePos) {
	offset := int(p) - f.Base
	pos.Offset = offset
	pos.Filename, pos.Line, pos.Column = f.unpack(offset)
	return
}

func (f *Code) unpack(offset int) (filename string, line, column int) {
	filename = f.Name
	if i := searchInts(f.Lines, offset); i >= 0 {
		line, column = i+1, offset-f.Lines[i]+1
	}
	return
}

func searchInts(a []int, x int) int {
	// This function body is a manually inlined version of:
	//   return sort.Search(len(a), func(i int) bool { return a[i] > x }) - 1
	i, j := 0, len(a)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if a[h] <= x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i - 1
}
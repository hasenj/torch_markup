package node

import "bytes"

//go:generate stringer -type=Category
type Category uint

// Categories of nodes
const (
	CategoryBlock Category = iota
	CategoryInline
)

//go:generate stringer -type=Type
type Type uint

// Types of nodes
const (
	// blocks
	TypeLine Type = iota
	TypeWalled
	TypeHanging
	TypeFenced

	// inlines
	TypeText
	TypeUniform
	TypeEscaped
	TypeForward
)

// TypeCategory is used by parser to determine node category based on type.
func TypeCategory(typ Type) Category {
	if typ > 3 {
		return CategoryInline
	}
	return CategoryBlock
}

// Node represents an element.
type Node interface {
	Node() string
}

type Block interface {
	Node
	Block()
}

type Inline interface {
	Node
	Inline()
}

type ContentInlineChildren interface {
	Content
	InlineChildren
}

type Content interface {
	Content() []byte
}

type HeadBody interface {
	Head() []byte
	Body() []byte
}

type BlockChildren interface {
	BlockChildren() []Block
}

type InlineChildren interface {
	InlineChildren() []Inline
}

// BlocksToNodes converts blocks to nodes.
func BlocksToNodes(blocks []Block) []Node {
	nodes := make([]Node, len(blocks))
	for i, b := range blocks {
		nodes[i] = Node(b)
	}
	return nodes
}

// InlinesToNodes converts inlines to nodes.
func InlinesToNodes(inlines []Inline) []Node {
	nodes := make([]Node, len(inlines))
	for i, v := range inlines {
		nodes[i] = Node(v)
	}
	return nodes
}

type Line struct {
	Name     string
	Children []Inline
}

func (l Line) Node() string {
	return l.Name
}

func (l Line) Block() {}

func (l *Line) InlineChildren() []Inline {
	return l.Children
}

type Walled struct {
	Name     string
	Children []Block
}

func (w Walled) Node() string {
	return w.Name
}

func (w Walled) Block() {}

func (w *Walled) BlockChildren() []Block {
	return w.Children
}

type Hanging struct {
	Name     string
	Children []Block
}

func (h Hanging) Node() string {
	return h.Name
}

func (h Hanging) Block() {}

func (h *Hanging) BlockChildren() []Block {
	return h.Children
}

type Fenced struct {
	Name  string
	Lines [][]byte
}

func (f Fenced) Node() string {
	return f.Name
}

func (f Fenced) Block() {}

func (f Fenced) Head() []byte {
	if len(f.Lines) == 0 {
		return nil
	}
	return f.Lines[0]
}

func (f Fenced) Body() []byte {
	if len(f.Lines) == 0 {
		return nil
	}
	return bytes.Join(f.Lines[1:], []byte("\n"))
}

type Uniform struct {
	Name     string
	Children []Inline
}

func (u Uniform) Node() string {
	return u.Name
}

func (u Uniform) Inline() {}

func (u *Uniform) InlineChildren() []Inline {
	return u.Children
}

type Escaped struct {
	Name     string
	Content0 []byte
}

func (e Escaped) Node() string {
	return e.Name
}

func (e Escaped) Inline() {}

func (e *Escaped) Content() []byte {
	return e.Content0
}

type Forward struct {
	Name      string
	Content0  []byte
	Children0 []Inline
}

func (f Forward) Node() string {
	return f.Name
}

func (f Forward) Inline() {}

func (f *Forward) Content() []byte {
	return f.Content0
}

func (f *Forward) InlineChildren() []Inline {
	return f.Children0
}

// Text represents text—an atomic, inline node.
type Text []byte

// Node returns the node's name.
func (t Text) Node() string {
	return "Text"
}

func (t Text) Inline() {}

// Content returns the text.
func (t Text) Content() []byte {
	return t
}

// LineComment represents text—an atomic, inline node.
type LineComment []byte

// Node returns the node's name.
func (c LineComment) Node() string {
	return "LineComment"
}

func (c LineComment) Inline() {}

// Content returns the LineComment's text.
func (c LineComment) Content() []byte {
	return c
}

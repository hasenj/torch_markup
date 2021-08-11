package parser_test

import (
	"github.com/touchmarine/to/node"
	"github.com/touchmarine/to/parser"
	"github.com/touchmarine/to/stringifier"
	"testing"
	"unicode"
)

func TestTextBlock(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			" ",
			[]node.Node{},
		},
		{
			" \n ",
			[]node.Node{},
		},
		{
			"a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			"a\n",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			"a\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
			},
		},
		{
			"a\n\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"a\n\n\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"a\n\n\n\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"a**\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Strong", []node.Inline{
						node.Text(" b"),
					}},
				}},
			},
		},
		{
			"a\n>b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},

		// nested
		{
			">a\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			">a\n\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			">a\n\n\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},

		{
			">a\nb",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			">>a\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},

		{
			">a\n>\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},

		{
			"*a\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"*a\n*b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			"*a\n\n*b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			"*a\n\n\n*b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},

		{
			"*a\nb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"**a\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},

		// spacing
		{
			"a \nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
			},
		},
		{
			"a  \nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
			},
		},
		{
			"a\n b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
			},
		},
		{
			"a\n  b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
			},
		},
		{
			"*a\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"*a\n  b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// block escape
		{
			"a\n\\__",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a "),
					&node.Uniform{"Emphasis", nil},
				}},
			},
		},

		// inline escape
		{
			"a\n\\\\__",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a __"),
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestLine(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			" ",
			[]node.Node{},
		},
		{
			" \n ",
			[]node.Node{},
		},
		{
			" \n \n",
			[]node.Node{},
		},
		{
			"a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			"a_*",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a_*")}},
			},
		},

		{
			"\na",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			"\n\na",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			" \na",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			"\t\na",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			"a\n",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			"a\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
			},
		},
		{
			"a\n\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"a\n\n\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestWalled(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			">",
			[]node.Node{&node.Walled{"Blockquote", nil}},
		},
		{
			">\na",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			">>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
			},
		},
		{
			">\n>",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
			},
		},
		{
			">a\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			">\n>>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">a\n>>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					}},
				}},
			},
		},
		{
			">>\n>>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">>a\n>>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
					}},
				}},
			},
		},
		{
			">>\n>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">>a\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			">\n\n>",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
				&node.Walled{"Blockquote", nil},
			},
		},
		{
			">a\n\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			">a\n \n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},

		// spacing
		{
			" >",
			[]node.Node{&node.Walled{"Blockquote", nil}},
		},

		// regression
		{
			"> >",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">\t>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			"> > >",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.Walled{"Blockquote", nil},
					}},
				}},
			},
		},
		{
			">\n >",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
			},
		},
		{
			" >\n>",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
			},
		},
		{
			">a\n >b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a b"),
					}},
				}},
			},
		},

		{
			">\na",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			">\n>\na",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			">\n>\n>\na",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			">\n>\n>\n>\na",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestVerbatimWalled(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"/",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", nil},
			},
		},
		{
			"/a",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
			},
		},
		{
			"/a\n",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
			},
		},
		{
			"/a\n/b",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{
					[]byte("a"),
					[]byte("b"),
				}},
			},
		},
		{
			"/a\n/\n/b",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{
					[]byte("a"),
					nil,
					[]byte("b"),
				}},
			},
		},

		// no nested content allowed
		{
			"/>a",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte(">a")}},
			},
		},
		{
			"/**a",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("**a")}},
			},
		},
		{
			`/\**a`,
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte(`\**a`)}},
			},
		},
		{
			`/\\**a`,
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte(`\\**a`)}},
			},
		},

		// spacing
		{
			"/\n/",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{nil, nil}},
			},
		},
		{
			"/ \n/ ",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{
					[]byte(" "),
					[]byte(" "),
				}},
			},
		},
		{
			"/ a",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte(" a")}},
			},
		},
		{
			"/a\n/ b",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{
					[]byte("a"),
					[]byte(" b"),
				}},
			},
		},
		{
			">/ a\n>/ b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						[]byte(" a"),
						[]byte(" b"),
					}},
				}},
			},
		},
		{
			"*/ a\n / b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						[]byte(" a"),
						[]byte(" b"),
					}},
				}},
			},
		},
		{
			" / a",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte(" a")}},
			},
		},
		{
			" / \n / ",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{
					[]byte(" "),
					[]byte(" "),
				}},
			},
		},
		{
			"*/\n / b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						nil,
						[]byte(" b"),
					}},
				}},
			},
		},

		// continuation (stop)
		{
			"/a\nb",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"/a\n\n/b",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("b")}},
			},
		},
		{
			"/a\n>b",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			"/a\n*b",
			[]node.Node{
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			">/a\n/b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
				}},
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("b")}},
			},
		},
		{
			"*/a\n/b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
				}},
				&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("b")}},
			},
		},

		// nested
		{
			">/a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{[]byte("a")}},
				}},
			},
		},
		{
			">/a\n>/b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						[]byte("a"),
						[]byte("b"),
					}},
				}},
			},
		},
		{
			"*/a\n /b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						[]byte("a"),
						[]byte("b"),
					}},
				}},
			},
		},

		{
			">/a\n>/\n>/b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						[]byte("a"),
						nil,
						[]byte("b"),
					}},
				}},
			},
		},
		{
			"*/a\n /\n /b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						[]byte("a"),
						nil,
						[]byte("b"),
					}},
				}},
			},
		},
		{
			"*\n /a\n /\n /b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.VerbatimWalled{"BlockComment", [][]byte{
						[]byte("a"),
						nil,
						[]byte("b"),
					}},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestHanging(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"*",
			[]node.Node{&node.Hanging{"DescriptionList", nil}},
		},
		{
			"**",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},
		{
			"*a",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
			},
		},
		{
			"*\n*",
			[]node.Node{
				&node.Hanging{"DescriptionList", nil},
				&node.Hanging{"DescriptionList", nil},
			},
		},
		{
			"*\n\n*",
			[]node.Node{
				&node.Hanging{"DescriptionList", nil},
				&node.Hanging{"DescriptionList", nil},
			},
		},
		{
			"*\n\n\n*",
			[]node.Node{
				&node.Hanging{"DescriptionList", nil},
				&node.Hanging{"DescriptionList", nil},
			},
		},
		{
			"*\n\t\n*",
			[]node.Node{
				&node.Hanging{"DescriptionList", nil},
				&node.Hanging{"DescriptionList", nil},
			},
		},
		{
			"*a\n\n*b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			"*a\nb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"*a\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"*a\n  b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// spacing
		{
			" *a\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			" *a\n  b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// nested
		{
			"*\n *",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},
		{
			"*a\n *b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					}},
				}},
			},
		},
		{
			"**a\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			"**a\n  b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
					}},
				}},
			},
		},
		{
			"**a\n   b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
					}},
				}},
			},
		},

		{
			">*",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},
		{
			">*\n*",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
				&node.Hanging{"DescriptionList", nil},
			},
		},
		{
			">*\n>*",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", nil},
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},
		{
			">*\n> *",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.Hanging{"DescriptionList", nil},
					}},
				}},
			},
		},
		{
			">*\n> a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
				}},
			},
		},
		{
			"> *\n>  a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
				}},
			},
		},

		// nested+spacing
		{
			" *\n *",
			[]node.Node{
				&node.Hanging{"DescriptionList", nil},
				&node.Hanging{"DescriptionList", nil},
			},
		},
		{
			" *\n  *",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},

		// tab (equals 8 spaces in this regard)
		{
			"*a\n\tb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"\t*a\n\tb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"\t*a\n\t b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"\t*a\n \tb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"\t*a\n  \tb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		{
			"\t\t*a\n                b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"\t\t*a\n                 b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"                *a\n\t\tb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"               *a\n\t\tb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// nested+blank lines
		{
			"*\n *",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},
		//*
		//
		// *
		{
			"*\n\n *",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},
		{
			"*\n \t\n *",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			},
		},

		//**
		//
		//*
		{
			"**\n\n*",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
				&node.Hanging{"DescriptionList", nil},
			},
		},
		{
			"a\n\n**\n\n*",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
				}},
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
				&node.Hanging{"DescriptionList", nil},
			},
		},

		{
			"**\n\na",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
				}},
			},
		},
		{
			"**\n\n a",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", nil},
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"),
					}},
				}},
			},
		},
		{
			"**\n\n  a",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("a"),
						}},
					}},
				}},
			},
		},
		{
			"**\n\n  a\nb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("a"),
						}},
					}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("b"),
				}},
			},
		},

		// regression
		{
			"*\n >b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("b"),
						}},
					}},
				}},
			},
		},
		{
			"*>a\n >b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("a b"),
						}},
					}},
				}},
			},
		},

		//*
		//	*a
		//	 *b
		//	c
		{
			"*\n\t*a\n\t *b\n\tc",
			[]node.Node{&node.Hanging{"DescriptionList", []node.Block{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"),
					}},
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("b"),
						}},
					}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("c"),
				}},
			}}},
		},

		{
			"*\n  >*a",
			[]node.Node{&node.Hanging{"DescriptionList", []node.Block{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("a"),
						}},
					}},
				}},
			}}},
		},
		{
			"*\n\t>*a",
			[]node.Node{&node.Hanging{"DescriptionList", []node.Block{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("a"),
						}},
					}},
				}},
			}}},
		},
		//*
		//	>	*
		//	>
		{
			"*\n\t>\t*\n\t>",
			[]node.Node{&node.Hanging{"DescriptionList", []node.Block{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", nil},
				}},
			}}},
		},
		//  >*a
		// > *b
		{
			"  >*a\n > *b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
						&node.Hanging{"DescriptionList", []node.Block{
							&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
						}},
					}},
				}},
			},
		},
		{
			"  > *a\n >  *b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
						&node.Hanging{"DescriptionList", []node.Block{
							&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
						}},
					}},
				}},
			},
		},

		//>*
		//>
		//> *
		{
			">*\n>\n> *",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.Hanging{"DescriptionList", nil},
					}},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestRankedHanging(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"=",
			[]node.Node{&node.Hanging{"Title", nil}},
		},
		{
			"==",
			[]node.Node{&node.RankedHanging{"Heading", 2, nil}},
		},
		{
			"= =",
			[]node.Node{
				&node.Hanging{"Title", []node.Block{
					&node.Hanging{"Title", nil},
				}},
			},
		},
		{
			"== ==",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.RankedHanging{"Heading", 2, nil},
				}},
			},
		},
		{
			"==a",
			[]node.Node{&node.RankedHanging{"Heading", 2, []node.Block{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
				}},
			}}},
		},
		{
			"==\n==",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, nil},
				&node.RankedHanging{"Heading", 2, nil},
			},
		},

		{
			"==a\nb",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"==a\n b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"==a\n  b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"==a\n   b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// spacing
		{
			" ==a\n  b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			" ==a\n   b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// nested
		{
			"==\n  ==",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.RankedHanging{"Heading", 2, nil},
				}},
			},
		},
		{
			"==a\n  ==b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					&node.RankedHanging{"Heading", 2, []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					}},
				}},
			},
		},
		{
			"== ==a\n  b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.RankedHanging{"Heading", 2, []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			"== ==a\n     b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.RankedHanging{"Heading", 2, []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
					}},
				}},
			},
		},
		{
			"== ==a\n      b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.RankedHanging{"Heading", 2, []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
					}},
				}},
			},
		},

		{
			">==",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.RankedHanging{"Heading", 2, nil},
				}},
			},
		},
		{
			">==\n==",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.RankedHanging{"Heading", 2, nil},
				}},
				&node.RankedHanging{"Heading", 2, nil},
			},
		},
		{
			">==\n>==",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.RankedHanging{"Heading", 2, nil},
					&node.RankedHanging{"Heading", 2, nil},
				}},
			},
		},
		{
			">==\n>  ==",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.RankedHanging{"Heading", 2, []node.Block{
						&node.RankedHanging{"Heading", 2, nil},
					}},
				}},
			},
		},
		{
			">==\n>  a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.RankedHanging{"Heading", 2, []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
				}},
			},
		},
		{
			"> ==\n>   a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.RankedHanging{"Heading", 2, []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
				}},
			},
		},

		// nested+spacing
		{
			" ==\n ==",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, nil},
				&node.RankedHanging{"Heading", 2, nil},
			},
		},
		{
			" ==\n   ==",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.RankedHanging{"Heading", 2, nil},
				}},
			},
		},

		// tab (equals 8 spaces in this regard)
		{
			"==a\n\tb",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"\t==a\n\tb",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"\t==a\n\t  b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"\t==a\n  \tb",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"\t==a\n   \tb",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		{
			"\t\t==a\n                 b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"\t\t==a\n                  b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},
		{
			"                 ==a\n\t\tb",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"              ==a\n\t\tb",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// regression
		{
			"==\n  >b",
			[]node.Node{
				&node.RankedHanging{"Heading", 2, []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("b"),
						}},
					}},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestVerbatimLine(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			".image",
			[]node.Node{&node.VerbatimLine{"Image", nil}},
		},
		{
			".imagea",
			[]node.Node{&node.VerbatimLine{"Image", []byte("a")}},
		},
		{
			".image a",
			[]node.Node{&node.VerbatimLine{"Image", []byte(" a")}},
		},
		{
			".imagea ",
			[]node.Node{&node.VerbatimLine{"Image", []byte("a ")}},
		},
		{
			".image*",
			[]node.Node{&node.VerbatimLine{"Image", []byte("*")}},
		},
		{
			`.image\**`,
			[]node.Node{&node.VerbatimLine{"Image", []byte(`\**`)}},
		},

		{
			".image\n      a",
			[]node.Node{
				&node.VerbatimLine{"Image", nil},
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
				}},
			},
		},
		{
			">.image",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.VerbatimLine{"Image", nil},
				}},
			},
		},
		{
			">\n.image",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
				&node.VerbatimLine{"Image", nil},
			},
		},
		{
			".image\n>",
			[]node.Node{
				&node.VerbatimLine{"Image", nil},
				&node.Walled{"Blockquote", nil},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestHangingMulti(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"1",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("1"),
				}},
			},
		},
		{
			"1.",
			[]node.Node{&node.Hanging{"NumberedListItemDot", nil}},
		},
		{
			"1.1.",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.Hanging{"NumberedListItemDot", nil},
				}},
			},
		},
		{
			"1.a",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
			},
		},
		{
			"1.\n1.",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", nil},
				&node.Hanging{"NumberedListItemDot", nil},
			},
		},
		{
			"1.a\nb",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"1.a\n b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"1.a\n  b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// spacing
		{
			" 1.a\n b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			" 1.a\n  b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			" 1.a\n   b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a b")}},
				}},
			},
		},

		// nested
		{
			"1.\n  1.",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.Hanging{"NumberedListItemDot", nil},
				}},
			},
		},
		{
			"1.a\n  1.b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					&node.Hanging{"NumberedListItemDot", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					}},
				}},
			},
		},

		{
			">1.",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"NumberedListItemDot", nil},
				}},
			},
		},
		{
			">1.\n1.",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"NumberedListItemDot", nil},
				}},
				&node.Hanging{"NumberedListItemDot", nil},
			},
		},
		{
			">1.\n>1.",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"NumberedListItemDot", nil},
					&node.Hanging{"NumberedListItemDot", nil},
				}},
			},
		},
		{
			">1.\n>  1.",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"NumberedListItemDot", []node.Block{
						&node.Hanging{"NumberedListItemDot", nil},
					}},
				}},
			},
		},
		{
			">1.\n>  a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"NumberedListItemDot", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
				}},
			},
		},
		{
			"> 1.\n>   a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"NumberedListItemDot", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
					}},
				}},
			},
		},

		{
			"1.-\n\n1.",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.Hanging{"ListItemDot", nil},
				}},
				&node.Hanging{"NumberedListItemDot", nil},
			},
		},

		// regression
		{
			"1.\n  >b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("b"),
						}},
					}},
				}},
			},
		},
		{
			"1.>a\n  >b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("a b"),
						}},
					}},
				}},
			},
		},
		{
			"1. >a\n   >b",
			[]node.Node{
				&node.Hanging{"NumberedListItemDot", []node.Block{
					&node.Walled{"Blockquote", []node.Block{
						&node.BasicBlock{"TextBlock", []node.Inline{
							node.Text("a b"),
						}},
					}},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestFenced(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"``",
			[]node.Node{&node.Fenced{"CodeBlock", nil, nil}},
		},
		{
			"``a",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{[]byte("a")}, nil},
			},
		},
		{
			"``a``",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{[]byte("a``")}, nil},
			},
		},
		{
			"``\na",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
			},
		},
		{
			"``\n a",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{nil, []byte(" a")}, nil},
			},
		},
		{
			"``\n\na",
			[]node.Node{
				&node.Fenced{
					"CodeBlock",
					[][]byte{
						nil,
						nil,
						[]byte("a")},
					nil,
				},
			},
		},
		{
			"````",
			[]node.Node{
				&node.Fenced{"CodeBlock", nil, nil},
			},
		},
		{
			"``\n``",
			[]node.Node{
				&node.Fenced{"CodeBlock", nil, nil},
			},
		},
		{
			"```\n```",
			[]node.Node{
				&node.Fenced{"CodeBlock", nil, nil},
			},
		},
		{
			"```\n``\n```",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{nil, []byte("``")}, nil},
			},
		},
		{
			"```\n`````",
			[]node.Node{
				&node.Fenced{"CodeBlock", nil, []byte("``")},
			},
		},
		{
			"``\n``a",
			[]node.Node{
				&node.Fenced{"CodeBlock", nil, []byte("a")},
			},
		},

		// nesting
		{
			"``\n>",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{nil, []byte(">")}, nil},
			},
		},
		{
			">``\na",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", nil, nil},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			">``\n>a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
		{
			">``\n>a\n>``",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
		{
			">``\n>a\n>``b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, []byte("b")},
				}},
			},
		},
		{
			">``\n>a\n>``\nb",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},

		// nesting+spacing
		{
			"> ``\n>a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
		{
			"> ``\n> a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
		{
			"> ``\n>  a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte(" a")}, nil},
				}},
			},
		},
		{
			">  ``\n> a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
		{
			">  ``\n>  a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},

		// tab
		{
			">\t``\n>a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
		{
			">\t``\n>        a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
		{
			">\t``\n>         a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte(" a")}, nil},
				}},
			},
		},
		{
			">\t``\n>            a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("    a")}, nil},
				}},
			},
		},
		{
			"> ``\n>\ta",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("       a")}, nil},
				}},
			},
		},
		{
			"> \t``\n>\t a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Fenced{"CodeBlock", [][]byte{nil, []byte("a")}, nil},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestSpacing(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"\n>",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
			},
		},

		// space
		{
			" >",
			[]node.Node{&node.Walled{"Blockquote", nil}},
		},
		{
			"> >",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">  >",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			"> ",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
			},
		},
		{
			"> a",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
			},
		},

		// tab
		{
			"\t>",
			[]node.Node{&node.Walled{"Blockquote", nil}},
		},
		{
			">\t>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">\t\t>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">\t",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
			},
		},
		{
			">\ta",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
			},
		},

		// space+tab
		{
			" \t>",
			[]node.Node{&node.Walled{"Blockquote", nil}},
		},
		{
			"> \t>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			">  \t>",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Walled{"Blockquote", nil},
				}},
			},
		},
		{
			"> \t",
			[]node.Node{
				&node.Walled{"Blockquote", nil},
			},
		},
		{
			"> \ta",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestUniform(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"__",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", nil},
				}},
			},
		},
		{
			"____",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", nil},
				}},
			},
		},
		{
			"__a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						node.Text("a"),
					}},
				}},
			},
		},
		{
			"__a__",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						node.Text("a"),
					}},
				}},
			},
		},
		{
			"__a__b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						node.Text("a"),
					}},
					node.Text("b"),
				}},
			},
		},

		// left-right delimiter
		{
			"((",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Group", nil},
				}},
			},
		},
		{
			"(())",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Group", nil},
				}},
			},
		},
		{
			"((a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Group", []node.Inline{
						node.Text("a"),
					}},
				}},
			},
		},
		{
			"((a))",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Group", []node.Inline{
						node.Text("a"),
					}},
				}},
			},
		},
		{
			"((a))b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Group", []node.Inline{
						node.Text("a"),
					}},
					node.Text("b"),
				}},
			},
		},

		// across lines
		{
			"a__\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", []node.Inline{
						node.Text(" b"),
					}},
				}},
			},
		},
		{
			"a__\n>b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", nil},
				}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			">a__\nb",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"), &node.Uniform{"Emphasis", nil},
					}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"a__\n%b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", nil},
				}},
				&node.Hat{
					[][]byte{[]byte("b")},
					nil,
				},
			},
		},

		// across line spacing
		{
			"a__ \nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", []node.Inline{node.Text(" b")}},
				}},
			},
		},
		{
			"a__  \nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", []node.Inline{node.Text(" b")}},
				}},
			},
		},
		{
			"a__\n b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", []node.Inline{node.Text(" b")}},
				}},
			},
		},
		{
			"a__\n  b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", []node.Inline{node.Text(" b")}},
				}},
			},
		},
		{
			"*a__\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"),
						&node.Uniform{"Emphasis", []node.Inline{node.Text(" b")}},
					}},
				}},
			},
		},
		{
			"*a__\n  b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"),
						&node.Uniform{"Emphasis", []node.Inline{node.Text(" b")}},
					}},
				}},
			},
		},

		// nested
		{
			"__**",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Uniform{"Strong", nil},
					}},
				}},
			},
		},
		{
			"__**a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Uniform{"Strong", []node.Inline{
							node.Text("a"),
						}},
					}},
				}},
			},
		},
		{
			"__**a**b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Uniform{"Strong", []node.Inline{
							node.Text("a"),
						}},
						node.Text("b"),
					}},
				}},
			},
		},
		{
			"__**a__b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Uniform{"Strong", []node.Inline{
							node.Text("a"),
						}},
					}},
					node.Text("b"),
				}},
			},
		},
		{
			"__**a**b__c",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Uniform{"Strong", []node.Inline{
							node.Text("a"),
						}},
						node.Text("b"),
					}},
					node.Text("c"),
				}},
			},
		},

		// nested across lines
		{
			"__**\na",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Uniform{"Strong", []node.Inline{
							node.Text(" a"),
						}},
					}},
				}},
			},
		},
		{
			"__**a\nb**__c",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Uniform{"Strong", []node.Inline{
							node.Text("a b"),
						}},
					}},
					node.Text("c"),
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestEscaped(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"a``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", nil},
				}},
			},
		},
		{
			"a```",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte("`")},
				}},
			},
		},
		{
			"a````",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", nil},
				}},
			},
		},
		{
			"a`````",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", nil},
					node.Text("`"),
				}},
			},
		},
		{
			"a``\\```",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte("```")},
				}},
			},
		},
		{
			"a``\\`",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte("`")},
				}},
			},
		},
		{
			"a``\\`\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte("`")},
				}},
			},
		},
		{
			"a``\\``\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte("``")},
				}},
			},
		},
		{
			"a``\\```\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte("```")},
				}},
			},
		},

		// left-right delim
		{
			"a[[",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Link", nil},
				}},
			},
		},
		{
			"a[[[",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Link", []byte("[")},
				}},
			},
		},
		{
			"a[[]]",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Link", nil},
				}},
			},
		},

		{
			"a\\````",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a`"),
					&node.Escaped{"Code", []byte("`")},
				}},
			},
		},
		{
			"a\\[[]]",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a[[]]"),
				}},
			},
		},

		// across lines
		{
			"a``\nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte(" b")},
				}},
			},
		},
		{
			"a``\n>b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", nil},
				}},
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				}},
			},
		},
		{
			">a``\nb",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"), &node.Escaped{"Code", nil},
					}},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"a``\n%b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", nil},
				}},
				&node.Hat{
					[][]byte{[]byte("b")},
					nil,
				},
			},
		},

		// across line spacing
		{
			"a`` \nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte(" b")},
				}},
			},
		},
		{
			"a``  \nb",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte(" b")},
				}},
			},
		},
		{
			"a``\n b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte(" b")},
				}},
			},
		},
		{
			"a``\n  b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte(" b")},
				}},
			},
		},
		{
			"*a``\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"),
						&node.Escaped{"Code", []byte(" b")},
					}},
				}},
			},
		},
		{
			"*a``\n  b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.BasicBlock{"TextBlock", []node.Inline{
						node.Text("a"),
						&node.Escaped{"Code", []byte(" b")},
					}},
				}},
			},
		},

		// nested elements are not allowed
		{
			"a``__``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Escaped{"Code", []byte("__")},
				}},
			},
		},
		{
			"a__``__b``c",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
					&node.Uniform{"Emphasis", []node.Inline{
						&node.Escaped{"Code", []byte("__b")},
						node.Text("c"),
					}},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestHat(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			"%a",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a")},
					nil,
				},
			},
		},

		{
			"%a\nb",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a")},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				},
			},
		},
		{
			"%a\n\nb",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a")},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				},
			},
		},
		{
			"%a\nb\n%c",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a")},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				},
				&node.Hat{
					[][]byte{[]byte("c")},
					nil,
				},
			},
		},
		{
			"%a\n%b\nc",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a"), []byte("b")},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("c")}},
				},
			},
		},
		{
			"%a\n%\nc",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a"), []byte("")},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("c")}},
				},
			},
		},

		{
			"%a\n>",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a")},
					&node.Walled{"Blockquote", nil},
				},
			},
		},
		{
			"%a\n*",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a")},
					&node.Hanging{"DescriptionList", nil},
				},
			},
		},

		{
			">%a\n>b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hat{
						[][]byte{[]byte("a")},
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					},
				}},
			},
		},
		{
			"*%a\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hat{
						[][]byte{[]byte("a")},
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					},
				}},
			},
		},

		{
			">%a\nb",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hat{
						[][]byte{[]byte("a")},
						nil,
					},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		{
			"*%a\nb",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hat{
						[][]byte{[]byte("a")},
						nil,
					},
				}},
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
			},
		},
		//*%a
		//
		// b
		{
			"*%a\n\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hat{
						[][]byte{[]byte("a")},
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					},
				}},
			},
		},
		//>*%a
		//>
		//> b
		{
			">*%a\n>\n> b",
			[]node.Node{
				&node.Walled{"Blockquote", []node.Block{
					&node.Hanging{"DescriptionList", []node.Block{
						&node.Hat{
							[][]byte{[]byte("a")},
							&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
						},
					}},
				}},
			},
		},

		{
			"%a\n\nb",
			[]node.Node{
				&node.Hat{
					[][]byte{[]byte("a")},
					&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
				},
			},
		},

		//*
		// %a
		//
		// b
		{
			"*\n %a\n\n b",
			[]node.Node{
				&node.Hanging{"DescriptionList", []node.Block{
					&node.Hat{
						[][]byte{[]byte("a")},
						&node.BasicBlock{"TextBlock", []node.Inline{node.Text("b")}},
					},
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestBlockEscape(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			`\`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", nil},
			},
		},
		{
			`\\`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(`\`)}},
			},
		},
		{
			`\a`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a")}},
			},
		},
		{
			`\>`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(">")}},
			},
		},
		{
			`\\>`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(`\>`)}},
			},
		},
		{
			"\\\n\\",
			[]node.Node{
				&node.BasicBlock{"TextBlock", nil},
			},
		},
		{
			"\\a\n\\b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(`a b`)}},
			},
		},
		{
			"\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{&node.Escaped{"Code", nil}}},
			},
		},
		{
			"\\\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("``"),
				}},
			},
		},

		{
			`\**`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{&node.Uniform{"Strong", nil}}},
			},
		},

		{
			"``\\**\n\\**",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{[]byte("\\**"), []byte("\\**")}, nil},
			},
		},

		// verbatim elements
		{
			`\.image`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(".image")}},
			},
		},
		{
			`.image\.image`,
			[]node.Node{
				&node.VerbatimLine{"Image", []byte(`\.image`)},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestInlineEscape(t *testing.T) {
	cases := []struct {
		in  string
		out []node.Node
	}{
		{
			`\`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", nil},
			},
		},
		{
			`\\`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(`\`)}},
			},
		},
		{
			`\\\`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(`\`)}},
			},
		},
		{
			`\\a`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(`\a`)}},
			},
		},
		{
			`\\**`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("**")}},
			},
		},
		{
			`\\\**`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text(`\`),
					&node.Uniform{"Strong", nil},
				}},
			},
		},
		{
			"\\\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("``")}},
			},
		},
		{
			"\\\\\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text(`\`),
					&node.Escaped{"Code", nil},
				}},
			},
		},
		{
			`\\<`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text(`\<`)}},
			},
		},
		{
			`\\\<`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text(`\<`),
				}},
			},
		},

		{
			`a\**`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{node.Text("a**")}},
			},
		},
		{
			`\**\**`,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Uniform{"Strong", []node.Inline{
						node.Text("**"),
					}},
				}},
			},
		},
		{
			"\\``\\``",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					&node.Escaped{"Code", []byte("``")},
				}},
			},
		},
		{
			"``\\**\n\\**",
			[]node.Node{
				&node.Fenced{"CodeBlock", [][]byte{[]byte(`\**`), []byte(`\**`)}, nil},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			test(t, c.in, c.out, nil)
		})
	}
}

func TestInvalidUTF8Encoding(t *testing.T) {
	const fcb = "\x80" // first continuation byte

	cases := []struct {
		name string
		in   string
		out  []node.Node
	}{
		{
			"at the beginning",
			fcb + "a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text(string(unicode.ReplacementChar) + "a"),
				},
				}},
		},
		{
			"in the middle",
			"a" + fcb + "b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a" + string(unicode.ReplacementChar) + "b"),
				},
				}},
		},
		{
			"in the end",
			"a" + fcb,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a" + string(unicode.ReplacementChar)),
				},
				}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			test(t, c.in, c.out, []error{parser.ErrInvalidUTF8Encoding})
		})
	}
}

func TestNULL(t *testing.T) {
	const null = "\u0000"

	cases := []struct {
		name string
		in   string
		out  []node.Node
	}{
		{
			"at the beginning",
			null + "a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text(string(unicode.ReplacementChar) + "a"),
				},
				}},
		},
		{
			"in the middle",
			"a" + null + "b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a" + string(unicode.ReplacementChar) + "b"),
				},
				}},
		},
		{
			"in the end",
			"a" + null,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a" + string(unicode.ReplacementChar)),
				},
				}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			test(t, c.in, c.out, []error{parser.ErrIllegalNULL})
		})
	}
}

func TestBOM(t *testing.T) {
	const bom = "\uFEFF"

	t.Run("at the beginning", func(t *testing.T) {
		test(
			t,
			bom+"a",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a"),
				}},
			},
			nil,
		)
	})

	cases := []struct {
		name string
		in   string
		out  []node.Node
	}{
		{
			"in the middle",
			"a" + bom + "b",
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a" + string(unicode.ReplacementChar) + "b"),
				}},
			},
		},
		{
			"in the end",
			"a" + bom,
			[]node.Node{
				&node.BasicBlock{"TextBlock", []node.Inline{
					node.Text("a" + string(unicode.ReplacementChar)),
				}},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			test(t, c.in, c.out, []error{parser.ErrIllegalBOM})
		})
	}
}

// test compares the string representation of nodes generated by the parser from
// the argument in and the nodes of the argument out. Expected error must be
// encountered once; test calls t.Error() if it is encountered multiple times or
// if it is never encountered.
func test(t *testing.T, in string, out []node.Node, expectedErrors []error) {
	nodes, errs := parser.Parse([]byte(in))

	if expectedErrors == nil {
		for _, err := range errs {
			t.Errorf("got error %q", err)
		}
	} else {
		leftErrs := expectedErrors // errors we have not encountered yet
		for _, err := range errs {
			if i := errorIndex(leftErrs, err); i > -1 {
				// remove error
				leftErrs = append(leftErrs[:i], leftErrs[i+1:]...)
				continue
			}

			t.Errorf("got error %q", err)
		}

		// if some expected errors were not encountered
		for _, le := range leftErrs {
			t.Errorf("want error %q", le)
		}
	}

	got, want := stringifier.Stringify(node.BlocksToNodes(nodes)...), stringifier.Stringify(out...)
	if got != want {
		t.Errorf("\ngot\n%s\nwant\n%s", got, want)
	}
}

func errorIndex(errors []error, err error) int {
	for i, e := range errors {
		if err == e {
			return i
		}
	}
	return -1
}

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/yuin/goldmark/ast"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func main() {
	files, _ := filepath.Glob("cordeis/*.md")

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	for _, file := range files {
		str, _ := ioutil.ReadFile(file)
		context := parser.NewContext()

		n := markdown.Parser().Parse(text.NewReader(str), parser.WithContext(context))
		metaData := meta.Get(context)
		fmt.Printf("%#v\n", metaData)

		for _, c := range children(n) {
			i := 0
			for i < c.Lines().Len() {

				fmt.Printf("%#v\n", c.Lines().At(i).Value())
				i++
			}
		}

	}
}

func children(n ast.Node) []ast.Node {
	c := make([]ast.Node, n.ChildCount())

	if len(c) == 0 {
		return c
	}

	i := 0
	s := n.FirstChild()
	for i < n.ChildCount() {
		c[i] = s
		s.NextSibling()
		i = i + 1
	}
	return c
}

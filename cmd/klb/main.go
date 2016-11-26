package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/NeowayLabs/nash/ast"
	"github.com/NeowayLabs/nash/parser"
)

func printDoc(docs []*ast.CommentNode, fn *ast.FnDeclNode) {
	fmt.Printf("fn %s(%s)\n", fn.Name(), strings.Join(fn.Args(), ", "))

	for _, doc := range docs {
		fmt.Printf("\t%s\n", doc.String()[2:])
	}
}

func lookFn(fname string, pack string, fun string) bool {
	content, err := ioutil.ReadFile(fname)

	if err != nil {
		panic(err)
	}

	parser := parser.NewParser(fname, string(content))

	tree, err := parser.Parse()

	if err != nil {
		return false
	}

	nodelen := len(tree.Root.Nodes)

	for i, j := 0, 1; j < nodelen; i, j = i+1, j+1 {
		var comments []*ast.CommentNode

		node := tree.Root.Nodes[i]
		next := tree.Root.Nodes[j]

		if node.Type() == ast.NodeComment {
			comments = append(comments, node.(*ast.CommentNode))
			last := node

			// process comments
			for i = i + 1; i < nodelen-1; i++ {
				node = tree.Root.Nodes[i]

				if node.Type() == ast.NodeComment &&
					node.Line() == last.Line()+1 {
					comments = append(comments, node.(*ast.CommentNode))
					last = node
				} else {
					break
				}
			}

			j = i
			i--
			next = tree.Root.Nodes[j]

			if next.Line() != last.Line()+1 {
				comments = []*ast.CommentNode{}
			}
		} else if node.Type() == ast.NodeFnDecl {
			fn := node.(*ast.FnDeclNode)

			if fn.Name() == fun {
				printDoc([]*ast.CommentNode{}, fn)
				return true
			}

			continue
		} else {
			continue
		}

		if next.Type() == ast.NodeFnDecl {
			fn := next.(*ast.FnDeclNode)

			if fn.Name() == fun {
				printDoc(comments, next.(*ast.FnDeclNode))

				return true
			}
		}
	}

	return false
}

func usage() {
	fmt.Printf("%s doc <package>.<fn name>\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	packfn := os.Args[2]
	parts := strings.Split(packfn, ".")

	if len(parts) != 2 {
		usage()
	}

	pack := parts[0]
	funName := parts[1]

	nashpath := os.Getenv("NASHPATH")

	if nashpath == "" {
		homepath := os.Getenv("HOME")

		if homepath == "" {
			fmt.Fprintf(os.Stderr, "NASHPATH not set...\n")
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "NASHPATH not set. Using ~/.nash\n")

		nashpath = homepath + "/.nash"
	}

	err := filepath.Walk(nashpath+"/lib", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		dirpath := filepath.Dir(path)
		dirname := filepath.Base(dirpath)
		ext := filepath.Ext(path)

		if ext != "" && ext != ".sh" && dirname != pack {
			return nil
		}

		if found := lookFn(path, pack, funName); found {
			return errors.New("found")
		}

		return nil
	})

	if err != nil && err.Error() != "found" {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/NeowayLabs/nash/ast"
	"github.com/NeowayLabs/nash/parser"
)

func printDoc(stdout, _ io.Writer, docs []*ast.CommentNode, fn *ast.FnDeclNode) {
	fmt.Fprintf(stdout, "fn %s(%s)\n", fn.Name(), strings.Join(fn.Args(), ", "))

	for _, doc := range docs {
		fmt.Fprintf(stdout, "\t%s\n", doc.String()[2:])
	}
}

func lookFn(stdout, stderr io.Writer, fname string, pack string, fun string) bool {
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
				printDoc(stdout, stderr, []*ast.CommentNode{}, fn)
				return true
			}

			continue
		} else {
			continue
		}

		if next.Type() == ast.NodeFnDecl {
			fn := next.(*ast.FnDeclNode)

			if fn.Name() == fun {
				printDoc(stdout, stderr, comments, next.(*ast.FnDeclNode))

				return true
			}
		}
	}

	return false
}

func docUsage(out io.Writer) {
	fmt.Fprintf(out, "Usage: %s doc <package>.<fn name>\n", filepath.Base(os.Args[0]))
}

func doc(stdout, stderr io.Writer, args []string) error {
	if len(args) < 2 {
		docUsage(stderr)
		return nil
	}

	packfn := args[1]
	parts := strings.Split(packfn, ".")

	if len(parts) != 2 {
		docUsage(stderr)
		return nil
	}

	pack := parts[0]
	funName := parts[1]

	nashpath := os.Getenv("NASHPATH")

	if nashpath == "" {
		homepath := os.Getenv("HOME")

		if homepath == "" {
			return fmt.Errorf("NASHPATH not set...\n")
		}

		fmt.Fprintf(stderr, "NASHPATH not set. Using ~/.nash\n")

		nashpath = homepath + "/.nash"
	}

	return filepath.Walk(nashpath+"/lib", func(path string, info os.FileInfo, err error) error {
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

		lookFn(stdout, stderr, path, pack, funName)

		return nil
	})
}

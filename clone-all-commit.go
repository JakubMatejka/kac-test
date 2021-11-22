package main

import (
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
	"strings"
)

func main() {
	storer := memory.NewStorage()
	fs := memfs.New()

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://github.com/JakubMatejka/kac-test",
	})

	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}

	w, err2 := r.Worktree()
	if err2 != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}

	// ... checking out to commit
	err = w.Checkout(&git.CheckoutOptions{

	})
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}

	ref, err := r.Head()
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}

	// ... retrieve the tree from the commit
	tree, err := commit.Tree()
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}

	// ... get the files iterator and print the file
	tree.Files().ForEach(func(f *object.File) error {
		if strings.HasPrefix(f.Name, "template1/") {
			fmt.Printf("%s\n", f.Name)
		}
		return nil
	})
}

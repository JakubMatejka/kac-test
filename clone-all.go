package main

import (
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
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

	files, err3 := w.Filesystem.ReadDir(".")
	if err3 != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.Exit(1)
	}
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("%s", files))
}

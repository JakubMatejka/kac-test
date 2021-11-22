package main

import (
	"fmt"
	"github.com/libgit2/git2go/v33"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	tempDir, err := ioutil.TempDir(os.TempDir(), "api-")

	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
	repo, err := git.Clone(
		"https://github.com/JakubMatejka/kac-test",
		tempDir,
		&git.CloneOptions{CheckoutOptions: git.CheckoutOptions{Strategy: git.CheckoutNone}},
	)
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
	err = repo.CheckoutHead(&git.CheckoutOptions{Paths: []string{"template1"}})
	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}

	files, err := os.ReadDir(tempDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	_"github.com/go-git/go-git/plumbing/object"
	"github.com/go-git/go-git/v5"
)

func main() {
	dir := "./repo-temp"
	
	os.RemoveAll(dir)

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/Jemilson-Luis/sat",
		Progress: os.Stdout,
	})

	if err != nil {
		panic(err.Error())
	}

	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("Erro ao pegar HEAD: %v", err)
	}
	
	commits, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatalf("Erro ao obter commits: %v", err)
	}

	fmt.Print(commits)

}

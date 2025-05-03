package main

import (
	"encoding/json"
	_ "fmt"
	_ "log"
	"net/http"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/rs/cors"
)


type CommitDTO struct {
	
}


func main() {
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"}, // Permite essas origens
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"}, // Métodos permitidos
		AllowedHeaders: []string{"Content-Type", "Authorization"}, // Cabeçalhos permitidos
		AllowCredentials: true, // Permitir envio de credenciais (cookies, etc.)
	})

	handler := corsOptions.Handler(http.DefaultServeMux)



	http.HandleFunc("/getCommit", GetCommit)
	http.ListenAndServe(":8000", handler)
}



func GetCommit(w http.ResponseWriter, r *http.Request) {
	dir := "./repo-temp"

	os.RemoveAll(dir)

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/Jemilson-Luis/sat",
		Progress: os.Stdout,
	})
	if err != nil {
		http.Error(w, "Erro ao clonar o repositório: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ref, err := repo.Head()
	if err != nil {
		http.Error(w, "Erro ao pegar HEAD: "+err.Error(), http.StatusInternalServerError)
		return
	}

	commits, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		http.Error(w, "Erro ao obter commits: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var commitTimes []string

	err = commits.ForEach(func(c *object.Commit) error {
		commitTimes = append(commitTimes, c.Author.When.Format("2006-01-02 15:04:05"))
		return nil
	})

	if err != nil {
		http.Error(w, "Erro ao processar os commits: "+err.Error(), http.StatusInternalServerError)
		return
	}


	if err := json.NewEncoder(w).Encode(commitTimes); err != nil {
		http.Error(w, "Erro ao codificar os commits em JSON", http.StatusInternalServerError)
	}
}


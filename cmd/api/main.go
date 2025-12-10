package main

import (
	"log"
	"net/http"

	apihttp "onlineshop-aura/internal/http" // 👈 Өөрийн module нэрээр солиорой
)

func main() {
	r := apihttp.NewRouter() // 👈 @ NewRouter() эндээс ирж байгаа

	addr := ":8080"
	log.Printf("API listening on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

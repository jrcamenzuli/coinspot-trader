package subscriber

import (
	"log"
	"net/http"
)

func Start() {
	fs := http.FileServer(http.Dir("subscriber/frontend"))
	http.Handle("/", fs)

	log.Println("The frontend is listening on :8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

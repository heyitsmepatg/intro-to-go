// Kudos to Mohamed Labouardy, from whom I drew inspiration in
// building this tutorial. Source: https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Hacker represents a unique specimen of the human species.
// They congregate in large numbers on weekends with little sleep
// and copious amounts of caffeine. Despite their oddities, the
// Hacker is a loving, welcoming, and friendly kind.
type Hacker struct {
	// This data structure is sad!
}

// GetAllHackers basic GET api endpoint
func GetAllHackers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetAllHackers is not implemented yet !")
}

// GetHackerByID basic GET api endpoint
func GetHackerByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetHackerByID is not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hackers", GetAllHackers).Methods("GET")
	r.HandleFunc("/hackers/{id}", GetHackerByID).Methods("GET")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}

// Kudos to Mohamed Labouardy, from whom I drew inspiration in
// building this tutorial. Source: https://hackernoon.com/build-restful-api-in-go-and-mongodb-5e7f2ec4be94

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Hacker represents a unique specimen of the human species.
// They congregate in large numbers on weekends with little sleep
// and copious amounts of caffeine. Despite their oddities, the
// Hacker is a loving, welcoming, and friendly kind.
type Hacker struct {
	Name             string `json:"name"`
	ID               int    `json:"id"`
	FavoriteLanguage string `json:"favorite-language"`
}

// Hackers is a struct which contains
// an array of Hacker
type Hackers struct {
	Hackers []Hacker `json:"hackers"`
}

// getHackers reads a local JSON file and returns all hackers
func getHackers() Hackers {
	// Open our json
	jsonFile, err := os.Open("hackers.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	fmt.Println("Successfully Opened hackers.json")

	// Declare our slice of hackers
	var hackers Hackers

	// read our json file into a byte array
	// _ means we are intentionally ignoring the err return value -> Not best	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	b, _ := ioutil.ReadAll(jsonFile)

	// Unmarshal b into hackers,defined above
	// Notice the pointer &
	json.Unmarshal(b, &hackers)

	fmt.Println(hackers)

	return hackers
}

// GetAllHackers basic GET api endpoint
func GetAllHackers(w http.ResponseWriter, r *http.Request) {
	hackers := getHackers()

	jsonResponse, err := json.Marshal(hackers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// GetHackerByID basic GET api endpoint
func GetHackerByID(w http.ResponseWriter, r *http.Request) {
	// Multiple steps in one! We are:
	// 0.) Using the r Request objects to get more information about our request
	// 1.) Getting a map from mux that represents the query parameters
	// 2.) Calling the specific key/value pair where key is "id"
	// 3.) Converting the value from a string to integer so we can compare with the Hacker.ID
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	hackers := getHackers()

	for _, hacker := range hackers.Hackers {
		if hacker.ID == id {
			jsonResponse, err := json.Marshal(hacker)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write(jsonResponse)
			return
		}
	}
	// Tell our user this hacker DNE in our db!
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("Cannot find hacker of id: %v", id)))
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

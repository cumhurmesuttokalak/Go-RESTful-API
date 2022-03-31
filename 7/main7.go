package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string `json:"message"`
}
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	apiRoot := "/api"

	//.../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkErr(err)
		//w.Header().Set("Content-Type", "application-json")
		fmt.Fprintf(w, string(output))
	})

	//.../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "Cumhur", LastName: "TOKALAK", Age: 24},
			User{ID: 2, FirstName: "Mesut", LastName: "TOKALAK", Age: 26},
			User{ID: 3, FirstName: "MesMesut", LastName: "TOKALAK", Age: 29},
		}
		message := users
		output, err := json.Marshal(message)
		checkErr(err)
		fmt.Fprintf(w, string(output))
	})

	// ... /api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{3, "SemihNebi", "EmirMelih", 20}
		message := user
		output, err := json.Marshal(message)
		checkErr(err)
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":8081", nil)
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("Fatal Error:", err.Error())
		os.Exit(1)
	}
}

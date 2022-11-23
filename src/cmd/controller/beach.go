package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/franlauriano/docker-compose/app/beach"
)

// BeachHandle responsible for creating and listing beaches
func BeachHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[beaches] %s %s", r.Method, r.URL.Path)

	responseBody := []byte{}
	statusCode := http.StatusMethodNotAllowed

	switch r.Method {
	case "GET":
		responseBody, statusCode = foods(w)
	case "POST":
		responseBody, statusCode = createFood(w, r)
	}

	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

func foods(w http.ResponseWriter) ([]byte, int) {
	foods, err := beach.List()
	if err != nil {
		log.Printf("Error on list beaches. %s", err)
	}

	body, err := json.Marshal(foods)
	if err != nil {
		log.Printf("Error on marshall data structure to JSON. %s", err)
		return []byte{}, http.StatusInternalServerError
	}

	return body, http.StatusOK
}

func createFood(w http.ResponseWriter, r *http.Request) ([]byte, int) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on read request body. %s", err)
		return []byte{}, http.StatusInternalServerError
	}

	var beachModel beach.Beach
	err = json.Unmarshal(body, &beachModel)
	if err != nil {
		log.Printf("Error on unmarshall JSON to data structure. %s", err)
		return []byte{}, http.StatusBadRequest
	}

	if err := beachModel.Create(); err != nil {
		log.Printf("Error on create beach. %s", err)
		return []byte{}, http.StatusInternalServerError
	}

	return body, http.StatusOK
}

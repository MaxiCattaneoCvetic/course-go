package main

import (
	"encoding/json" // con eso transformamos una estructura a json
	"fmt"
	"log"
	"net/http"
)

func init() {
	// cuando se ejecuta la aplicacion esto es lo primero que se ejecuta (antes del main)
	users = []User{
		{
			ID:        1,
			FirstName: "Cristian",
			LastName:  "Vega",
			Email:     "j9Z8x@example.com",
		},
		{
			ID:        2,
			FirstName: "Juan",
			LastName:  "Peres",
			Email:     "juanperez@example.com",
		},
		{
			ID:        3,
			FirstName: "Paco",
			LastName:  "Costa",
			Email:     "Costa@example.com",
		},
	}
	MaxID = 3
}

func main() {

	http.HandleFunc("/users", UserServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserServer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetAllUsers(w)
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var u User
		if err := decode.Decode(&u); err != nil {
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		PostUser(w, u)
	default:
		InvalidMethod(w)
	}
}

/* METODO GET*/
func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users []User
var MaxID uint64

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users) // me transforma 1 entidad a un json
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, value)
}

/* METODO POST*/

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)

}

func PostUser(w http.ResponseWriter, data interface{}) {
	user := data.(User) // casteo el valor de data a un User
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Faltan datos")
		return
	}

	MaxID++
	user.ID = MaxID
	// agregamos el user a los usuarios que ya tenemos
	users = append(users, user)
	DataResponse(w, http.StatusCreated, user)

}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "Method doesnt exist"}`, status)
}

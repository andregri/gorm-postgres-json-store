package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/andregri/gorm-postgres-json-store/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

func (driver *DBClient) GetUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	vars := mux.Vars(r)

	// handle response details
	driver.db.First(&user, vars["id"])
	var userData interface{}

	// unmarshal json string to interface
	json.Unmarshal([]byte(user.Data), &userData)

	var response = UserResponse{User: user, Data: userData}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	respJson, _ := json.Marshal(response)
	w.Write(respJson)
}

func (driver *DBClient) PostUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}

	// insert to db
	postBody, _ := ioutil.ReadAll(r.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)

	// build response
	responseMap := map[string]interface{}{"id": user.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic("Error during table creation")
	}

	dbClient := &DBClient{db: db}

	r := mux.NewRouter()

	r.HandleFunc("/v1", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("api v1"))
	})
	r.HandleFunc("/v1/user/{id:[a-zA-Z0-9]*}", dbClient.GetUser).Methods("GET")
	r.HandleFunc("/v1/user", dbClient.PostUser).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Panic(srv.ListenAndServe())
}

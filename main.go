package main

import (
	"fmt"
	"os"
	"log"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"encoding/json"
	//"strconv"
	"github.com/heroku/cmanager/entity"
	"strconv"
	"flag"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var dir string
	flag.StringVar(&dir, "dir", "static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	corsObj := handlers.AllowedOrigins([]string{"*"})

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	router.HandleFunc("/car", CarListAction).Methods("GET")
	router.HandleFunc("/car/{plate}", GetCarAction).Methods("GET")
	router.HandleFunc("/car/{plate}", CreateOrUpdateAction).Methods("POST")
	router.HandleFunc("/car/{plate}", DeleteAction).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj)(router)))
}

func CarListAction(w http.ResponseWriter, r *http.Request) {
	//traceAllVars(w, r)
	json.NewEncoder(w).Encode(entity.GetAllCars())
}

func GetCarAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := entity.GetCarByPlate(vars["plate"])
	json.NewEncoder(w).Encode(result)
}

func CreateOrUpdateAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		panic(err)
	}
	entity.UpsertCar(r.FormValue("model"), year, vars["plate"], r.FormValue("brand"))
}

func DeleteAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entity.DeleteCar(vars["plate"])
}

func TraceAllVars(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "%#v", vars)
}
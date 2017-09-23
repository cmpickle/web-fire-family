package web_fire_family


import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
 router := mux.NewRouter()
 log.Fatal(http.ListenAndServe(":8000", router))
}
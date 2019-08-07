package main
import (
    "fmt"
     "log"
     "net/http"
     //"encoding/json"
     "github.com/gorilla/mux"
    "category"
    "queries"
)

func main() {
	handleRequests()
	fmt.Println("Rest Api's")

}
func handler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Success!\n"))
}
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    //myRouter.HandleFunc("/", handler)
    myRouter.HandleFunc("/category/{category}", category.CategoryData).Methods("GET")
    myRouter.HandleFunc("/queries/{type}", queries.GetQueries).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

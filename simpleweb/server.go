package main
import (
	"fmt"
	"log"
	"net/http"
	//"time"

	"github.com/gorilla/mux"
)


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Welcome to yamaszone!")
		})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// s := &http.Server{
	// 	Addr:           ":8081",
	// 	//Handler:        myHandler,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	log.Fatal(http.ListenAndServe(":8081", r))
}

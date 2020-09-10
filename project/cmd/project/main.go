package main


import (
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/server"
	"log"
	"net/http"
)

func main(){
	s := server.New()

	log.Fatal(http.ListenAndServe(":8000", s.Router()))
}
package main


import (
	"github.com/abnergarcia1/voxie-engineering-test/project/pkg/project/server"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

func main(){
	s := server.New()

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS","DELETE"})

	log.Fatal(http.ListenAndServe(":8000",handlers.CORS(originsOk, headersOk, methodsOk)(s.Router())))
}
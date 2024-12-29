package main
// entry point
import (
    "fmt"
    "log"
    "github.com/movie-api/router"
    "net/http"
)

func main() {
	// initialising router
    r := router.InitRouter()
    fmt.Println("Server running on port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}

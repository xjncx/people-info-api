package m

import (
	"log"
	"net/http"

	"github.com/xjncx/people-info-api/internal/api"
	"github.com/xjncx/people-info-api/internal/service"
)

func main() {
	// Пока сервис с nil (заглушка)

	var repo service.PersonRepository // пока nil
	service := service.NewPersonService(repo)
	handler := &api.Handler{PersonService: service}

	r := api.NewRouter(handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

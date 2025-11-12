package routes

import (
	"net/http"
	"github.com/VarunSharma3520/go-api/internal/controllers"
)


var ApiV1Mux = http.NewServeMux()

func init() {
	ApiV1Mux.HandleFunc("/create-todo", controllers.CreateTodoController)
	ApiV1Mux.HandleFunc("/read-todo", controllers.ReadTodoController)
	ApiV1Mux.HandleFunc("/update-todo", controllers.UpdateTodoController)
	ApiV1Mux.HandleFunc("/delete-todo", controllers.DeleteTodoController)
}

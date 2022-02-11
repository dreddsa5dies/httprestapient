package router

import (
	"net/http"

	"github.com/dreddsa5dies/httprestapient/controller"

	"github.com/gorilla/mux"
)

// all route
func registerMatrixAttackRouter(r *mux.Router) {
	userRouter := r.PathPrefix("/api").Subrouter()
	userRouter.HandleFunc("/", controller.MatrixAttacksGetAllController).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", controller.MatrixAttackGetByIDController).Methods(http.MethodGet)
	userRouter.HandleFunc("/", controller.MatrixAttackCreateController).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", controller.MatrixAttackUpdateController).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", controller.MatrixAttackDeleteController).Methods(http.MethodDelete)
}

// register all available route
func RegisterRouter(r *mux.Router) {
	registerMatrixAttackRouter(r)
}

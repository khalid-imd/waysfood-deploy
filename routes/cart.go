package routes

import (
	"fundamental-golang/handlers"
	"fundamental-golang/pkg/middleware"
	"fundamental-golang/pkg/mysql"
	"fundamental-golang/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/carts", middleware.Auth(h.FindCarts)).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart/{id}", h.UpdateCart).Methods("PATCH")
	r.HandleFunc("/cart/{id}", h.DeleteCart).Methods("DELETE")
}

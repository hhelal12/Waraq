package router

import "net/http"

func SetupRoutes(mux *http.ServeMux, api *API) {

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			api.UserHandler.GetAllUsers(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	
}

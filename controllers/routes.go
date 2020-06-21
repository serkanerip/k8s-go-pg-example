package controllers

func (s *Server) initilizeRoutes() {
	s.Router.HandleFunc("/", s.Home).Methods("GET")
	s.Router.HandleFunc("/users", s.ListUsers).Methods("GET")
	s.Router.HandleFunc("/users", s.CreateUser).Methods("POST")
}

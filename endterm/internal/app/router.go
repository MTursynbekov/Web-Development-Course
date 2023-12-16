package app

func (s *Server) Route() {
	s.app.Post("/signup", s.SignupHandler)
	s.app.Post("/signin", s.SigninHandler)

	api := s.app.Group("/api")

	api.Get("/users/:userId/posts", s.GetUsersPostsHandler)
	api.Get("/users/:userId/posts/:postId", s.GetPostHandler)
	api.Post("/posts", s.CreatePostHandler)
	api.Put("/posts/:postId", s.UpdatePostHandler)
	api.Delete("/posts/:postId", s.DeletePostHandler)

}

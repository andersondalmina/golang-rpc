package server

import (
	"github.com/andersondalmina/golang-rpc/persist"
	"github.com/andersondalmina/golang-rpc/services"
)

type Server int

func (s *Server) CreateBook(params map[string]string, reply *persist.Book) error {
	var err error
	*reply, err = services.CreateBook(params)
	return err
}

func (s *Server) UpdateBook(params map[string]string, reply *persist.Book) error {
	var err error
	*reply, err = services.UpdateBook(params)
	return err
}

func (s *Server) SearchBookByTitle(title string, reply *[]persist.Book) error {
	*reply = services.SearchBooksByTitle(title)
	return nil
}

func (s *Server) SearchBookByAuthor(author string, reply *[]persist.Book) error {
	*reply = services.SearchBooksByAuthor(author)
	return nil
}

func (s *Server) SearchBookByEdition(edition string, reply *[]persist.Book) error {
	*reply = services.SearchBooksByEdition(edition)
	return nil
}

func (s *Server) SearchBookByYear(year string, reply *[]persist.Book) error {
	*reply = services.SearchBooksByYear(year)
	return nil
}

func (s *Server) DeleteBookByTitle(title string, reply *persist.Book) error {
	return services.DeleteBooksByTitle(title)
}

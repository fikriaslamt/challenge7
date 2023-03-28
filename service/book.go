package service

import (
	"sesi_8/model"
)

type BookService interface {
	GetAllBook() ([]model.Book, error)
	GetBookById(int) (model.Book, error)
	AddBook(model.Book) (model.Book, error)
	UpdateBook(int, model.Book) error
	DeleteBook(int) error
}

func (s *Service) GetAllBook() ([]model.Book, error) {
	res, err := s.repo.GetAllBook()
	if err != nil {
		return []model.Book{}, err
	}

	return res, nil
}
func (s *Service) GetBookById(id int) (model.Book, error) {
	res, err := s.repo.GetBookById(id)
	if err != nil {

		return model.Book{}, err
	}

	return res, nil
}
func (s *Service) AddBook(data model.Book) (model.Book, error) {
	res, err := s.repo.AddBook(data)
	if err != nil {
		return model.Book{}, err
	}
	return res, nil
}
func (s *Service) UpdateBook(id int, data model.Book) error {
	err := s.repo.UpdateBook(id, data)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) DeleteBook(id int) error {
	err := s.repo.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

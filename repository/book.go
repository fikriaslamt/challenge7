package repository

import (
	"database/sql"
	"errors"
	"sesi_8/model"
	"sesi_8/repository/query"
)

// interface employee
type BookRepo interface {
	GetAllBook() ([]model.Book, error)
	GetBookById(int) (model.Book, error)
	AddBook(model.Book) (model.Book, error)
	UpdateBook(int, model.Book) error
	DeleteBook(int) error
}

func (r Repo) GetAllBook() ([]model.Book, error) {
	res := []model.Book{}
	var err error

	rows, err := r.db.Query(query.GetAllBook)
	if err != nil {
		return []model.Book{}, err
	}

	defer rows.Close()

	for rows.Next() {
		book := model.Book{}

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			return []model.Book{}, err
		}
		res = append(res, book)

	}

	return res, nil

}

func (r Repo) GetBookById(id int) (model.Book, error) {
	res := model.Book{}

	rows := r.db.QueryRow(query.GetBookById, id)

	err := rows.Scan(&res.ID, &res.Title, &res.Author, &res.Desc)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Book{}, errors.New("DATA NOT FOUND")
		}
		return model.Book{}, err
	}

	return res, nil
}

func (r Repo) AddBook(data model.Book) (model.Book, error) {
	var res = model.Book{}

	err := r.db.QueryRow(query.AddBook, data.Title, data.Author, data.Desc).Scan(&res.ID, &res.Title, &res.Author, &res.Desc)
	if err != nil {
		return model.Book{}, err
	}

	return res, nil
}

func (r Repo) UpdateBook(id int, data model.Book) error {
	res, err := r.db.Exec(query.UpdateBook, id, data.Title, data.Author, data.Desc)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("DATA NOT FOUND")
	}
	return nil
}

func (r Repo) DeleteBook(id int) error {
	res, err := r.db.Exec(query.DeleteBook, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("DATA NOT FOUND")
	}
	return nil
}

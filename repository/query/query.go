package query

const (
	GetAllBook = `
	SELECT * FROM postgres.book ;
	`
	GetBookById = `
	SELECT * FROM postgres.book
	WHERE id = $1;
	`
	AddBook = `
	INSERT INTO postgres.book (title, author, "desc")
	VALUES ($1, $2, $3) RETURNING *;
	`
	UpdateBook = `
	UPDATE postgres.book
	SET title = $2, author = $3 , "desc" = $4
	WHERE id = $1 ; 
	`
	DeleteBook = `
	DELETE FROM postgres.book
	WHERE id = $1 ; 
	`
)

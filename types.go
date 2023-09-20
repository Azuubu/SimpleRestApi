package main

type Book struct {
	ID         string      `json:"id"`
	Title      string      `json:"book_title"`
	Pages      int64       `json:"book_pages"`
	Genre      string      `json:"book_genre"`
	BookAuthor *BookAuthor `json:"author"`
}

type BookAuthor struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Nationality string `json:"nationality"`
	IsAlive     bool   `json:"is_alive"`
}

package books

type Service interface {
	GetBooks() []Book
	GetBook(id string) Book
	CreateBook(libro Book) []Book
	//updateBook (libro Book) Book
	//deleteBook (libro Book) []Book
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func GetBooks() []Book {
	// Hardcoded data - @todo: add database
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	return books
}

// Get single book
func GetBook(id string) Book{
	if id == "1" {
		return Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}}
	} else {
		return Book{}
	}
}

// Add new book
func CreateBook(libro Book) []Book {
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, Book{
		ID:     libro.ID,
		Isbn:   libro.Isbn,
		Title:  libro.Title,
		Author: libro.Author,
	})
	return books
}

// Update book
/*func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}*/

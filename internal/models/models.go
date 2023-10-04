package models

type Book struct {
	ID     int
	TITLE  string
	AUTHOR string
	PRICE  int
}

// dummy-data
var Book1 = Book{1, "Book number 1", "Kouhadi Bakr", 19}
var Book2 = Book{2, "Book number 2", "Bryan Aboubakr", 149}
var Book3 = Book{3, "Book number 3", "Omar Kouhadi", 399}
var Book4 = Book{4, "Book number 4", "Sheikh zoubir Kouhadi", 99}
var Book5 = Book{5, "Book number 5", "Mark Twain", 49}

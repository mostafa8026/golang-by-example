package models

// Book struct
type Book struct {
	Id     int
	Title  string
	Author string
	Price  float32
}

// AllBooks get all books from table
func (db *DB) AllBooks() ([]*Book, error) {
	rows, err := db.Query("Select * from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Id, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

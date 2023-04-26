package store

import (
	"database/sql"
	"github.com/laurianderson/bootcamp_go_repository/internal/domain"
)

// Querys
const (
	QueryGetByID         = `SELECT id, name, quantity, code_value, is_published, expiration, price FROM products WHERE id = ?; `
	QueryCreate          = `INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?);`
	QueryUpdateByID      = `UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?;`
	QueryDeleteByID      = `DELETE FROM products WHERE id = ?;`
	QueryCodeValueExists = `SELECT id FROM products WHERE code_value = ?;`
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Read(id int) (domain.Product, error) {
	var product domain.Product
	//prepare statement
	row := s.db.QueryRow(QueryGetByID, id)

	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *sqlStore) Create(product domain.Product) error {
	//prepare statement
	stmt, err := s.db.Prepare(QueryCreate)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Update(product domain.Product) error {
	//prepare statement
	stmt, err := s.db.Prepare(QueryUpdateByID)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {
	//prepare statement
	stmt, err := s.db.Prepare(QueryDeleteByID)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Exists(codeValue string) bool {
	var exists bool
	var id int
	//prepare statement
	row := s.db.QueryRow(QueryCodeValueExists, codeValue)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

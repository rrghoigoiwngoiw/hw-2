package database

import (
	"database/sql"
	"fmt"
)

type DB struct {
	*sql.DB
}

func NewConnection(dbname, user, password, host string, port int) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

func (db *DB) CreateUser(username, password, email string) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, username, password, email)
	return err
}

func (db *DB) UpdateUser(id int, username, password, email string) error {
	query := `UPDATE users SET username=$1, password=$2, email=$3 WHERE id=$4`
	_, err := db.Exec(query, username, password, email, id)
	return err
}

func (db *DB) GetUserByID(id int) (*User, error) {
	query := `SELECT id, username, password, email FROM users WHERE id = $1`
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (db *DB) CreateProduct(name string, price float64) error {
	query := `INSERT INTO products (name, price) VALUES ($1, $2)`
	_, err := db.Exec(query, name, price)
	return err
}

func (db *DB) UpdateProduct(id int, name string, price float64) error {
	query := `UPDATE product SET name=$1, price=$2 WHERE id=$3`
	_, err := db.Exec(query, name, price, id)
	return err
}

func (db *DB) DeleteProduct(id int) error {
	query := `DELETE FROM Products WHERE id = $1;`
	_, err := db.Exec(query, id)
	return err
}

func (db *DB) GetPriceByName(id int) (*Product, error) {
	query := `SELECT price FROM products WHERE id = $1`
	row := db.QueryRow(query, id)

	var product Product
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &product, nil
}

type Orders struct {
	ID          int
	UserID      int
	OrderDate   string
	TotalAmount float64
}

func (db *DB) CreateOrder(userID int, orderDate string, totalAmount float64) error {
	query := `INSERT INTO orders (user_id, order_date, total_amount) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, userID, orderDate, totalAmount)
	return err
}

func (db *DB) DeleteOrder(id int) error {
	query := `DELETE FROM orders WHERE id = $1;`
	_, err := db.Exec(query, id)
	return err
}

func (db *DB) GetOrderByID(id int) (*Orders, error) {
	query := `SELECT id, user_id, order_date, total_amount FROM orders WHERE id = $1`
	row := db.QueryRow(query, id)

	var order Orders
	err := row.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &order, nil
}

func (db *DB) GetOrdersByUserID(userID int) ([]Orders, error) {
	query := `SELECT id, user_id, order_date, total_amount FROM orders WHERE user_id = $1`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Orders
	var scanErr error

	for rows.Next() {
		var order Orders
		scanErr = rows.Scan(&order.ID, &order.UserID, &order.OrderDate, &order.TotalAmount)
		if scanErr != nil {
			return nil, scanErr
		}
		orders = append(orders, order)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}

	return orders, nil
}

func (db *DB) AddProductToOrder(orderID, productID, quantity int) error {
	query := `INSERT INTO OrderProducts (order_id, product_id, quantity) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, orderID, productID, quantity)
	return err
}

func (db *DB) RemoveProductFromOrder(orderID, productID int) error {
	query := `DELETE FROM OrderProducts WHERE order_id = $1 AND product_id = $2;`
	_, err := db.Exec(query, orderID, productID)
	return err
}

func (db *DB) CreateOrderWithProducts(userID int, orderDate string, totalAmount float64, products []Product) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	orderQuery := `INSERT INTO Orders (user_id, order_date, total_amount) VALUES ($1, $2, $3) RETURNING id;`
	var orderID int
	err = tx.QueryRow(orderQuery, userID, orderDate, totalAmount).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, product := range products {
		productQuery := `INSERT INTO OrderProducts (order_id, product_id, quantity) VALUES ($1, $2, $3);`
		_, err = tx.Exec(productQuery, orderID, product.ID, 1)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

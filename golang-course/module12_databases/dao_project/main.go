package main

import (
	"context"
	"database/sql"
	"fmt"
)

// ─────── DAO (Data Access Object) PATTERN ───────
// A great way to structure database access in Go is the DAO or Repository pattern.
// You define an interface for your database operations, and then a concrete struct
// that implements that interface using `*sql.DB`.
// This makes it incredibly easy to swap out the real database with a mock for testing!

// 1. Domain Model
type Product struct {
	ID    int
	Name  string
	Price float64
}

// 2. The Interface (Used by the business logic)
type ProductRepository interface {
	GetByID(ctx context.Context, id int) (Product, error)
	Create(ctx context.Context, p *Product) error
}

// 3. The SQL Implementation
type SQLProductRepository struct {
	DB *sql.DB
}

func (r *SQLProductRepository) GetByID(ctx context.Context, id int) (Product, error) {
	var p Product
	query := "SELECT id, name, price FROM products WHERE id = $1"
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *SQLProductRepository) Create(ctx context.Context, p *Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id"
	// Postgres supports RETURNING id. MySQL would use LastInsertId() on the sql.Result.
	err := r.DB.QueryRowContext(ctx, query, p.Name, p.Price).Scan(&p.ID)
	return err
}

// 4. The Business Logic Service
type ProductService struct {
	Repo ProductRepository // Depends ONLY on the interface!
}

func (s *ProductService) RegisterNewProduct(ctx context.Context, name string, price float64) (*Product, error) {
	if price <= 0 {
		return nil, fmt.Errorf("price must be greater than zero")
	}
	
	p := &Product{Name: name, Price: price}
	err := s.Repo.Create(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("failed to save product: %w", err)
	}
	
	return p, nil
}

func main() {
	fmt.Println("--- DAO / Repository Pattern Demo ---")
	fmt.Println("This pattern decouples business logic from database technology.")
	
	// In reality:
	// db, _ := sql.Open("postgres", "...")
	// repo := &SQLProductRepository{DB: db}
	// svc := &ProductService{Repo: repo}
	// svc.RegisterNewProduct(...)
}

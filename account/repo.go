package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
)

var (
	RepoErr             = errors.New("Unable to handle Repo Request")
	ErrIdNotFound       = errors.New("Id not found")
	ErrPhonenumNotFound = errors.New("Phone num is not found")
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "postgres"),
	}, nil
}

func (repo *repo) InitialCustomer(ctx context.Context) {
	stmt := `DO $$
	BEGIN
	CREATE TABLE IF NOT EXISTS customers (
	id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 124567895123 CACHE 1 ),
	email character varying COLLATE pg_catalog."default",
	phone character varying COLLATE pg_catalog."default",
	CONSTRAINT customers_pkey PRIMARY KEY (id);
	END;
	$$;`
	repo.db.Exec(stmt)
}

func (repo *repo) CreateCustomer(ctx context.Context, customer Customer) error {
	stmt := `DO $$
	BEGIN
	CREATE TABLE IF NOT EXISTS customers (
	id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 124567895123 CACHE 1 ),
	email character varying COLLATE pg_catalog."default",
	phone character varying COLLATE pg_catalog."default",
	CONSTRAINT customers_pkey PRIMARY KEY (id)
	);
	END;
	$$;`
	_, err := repo.db.Exec(stmt)
	fmt.Println(err)

	_, err = repo.db.ExecContext(ctx, "INSERT INTO Customers(email, phone) VALUES ($1, $2)", customer.Email, customer.Phone)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Error occured inside CreateCustomer in repo")
		return err
	} else {
		fmt.Println("User Created:", customer.Email)
	}
	return nil
}
func (repo *repo) GetCustomerById(ctx context.Context, id int64) (interface{}, error) {
	customer := Customer{}

	err := repo.db.QueryRowContext(ctx, "SELECT c.id,c.email,c.phone FROM Customers as c where c.id = $1", id).Scan(&customer.ID, &customer.Email, &customer.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}
		return customer, err
	}
	return customer, nil
}
func (repo *repo) GetAllCustomers(ctx context.Context) (interface{}, error) {
	customer := Customer{}
	var res []interface{}
	rows, err := repo.db.QueryContext(ctx, "SELECT c.id,c.email,c.phone FROM Customers as c ")
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, ErrIdNotFound
		}
		return customer, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Email, &customer.Phone)
		res = append([]interface{}{customer}, res...)
	}
	return res, nil
}
func (repo *repo) DeleteCustomer(ctx context.Context, id int64) (string, error) {
	res, err := repo.db.ExecContext(ctx, "DELETE FROM Customers WHERE id = $1 ", id)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	} else if rowCnt == 0 {
		return "", ErrIdNotFound
	}
	return "Successfully deleted ", nil
}
func (repo *repo) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	res, err := repo.db.ExecContext(ctx, "UPDATE Customers SET email=$1 , phone = $2 WHERE id = $3", customer.Email, customer.Phone, customer.ID)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", ErrIdNotFound
	}

	return "successfully updated", err
}

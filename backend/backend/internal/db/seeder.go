package db

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/durianpay/fullstack-boilerplate/helper"
	"golang.org/x/crypto/bcrypt"
)

func Seed(db *sql.DB) error {
	rand.Seed(time.Now().UnixNano())

	if err := seedUsers(db); err != nil {
		return err
	}

	if err := seedPayments(db); err != nil {
		return err
	}

	return nil
}

func seedUsers(db *sql.DB) error {
	var cnt int
	db.QueryRow("SELECT COUNT(1) FROM users").Scan(&cnt)

	if cnt > 0 {
		return nil
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	_, err := db.Exec(`
		INSERT INTO users(email,password_hash,role) VALUES
		('cs@test.com', ?, 'cs'),
		('operation@test.com', ?, 'operation')
	`, string(hash), string(hash))

	return err
}

func seedPayments(db *sql.DB) error {
	var cnt int
	db.QueryRow("SELECT COUNT(1) FROM payments").Scan(&cnt)

	if cnt > 0 {
		return nil
	}

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(`
		INSERT INTO payments(payment_id,merchant,status,amount,created_at)
		VALUES (?, ?, ?, ?, ?)
	`)

	defer stmt.Close()

	for i := 0; i < 500; i++ {
		_, err := stmt.Exec(
			helper.RandomPaymentID(),
			helper.RandomMerchant(),
			helper.RandomStatus(),
			rand.Intn(90000)+10000,
			helper.RandomCreatedAt(),
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

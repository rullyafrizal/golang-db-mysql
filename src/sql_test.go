package src

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsertData(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	//var balance int32 = 150000
	var rating float32 = 4.3

	currentTime := time.Now()
	var created_at time.Time = currentTime
	var birth_date time.Time = currentTime

	var script string = "INSERT INTO customers(id, name, rating, created_at, birth_date, is_married) VALUES('C2', 'Rully Afrizal', " + strconv.FormatFloat(float64(rating), 'f', 2, 32) + ", '" + created_at.Format("2006-01-02 19:00:00") + "', '" + birth_date.Format("2006-01-02") + "', 0);"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted data successfully\n")
}

func TestQueryData(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var script string = "SELECT * FROM customers;"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	// iterasi dan tampilkan data
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating float32
		var created_at, birth_date time.Time
		var is_married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &is_married)

		if err != nil {
			panic(err)
		}

		fmt.Println("---------------------")
		fmt.Println("ID: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		if balance.Valid {
			fmt.Println("Balance: ", balance.Int32)
		}
		fmt.Println("Rating: ", rating)
		fmt.Println("Created At: ", created_at)
		fmt.Println("Birth Date: ", birth_date)
		fmt.Println("Is Married: ", is_married)
		fmt.Println("---------------------")
	}

	defer rows.Close()
}

func TestQueryWithParameter(t *testing.T) {
	var username string = "rully"
	var password string = "rully123"

	db := GetConnection()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	query := "SELECT * FROM users WHERE username = ? AND password = ? LIMIT 1;"
	rows, err := db.QueryContext(ctx, query, username, password)

	defer rows.Close()

	if err != nil {
		panic(err)
	}

	// iterasi dan tampilkan data
	if rows.Next() {
		var username string
		var password string

		err := rows.Scan(&username, &password)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login")
	} else {
		fmt.Println("Gagal login")
	}
}

func TestInsertWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var username string = "rully"
	var password string = "rully123"

	var script string = "INSERT INTO users(username, password) VALUES(?, ?);"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully insert data\n")
}

func TestLastInserId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	email := "rully@gmail.com"
	comment := "Ini adalah komentar"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?);"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertedId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Last Inserted Id: ", insertedId)

}

func TestPrepareStatement(t *testing.T) {
	// jika query yang sama berulang-ulang lebih baik pakai prepare statement
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?);"

	// buat dulu prepare context-nya
	stmt, err := db.PrepareContext(ctx, script)

	for i := 0; i < 10; i++ {
		email := "rully" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini adalah komentar ke-" + strconv.Itoa(i)

		// execute query tanpa script sql
		// karena script sql sudah di-binding dengan prepare statement sebelumnya
		result, err := stmt.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id :", id)
	}

	if err != nil {
		panic(err)
	}

	defer stmt.Close()
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?);"

	for i := 0; i < 10; i++ {
		email := "rully" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini adalah komentar ke-" + strconv.Itoa(i)

		// execute query tanpa script sql
		// karena script sql sudah di-binding dengan prepare statement sebelumnya
		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id :", id)
	}

	err = tx.Commit()

	if err != nil {
		panic(err)
	}
}

package repositories

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "golang-database-mysql/src"
	"golang-database-mysql/src/models"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(GetConnection())

	ctx := context.Background()
	comment := models.Comment{
		Email:   "repository@test.com",
		Comment: "Ini adalah komen dari repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 24)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(GetConnection())

	ctx := context.Background()

	results, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, v := range results {
		fmt.Println("id :", v.Id)
		fmt.Println("email :", v.Email)
		fmt.Println("comment :", v.Comment)
		fmt.Printf("\n")
	}
}

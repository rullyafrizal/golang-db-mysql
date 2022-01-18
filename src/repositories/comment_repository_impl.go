package repositories

import (
	"context"
	"database/sql"
	"errors"
	"golang-database-mysql/src/models"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (c *CommentRepositoryImpl) Insert(ctx context.Context, comment models.Comment) (models.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := c.DB.ExecContext(ctx, script, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)

	return comment, nil
}

func (c *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (models.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id=? LIMIT 1;"
	rows, err := c.DB.QueryContext(ctx, script, id)

	comment := models.Comment{}

	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		return comment, nil
	}

	return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
}

func (c *CommentRepositoryImpl) FindAll(ctx context.Context) ([]models.Comment, error) {
	script := "SELECT id, email, comment FROM comments;"
	rows, err := c.DB.QueryContext(ctx, script)

	comments := []models.Comment{}

	if err != nil {
		return comments, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment models.Comment

		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		if err != nil {
			panic(err)
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

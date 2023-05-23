package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/fixztter/rest-fundamentals/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, password) values($1, $2, $3)", user.Id, user.Email, user.Password)
	return err
}

func (r *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id=$1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email=$1", email)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) values($1, $2, $3)", post.Id, post.PostContent, post.UserId)
	return err
}

func (r *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, post_content, created_at, user_id FROM posts WHERE id=$1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var post = models.Post{}
	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.PostContent, &post.CreatedAt, &post.UserId); err == nil {
			return &post, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := r.db.ExecContext(ctx, "UPDATE posts SET post_content=$1 WHERE id=$2 AND user_id=$3", &post.PostContent, post.Id, post.UserId)
	return err
}

func (r *PostgresRepository) DeletePost(ctx context.Context, id string, userId string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id=$1 AND user_id=$2", id, userId)
	return err
}

func (r *PostgresRepository) ListPost(ctx context.Context, page uint64) ([]*models.Post, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, post_content, user_id, created_at FROM posts LIMIT $1 OFFSET $2", 5, page*2)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var posts []*models.Post
	for rows.Next() {
		var post = models.Post{}
		if err = rows.Scan(&post.Id, &post.PostContent, &post.UserId, &post.CreatedAt); err == nil {
			posts = append(posts, &post)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}

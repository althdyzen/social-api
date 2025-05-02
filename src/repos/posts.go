package repos

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func (repo *Posts) Get(userID uint64) ([]models.Post, error) {
	lines, erro := repo.db.Query(`
		SELECT DISTINCT p.*, u.nick, u.name, u.id = ? FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers f ON p.author_id = f.id_user
		WHERE u.id = ? OR f.id_follower = ?
		ORDER BY 1 DESC`,
		userID,
		userID,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if erro := lines.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
			&post.AuthorName,
			&post.IsAuthor,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil

}

func (repo *Posts) GetMe(userID uint64) ([]models.Post, error) {
	lines, erro := repo.db.Query(`
		SELECT p.* FROM posts p INNER JOIN users u ON u.id = p.author_id WHERE u.id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if erro := lines.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil

}

func (repo *Posts) GetByID(postID uint64) (models.Post, error) {
	lines, erro := repo.db.Query(`
		SELECT p.*, u.nick , u.name
		FROM posts p 
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.id = ?`,
		postID,
	)
	if erro != nil {
		return models.Post{}, erro
	}
	defer lines.Close()

	if !lines.Next() {
		return models.Post{}, erro
	}

	var post models.Post

	if erro := lines.Scan(
		&post.ID,
		&post.Title,
		&post.Body,
		&post.AuthorId,
		&post.Likes,
		&post.CreatedAt,
		&post.AuthorNick,
		&post.AuthorName,
	); erro != nil {
		return models.Post{}, erro
	}

	return post, nil
}

func (repo *Posts) New(post models.Post) (uint64, error) {
	statement, erro := repo.db.Prepare("INSERT INTO posts (title, body, author_id) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(post.Title, post.Body, post.AuthorId)
	if erro != nil {
		return 0, erro
	}

	lastInsertedID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertedID), nil
}

func (repo *Posts) Update(postID uint64, post models.Post) error {
	statement, erro := repo.db.Prepare("UPDATE posts SET title = ?, body = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(post.Title, post.Body, postID); erro != nil {
		return erro
	}

	return nil
}

func (repo *Posts) Delete(postID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM posts WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

func (repo *Posts) GetByUser(userID uint64) ([]models.Post, error) {
	lines, erro := repo.db.Query(`
		SELECT p.*, u.nick, u.name FROM posts p
		JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if erro := lines.Scan(
			&post.ID,
			&post.Title,
			&post.Body,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
			&post.AuthorName,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *Posts) Like(postID uint64) error {
	statement, erro := repo.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

func (repo *Posts) Dislike(postID uint64) error {
	statement, erro := repo.db.Prepare(`
		UPDATE posts SET likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0
		END
		WHERE id = ?`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

func NewPostRepo(db *sql.DB) *Posts {
	return &Posts{db}
}

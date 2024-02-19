package repositories

import (
	"database/sql"
	"goapi/src/models"
)

// Posts represents a posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository creates a new post repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create inserts a post into a database
func (repository Posts) Create(Post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into Posts (title, content, author_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(Post.Title, Post.Content, Post.AuthorID)
	if err != nil {
		return 0, err
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(LastInsertId), nil
}

// SearchByID brings a single post to the database
func (repository Posts) SearchByID(PostID uint64) (models.Post, error) {
	line, err := repository.db.Query(`
	select p.*, u.nick from 
	Posts p inner join users u
	on u.id = p.author_id where p.id = ?`,
		PostID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var Post models.Post

	if line.Next() {
		if err = line.Scan(
			&Post.ID,
			&Post.Title,
			&Post.Content,
			&Post.AuthorID,
			&Post.Likes,
			&Post.CreatedAt,
			&Post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return Post, nil
}

// Search brings a post of followed users and of the user that made the request
func (repository Posts) Search(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
	select distinct p.*, u.nick from Posts p 
	inner join users u on u.id = p.author_id 
	inner join followers f on p.author_id = f.user_id 
	where u.id = ? or f.follower_id = ?
	order by 1 desc`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var Posts []models.Post

	for lines.Next() {
		var Post models.Post

		if err = lines.Scan(
			&Post.ID,
			&Post.Title,
			&Post.Content,
			&Post.AuthorID,
			&Post.Likes,
			&Post.CreatedAt,
			&Post.AuthorNick,
		); err != nil {
			return nil, err
		}

		Posts = append(Posts, Post)
	}

	return Posts, nil
}

// Update changes data of a post on the database
func (repository Posts) Update(PostID uint64, Post models.Post) error {
	statement, err := repository.db.Prepare("update Posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(Post.Title, Post.Content, PostID); err != nil {
		return err
	}

	return nil
}

// Delete exclui uma publicação do banco de dados
func (repository Posts) Delete(PostID uint64) error {
	statement, err := repository.db.Prepare("delete from Posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(PostID); err != nil {
		return err
	}

	return nil
}

// SearchByUser gets all posts of a specific user
func (repository Posts) SearchByUser(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		select p.*, u.nick from Posts p
		join users u on u.id = p.author_id
		where p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var Posts []models.Post

	for lines.Next() {
		var Post models.Post

		if err = lines.Scan(
			&Post.ID,
			&Post.Title,
			&Post.Content,
			&Post.AuthorID,
			&Post.Likes,
			&Post.CreatedAt,
			&Post.AuthorNick,
		); err != nil {
			return nil, err
		}

		Posts = append(Posts, Post)
	}

	return Posts, nil
}

// Like adds a like to a post
func (repository Posts) Like(PostID uint64) error {
	statement, erro := repository.db.Prepare("update Posts set likes = likes + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(PostID); erro != nil {
		return erro
	}

	return nil
}

// Dislike subtracts a like of a post
func (repository Posts) Dislike(PostID uint64) error {
	statement, err := repository.db.Prepare(`
		update Posts set likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		where id = ?
	`)
	if err != nil {
		return err
	}

	if _, err = statement.Exec(PostID); err != nil {
		return err
	}

	return nil
}

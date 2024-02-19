package repositories

import (
	"database/sql"
	"fmt"
	"goapi/src/models"
)

// this represents a users repository
type Users struct {
	db *sql.DB
}

// This func creates a new users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// This adds user to db
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (nameuser, nick, email, pw) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err

	}
	defer statement.Close()

	result, err := statement.Exec(user.Nameuser, user.Nick, user.Email, user.Pw)
	if err != nil {
		return 0, err
	}
	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil

}

// Search returns all users that meet the name or nick filter request
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick

	lines, err := repository.db.Query(
		"select id, nameuser, nick, email, createdAt from users where nameuser LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err

	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Nameuser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}
func (repository Users) SearchForID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, nameuser, nick, email, createdAt from users where id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err

	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Nameuser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// Update alters info of a user in db
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set nameuser= ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Nameuser, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil

}

// Delete deletes a user from the db
func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from users where id = ?")
	if err != nil {
		return err

	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}
func (repository Users) SearchForEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, pw from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Pw); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
func (repository Users) Follow(userId, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerID); err != nil {
		return err
	}

	return nil
}

// UnFollow allows a user to stop following another
func (repository Users) UnFollow(userId, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"Delete from followers (user_id, follower_id) values (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerID); err != nil {
		return err
	}

	return nil
}

// SearchFollowers gets all followers a user has
func (repository Users) SearchFollowers(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.nameuser, u.nick, u.email, u.createdAt
		from users u inner join followers f on u.id = f.follower_id where f.user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Nameuser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchFollowing gets all followers from a user
func (repository Users) SearchFollowing(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
			select u.id, u.nameuser, u.nick, u.email, u.createdAt
			from Users u inner join followers f on u.id = f.user_id where f.follower_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Nameuser,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchPw gets a user pw by the ID
func (repository Users) SearchPw(userID uint64) (string, error) {
	line, erro := repository.db.Query("select pw from users where id = ?", userID)
	if erro != nil {
		return "", erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Pw); erro != nil {
			return "", erro
		}
	}

	return user.Pw, nil
}
func (repository Users) UpdatePw(userID uint64, pw string) error {
	statement, err := repository.db.Prepare("update users set pw = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(pw, userID); err != nil {
		return err
	}

	return nil
}

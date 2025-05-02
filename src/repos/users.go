package repos

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func (repo users) New(usuario models.User) (uint64, error) {
	statement, erro := repo.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(usuario.Name, usuario.Nick, usuario.Email, usuario.Password)
	if erro != nil {
		return 0, erro
	}

	lastIdInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIdInserted), nil
}

func (repo users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repo.db.Query("SELECT id, name, nick FROM users WHERE name LIKE ? OR nick LIKE ? LIMIT 5", nameOrNick, nameOrNick)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro := lines.Scan(&user.ID, &user.Name, &user.Nick); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil

}

func (repo users) GetByID(userID, userIDInToken uint64) (models.User, error) {
	lines, erro := repo.db.Query(`
		SELECT 
		    id, 
		    name, 
		    nick,
		    id = ?,
		    (SELECT COUNT(*) FROM followers WHERE id_user = ? AND id_follower = ?) as segue,
		    (SELECT COUNT(*) FROM followers WHERE id_user = ?) AS total_seguidores,
		    (SELECT COUNT(*) FROM followers WHERE id_follower = ?) AS total_seguindo
		FROM 
		    users
		WHERE 
		    id = ?;`,
		userIDInToken,
		userID,
		userIDInToken,
		userID,
		userID,
		userID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if !lines.Next() {
		return models.User{}, fmt.Errorf("usuário com id não encontrado")
	}

	if erro := lines.Scan(
		&user.ID,
		&user.Name,
		&user.Nick,
		&user.IsOwner,
		&user.IsFollowing,
		&user.Followers,
		&user.Following,
	); erro != nil {
		return models.User{}, erro
	}

	return user, nil
}

func (repo users) Update(ID uint64, user models.User) error {
	statement, erro := repo.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repo users) Delete(ID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repo users) GetByEmail(email string) (models.User, error) {
	line, erro := repo.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repo users) Follow(followerID, userID uint64) error {
	statement, erro := repo.db.Prepare("INSERT IGNORE INTO followers (id_user, id_follower) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

func (repo users) Unfollow(followerID, userID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM followers WHERE id_user = ? AND id_follower = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

func (repo users) GetFollowersByID(userID uint64) ([]models.User, error) {
	lines, erro := repo.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt
		FROM users u INNER JOIN followers f ON u.id = f.id_follower WHERE f.id_user = ?`,
		userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo users) GetFollowingByID(userID uint64) ([]models.User, error) {
	lines, erro := repo.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.createdAt
		FROM users u INNER JOIN followers f ON u.id = f.id_user WHERE f.id_follower = ?`,
		userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo users) GetPasswordByID(userID uint64) (string, error) {
	line, erro := repo.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if erro != nil {
		return "", erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}

func (repo users) UpdatePassword(userID uint64, password string) error {
	statement, erro := repo.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(password, userID); erro != nil {
		return erro
	}

	return nil
}

func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

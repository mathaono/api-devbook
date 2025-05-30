package repositories

import (
	"api/src/models"
	"database/sql"
	"log"
)

// Representa um repositório de usuários
type users struct {
	db *sql.DB
}

// Criação de um repositório de usuários
func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

// Insere um usuário no banco de dados
func (repoUser users) Create(user models.User) (int64, error) {
	statement, err := repoUser.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES ($1, $2, $3, $4);")
	if err != nil {
		log.Printf("❌ [ERROR DB] - Failed to connect DB: %v", err)
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		log.Printf("❌ [ERROR DB] - Failed to insert user on DB: %v", err)
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	log.Printf("[SUCCESS DB] - User inserted on DB: %d", id)

	return id, nil
}

// Trás todos os usuários da base de dados que atendem aos filtros
func (repoUser users) Search(userName string) ([]models.User, error) {
	userName = "%" + userName + "%"

	rows, err := repoUser.db.Query("SELECT id, name, nickname, email FROM users WHERE name ILIKE $1 OR nickname ILIKE $2", userName, userName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Trás um usuário do banco pelo ID
func (repoUser users) SearchByID(ID uint64) (models.User, error) {
	rows, err := repoUser.db.Query("SELECT id, name, nickname, email, created_at FROM users WHERE id = $1", ID)
	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Atualiza as infos de um usuário no banco de dados
func (repoUser users) Update(ID uint64, user models.User) error {
	statement, err := repoUser.db.Prepare(
		"UPDATE users SET name = $1, nickname = $2, email = $3 WHERE id = $4")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nickname, user.Email, ID)
	if err != nil {
		return err
	}

	return nil
}

// Exclui as infos de um usuário no banco de dados
func (repoUser users) Delete(ID uint64) error {
	statement, err := repoUser.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

// Busca um usuário no banco de dados pelo email e retorna ID e senha com Hash
func (repoUser users) SearchByEmail(email string) (models.User, error) {
	row, err := repoUser.db.Query("SELECT id, password FROM users WHERE email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		err = row.Scan(&user.ID, &user.Password)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Permite que um usuário siga outro
func (repoUser users) Follow(userID, followerID uint64) error {
	statement, err := repoUser.db.Prepare(`INSERT INTO followers (user_id, follow_id) VALUES ($1, $2) ON CONFLICT (user_id, follow_id) DO NOTHING`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Permite que um usuário deixe de seguir outro
func (repoUser users) Unfollow(userID, followerID uint64) error {
	statement, err := repoUser.db.Prepare(`DELETE FROM followers WHERE user_id = $1 AND follow_id = $2`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Permite buscar seguidores de um usuário no banco
func (repoUser users) SearchFollowers(userID uint64) ([]models.User, error) {
	rows, err := repoUser.db.Query(`SELECT u.id, u.name, u.nickname, u.email
		FROM users u 
		INNER JOIN followers f ON u.id = f.follow_id 
		WHERE f.user_id = $1`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Permite buscar os usuários que o usuário da request está seguindo
func (repoUser users) SearchFollowing(userID uint64) ([]models.User, error) {
	rows, err := repoUser.db.Query(`SELECT u.id, u.name, u.nickname, u.email
		FROM users u INNER JOIN followers f ON u.id = f.user_id
		WHERE f.follow_id = $1`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repoUser users) SearchPass(userID uint64) (string, error) {
	row, err := repoUser.db.Query(`SELECT password from users WHERE id = $1`, userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

// Altera a senha do usuário da request
func (repoUser users) UpdatePass(userID uint64, pass string) error {
	statement, err := repoUser.db.Prepare(`UPDATE users SET password = $1 WHERE id = $2`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(pass, userID); err != nil {
		return err
	}

	return nil
}

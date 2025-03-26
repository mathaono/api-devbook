package repositories

import (
	"api/src/models"
	"database/sql"
)

// Representa um repositório de publicações
type publications struct {
	db *sql.DB
}

// Criação de um repositório de publicações
func NewPublicationRepository(db *sql.DB) *publications {
	return &publications{db}
}

// Insere uma nova publicação no banco de dados
func (repoPub publications) Create(pub models.Publication) (int64, error) {
	statement, err := repoPub.db.Prepare(`INSERT INTO publications (title, content, user_id) VALUES ($1, $2, $3) RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var lastInsertID int64
	err = statement.QueryRow(pub.Title, pub.Content, pub.UserID).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// Busca todas as publicações, tanto do próprio usuário quanto as publicações de usuários seguidos
func (repoPub publications) SearchPublications(pubID uint64) ([]models.Publication, error) {
	rows, err := repoPub.db.Query(`
		SELECT p.*, u.nickname FROM	publications p 
		INNER JOIN users u ON u.id = p.user_id
		INNER JOIN followers f ON p.user_id = f.user_id
		WHERE u.id = $1 OR f.follow_id = $2 ORDER BY 1 ASC`, pubID, pubID)
	if err != nil {
		return nil, err
	}

	var pubs []models.Publication

	for rows.Next() {
		var publication models.Publication
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.UserID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UserNick,
		); err != nil {
			return nil, err
		}

		pubs = append(pubs, publication)
	}

	return pubs, nil
}

// Busca uma publicação por ID
func (repoPub publications) SearchPublicationByID(pubID uint64) (models.Publication, error) {
	rows, err := repoPub.db.Query(`
		SELECT p.*, u.nickname FROM	publications p 
		INNER JOIN users u
		ON u.id = p.user_id WHERE p.id = $1`, pubID)
	if err != nil {
		return models.Publication{}, err
	}

	defer rows.Close()

	var publication models.Publication

	if rows.Next() {
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.UserID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UserNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

// Permite atualizar os dados de uma publicação do usuário
func (repoPub publications) Update(pubID uint64, publication models.Publication) error {
	statement, err := repoPub.db.Prepare(`UPDATE publications SET title = $1, content = $2 WHERE id = $3`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, pubID); err != nil {
		return err
	}

	return nil
}

// Permite que o usuário exclua uma de suas publicações do banco de dados
func (repoPub publications) Delete(pubID uint64) error {
	statement, err := repoPub.db.Prepare(`DELETE FROM publications WHERE id = $1`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(pubID); err != nil {
		return err
	}

	return nil
}

// Permite o usuário trazer todas as publicações no banco de dados de um usuário específico
func (repoPub publications) SearchPublicationByUser(userID uint64) ([]models.Publication, error) {
	rows, err := repoPub.db.Query(`
		SELECT p.*, u.nickname FROM publications p
		JOIN users u on u.id = p.user_id
		WHERE p.user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.UserID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.UserNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Adiciona uma curtida na publicação no banco de dados
func (repoPub publications) Like(publicationID uint64) error {
	statement, err := repoPub.db.Prepare(`UPDATE publications SET likes = Likes + 1 WHERE id = $1`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

func (repoPub publications) Dislike(publicationID uint64) error {
	statement, err := repoPub.db.Prepare(`UPDATE publications SET likes = likes - 1 WHERE id = $1 AND likes > 0`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

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

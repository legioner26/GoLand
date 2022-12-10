package post

import (
	"context"
	"database/sql"
	models "skillbox/30-31/models"
	pRepo "skillbox/30-31/repository"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLPostRepo(Conn *sql.DB) pRepo.PostRepo {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

type mysqlPostRepo struct {
	Conn *sql.DB
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) (*models.RequestCreate, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.RequestCreate, 0)
	for rows.Next() {
		data := new(models.RequestCreate)

		err := rows.Scan(
			&data.Name,
			&data.Age,
			&data.Friends,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, Name string) (*models.RequestCreate, error) {
	query := "Select Name, Age, Friends From friends where Name=?"

	return m.fetch(ctx, query, Name)
}

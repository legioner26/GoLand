package post

import (
	"context"
	"database/sql"
	"log"
	"os"
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

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.RequestCreate, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("Получаем аргумент %s", args...)
	infoLog.Printf("Получаем строку запроса %s", query)

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	infoLog.Printf("Получаем ответ с запроса %s", rows)
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

func (m *mysqlPostRepo) Fetch(ctx context.Context, Names string) (*models.RequestCreate, error) {

	query := "Select name, age, friends From friends where name=?"

	rows, err := m.fetch(ctx, query, Names)
	if err != nil {
		return nil, err
	}

	payload := &models.RequestCreate{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

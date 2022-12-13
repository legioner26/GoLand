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

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.RequestSelect, error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("Получаем аргумент %s", args...)
	infoLog.Printf("Получаем строку запроса %s", query)

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payload := make([]*models.RequestSelect, 0)
	for rows.Next() {
		data := new(models.RequestSelect)

		err := rows.Scan(
			&data.Name,
			&data.Age,
			&data.Friend,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
		infoLog.Printf("данные %s", payload)
	}
	return payload, nil

}

func (m *mysqlPostRepo) Fetch(ctx context.Context, Names string) ([]*models.RequestSelect, error) {

	query := "Select name, age, friend From friends where name=?"

	return m.fetch(ctx, query, Names)
}

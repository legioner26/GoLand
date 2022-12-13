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
func (m *mysqlPostRepo) Create(ctx context.Context, p *models.RequestSelect) (int64, error) {

	query := "Insert friends SET name=?, age=?, friend=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.Age, p.Friend)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
func (m *mysqlPostRepo) Update(ctx context.Context, p *models.RequestMakeFriend) (*models.RequestMakeFriend, error) {
	query := "Update friends set name=?, friend=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.Friend,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlPostRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From friends Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

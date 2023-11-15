package main

import "database/sql"

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS vacancies (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255),
			company VARCHAR(255),
			location VARCHAR(255),
			description TEXT
		);

		CREATE TABLE IF NOT EXISTS search_history (
			id SERIAL PRIMARY KEY,
			query VARCHAR(255),
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

	return &SQLRepository{db: db}
}

func (r *SQLRepository) SearchVacancy(query string) ([]Vacancy, error) {
	rows, err := r.db.Query(`
		SELECT id, title, company, location, description
		FROM vacancies
		WHERE title ILIKE $1
	`, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []Vacancy
	for rows.Next() {
		var v Vacancy
		err := rows.Scan(&v.ID, &v.Title, &v.Company, &v.Location, &v.Description)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, v)
	}

	return vacancies, nil
}

func (r *SQLRepository) GetVacancy(id string) (*Vacancy, error) {
	var v Vacancy
	err := r.db.QueryRow(`
		SELECT id, title, company, location, description
		FROM vacancies
		WHERE id = $1
	`, id).Scan(&v.ID, &v.Title, &v.Company, &v.Location, &v.Description)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (r *SQLRepository) ListVacancies() ([]Vacancy, error) {
	rows, err := r.db.Query(`
		SELECT id, title, company, location, description
		FROM vacancies
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []Vacancy
	for rows.Next() {
		var v Vacancy
		err := rows.Scan(&v.ID, &v.Title, &v.Company, &v.Location, &v.Description)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, v)
	}

	return vacancies, nil
}

func (r *SQLRepository) SaveVacancy(vacancy Vacancy) error {
	_, err := r.db.Exec(`
		INSERT INTO vacancies (title, company, location, description)
		VALUES ($1, $2, $3, $4)
	`, vacancy.Title, vacancy.Company, vacancy.Location, vacancy.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLRepository) DeleteVacancy(id string) error {
	_, err := r.db.Exec(`
		DELETE FROM vacancies
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLRepository) SaveSearchHistory(query string) error {
	_, err := r.db.Exec(`
		INSERT INTO search_history (query)
		VALUES ($1)
	`, query)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLRepository) ListSearchHistory() ([]SearchHistory, error) {
	rows, err := r.db.Query(`
		SELECT id, query, timestamp
		FROM search_history
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []SearchHistory
	for rows.Next() {
		var h SearchHistory
		err := rows.Scan(&h.ID, &h.Query, &h.Timestamp)
		if err != nil {
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}

func (r *SQLRepository) DeleteSearchHistory(id int) error {
	_, err := r.db.Exec(`
		DELETE FROM search_history
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	return nil
}

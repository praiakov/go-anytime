package core

import "database/sql"

type UseCase interface {
	GetPeople() ([]*Person, error)
	// GetPerson(id int64) (*Person, error)
	// CreatePerson(p *Person) error
	// DeletePerson(id int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) GetPeople() ([]*Person, error) {
	var result []*Person

	rows, err := s.DB.Query("select id, firtname, lastname from peoples")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var p Person
		err = rows.Scan(&p.ID, &p.Firstname, &p.Lastname)
		if err != nil {
			return nil, err
		}

		result = append(result, &p)
	}

	return result, nil
}

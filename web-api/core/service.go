package core

import "database/sql"

type UseCase interface {
	GetPeople() ([]*Person, error)
	GetPerson(id int64) (*Person, error)
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

	rows, err := s.DB.Query("SELECT p.id, p.firtname, p.lastname, a.city, a.name_state FROM peoples AS p INNER JOIN adresses as a ON p.id = a.person_id")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var p Person
		var a Address
		p.Address = &a

		err = rows.Scan(&p.ID, &p.Firstname, &p.Lastname, &p.Address.City, &p.Address.State)
		if err != nil {
			return nil, err
		}

		result = append(result, &p)
	}

	return result, nil
}

func (s *Service) GetPerson(id int64) (*Person, error) {
	var p Person
	var a Address
	p.Address = &a

	stmt, err := s.DB.Prepare("SELECT p.id, p.firtname, p.lastname, a.city, a.name_state FROM peoples AS p INNER JOIN adresses AS a	ON p.id = a.id WHERE p.id = $1")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Firstname, &p.Lastname, &p.Address.City, &p.Address.State)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

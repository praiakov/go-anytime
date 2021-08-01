package core

import (
	"database/sql"
	"fmt"
)

type UseCase interface {
	GetPeople() ([]*Person, error)
	GetPerson(id int64) (*Person, error)
	CreatePerson(p *Person) error
	UpdatePerson(p *Person) error
	DeletePerson(id int64) error
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

	stmt, err := s.DB.Prepare("SELECT p.id, p.firtname, p.lastname, a.city, a.name_state FROM peoples AS p INNER JOIN adresses AS a	ON p.id = a.person_id WHERE p.id = $1")

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

func (s *Service) CreatePerson(p *Person) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("WITH person AS (INSERT INTO peoples(firtname,lastname) VALUES($1, $2) RETURNING id) INSERT INTO adresses (city, name_state, person_id) VALUES ($3, $4, (select id from person))")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(p.Firstname, p.Lastname, p.Address.City, p.Address.State)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) DeletePerson(id int64) error {
	if id == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM peoples where id=$1", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) UpdatePerson(p *Person) error {
	if p.ID == "" {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("WITH person AS (UPDATE peoples SET firtname= $1, lastname= $2 WHERE id = $3 RETURNING id) UPDATE adresses AS a SET city = $4, name_state = $5 WHERE a.person_id = (select id from person)")

	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()
	_, err = stmt.Exec(p.Firstname, p.Lastname, p.ID, p.Address.City, p.Address.State)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil

}

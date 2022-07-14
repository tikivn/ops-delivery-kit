package model

import (
	"encoding/json"
	"errors"
	"time"
)

type UID struct {
	ID    string    `json:"id"`
	SUB   string    `json:"sub"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Exp   time.Time `json:"exp"`
}

var errorNotValid = errors.New("")

func (p *UID) Valid() error {
	if time.Now().After(p.Exp) {
		return errorNotValid
	}

	return nil
}

func (p *UID) UnmarshalJSON(data []byte) error {
	uidTemp := struct {
		ID    string `json:"id"`
		Sub   string `json:"sub"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Trn   string `json:"trn"`
		Exp   int64  `json:"exp"`
	}{}

	if err := json.Unmarshal(data, &uidTemp); err != nil {
		return err
	}

	p.ID = uidTemp.ID
	if p.ID == "" {
		p.ID = uidTemp.Trn
	}

	if p.ID == "" {
		p.ID = uidTemp.Sub
	}
	p.Name = uidTemp.Name
	p.Email = uidTemp.Email
	p.SUB = uidTemp.Sub
	p.Exp = time.Unix(uidTemp.Exp, 0)

	return nil
}

func (p *UID) MarshalJSON() ([]byte, error) {
	uidTemp := struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Sub   string `json:"sub"`
		Exp   int64  `json:"exp"`
	}{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
		Sub:   p.SUB,
		Exp:   p.Exp.Unix(),
	}

	return json.Marshal(uidTemp)
}

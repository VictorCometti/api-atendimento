package models

import "time"

type Service struct {
	ID                int       `json:"-"`
	ServiceDate       time.Time `json:"service_date,omitempty"`
	ServiceIdentifier int       `json:"service_identifier,omitempty"`
	Exams             []Exam    `json:"exams,omitempty"`
}

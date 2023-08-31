package models

import "time"

type Exam struct {
	ID                  int    `json:"-"`
	PatientName         string `json:"patient_name,omitempty"`
	ExamDescription     string `json:"exam_description,omitempty"`
	RequestingPhysician string `json:"requesting_physician,omitempty"`
}

type ExamDTO struct {
	ServiceDate         time.Time
	ServiceIdentifier   int
	Exams               []Exam `json:"exams,omitempty"`
	PatientName         string
	ExamDescription     string
	RequestingPhysician string
}

package models

type Exam struct {
	ID                  int    `json:"-"`
	PatientName         string `json:"patient_name,omitempty"`
	ExamDescription     string `json:"exam_description,omitempty"`
	RequestingPhysician string `json:"requesting_physician,omitempty"`
}

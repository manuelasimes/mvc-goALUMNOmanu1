package dto

type SubjectDto struct {
	Name      string `json:"name_subject"`
	StudentId int    `json:"student_id,omitempty"`
}

type SubjectsDto []SubjectDto

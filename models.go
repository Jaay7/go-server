package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           *string            `json:"name" validate:"required,min=2,max=100"`
	Salary         *string            `json:"salary" validate:"required"`
	Description    *string            `json:"description" validate:"required"`
	JobType        *string            `json:"jobtype" validate:"required"`
	Qualifications *string            `json:"qualifications" validate:"required"`
	Eligibility    *string            `json:"eligibility" validate:"required"`
	Requirements   *string            `json:"requirements" validate:"required"`
}

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

type Student struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	CollegeID        *string            `json:"collegeid" validate:"required"`
	Password         *string            `json:"password" validate:"required"`
	FirstName        *string            `json:"firstname" validate:"required"`
	MiddleName       *string            `json:"middlename" validate:"required"`
	LastName         *string            `json:"lastname" validate:"required"`
	DateOfBirth      *string            `json:"dateofbirth" validate:"required"`
	Gender           *string            `json:"gender" validate:"required"`
	FatherName       *string            `json:"fathername" validate:"required"`
	MotherName       *string            `json:"mothername" validate:"required"`
	MotherMaidenName *string            `json:"mothermaidenname" validate:"required"`
	BloodGroup       *string            `json:"bloodgroup" validate:"required"`
	MartialStatus    *string            `json:"martialstatus" validate:"required"`
	MotherTongue     *string            `json:"mothertongue" validate:"required"`
	Caste            *string            `json:"caste" validate:"required"`
	PersonalEmail    *string            `json:"personalemail" validate:"required"`
	Identification   *string            `json:"identification" validate:"required"`
	Disability       *string            `json:"disability" validate:"required"`
	PlaceOfBirth     *string            `json:"placeofbirth" validate:"required"`
	Heigth           *string            `json:"heigth" validate:"required"`
	Weight           *string            `json:"weigth" validate:"required"`
	Religion         *string            `json:"religion" validate:"required"`
	Nationality      *string            `json:"nationality" validate:"required"`
	AdmissionDate    *string            `json:"admissiondate" validate:"required"`
	MajorDegree      *string            `json:"majordegree" validate:"required"`
	Reference        *string            `json:"reference" validate:"required"`
	Program          *string            `json:"program" validate:"required"`
	Regulation       *string            `json:"regulation" validate:"required"`
	Pending          *bool              `json:"pending" validate:"required"`
}

type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CollegeID   *string            `json:"collegeid" validate:"required"`
	AddressType *string            `json:"addresstype" validate:"required"`
	Doorno      *string            `json:"doorno" validate:"required"`
	Street      *string            `json:"street" validate:"required"`
	Landmark    *string            `json:"landmark" validate:"required"`
	City        *string            `json:"city" validate:"required"`
	District    *string            `json:"district" validate:"required"`
	State       *string            `json:"state" validate:"required"`
	Country     *string            `json:"country" validate:"required"`
}

type Contacts struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CollegeID     *string            `json:"collegeid" validate:"required"`
	ContactType   *string            `json:"contacttype" validate:"required"`
	ContactPerson *string            `json:"contactperson" validate:"required"`
	PhoneNumber   *string            `json:"phonenumber" validate:"required"`
}

type Identity struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CollegeID      *string            `json:"collegeid" validate:"required"`
	IdentityType   *string            `json:"identitytype" validate:"required"`
	IdentityNumber *string            `json:"identitynumber" validate:"required"`
	IssuedOn       *string            `json:"issuedon" validate:"required"`
	DateOfExpiry   *string            `json:"dateofexpiry" validate:"required"`
	PlaceOfIssue   *string            `json:"placeofissue" validate:"required"`
}

type Qualifications struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CollegeID     *string            `json:"collegeid" validate:"required"`
	Qualification *string            `json:"qualification" validate:"required"`
	Board         *string            `json:"board" validate:"required"`
	EduName       *string            `json:"eduname" validate:"required"`
	CGPA          *string            `json:"cgpa" validate:"required"`
	YearOfPassing *string            `json:"yearofpassing" validate:"required"`
	Specilization *string            `json:"specilization" validate:"required"`
}

type Courses struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CollegeID    *string            `json:"collegeid" validate:"required"`
	Year         *string            `json:"year" validate:"required"`
	AcademicYear *string            `json:"academicyear" validate:"required"`
	Semester     *string            `json:"semester" validate:"required"`
	CourseCode   *string            `json:"coursecode" validate:"required"`
	CourseDesc   *string            `json:"coursedesc" validate:"required"`
	LTPS         *string            `json:"ltps" validate:"required"`
	Section      *string            `json:"section" validate:"required"`
	FacultyName  *string            `json:"facultyname" validate:"required"`
}

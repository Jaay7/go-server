package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	uri := os.Getenv("MONGODB_URI")
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())

	router.POST("/createjob", func(c *gin.Context) {
		var job Job

		if err := c.BindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := client.Database("gomongo").Collection("job")

		result, err := collection.InsertOne(ctx, job)

		if err != nil {
			msg := fmt.Sprintf("Job was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	router.GET("/jobs", func(c *gin.Context) {
		jobs := []Job{}

		collection := client.Database("gomongo").Collection("job")

		cursor, err := collection.Find(context.TODO(), bson.M{})

		if err != nil {
			msg := fmt.Sprintf("no jobs found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		for cursor.Next(context.TODO()) {
			var job Job
			cursor.Decode(&job)
			jobs = append(jobs, job)
		}
		defer cancel()

		c.JSON(http.StatusOK, gin.H{
			"data": jobs,
		})
		return
	})

	router.GET("/job/:jobId", func(c *gin.Context) {
		jobId := c.Param("jobId")

		job := Job{}
		collection := client.Database("gomongo").Collection("job")

		objId, _ := primitive.ObjectIDFromHex(jobId)
		err := collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&job)

		if err != nil {
			msg := fmt.Sprintf("current job not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, gin.H{
			"data": job,
		})
	})

	router.GET("/jobs/:jobtype", func(c *gin.Context) {
		jobs := []Job{}

		jobtype := c.Param("jobtype")
		collection := client.Database("gomongo").Collection("job")

		// filter := bson.D{{"jobtype", jobtype}}
		cursor, err := collection.Find(context.TODO(), bson.D{{"jobtype", jobtype}})

		if err != nil {
			msg := fmt.Sprintf("no jobs found under this section")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		for cursor.Next(context.TODO()) {
			var job Job
			cursor.Decode(&job)
			jobs = append(jobs, job)
		}
		defer cancel()

		c.JSON(http.StatusOK, gin.H{
			"data": jobs,
		})

	})

	// student routes
	router.POST("/login", func(c *gin.Context) {
		// var student Student
		stu := Student{}

		CollegeID := c.Query("collegeid")
		Password := c.Query("password")

		collection := client.Database("gomongo").Collection("student")

		err := collection.FindOne(context.TODO(), bson.M{"collegeid": CollegeID}).Decode(&stu)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Student not found",
			})
			return
		}

		if Password != *stu.Password {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Incorrect Password",
			})
			return
		}

		token, err := CreateToken(*stu.CollegeID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  http.StatusUnprocessableEntity,
				"message": "Could not create token",
			})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "Login successful",
			"collegeid": stu.CollegeID,
			"pending":   stu.Pending,
			"token":     token,
		})

	})

	router.POST("/addstudent", func(c *gin.Context) {
		var student Student

		if err := c.BindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		v := rand.Intn(999999999-100000000) + 100000000

		CollegeID := strconv.Itoa(v)
		Password := CollegeID
		FirstName := student.FirstName
		MiddleName := student.MiddleName
		LastName := student.LastName
		DateOfBirth := student.DateOfBirth
		Gender := student.Gender
		FatherName := student.FatherName
		MotherName := student.MotherName
		MotherMaidenName := student.MotherMaidenName
		BloodGroup := student.BloodGroup
		MartialStatus := student.MartialStatus
		MotherTongue := student.MotherTongue
		Caste := student.Caste
		PersonalEmail := student.PersonalEmail
		Identification := student.Identification
		Disability := student.Disability
		PlaceOfBirth := student.PlaceOfBirth
		Heigth := student.Heigth
		Weight := student.Weight
		Religion := student.Religion
		Nationality := student.Nationality
		AdmissionDate := student.AdmissionDate
		MajorDegree := student.MajorDegree
		Reference := student.Reference
		Program := student.Program
		Regulation := student.Regulation
		Pending := true

		newStudent := Student{
			CollegeID:        &CollegeID,
			Password:         &Password,
			FirstName:        FirstName,
			MiddleName:       MiddleName,
			LastName:         LastName,
			DateOfBirth:      DateOfBirth,
			Gender:           Gender,
			FatherName:       FatherName,
			MotherName:       MotherName,
			MotherMaidenName: MotherMaidenName,
			BloodGroup:       BloodGroup,
			MartialStatus:    MartialStatus,
			MotherTongue:     MotherTongue,
			Caste:            Caste,
			PersonalEmail:    PersonalEmail,
			Identification:   Identification,
			Disability:       Disability,
			PlaceOfBirth:     PlaceOfBirth,
			Heigth:           Heigth,
			Weight:           Weight,
			Religion:         Religion,
			Nationality:      Nationality,
			AdmissionDate:    AdmissionDate,
			MajorDegree:      MajorDegree,
			Reference:        Reference,
			Program:          Program,
			Regulation:       Regulation,
			Pending:          &Pending,
		}

		collection := client.Database("gomongo").Collection("student")

		result, err := collection.InsertOne(context.TODO(), newStudent)

		if err != nil {
			msg := fmt.Sprintf("student was not added")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)

	})

	router.GET("/studentDetails/:collegeid", func(c *gin.Context) {
		collegeid := c.Param("collegeid")

		student := Student{}
		// address := Address{}
		contact := Contacts{}
		// identity := Identity{}
		// qualification := Qualifications{}

		collection := client.Database("gomongo").Collection("student")
		// collection2 := client.Database("gomongo").Collection("address")
		collection3 := client.Database("gomongo").Collection("contacts")
		// collection4 := client.Database("gomongo").Collection("identity")
		// collection5 := client.Database("gomongo").Collection("qualifications")

		err := collection.FindOne(context.TODO(), bson.M{"collegeid": collegeid}).Decode(&student)
		// err2 := collection2.FindOne(context.TODO(), bson.M{"collegeid": collegeid}).Decode(&address)
		err3 := collection3.FindOne(context.TODO(), bson.M{"collegeid": collegeid}).Decode(&contact)
		// err4 := collection4.FindOne(context.TODO(), bson.M{"collegeid": collegeid}).Decode(&identity)
		// err5 := collection5.FindOne(context.TODO(), bson.M{"collegeid": collegeid}).Decode(&qualification)

		if err != nil {
			msg := fmt.Sprintf("student not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		if err3 != nil {
			fmt.Sprintf("details not found")
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, gin.H{
			"status":         http.StatusOK,
			"studentDetails": student,
			// "studentAddress":       address,
			"studentContact": contact,
			// "studentIdentity":      identity,
			// "studentQualification": qualification,
		})
	})

	router.POST("/changePassword/:collegeid", func(c *gin.Context) {
		var student Student

		isForgot := c.Query("forgot")

		collegeid := c.Param("collegeid")
		if err := c.BindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		password := student.Password

		if isForgot == "yes" {
			c.JSON(http.StatusOK, gin.H{
				"message": "under process need to send a verification email",
			})
		} else {
			newData := bson.M{
				"$set": bson.M{
					"password": password,
					"pending":  false,
				},
			}
			collection := client.Database("gomongo").Collection("student")

			_, err := collection.UpdateOne(context.TODO(), bson.M{"collegeid": collegeid}, newData)

			if err != nil {
				log.Printf("Error, Reason: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  500,
					"message": "Something went wrong",
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "password successfully updated",
			})
		}
	})

	//address routes
	router.POST("/addaddress/:stuId", func(c *gin.Context) {
		stuId := c.Param("stuId")
		var address Address

		if err := c.BindJSON(&address); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CollegeID := stuId
		AddressType := address.AddressType
		Doorno := address.Doorno
		Street := address.Street
		Landmark := address.Landmark
		City := address.City
		District := address.District
		State := address.State
		Country := address.Country

		newAddess := Address{
			CollegeID:   &CollegeID,
			AddressType: AddressType,
			Doorno:      Doorno,
			Street:      Street,
			Landmark:    Landmark,
			City:        City,
			District:    District,
			State:       State,
			Country:     Country,
		}

		collection := client.Database("gomongo").Collection("address")

		result, err := collection.InsertOne(context.TODO(), newAddess)

		if err != nil {
			msg := fmt.Sprintf("contact of student was not added")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	// contact routes
	router.POST("/addcontact/:stuId", func(c *gin.Context) {
		stuId := c.Param("stuId")
		var contact Contacts

		if err := c.BindJSON(&contact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CollegeID := stuId
		ContactType := contact.ContactType
		ContactPerson := contact.ContactPerson
		PhoneNumber := contact.PhoneNumber

		newContact := Contacts{
			CollegeID:     &CollegeID,
			ContactType:   ContactType,
			ContactPerson: ContactPerson,
			PhoneNumber:   PhoneNumber,
		}

		collection := client.Database("gomongo").Collection("contacts")

		result, err := collection.InsertOne(context.TODO(), newContact)

		if err != nil {
			msg := fmt.Sprintf("contact of student was not added")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	//identity routes
	router.POST("/additentity/:stuId", func(c *gin.Context) {
		stuId := c.Param("stuId")
		var identity Identity

		if err := c.BindJSON(&identity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CollegeID := stuId
		IdentityType := identity.IdentityType
		IdentityNumber := identity.IdentityNumber
		IssuedOn := identity.IssuedOn
		DateOfExpiry := identity.DateOfExpiry
		PlaceOfIssue := identity.PlaceOfIssue

		newIdentity := Identity{
			CollegeID:      &CollegeID,
			IdentityType:   IdentityType,
			IdentityNumber: IdentityNumber,
			IssuedOn:       IssuedOn,
			DateOfExpiry:   DateOfExpiry,
			PlaceOfIssue:   PlaceOfIssue,
		}

		collection := client.Database("gomongo").Collection("identity")

		result, err := collection.InsertOne(context.TODO(), newIdentity)

		if err != nil {
			msg := fmt.Sprintf("identity of student was not added")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	//qualification routes
	router.POST("/addqualification/:stuId", func(c *gin.Context) {
		stuId := c.Param("stuId")
		var qualifications Qualifications

		if err := c.BindJSON(&qualifications); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CollegeID := stuId
		Qualification := qualifications.Qualification
		Board := qualifications.Board
		EduName := qualifications.EduName
		CGPA := qualifications.CGPA
		YearOfPassing := qualifications.YearOfPassing
		Specilization := qualifications.Specilization

		newQualifications := Qualifications{
			CollegeID:     &CollegeID,
			Qualification: Qualification,
			Board:         Board,
			EduName:       EduName,
			CGPA:          CGPA,
			YearOfPassing: YearOfPassing,
			Specilization: Specilization,
		}

		collection := client.Database("gomongo").Collection("qualifications")

		result, err := collection.InsertOne(context.TODO(), newQualifications)

		if err != nil {
			msg := fmt.Sprintf("qualifications of student was not added")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	//course router
	router.POST("/addcourse/:stuId", func(c *gin.Context) {
		stuId := c.Param("stuId")
		var courses Courses

		if err := c.BindJSON(&courses); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		CollegeID := stuId
		Year := courses.Year
		AcademicYear := courses.AcademicYear
		Semester := courses.Semester
		CourseCode := courses.CourseCode
		CourseDesc := courses.CourseDesc
		LTPS := courses.LTPS
		Section := courses.Section
		FacultyName := courses.FacultyName

		newCourse := Courses{
			CollegeID:    &CollegeID,
			Year:         Year,
			AcademicYear: AcademicYear,
			Semester:     Semester,
			CourseCode:   CourseCode,
			CourseDesc:   CourseDesc,
			LTPS:         LTPS,
			Section:      Section,
			FacultyName:  FacultyName,
		}

		collection := client.Database("gomongo").Collection("courses")

		result, err := collection.InsertOne(context.TODO(), newCourse)

		if err != nil {
			msg := fmt.Sprintf("course of student was not added")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	router.Run(":" + port)
}

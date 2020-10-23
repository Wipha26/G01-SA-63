package main

import (

	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wipha26/app/controllers"
	_ "github.com/wipha26/app/docs"
	"github.com/wipha26/app/ent"
)

type Patients struct {
	Patient []Patient
}

type Patient struct {
	Firstname string
	Lastname  string
	Cardid    string
	Allergic  string
	Age       int
}

type Users struct {
	User []User
}

type User struct {
	Email    string
	Name     string
	Password string
}

type Drugs struct {
	Drug []Drug
}

type Drug struct {
	Name string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default(
		
	))

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDispenseController(v1, client)
	controllers.NewDrugController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewUserController(v1, client)

	// // Set User
	// users := Users{
	// 	User: []User{
	// 		User{"สมหมาย ใจเกินร้อย", "sommai@gmail.com", "1212312121"},
	// 	},
	// }

	// for _, u := range users.User {
	// 	client.User.
	// 		Create().
	// 		SetName(u.Name).
	// 		SetEmail(u.Email).
	// 		SetPassword(u.Password).
	// 		Save(context.Background())
	// }

	// Set drug
	drugs := Drugs{
		Drug: []Drug{
			Drug{"Abacavir (อะบาคาเวียร์)"},
			Drug{"Acetazolamide (อะเซตาโซลาไมด์)"},
			Drug{"Botulinum Toxin (โบทูไลนัม ท็อกซิน)"},
			Drug{"Budesonide (บูเดโซไนด์)"},

		},
	}

	for _, d := range drugs.Drug {
		client.Drug.
			Create().
			SetName(d.Name).
			Save(context.Background())
	}

	// Set Patient
	patients := Patients{
		Patient: []Patient{
			Patient{"นางสาววิภา", "ขุนหมื่น", "1119500006003","ไม่มี", 10},
			Patient{"นายเพชร", "รูปสุข", "1111111111111","แพ้ยาพารา", 20},
		},
	}

	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetFirstname(p.Firstname).
			SetLastname(p.Lastname).
			SetCardid(p.Cardid).
			SetAllergic(p.Allergic).
			SetAge(p.Age).
			Save(context.Background())
	}

	// Set User
	users := Users{
		User: []User{
			User{"pooh@gmail.com", "นายสุขดี  มีชัย", "00000000"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.Email).
			SetName(u.Name).
			SetPassword(u.Password).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}

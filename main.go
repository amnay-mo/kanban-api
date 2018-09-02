package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/amnay-mo/kanban-api/controller"
	"github.com/amnay-mo/kanban-api/datastore"
	"github.com/amnay-mo/kanban-api/middleware"
	"github.com/amnay-mo/kanban-api/model"

	"github.com/julienschmidt/httprouter"
)

type app struct {
	router *httprouter.Router
	port   int
}

func setup() *app {

	// setup datastore
	mongoHost := os.Getenv("MONGO_HOST")
	if mongoHost == "" {
		mongoHost = "127.0.0.1:27017"
	}
	log.Printf("Using Mongo %s", mongoHost)
	ds, _ := datastore.NewMongoTasks("mongodb://" + mongoHost)
	model.SetDatastore(ds)

	// setup jwt secret key
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "thisisbad"
	}
	middleware.SetSecretKey(secretKey)

	// setup app
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "9000"
	}
	a := new(app)
	a.port, _ = strconv.Atoi(appPort)
	a.router = httprouter.New()
	a.router.GET("/api/v1/tasks", middleware.AuthMiddleware(controller.GetTasks))
	a.router.POST("/api/v1/tasks", middleware.AuthMiddleware(controller.AddTask))
	a.router.DELETE("/api/v1/tasks/:task_id", middleware.AuthMiddleware(controller.DeleteTask))
	a.router.PATCH("/api/v1/tasks/:task_id", middleware.AuthMiddleware(controller.UpdateTask))
	a.router.POST("/api/v1/signup", controller.SignUp)
	a.router.POST("/api/v1/auth", controller.Authenticate)
	return a
}

func (a *app) run() error {
	log.Printf("Listening on port %d", a.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", a.port), middleware.CORSMiddleware{Next: middleware.LoggerMiddleware{Next: a.router}})
}

func main() {
	app := setup()

	err := app.run()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"GinLearning/database"
	"GinLearning/internal"
	"GinLearning/middleware"
	"GinLearning/structs"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gopkg.in/olahol/melody.v1"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DbConfig database.Config `yaml:"database"`
}

func setupLogging() {
	f, _ := os.Create("middleware/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogging()

	config := Config{}
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(b, &config)
	dc := config.DbConfig

	router := gin.Default()

	/////////////////////////////////
	// 實作匿名聊天室
	m := melody.New()
	router.LoadHTMLGlob("template/html/*")
	router.Static("/assets", "./template/assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleConnect(func(session *melody.Session) {
		id := session.Request.URL.Query().Get("id")
		m.Broadcast(structs.NewMessage("other", id, "加入聊天室").GetByteMessage())
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})
	m.HandleClose(func(session *melody.Session, i int, s string) error {
		id := session.Request.URL.Query().Get("id")
		m.Broadcast(structs.NewMessage("other", id, "離開聊天室").GetByteMessage())
		return nil
	})
	/////////////////////////////////

	// 註冊 Validator 的 Func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middleware.IsValidPassword)
		v.RegisterStructValidation(middleware.UserList, structs.Users{})
	}

	router.Use(gin.Recovery(), middleware.Logger())
	v1 := router.Group("/v1")
	internal.AddUserRouter(v1)
	go database.ConnectMySQL(dc.UserName, dc.Password, dc.Server, dc.Port, dc.DatabaseName)
	go database.ConnectMongo()

	router.Run(":8080")
}

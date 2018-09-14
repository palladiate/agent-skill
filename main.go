package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.webapps.rr.com/nspangler/agentSkill/docs"
	"github.webapps.rr.com/nspangler/agentSkill/models"
)

// Peek Agent Info Service

// @version 2.0
// @description Service for return Agent login information for Avaya
// @termsOfService

// @contact.name Nicholas Spangler
// @contact.url
// @contact.email nicholas.spangler@charter.com

// @host localhost:9005
// @BasePath /peek

var ds models.Datastore

const CONN = "nspangler:d00dsp34k@tcp(cmspapp03.twcable.com:3306)/ossi"

func main() {
	ds.Init(CONN)
	defer ds.Close()

	r := InitRouter()
	r.LoadHTMLGlob("templates/*")

	loginService := r.Group("/peek")
	{
		loginService.GET("/",Index)
		loginService.Static("/assets", "./assets")

		v2 := loginService.Group("/v2")
		{

			v2.GET("/switches", ReturnSwitches)
			v2.GET("/switches/agents", ReturnAgents)
			v2.GET("/login/:login", SearchLogin)

		}

		// set Swagger route
		loginService.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}


	r.Run(":9005")
}

func Index(c *gin.Context) {
	c.HTML(200, "index",
		gin.H{
			"title": "Peek Agent Information Service",
		},
	)
}

func InitRouter() (r *gin.Engine) {
	r = gin.Default()
	return
}

func ReturnSwitches(c *gin.Context) {
	c.JSON(200, ds.Switches() )
}

func ReturnAgents(c *gin.Context) {

	c.JSON(200, nil)
}


// Search Login Service

// @Summary Finds a PID and returns the full account detail object
// @Description This service provides full Avaya switch and skill information when provided a PID.
// @Accept json
// @Product json
// @Param login path string true "PID"
// @Success 200 {object} models.Agent
// @Failure 404 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /v2/login/{login} [get]
func SearchLogin(c *gin.Context) {
	agent, err := ds.GetAgent(c.Param("login"))
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = ds.AgentInfo(&agent)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, agent)
	return
}

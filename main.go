package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.webapps.rr.com/nspangler/agentSkill/models"
)

var ds models.Datastore

const CONN = "nspangler:d00dsp34k@tcp(cmspapp03.twcable.com:3306)/ossi"

func main() {
	ds.Init(CONN)
	defer ds.Close()

	r := InitRouter()

	/*
	skillService := r.Group("/peek/v2/switch")
	{
		skillService.GET("/",ReturnSwitches)
		skillService.GET("/:switch",ReturnAgents)
		skillService.GET("/:switch/:agent",ReturnAgentInformation)
	}
	*/

	loginService := r.Group("/peek/v2/login")
	{
		loginService.GET("/:login", SearchLogin)
	}

	r.Run(":9005")
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
// @Description This service provides full Avaya switch and skill information when provided a
// Charter PID
// @Accept json
// @Product json
// @Param Login path string true "PID"
// @Success 200 {object} models.Agent
// @Failure 404 {object} error
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router / [GET]
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

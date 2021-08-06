package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"live-config/server/db"
	"live-config/server/domain"
	"live-config/server/redis"
	"net/http"
	"strconv"
)

type PropertyController struct {
	Db     *gorm.DB
	Broker *redis.MessageBroker
}

func New() PropertyController {
	dbInstance, err := db.Init()

	if err != nil {
		panic(err)
	}

	broker := &redis.MessageBroker{Redis: redis.Init()}

	return PropertyController{Db: dbInstance, Broker: broker}
}

func (p *PropertyController) GetByApplicationProfileAndLabel(c *gin.Context) {
	app := c.Param("application")
	profile := c.Param("profile")
	label := c.Param("label")

	var props []domain.Property
	s := p.Db.Select("*").Table("properties")
	s.Where("application = ? AND profile = ? AND label = ?", app, profile, label).Find(&props)

	if len(props) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": "Could not find any property"})
		return
	}

	var jsonResponse = map[string]interface{}{}

	for _, prop := range props {
		jsonResponse[prop.Key] = prop.Value
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (p *PropertyController) Create(c *gin.Context) {
	prop := new(domain.Property)
	err := c.BindJSON(&prop)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Could not map request body"})
		return
	}

	create := p.Db.Create(&prop)
	err = create.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Could not save property into DB"})
		return
	}

	p.Broker.PublishMessage(prop)

	c.JSON(http.StatusOK, prop)
}

func (p *PropertyController) Modify(c *gin.Context) {
	id := c.Param("id")
	prop := new(domain.Property)
	err := c.BindJSON(&prop)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Could not map request body"})
		return
	}

	parseUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Could not parse property Id"})
		return
	}

	prop.Id = domain.ReadOnlyId(parseUint)

	//TODO create a lock on the row to avoid concurrency issue
	create := p.Db.Save(&prop)
	err = create.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": "Could not update property into DB"})
		return
	}

	p.Broker.PublishMessage(prop)

	c.JSON(http.StatusOK, prop)
}

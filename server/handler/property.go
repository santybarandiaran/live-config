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
	Db *gorm.DB
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
		c.JSON(http.StatusNotFound, "Could not find any property")
		return
	}

	c.JSON(http.StatusOK, props)
}

func (p *PropertyController) Create(c *gin.Context) {
	prop := new (domain.Property)
	err := c.BindJSON(&prop)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Could not map request body")
		return
	}

	create := p.Db.Create(&prop)
	err = create.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not save property into DB")
		return
	}

	p.Broker.PublishMessage(prop)

	c.JSON(http.StatusOK, prop)
}

func (p *PropertyController) Modify(c *gin.Context) {
	id := c.Param("id")
	prop := new (domain.Property)
	err := c.BindJSON(&prop)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Could not map request body")
		return
	}

	parseUint, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not parse property Id")
		return
	}

	prop.Id = domain.ReadOnlyId(parseUint)

	//TODO create a lock on the row to avoid concurrency issue
	create := p.Db.Save(&prop)
	err = create.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not update property into DB")
		return
	}

	p.Broker.PublishMessage(prop)

	c.JSON(http.StatusOK, prop)
}

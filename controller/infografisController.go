package controller

import (
	"net/http"
	"portal-data-kalsel/model"
	"portal-data-kalsel/service"
	"portal-data-kalsel/util"
	"time"

	"github.com/gin-gonic/gin"
)

// Portal-Data-Kalsel godoc
// @Tags Infografis
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param parameter body model.InfografisView true "string"
// @Router /publikasi/infografis/create [post]
func InsertInfografisController(c *gin.Context) {
	var infografis model.Infografis
	if err := c.ShouldBindJSON(&infografis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	infografis.ID = util.GenerateID()
	infografis.CreatedAt = time.Now().UnixMilli()
	infografis.UpdatedAt = time.Now().UnixMilli()

	if err := service.InsertInfografis(infografis); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Message": "Infografis created"})
}

// Portal-Data-Kalsel godoc
// @Tags Infografis
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param id path string true "string"
// @Param parameter body model.InfografisView true "string"
// @Router /publikasi/infografis/update/{id} [put]
func UpdateInfografisController(c *gin.Context) {
	id := c.Param("id")

	var infografis model.Infografis
	if err := c.ShouldBindJSON(&infografis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	infografis.ID = id
	if err := service.UpdateInfografis(infografis); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Update succes"})
}

// Portal-Data-Kalsel
// @Tags Infografis
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param id path string true "string"
// @Router /publikasi/infografis/delete/{id} [delete]
func DeleteInfografisController(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteInfografis(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Delete succes"})
}

// Portal-Data-Kalsel godoc
// @Tags Infografis
// @Summary
// @Description
// @Produce json
// @Param id path string true "string"
// @Router /publikasi/infografis/get-one/{id} [get]
func GetInfografisController(c *gin.Context) {
	id := c.Param("id")

	infografis, err := service.GetInfografis(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, infografis)
}

// Portal-Data-Kalsel godoc
// @Tags Infografis
// @Summary
// @Description
// @Produce json
// @Param parameter body model.Request_Pagination true "string"
// @Router /publikasi/infografis/get-all [post]
func GetAllInfografisController(c *gin.Context) {
	var filter = model.Request_Pagination{}
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	infografiss, err := service.GetAllInfografis(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, infografiss)
}

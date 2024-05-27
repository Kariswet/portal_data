package controller

import (
	"log"
	"net/http"
	"portal-data-kalsel/model"
	"portal-data-kalsel/service"
	"portal-data-kalsel/util"
	"time"

	"github.com/gin-gonic/gin"
)

// Portal-Data-Kalsel godoc
// @Tags Artikel
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param parameter body model.ArtikelView true "string"
// @Router /publikasi/artikel/create [post]
func InsertArtikelController(c *gin.Context) {
	var artikel model.Artikel
	if err := c.ShouldBindJSON(&artikel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	artikel.Id = util.GenerateID()
	artikel.Tanggal_publikasi = time.Now().UnixMilli()
	artikel.Created_at = time.Now().UnixMilli()
	artikel.Updated_at = time.Now().UnixMilli()

	if err := service.InsertArtikel(artikel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Message": "Artikel created"})
}

// Portal-Data-Kalsel godoc
// @Tags Artikel
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param id path string true "string"
// @Param parameter body model.ArtikelView true "string"
// @Router /publikasi/artikel/update/{id} [put]
func UpdateArtikelController(c *gin.Context) {
	id := c.Param("id")

	var artikel model.Artikel
	if err := c.ShouldBindJSON(&artikel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	artikel.Id = id
	if err := service.UpdateArtikel(artikel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Update Succes"})
}

// Portal-Data-Kalsel godoc
// @Tags Artikel
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param id path string true "string"
// @Router /publikasi/artikel/delete/{id} [delete]
func DeleteArtikelController(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteArtikel(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Delete Succes"})
}

// Portal-Data-Kalsel godoc
// @Tags Artikel
// @Summary
// @Description
// @Produce json
// @Param id path string true "string"
// @Router /publikasi/artikel/get-one/{id} [get]
func GetArtikelController(c *gin.Context) {
	id := c.Param("id")

	log.Println("id = " + id)
	artikel, err := service.GetArtikel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artikel)
}

// Portal-Data-Kalsel godoc
// @Tags Artikel
// @Summary
// @Description
// @Produce json
// @Param parameter body model.Request_Pagination true "string"
// @Router /publikasi/artikel/get-all [post]
func GetAllArtikelController(c *gin.Context) {
	var filter model.Request_Pagination
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	artikels, err := service.GetAllArtikel(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, artikels)
}

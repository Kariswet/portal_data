package service

import (
	"fmt"
	"log"
	"portal-data-kalsel/database"
	"portal-data-kalsel/model"
)

func InsertArtikel(artikel model.Artikel) error {
	return database.DB.Create(&artikel).Error
}

func UpdateArtikel(artikel model.Artikel) error {
	if err := database.DB.Model(&model.Artikel{}).Where("id = ?", artikel.Id).Updates(&artikel).Error; err != nil {
		fmt.Println(err)
		return err
	}
	// result := database.DB.Table().Update()
	return nil
}

func DeleteArtikel(id string) error {
	return database.DB.Delete(&model.Artikel{}, "id = ?", id).Error
}

func GetArtikel(id string) (model.Artikel, error) {
	var artikel model.Artikel
	err := database.DB.First(&artikel, "id = ?", id).Error
	log.Println(artikel)
	return artikel, err
}

func GetAllArtikel(param model.Request_Pagination) (artikel []map[string]interface{}, err error) {
	query := "SELECT * FROM artikels "

	// order and oerder by
	if param.OrderBy != "" {
		query += fmt.Sprintf("ORDER BY %s %s", param.OrderBy, param.Order)
	}

	// Pagination
	offset := (param.Page - 1) * param.Size
	if param.Size != 0 {
		query += fmt.Sprintf("LIMIT %d OFFSET %d", param.Size, offset)
	}

	log.Println(query)
	if err := database.DB.Raw(query).Scan(&artikel).Error; err != nil {
		log.Println(err)
	}
	return
}

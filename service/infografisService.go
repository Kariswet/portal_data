package service

import (
	"fmt"
	"log"
	"portal-data-kalsel/database"
	"portal-data-kalsel/model"
)

func InsertInfografis(infografis model.Infografis) error {
	return database.DB.Create(&infografis).Error
}

func UpdateInfografis(infografis model.Infografis) error {
	if err := database.DB.Model(&model.Infografis{}).Where("id = ?", infografis.ID).Updates(infografis).Error; err != nil {
		return err
	}
	return nil
}

func DeleteInfografis(id string) error {
	return database.DB.Delete(&model.Infografis{}, "id = ?", id).Error
}

func GetInfografis(id string) (model.Infografis, error) {
	var infografis model.Infografis
	err := database.DB.First(&infografis, "id = ?", id).Error
	return infografis, err
}

func GetAllInfografis(param model.Request_Pagination) (infografis []model.Infografis, err error) {
	query := "select * from infografis "

	// order and oerder by
	if param.OrderBy != "" {
		query += fmt.Sprintf("ORDER BY %s %s", param.OrderBy, param.Order)
	}

	// Pagination
	offset := (param.Page - 1) * param.Size
	if param.Size != 0 {
		query += fmt.Sprintf("LIMIT %d OFFSET %d", param.Size, offset)
	}
	// var infografis []map[string]any
	if err := database.DB.Raw(query).Scan(&infografis).Error; err != nil {
		log.Println(err)
	}

	return
}

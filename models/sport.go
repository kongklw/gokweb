package models

import (
	"gokweb/dao"
)

type Sports struct {
	ID         uint
	Name       string
	Country    string
	Popularity int64
}

func (Sports) TableName() string {
	return "tb_sports"
}

func GetSportList() (Sports, error) {
	var sport Sports
	err := dao.Db.First(&sport).Error

	return sport, err

}

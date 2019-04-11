
package models

import (
	// "encoding/json"
	// "github.com/astaxie/beego/logs"
	// "github.com/jinzhu/gorm"
	// "strconv"
)

type Cinema struct {
	ID int `gorm:"column:id" json:"id"`
	AccountOwner        string `gorm:"column:accountOwner" json:"accountOwner"`
	Address       string `gorm:"column:address" json:"address"`
	Area  string `gorm:"column:area" json:"area"`
	BankAccount        string `gorm:"column:bank_account" json:"bank_account"`
	BankName       string `gorm:"column:bank_name" json:"bank_name"`
	CinemaId  string `gorm:"column:cinema_id" json:"cinema_id"`
	CinemaName string `gorm:"column:cinema_name" json:"cinema_name"`
	City         string `gorm:"column:city" json:"city"`
	ContactName         string `gorm:"column:contact_name" json:"contact_name"`
	ContactPhone         string `gorm:"column:contact_phone" json:"contact_phone"`
	Province         string `gorm:"column:province" json:"province"`
	Region         string `gorm:"column:region" json:"region"`
	MinPrice         string `gorm:"column:min_price" json:"min_price"`
	MinAudience         string `gorm:"column:min_audience" json:"min_audience"`
	TotalNumber         string `gorm:"column:total_number" json:"total_number"`
	State         string `gorm:"column:state" json:"state"`

}

func (cinema *Cinema) CinemasList() (*[]Cinema, error) {
	var cinemaList []Cinema

	err := db.Model(&cinemaList).Scan(&cinemaList).Error
	if err != nil {
		return nil, err
	}
	//Set(key, problemList, 3600)
	return &cinemaList, nil
}

func (cinema *Cinema) CinemaDetail() (*Cinema, error) {
	err := db.Model(&Cinema{}).Where(&cinema).Scan(&cinema).Error
	if err != nil {
		return nil, err
	}
	return cinema, nil
}

func GetCinemaDetail(pid int) (*Cinema, error) {
	var cinema Cinema 
	err := db.Model(&Cinema{}).Where(&cinema).Scan(&cinema).Error
	if err != nil {
		return nil, err
	}
	return &cinema, nil
}

func (cinema *Cinema) CreateAndUpdateCinema() error {
	return db.Model(&Cinema{}).Update(&cinema).Error
}
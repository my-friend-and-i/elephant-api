
package models

import (
	// "encoding/json"
	// "github.com/astaxie/beego/logs"
	// "github.com/jinzhu/gorm"
	// "strconv"
)

type Movie struct {
	ID int `gorm:"column:id" json:"id"`
	Actors        string `gorm:"column:actors" json:"actors"`
	Brief       string `gorm:"column:brief" json:"brief"`
	Caption  string `gorm:"column:caption" json:"caption"`
	DefaultCover        string `gorm:"column:default_cover" json:"default_cover"`
	Directors       string `gorm:"column:directors" json:"directors"`
	FundMiniFlag  string `gorm:"column:fund_mini_flag" json:"fund_mini_flag"`
	FundNumLimit string `gorm:"column:fund_num_limit" json:"fund_num_limit"`
	GuideContent         string `gorm:"column:guide_content" json:"guide_content"`
	GuideScore         string `gorm:"column:guide_score" json:"guide_score"`
	GuideTitle         string `gorm:"column:guide_title" json:"guide_title"`
	GuideUrl         string `gorm:"column:guide_url" json:"guide_url"`
	MovieId         string `gorm:"column:movie_id" json:"movie_id"`
	Plot         string `gorm:"column:plot" json:"plot"`
	Poster         string `gorm:"column:poster" json:"poster"`
	Rating         string `gorm:"column:rating" json:"rating"`
	Title         string `gorm:"column:title" json:"title"`
	TitleSlogan         string `gorm:"column:title_slogan" json:"title_slogan"`
	Trailer         string `gorm:"column:trailer" json:"trailer"`
	WantSees         string `gorm:"column:want_sees" json:"want_sees"`
	State         string `gorm:"column:state" json:"state"`
}

func (movie *Movie) MoviesList() (*[]Movie, error) {
	var movieList []Movie

	err := db.Model(&movieList).Scan(&movieList).Error
	if err != nil {
		return nil, err
	}
	//Set(key, problemList, 3600)
	return &movieList, nil
}

func (movie *Movie) MovieDetail() (*Movie, error) {
	err := db.Model(&Movie{}).Where(&movie).Scan(&movie).Error
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func GetMovieDetail(pid int) (*Movie, error) {
	var movie Movie
	err := db.Model(&Movie{}).Where(&movie).Scan(&movie).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (movie *Movie) CreateAndUpdateMovie() error {
	return db.Model(&Movie{}).Update(&movie).Error
}
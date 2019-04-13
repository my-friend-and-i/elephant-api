package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"elephant-api/common/errCode"
	"elephant-api/models"
	"strconv"
)

type CinemaRequest struct {
	AccountOwner string `form:"accountOwner" json:"accountOwner"`
	Address string `form:"address" json:"address"`
	Area string `form:"area" json:"area"`
	BankAccount string `form:"bankAccount" json:"bankAccount" `
	BankName string `form:"bankName" json:"bankName"`
	CinemaId int `form:"cinemaId" json:"cinemaId"`
	CinemaName string `form:"cinemaName" json:"cinemaName"`
	City string `form:"city" json:"city"`
	ContactName string `form:"contactName" json:"contactName"`
	ContactPhone string `form:"contactPhone" json:"contactPhone`
	Province string `form:"province" json:"province"`
	Region string `form:"region" json:"region"`
	MinPrice int `form:"minPrice" json:"minPrice`
	MinAudience int `form:"minAudience" json:"minAudience`
	TotalNumber int `form:"totalNumber" json:"totalNumber"`
	State int `form:"state" json:"state"`         
}

// @Summary Cinema List
// @Produce json
// @Router /api/v1/cinema/list [get]
func GetCinemaList(c *gin.Context) {
	var ciemaList models.Cinema
	data, err := ciemaList.GetCinemaList()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary Cinema Detail
// @Produce json
// @Param cinema_id query int true "cinema_id"
// @Router /api/v1/cinema/detail [get]
func GetCinemaDetail(c *gin.Context) {
	pid := c.Query("cinema_id")
	if pid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	cinemaID, err := strconv.Atoi(pid)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	cinema := models.Cinema{
		ID: cinemaID,
	}
	data, err := cinema.CinemaDetail()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}


func PostCreateCinema(c *gin.Context) {
	req := CinemaRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	cinema := models.Cinema{
		AccountOwner:        req.AccountOwner,
		Address:       req.Address,
		Area:  req.Area,
		BankAccount: req.BankAccount,
		BankName: req.BankName,
		CinemaId:       req.CinemaId,
		CinemaName:       req.CinemaName,
		City:       req.City,
		ContactName:       req.ContactName,
		ContactPhone:       req.ContactPhone,
		Region:       req.Region,
		Province:       req.Province,
		MinPrice:       req.MinPrice,
		MinAudience:       req.MinAudience,
		TotalNumber:       req.TotalNumber,
		State:       req.State,
	}

	fmt.Println(cinema)

	if err := cinema.CreateCinema(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func PostUpdateCinema(c *gin.Context) {
	req := CinemaRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	cinema := models.Cinema{
		ID:           req.ID,
		AccountOwner:        req.AccountOwner,
		Address:       req.Address,
		Area:  req.Area,
		BankAccount: req.BankAccount,
		BankName: req.BankName,
		CinemaId:       req.CinemaId,
		CinemaName:       req.CinemaName,
		City:       req.City,
		ContactName:       req.ContactName,
		ContactPhone:       req.ContactPhone,
		Region:       req.Region,
		Province:       req.Province,
		MinPrice:       req.MinPrice,
		MinAudience:       req.MinAudience,
		TotalNumber:       req.TotalNumber,
		State:       req.State,
	}

	fmt.Println(cinema)

	if err := cinema.UpdateCinema(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func PostDeleteCinema(c *gin.Context) {
	req := CinemaRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	cinema := models.Cinema{
		ID:           req.ID,
	}

	fmt.Println(movie)

	if err := movie.DeleteCinema(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}
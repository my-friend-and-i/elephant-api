package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"elephant-api/common/errCode"
	"elephant-api/models"
	"strconv"
)

type MovieRequest struct {
	Actors        string `form:"actors" json:"actors"`
	Brief       string `form:"brief" json:"brief" biding:"-"`
	Caption  string `form:"caption" json:"caption"`
	DefaultCover        string `form:"default_cover" json:"default_cover"`
	Directors       string `form:"directors" json:"directors"`
	FundMiniFlag  string `form:"fund_mini_flag" json:"fund_mini_flag"`
	FundNumLimit string `form:"fund_num_limit" json:"fund_num_limit"`
	GuideContent         string `form:"guide_content" json:"guide_content" biding:"-"`
	GuideScore         string `form:"guide_score" json:"guide_score" biding:"-"`
	GuideTitle         string `form:"guide_title" json:"guide_title" biding:"-"`
	GuideUrl         string `form:"guide_url" json:"guide_url" biding:"-"`
	MovieId         string `form:"movie_id" json:"movie_id" biding:"-"`
	Plot         string `form:"plot" json:"plot" biding:"-"`
	Poster         string `form:"poster" json:"poster" biding:"-"`
	Rating         string `form:"rating" json:"rating" biding:"-"`
	Title         string `form:"title" json:"title" biding:"-"`
	TitleSlogan         string `form:"title_slogan" json:"title_slogan" biding:"-"`
	Trailer         string `form:"trailer" json:"trailer" biding:"-"`
	WantSees         string `form:"want_sees" json:"want_sees" biding:"-"`
	State         string `form:"state" json:"state" biding:"-"`
}

// @Summary Movies List
// @Produce json
// @Router /api/v1/movie/list [get]
func GetMovieList(c *gin.Context) {
	var movieList models.Movie
	data, err := movieList.GetMovieList()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}

// @Summary Movie Detail
// @Produce json
// @Param movie_id query int true "movie_id"
// @Router /api/v1/movie/detail [get]
func GetMovieDetail(c *gin.Context) {
	pid := c.Query("movie_id")
	if pid == "" {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	movieID, err := strconv.Atoi(pid)
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	movie := models.Movie{
		ID: movieID,
	}
	data, err := movie.MovieDetail()
	if err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}
	Response(c, http.StatusOK, errCode.SUCCESS, data)
}


func PostCreateProblem(c *gin.Context) {
	req := MovieRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	movie := models.Movie{
		Actors:        req.Actors,
		Brief:       req.Brief,
		Caption:  req.Actors,
		DefaultCover: req.DefaultCover,
		Directors: req.Directors,
		FundMiniFlag:       req.FundMiniFlag,
		FundNumLimit:       req.FundNumLimit,
		GuideContent:       req.GuideContent,
		GuideScore:       req.GuideScore,
		GuideTitle:      req.GuideTitle,
		GuideUrl:       req.GuideUrl,
		MovieId:       req.MovieId,
		Plot:       req.Plot,
		Poster:       req.Poster,
		Rating:       req.Rating,
		Title:       req.Title,
		TitleSlogan:       req.TitleSlogan,
		Trailer:       req.Trailer,
		WantSees:       req.WantSees,
		State:       req.State,
	}

	fmt.Println(movie)

	if err := movie.CreateMovie(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func PostUpdateMovie(c *gin.Context) {
	req := MovieRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	movie := models.Movie{
		ID:           req.ID,
		Actors:        req.Actors,
		Brief:       req.Brief,
		Caption:  req.Actors,
		DefaultCover: req.DefaultCover,
		Directors: req.Directors,
		FundMiniFlag:       req.FundMiniFlag,
		FundNumLimit:       req.FundNumLimit,
		GuideContent:       req.GuideContent,
		GuideScore:       req.GuideScore,
		GuideTitle:      req.GuideTitle,
		GuideUrl:       req.GuideUrl,
		MovieId:       req.MovieId,
		Plot:       req.Plot,
		Poster:       req.Poster,
		Rating:       req.Rating,
		Title:       req.Title,
		TitleSlogan:       req.TitleSlogan,
		Trailer:       req.Trailer,
		WantSees:       req.WantSees,
		State:       req.State,
	}

	fmt.Println(movie)

	if err := movie.UpdateMovie(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}

func PostDeleteMovie(c *gin.Context) {
	req := MovieRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		Response(c, http.StatusBadRequest, errCode.BADREQUEST, nil)
		return
	}

	movie := models.Movie{
		ID:           req.ID,
	}

	fmt.Println(movie)

	if err := movie.DeleteMovie(); err != nil {
		Response(c, http.StatusInternalServerError, errCode.ERROR, nil)
		return
	}

	Response(c, http.StatusOK, errCode.SUCCESS, nil)
}
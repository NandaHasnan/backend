package controllers

import (
	"backend/lib"
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type task struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
}

type Movie struct {
	Id        int    `json:"id"`
	Title     string `json:"title" form:"title"`
	Image     string `json:"image" form:"image"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

// var Data []Movie = []Movie{
// 	{
// 		Id:        1,
// 		Title:     "Dilan 1990",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        2,
// 		Title:     "Habibie & Ainun",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        3,
// 		Title:     "Ayat-ayat Cinta",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        4,
// 		Title:     "Ada Apa dengan Cinta? 2",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/4.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        5,
// 		Title:     "THE GODFATHER",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/5.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        6,
// 		Title:     "THE SHAWSHANK REDEMPTION",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/6.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        7,
// 		Title:     "MEMORIES OF MURDER",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/7.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        8,
// 		Title:     "THE DAVINCI CODE",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/8.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        9,
// 		Title:     "PARASITE",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        10,
// 		Title:     "Dilan 1990",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        11,
// 		Title:     "Habibie & Ainun",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        12,
// 		Title:     "Ayat-ayat Cinta",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        13,
// 		Title:     "Ada Apa dengan Cinta? 2",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/4.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        14,
// 		Title:     "THE GODFATHER",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/5.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        15,
// 		Title:     "THE SHAWSHANK REDEMPTION",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/6.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        16,
// 		Title:     "MEMORIES OF MURDER",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/7.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        17,
// 		Title:     "THE DAVINCI CODE",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/8.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// 	{
// 		Id:        18,
// 		Title:     "PARASITE",
// 		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
// 		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
// 	},
// }

// func AllMovies(ctx *gin.Context) {

// 	title := ctx.Query("title")
// 	limitParams := ctx.Query("limit")
// 	pageParams := ctx.Query("page")

// 	page, _ := strconv.Atoi(pageParams)
// 	limit, _ := strconv.Atoi(limitParams)

// 	var filteredData []Movie
// 	if title != "" {
// 		for _, movie := range Data {
// 			if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(title)) {
// 				filteredData = append(filteredData, movie)
// 			}
// 		}

// 		if len(filteredData) == 0 {
// 			ctx.JSON(http.StatusNotFound, task{
// 				Success: false,
// 				Message: "Title not found",
// 			})
// 			return
// 		}
// 	} else {
// 		filteredData = Data
// 	}

// 	sort.Slice(filteredData, func(i, j int) bool {
// 		return strings.ToLower(filteredData[i].Title) < strings.ToLower(filteredData[j].Title)
// 	})

// 	if page <= 0 {
// 		page = 1
// 	}
// 	if limit <= 0 {
// 		limit = 10
// 	}
// 	offset := (page - 1) * limit

// 	if offset >= len(filteredData) {
// 		ctx.JSON(http.StatusOK, task{
// 			Success: false,
// 			Message: "Paging invalid",
// 		})
// 		return
// 	}

// 	end := offset + limit
// 	if end > len(filteredData) {
// 		end = len(filteredData)
// 	}

// 	pagedData := filteredData[offset:end]

// 	ctx.JSON(http.StatusOK, task{
// 		Success: true,
// 		Message: "Search and paging successful",
// 		Result:  pagedData,
// 	})
// }

// ctx.JSON(http.StatusOK, task{
// 	Success: true,
// 	Message: "all movie",
// 	Result:  Data,
// })

func IdMovies(ctx *gin.Context) {

	iddb, _ := strconv.Atoi(ctx.Param("id"))
	movie := models.MovieById(iddb)
	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "Detail Movie",
		Result:  movie,
	})

}

func AddMovies(ctx *gin.Context) {

	var newMovie models.Movie
	ctx.ShouldBind(&newMovie)
	f, _ := ctx.MultipartForm()
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse{
			Success: false,
			Message: "File size is large",
		})
		return
	}

	newMovie.Synopsis = f.Value["synopsis"][0]

	if file.Filename != "" {
		filename := uuid.New().String()
		splitFile := strings.Split(file.Filename, ".")
		filex := splitFile[len(splitFile)-1]
		// if filex != ".jpg" && filex != ".png" {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Only .jpg and .png files are allowed"})
		// 	return
		// }
		filestored := fmt.Sprintf("%s.%s", filename, filex)
		ctx.SaveUploadedFile(file, fmt.Sprintf("upload/movies/%s", filestored))
		newMovie.Image_movie = filestored
	}
	movie := models.InsertMovie(newMovie)

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "add movie sukses",
		Result:  movie,
	})

}

func OrderMovies(ctx *gin.Context) {

	var orderMovie models.Order
	ctx.ShouldBind(&orderMovie)

	order := models.OrderTicket(orderMovie)

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "add movie sukses",
		Result:  order,
	})

}

// func EditMovies(ctx *gin.Context) {

// 	iddb, _ := strconv.Atoi(ctx.Param("id"))
// 	movies := models.MovieById(iddb)
// 	if movies == (models.Movie{}) {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse{
// 			Success: false,
// 			Message: "invalid edit movie",
// 			Result:  iddb,
// 		})
// 		return
// 	}

// 	ctx.ShouldBind(&movies)

// 	f, _ := ctx.MultipartForm()
// 	file, err := ctx.FormFile("image")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse{
// 			Success: false,
// 			Message: "File size is large",
// 		})
// 		return
// 	}

// 	movies.Synopsis = f.Value["synopsis"][0]

// 	if file.Filename != "" {
// 		filename := uuid.New().String()
// 		splitFile := strings.Split(file.Filename, ".")
// 		filex := splitFile[len(splitFile)-1]
// 		if filex != ".jpg" && filex != ".png" {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Only .jpg and .png files are allowed"})
// 			return
// 		}
// 		filestored := fmt.Sprintf("%s.%s", filename, filex)
// 		ctx.SaveUploadedFile(file, fmt.Sprintf("upload/movies/%s", filestored))
// 		movies.Image = filestored
// 	}

// 	UpdateMovie := models.UpdateMovie(movies)

// 	ctx.JSON(http.StatusOK, TaskResponse{
// 		Success: true,
// 		Message: "Update Movie sukses",
// 		Result:  UpdateMovie,
// 	})

// }

// func DeleteMovies(ctx *gin.Context) {
// 	iddb, _ := strconv.Atoi(ctx.Param("id"))
// 	movie := models.MovieById(iddb)
// 	if movie == (models.Movie{}) {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse{
// 			Success: false,
// 			Message: "invalid delete movie",
// 			Result:  iddb,
// 		})
// 		return
// 	}

// 	ctx.ShouldBind(&movie)

// 	DeleteMovie := models.DeleteMovie(iddb)

// 	ctx.JSON(http.StatusOK, TaskResponse{
// 		Success: true,
// 		Message: "Delete Movie sukses",
// 		Result:  DeleteMovie,
// 	})
// }

// func SearchMovies(ctx *gin.Context) {
// 	var cari Movie

// 	ctx.ShouldBind(&cari)

// 	var movie Movie
// 	found := false
// 	// query :=
// 	for _, i := range Data {
// 		if strings.Contains(strings.ToLower(i.Title), strings.ToLower(cari.Title)) { //i.Title == cari.Title
// 			movie = i
// 			found = true
// 		}
// 	}

// 	if !found {
// 		ctx.JSON(http.StatusNotFound, task{
// 			Success: false,
// 			Message: "Title not found",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, task{
// 		Success: true,
// 		Message: "Search successful",
// 		Result:  movie,
// 	})
// }

// func Paging(ctx *gin.Context) {
// 	var paging Movie

// 	ctx.ShouldBind(&paging)

// 	title := ctx.Query("title")
// 	limitParams := ctx.Query("limit")
// 	pageParams := ctx.Query("page")

// 	page, _ := strconv.Atoi(pageParams)
// 	limit, _ := strconv.Atoi(limitParams)

// 	var filter []Movie
// 	// found := false
// 	if title != "" {
// 		for _, i := range Data {
// 			if strings.Contains(strings.ToLower(i.Title), strings.ToLower(title)) { //i.Title == cari.Title
// 				filter = append(filter, i)
// 			}
// 		}
// 	}

// 	if len(filter) == 0 {
// 		ctx.JSON(http.StatusNotFound, task{
// 			Success: false,
// 			Message: "Title not found",
// 		})
// 		return
// 	}

// 	offset := (page - 1) * limit

// 	if offset >= len(filter) {
// 		ctx.JSON(http.StatusNotFound, task{
// 			Success: false,
// 			Message: "All movie invalid",
// 		})
// 		return
// 	}

// 	end := offset + limit
// 	if end > len(filter) {
// 		end = len(filter)
// 	}

// 	pageData := filter[offset:end]

// 	ctx.JSON(http.StatusOK, task{
// 		Success: true,
// 		Message: "All movie",
// 		Result:  pageData,
// 	})
// }

func AllMovieDB(ctx *gin.Context) {
	pageParams, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limitParams, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	search := ctx.DefaultQuery("search", "")
	sort := ctx.DefaultQuery("sort", "ASC")

	if sort != "ASC" {
		sort = "DESC"
	}

	var movies models.Data2
	var count int

	get := lib.Redis().Get(context.Background(), ctx.Request.RequestURI)
	getCount := lib.Redis().Get(context.Background(), fmt.Sprintf("Count+%s", ctx.Request.RequestURI))

	if get.Val() != "" {
		dataRaw := []byte(get.Val())
		json.Unmarshal(dataRaw, &movies)
	} else {
		movies = models.MovieAll(pageParams, limitParams, search, sort)
		encoded, _ := json.Marshal(movies)
		lib.Redis().Set(context.Background(),
			ctx.Request.RequestURI, string(encoded),
			0,
		)
	}

	if getCount.Val() != "" {
		dataRaw := []byte(getCount.Val())
		json.Unmarshal(dataRaw, &count)
	} else {
		count = models.CountMovie(search)
		encoded, _ := json.Marshal(count)
		lib.Redis().Set(context.Background(),
			fmt.Sprintf("Count+%s", ctx.Request.RequestURI),
			string(encoded),
			0,
		)

	}

	totalPage := int(math.Ceil(float64(count) / float64(limitParams)))
	pageNext := pageParams + 1
	if pageNext > totalPage {
		pageNext = totalPage
	}

	pagePrev := pageParams - 1
	if pagePrev < 1 {
		pagePrev = 0
	}

	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "All Movie",
		PageInfo: PageInfo{
			TotalPage:   totalPage,
			TotalData:   count,
			NextPage:    limitParams,
			PrevPage:    pagePrev,
			CurrentPage: pageParams,
		},
		Result: movies,
	})

}

// (57, 1)
// (12, 2)
// (85, 3)
// (23, 4)
// (78, 5)
// (40, 6)
// (7, 1)
// (99, 2)
// (36, 3)
// (48, 4)
// (65, 5)
// (9, 6)
// (18, 2)
// (80, 3)
// (92, 1)
// (50, 5)
// (31, 2)
// (67, 3)
// (14, 2)
// (100, 2)
// (21, 1)
// (74, 6)
// (42, 4)
// (88, 5)
// (96, 4)
// (5, 3)
// (83, 2)
// (33, 2)
// (61, 2)
// (25, 3)
// (44, 5)
// (8, 6)
// (101, 2)
// (71, 4)
// (3, 5)
// (90, 4)
// (55, 3)
// (39, 3)
// (76, 3)
// (19, 2)
// (63, 2)
// (1, 3)
// (11, 3)
// (53, 5)
// (27, 3)
// (87, 4)
// (20, 4)
// (97, 4)
// (69, 3)
// (32, 2)
// (4, 2)
// (95, 1)
// (46, 5)
// (81, 2)
// (15, 5)
// (59, 3)
// (98, 2)
// (6, 2)
// (73, 2)
// (28, 2)
// (41, 4)
// (30, 4)
// (84, 4)
// (62, 5)
// (89, 3)
// (38, 2)
// (13, 1)
// (64, 1)
// (75, 5)
// (2, 5)
// (91, 3)
// (45, 2)
// (17, 1)
// (35, 3)
// (47, 5)
// (86, 6)
// (24, 3)
// (58, 6)
// (49, 6)
// (26, 5)
// (68, 3)
// (10, 5)
// (34, 3)
// (82, 2)
// (70, 5)
// (22, 6)
// (94, 4)
// (29, 3)
// (43, 2)
// (60, 2)
// (37, 5)
// (77, 6)
// (16, 5)
// (66, 4)
// (56, 2)
// (93, 4)
// (52, 4)
// (54, 4)
// (79, 4)
// (72, 4)
// (51, 5)

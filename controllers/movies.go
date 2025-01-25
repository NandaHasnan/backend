package controllers

import (
	"backend/lib"
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
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

// Movie godoc
// @Schemes Detail Movies
// @Description Get Detail Movie
// @Tags Movie
// @Accept json
// @Produce json
// @param id path int true "Detail Movies"
// @Success 200 {object} TaskResponse2{results=models.Movie_cinema}
// @Router /movies/{id} [get]
func IdMovies(ctx *gin.Context) {

	iddb, _ := strconv.Atoi(ctx.Param("id"))
	movie := models.MovieById(iddb)
	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Detail Movie",
		Result:  movie,
	})

}

// Movie godoc
// @Schemes
// @Description Create New Movie
// @Tags Movie
// @Accept mpfd
// @Produce json
// @param title formData string true "Title"
// @param image_movie formData file true "Image Movies"
// @param genre formData string true "Genre"
// @param release_date formData string true "Release Date"
// @param duration formData string true "Duration"
// @param director formData string true "Director"
// @param cast_actor formData string true "Cst Actor"
// @param synopsis formData string true "Synopsis"
// @Success 200 {object} TaskResponse2{result=models.Movie_body}
// @Security ApiKeyAuth
// @Router /movies/addmovie [post]
func AddMovies(ctx *gin.Context) {

	var newMovie models.Movie_body
	ctx.ShouldBind(&newMovie)
	// f, _ := ctx.MultipartForm()
	file, err := ctx.FormFile("image_movie")

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "No file provided",
		})
		return
	}

	maxFile := 2 * 1024 * 1024
	if file.Size > int64(maxFile) {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "File size is too large",
		})
		return
	}

	if file.Filename != "" {
		splitFile := strings.Split(file.Filename, ".")
		if len(splitFile) < 2 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid file format",
			})
			return
		}

		fileExtension := strings.ToLower(splitFile[len(splitFile)-1])

		allowedExtensions := map[string]bool{
			"jpg": true,
			"png": true,
		}

		if !allowedExtensions[fileExtension] {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Only .jpg and .png files are allowed",
			})
			return
		}

		filename := uuid.New().String()
		storedFile := fmt.Sprintf("%s.%s", filename, fileExtension)
		err := ctx.SaveUploadedFile(file, fmt.Sprintf("upload/movies/%s", storedFile))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, TaskResponse2{
				Success: false,
				Message: "Failed to save file",
			})
			return
		}

		newMovie.Image_movie = storedFile
	}

	movie, err := models.InsertMovie(newMovie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, task{
			Success: false,
			Message: err.Error(),
			Result:  nil,
		})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "add movie sukses",
		Result:  movie,
	})

}

// func OrderMovies(ctx *gin.Context) {

// 	var orderMovie models.OrderBody
// 	err := ctx.ShouldBind(&orderMovie)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	order, err := models.OrderTicket(orderMovie)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, task{
// 		Success: true,
// 		Message: "Order tiket sukses",
// 		Result:  order,
// 	})

// }

// Movie godoc
// @Schemes
// @Description Update Movie
// @Tags Movie
// @Accept mpfd
// @Produce json
// @param id path int true "Id Movies"
// @param title formData string true "Title"
// @param image_movie formData file true "Image Movies"
// @param genre formData string true "Genre"
// @param release_date formData string true "Release Date"
// @param duration formData string true "Duration"
// @param director formData string true "Director"
// @param cast_actor formData string true "Cst Actor"
// @param synopsis formData string true "Synopsis"
// @Success 200 {object} TaskResponse2{result=models.Movie_body}
// @Security ApiKeyAuth
// @Router /movies/{id} [patch]
func EditMovies(ctx *gin.Context) {
	iddb, _ := strconv.Atoi(ctx.Param("id"))
	_, err := models.MovieById2(iddb)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	var moviesBody models.Movie_body

	if err := ctx.ShouldBind(&moviesBody); err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "Invalid input data",
		})
		return
	}

	moviesBody.Id = iddb

	file, err := ctx.FormFile("image_movie")

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "No file provided",
		})
		return
	}

	maxFile := 2 * 1024 * 1024
	if file.Size > int64(maxFile) {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "File size is too large, max 2 mb",
		})
		return
	}

	if file.Filename != "" {
		splitFile := strings.Split(file.Filename, ".")
		if len(splitFile) < 2 {
			ctx.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "Invalid file format",
			})
			return
		}

		fileExtension := strings.ToLower(splitFile[len(splitFile)-1])

		allowedExtensions := map[string]bool{
			"jpg": true,
			"png": true,
		}

		if !allowedExtensions[fileExtension] {
			ctx.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "Only .jpg and .png files are allowed",
			})
			return
		}

		filename := uuid.New().String()
		storedFile := fmt.Sprintf("%s.%s", filename, fileExtension)
		err := ctx.SaveUploadedFile(file, fmt.Sprintf("upload/movies/%s", storedFile))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "Failed to save file",
			})
			return
		}

		moviesBody.Image_movie = storedFile
	}

	updatedMovie, err := models.Update(moviesBody)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Movie updated successfully",
		Result:  updatedMovie,
	})
}

// Movie godoc
// @Schemes
// @Description Delete Movie
// @Tags Movie
// @Accept json
// @Produce json
// @param id path int true "Id Movies"
// @Success 200 {object} TaskResponse2{result=models.Movie}
// @Security ApiKeyAuth
// @Router /movies/{id} [delete]
func DeleteMovies(ctx *gin.Context) {
	iddb, _ := strconv.Atoi(ctx.Param("id"))
	movie, _ := models.MovieById2(iddb)
	if movie == (models.Movie{}) {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "invalid delete movie",
			Result:  iddb,
		})
		return
	}

	ctx.ShouldBind(&movie)

	DeleteMovie := models.DeleteMovie(iddb)

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Delete Movie sukses",
		Result:  DeleteMovie,
	})
}

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

// Movie godoc
// @Schemes
// @Description Get All Movie
// @Tags Movie
// @Accept json
// @Produce json
// @param search query string false "Search Movies"
// @param page query int false "Page Movies"
// @param limit query int false "Limit Movies"
// @param sort query string false "Sort Movies"
// @Success 200 {object} TaskResponse{result=models.Allmovie}
// @Router /movies [get]
func AllMovieDB(ctx *gin.Context) {
	pageParams, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limitParams, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		log.Println(err)
	}
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

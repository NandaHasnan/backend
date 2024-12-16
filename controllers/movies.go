package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type task struct {
	Success bool   //`json:"success"`
	Message string //`json:"message"`
	Result  any    //`json:"result"`
}

type Movie struct {
	Id        int    `json:"id"`
	Title     string `json:"title" form:"title"`
	Image     string `json:"image" form:"image"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

var Data []Movie = []Movie{
	{
		Id:        1,
		Title:     "Dilan 1990",
		Image:     "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        2,
		Title:     "Habibie & Ainun",
		Image:     "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        3,
		Title:     "Ayat-ayat Cinta",
		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        4,
		Title:     "Ada Apa dengan Cinta? 2",
		Image:     "https://rickandmortyapi.com/api/character/avatar/4.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        5,
		Title:     "THE GODFATHER",
		Image:     "https://rickandmortyapi.com/api/character/avatar/5.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        6,
		Title:     "THE SHAWSHANK REDEMPTION",
		Image:     "https://rickandmortyapi.com/api/character/avatar/6.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        7,
		Title:     "MEMORIES OF MURDER",
		Image:     "https://rickandmortyapi.com/api/character/avatar/7.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        8,
		Title:     "THE DAVINCI CODE",
		Image:     "https://rickandmortyapi.com/api/character/avatar/8.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        9,
		Title:     "PARASITE",
		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        10,
		Title:     "Dilan 1990",
		Image:     "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        11,
		Title:     "Habibie & Ainun",
		Image:     "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        12,
		Title:     "Ayat-ayat Cinta",
		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        13,
		Title:     "Ada Apa dengan Cinta? 2",
		Image:     "https://rickandmortyapi.com/api/character/avatar/4.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        14,
		Title:     "THE GODFATHER",
		Image:     "https://rickandmortyapi.com/api/character/avatar/5.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        15,
		Title:     "THE SHAWSHANK REDEMPTION",
		Image:     "https://rickandmortyapi.com/api/character/avatar/6.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        16,
		Title:     "MEMORIES OF MURDER",
		Image:     "https://rickandmortyapi.com/api/character/avatar/7.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        17,
		Title:     "THE DAVINCI CODE",
		Image:     "https://rickandmortyapi.com/api/character/avatar/8.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        18,
		Title:     "PARASITE",
		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
}

func AllMovies(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "all movie",
		Result:  Data,
	})

}

func IdMovies(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	data := Data
	if err != nil {
		ctx.JSON(http.StatusBadRequest, task{
			Success: false,
			Message: "invalid all movie",
			Result:  id,
		})
		return
	}

	var movie Movie

	for _, data := range data {
		if data.Id == id {
			movie = data
		}
	}

	if movie == (Movie{}) {
		ctx.JSON(http.StatusNotFound, task{
			Success: false,
			Message: "all movie not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "all movie",
		Result:  movie,
	})

}

func AddMovies(ctx *gin.Context) {

	var newMovie Movie

	if err := ctx.ShouldBind(&newMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMovie.Id = len(Data) + 1
	Data = append(Data, newMovie)

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "add movie sukses",
		Result:  newMovie,
	})

}

func EditMovies(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	// data := Data
	if err != nil {
		ctx.JSON(http.StatusBadRequest, task{
			Success: false,
			Message: "invalid all movie",
			Result:  id,
		})
		return
	}

	var updateMovie Movie
	if err := ctx.ShouldBind(&updateMovie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, movie := range Data {
		if movie.Id == id {
			if updateMovie.Title != "" {
				Data[i].Title = updateMovie.Title
			}
			if updateMovie.Image != "" {
				Data[i].Image = updateMovie.Image
			}
			if updateMovie.Deskripsi != "" {
				Data[i].Deskripsi = updateMovie.Deskripsi
			}
		}

		ctx.JSON(http.StatusOK, task{
			Success: true,
			Message: "Update movie sukses",
			Result:  Data[i],
		})
	}

}

func DeleteMovies(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	// data := Data
	if err != nil {
		ctx.JSON(http.StatusBadRequest, task{
			Success: false,
			Message: "invalid all movie",
			Result:  id,
		})
		return
	}

	for i, movie := range Data {
		if movie.Id == id {
			Data = append(Data[:i], Data[i+1:]...)
		}
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "Delete movie sukses",
		Result:  Data,
	})

}

func SearchMovies(ctx *gin.Context) {
	var cari Movie

	ctx.ShouldBind(&cari)

	var movie Movie
	found := false
	// query :=
	for _, i := range Data {
		if strings.Contains(strings.ToLower(i.Title), strings.ToLower(cari.Title)) { //i.Title == cari.Title
			movie = i
			found = true
		}
	}

	if !found {
		ctx.JSON(http.StatusNotFound, task{
			Success: false,
			Message: "Title not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "Search successful",
		Result:  movie,
	})
}

func Paging(ctx *gin.Context) {
	var paging Movie

	ctx.ShouldBind(&paging)

	limitParams := ctx.DefaultQuery("limit", "5")
	pageParams := ctx.DefaultQuery("page", "1")

	page, _ := strconv.Atoi(pageParams)
	limit, _ := strconv.Atoi(limitParams)

	offset := (page - 1) * limit

	if offset > len(Data) {
		ctx.JSON(http.StatusOK, task{
			Success: false,
			Message: "Paging invalid",
		})
		return
	}

	end := offset + limit
	if end > len(Data) {
		end = len(Data)
	}

	pageLimit := Data[offset:end]

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "Paging successful",
		Result:  pageLimit,
	})
}

// func Sort(ctx *gin.Context) {
// 	sortedMovies := make([]Movie, len(Data))
// 	copy(sortedMovies, Data)

// 	sort.Slice(sortedMovies, func(i, j int) bool {
// 		return strings.ToLower(sortedMovies[i].Title) < strings.ToLower(sortedMovies[j].Title)
// 	})

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "Movies sorted by title",
// 		"result":  sortedMovies,
// 	})
// }

package movies

import (
	"net/http"
	"strconv"

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
		Title:     "image 1",
		Image:     "https://rickandmortyapi.com/api/character/avatar/1.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        2,
		Title:     "image 2",
		Image:     "https://rickandmortyapi.com/api/character/avatar/2.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        3,
		Title:     "image 3",
		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        4,
		Title:     "image 4",
		Image:     "https://rickandmortyapi.com/api/character/avatar/4.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        5,
		Title:     "image 5",
		Image:     "https://rickandmortyapi.com/api/character/avatar/5.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        6,
		Title:     "image 6",
		Image:     "https://rickandmortyapi.com/api/character/avatar/6.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        7,
		Title:     "image 7",
		Image:     "https://rickandmortyapi.com/api/character/avatar/7.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        8,
		Title:     "image 8",
		Image:     "https://rickandmortyapi.com/api/character/avatar/8.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
	{
		Id:        9,
		Title:     "image 9",
		Image:     "https://rickandmortyapi.com/api/character/avatar/3.jpeg",
		Deskripsi: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	},
}

func Movies() {
	router := gin.Default()

	router.GET("/movies", func(c *gin.Context) {
		c.JSON(http.StatusOK, task{
			Success: true,
			Message: "all image",
			Result:  Data,
		})
	})

	// router.GET("/auth", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, task2{
	// 		Message: "all users",
	// 		Result:  Data2,
	// 	})
	// })

	router.GET("/movies/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		data := Data
		if err != nil {
			ctx.JSON(http.StatusBadRequest, task{
				Success: false,
				Message: "invalid all image",
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
				Message: "all image not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, task{
			Success: true,
			Message: "all image",
			Result:  movie,
		})
	})

	// router.GET("/movies/search", func(c *gin.Context) {
	// 	title := c.Query("title")
	// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	// 	if page < 1 {
	// 		page = 1
	// 	}
	// 	if limit < 1 {
	// 		limit = 5
	// 	}

	// 	var searchTitle []Movie
	// 	for _, movie := range Data {
	// 		if title == "" || strings.Contains(movie.Title, title) {
	// 			searchTitle = append(searchTitle, movie)
	// 		}
	// 	}

	// 	start = (page-1) * limit
	// 	end = page + limit

	// 	if start

	// })

	//string.contain

	router.GET("/movie", func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
		for i, movie := range Data {
			if movie.Id == page {
				Data = append(Data[:i], Data[i+1:]...)
			}
		}
		for i, movie := range Data {
			if movie.Id == limit {
				Data = append(Data[:i], Data[i+1:]...)
			}
		}
	})

	router.POST("/movies", func(ctx *gin.Context) {
		var newMovie Movie

		if err := ctx.ShouldBindJSON(&newMovie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newMovie.Id = len(Data) + 1
		Data = append(Data, newMovie)

		ctx.JSON(http.StatusOK, task{
			Success: true,
			Message: "add image sukses",
			Result:  newMovie,
		})
	})

	router.PATCH("/movies/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		// data := Data
		if err != nil {
			ctx.JSON(http.StatusBadRequest, task{
				Success: false,
				Message: "invalid all image",
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
				Message: "Update image sukses",
				Result:  Data[i],
			})
		}
	})

	router.DELETE("/movies/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		// data := Data
		if err != nil {
			ctx.JSON(http.StatusBadRequest, task{
				Success: false,
				Message: "invalid all image",
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
			Message: "Delete image sukses",
			Result:  Data,
		})
	})

	router.Run("localhost:8888")
}

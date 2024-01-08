package handler

import (
	"fmt"
	"net/http"
	"scrape-film/helper"
	"scrape-film/model"
	"scrape-film/repository"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
)

type MovieHandler struct {
	MovieRepository repository.MovieRepository
}

func NewMovieHandler(movieRep repository.MovieRepository) *MovieHandler {
	return &MovieHandler{movieRep}
}

func (mh *MovieHandler) StartScrape(c *gin.Context) {
	browser := rod.New().MustConnect().NoDefaultDevice()
	page := browser.MustPage("https://tv7.lk21official.wiki/").MustWindowFullscreen()

	// Menunggu dan mengklik elemen untuk memuat data
	page.MustElement("body > main > section > div > section > footer > a").MustClick()

	time.Sleep(2 * time.Second)

	// Menunggu dan mendapatkan elemen data
	data := page.MustElement("#grid-wrapper").MustWaitLoad()

	// Ambil semua elemen dengan kelas "grid-title" di dalam elemen data
	cardFilm := data.MustElements(".mega-item")

	// Inisialisasi slice untuk menyimpan judul dan isi
	var movies []model.Movie

	for _, film := range cardFilm {
		var movieData model.Movie

		// ratingParent := film.MustParent().MustParent().MustElement(".grid-meta")

		// Ambil teks judul
		movieData.Title = film.MustElement(".grid-header .grid-title").MustText()
		// Ambil teks dari elemen-elemen terkait
		ratingElems := film.MustElements(".grid-meta .rating")

		// Periksa apakah ada elemen rating sebelum mencoba mengambil teksnya
		if len(ratingElems) > 0 {
			rating := ratingElems[0].MustText()
			movieData.Rating = rating
		} else {
			// Atur nilai rating menjadi string kosong atau nilai default
			movieData.Rating = ""
		}

		// Inisialisasi slice untuk menyimpan kategori
		var genres []string

		gridAction := film.MustElement(".grid-action")

		if gridAction != nil {
			categoriesElem := gridAction.MustElement(".grid-categories")

			categoryElems := categoriesElem.MustElements("a")

			// Iterasi melalui setiap elemen <a> dan ambil teksnya
			for _, categoryElem := range categoryElems {
				genre := categoryElem.MustText()
				genres = append(genres, genre)
			}
		}

		movieData.Genre = genres

		// Ambil elemen <figure> yang berisi gambar
		figureElem := film.MustElement(".grid-poster")

		// Periksa apakah elemen <figure> ditemukan
		if figureElem != nil {
			// Ambil elemen <source> di dalam elemen <figure>
			sourceElem := figureElem.MustElement("source")

			// Periksa apakah elemen <source> ditemukan
			if sourceElem != nil {
				// Ambil nilai atribut srcset dari elemen <source>
				srcset, err := sourceElem.Attribute("srcset")
				if err == nil {
					// Pilih URL yang diinginkan dari srcset
					imageURL := helper.ExtractWebPURL(*srcset)

					if !strings.HasPrefix(imageURL, "https:") {
						imageURL = "https:" + imageURL
					}

					// Set nilai image dalam map movieData
					movieData.URLImage = imageURL
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Elemen <source> tidak ditemukan.")
			}
		} else {
			fmt.Println("Elemen <figure> tidak ditemukan.")
		}

		movieData.Quality = film.MustElement(".grid-meta .quality").MustText()
		movieData.Duration = film.MustElement(".grid-meta .duration").MustText()

		// // Tambahkan informasi ke dalam movieData
		movie := model.Movie{
			Title:    movieData.Title,
			Rating:   movieData.Rating,
			Quality:  movieData.Quality,
			Duration: movieData.Duration,
			Genre:    movieData.Genre,
			URLImage: movieData.URLImage,
		}

		err := repository.NewMovieRepository().Save(movie)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// Tambahkan data film ke dalam slice movies
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "Success scrape data trend film", movies))
}

func (mh *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, err := repository.NewMovieRepository().FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseAPI(false, http.StatusInternalServerError, "Failed to get data", nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "Success get data movies", movies))
}

func (mh *MovieHandler) InsertData(c *gin.Context) {

}

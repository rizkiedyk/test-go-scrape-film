package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"
)

func main() {
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
	var movies []map[string]interface{}

	for _, film := range cardFilm {
		movieData := make(map[string]interface{})

		// ratingParent := film.MustParent().MustParent().MustElement(".grid-meta")

		// Ambil teks judul
		movieData["title"] = film.MustElement(".grid-header .grid-title").MustText()
		// Ambil teks dari elemen-elemen terkait
		ratingElems := film.MustElements(".grid-meta .rating")

		// Periksa apakah ada elemen rating sebelum mencoba mengambil teksnya
		if len(ratingElems) > 0 {
			rating := ratingElems[0].MustText()
			movieData["rating"] = rating
		} else {
			// Atur nilai rating menjadi string kosong atau nilai default
			movieData["rating"] = ""
		}

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
					imageURL := extractWebPURL(*srcset)

					if !strings.HasPrefix(imageURL, "https:") {
						imageURL = "https:" + imageURL
					}

					// Set nilai image dalam map movieData
					movieData["image"] = imageURL
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Elemen <source> tidak ditemukan.")
			}
		} else {
			fmt.Println("Elemen <figure> tidak ditemukan.")
		}

		quality := film.MustElement(".grid-meta .quality").MustText()
		duration := film.MustElement(".grid-meta .duration").MustText()

		// // Tambahkan informasi ke dalam map movieData
		// movieData["rating"] = rating
		movieData["quality"] = quality
		movieData["duration"] = duration

		// Tambahkan data film ke dalam slice movies
		movies = append(movies, movieData)
	}

	// Cetak hasil
	fmt.Println("Informasi Film:")
	for i, movie := range movies {
		fmt.Printf("%d. Judul: %s\n", i+1, movie["title"])
		fmt.Printf("   Rating: %s\n", movie["rating"])
		fmt.Printf("   Quality: %s\n", movie["quality"])
		fmt.Printf("   Duration: %s\n", movie["duration"])
		fmt.Printf("   Image: %s\n", movie["image"])
		fmt.Println()
	}

	// Menunggu sejenak sebelum menutup browser
	time.Sleep(5 * time.Hour)
}

func extractWebPURL(srcset string) string {
	// Logika ekstraksi URL webp dari srcset
	// Misalnya, ambil URL yang berakhir dengan ".webp"
	// Anda mungkin perlu menggunakan ekspresi reguler atau metode lain sesuai kebutuhan
	// Contoh sederhana: Ambil URL yang berakhir dengan ".webp"
	urls := strings.Split(srcset, ",")
	for _, url := range urls {
		if strings.Contains(url, ".webp") {
			return strings.TrimSpace(url)
		}
	}
	return ""
}

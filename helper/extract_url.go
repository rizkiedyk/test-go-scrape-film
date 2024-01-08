package helper

import "strings"

func ExtractWebPURL(srcset string) string {
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

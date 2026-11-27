package database

import (
	"testing"

	"bem-portfolio/backend/models"
)

func TestIsDefaultHeroContent(t *testing.T) {
	content := models.KontenWebsite{
		NamaTampilan:  "Nama Anda",
		LabelProfesi:  "PORTFOLIO",
		JudulUtama:    "Halo, saya Nama Anda",
		TeksUtamaHero: "Saya seorang web developer",
		TeksPendukung: "Saya senang membangun aplikasi yang bermanfaat.",
		Deskripsi:     "Selamat datang di portfolio pribadi saya.",
		TeksTombolCTA: "Hubungi saya",
		URLTujuanCTA:  "mailto:halo@example.com",
	}
	if !isDefaultHeroContent(content) {
		t.Fatal("default hero content was not recognized")
	}

	content.JudulUtama = "Hero yang sudah diedit"
	if isDefaultHeroContent(content) {
		t.Fatal("custom hero content was incorrectly recognized as default")
	}
}

func TestIsLegacyGroupHeroContent(t *testing.T) {
	content := models.KontenWebsite{
		NamaTampilan:    "Kelompok 1",
		LabelProfesi:    "KELOMPOK 1",
		JudulUtama:      "Belajar bareng, berkarya bareng.",
		TeksUtamaHero:   "Kami dari Kelompok 1, lagi belajar dan bikin proyek bersama.",
		TeksPendukung:   "Di sini kami kumpulkan tugas, proyek, dan pengalaman selama belajar.",
		TeksHeroLainnya: "Akademik × PTI",
		Deskripsi:       "Selamat datang! Lihat proyek dan perjalanan kami di sini.",
		TeksTombolCTA:   "Lihat proyek kami",
		URLTujuanCTA:    "#proyek",
	}
	if !isLegacyGroupHeroContent(content) {
		t.Fatal("legacy group hero content was not recognized")
	}

	content.JudulUtama = "Hero yang sudah diedit"
	if isLegacyGroupHeroContent(content) {
		t.Fatal("custom hero content was incorrectly recognized as legacy default")
	}
}

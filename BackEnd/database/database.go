package database

import (
	"bem-portfolio/backend/config"
	"bem-portfolio/backend/models"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("buka database: %w", err)
	}
	if err := db.AutoMigrate(
		&models.Autentikasi{}, &models.KontenWebsite{}, &models.Profil{}, &models.Ikon{},
		&models.MediaSosial{}, &models.Perjalanan{}, &models.Teknologi{}, &models.Proyek{},
		&models.Sertifikasi{}, &models.TeknologiProyek{}, &models.TeknologiProfil{},
	); err != nil {
		return nil, fmt.Errorf("migration: %w", err)
	}
	if err := seed(db, cfg); err != nil {
		return nil, err
	}
	return db, nil
}

func seed(db *gorm.DB, cfg config.Config) error {
	var auth models.Autentikasi
	err := db.First(&auth, 1).Error
	if err == gorm.ErrRecordNotFound {
		if len(cfg.AdminPassword) < 12 {
			return fmt.Errorf("ADMIN_PASSWORD wajib diisi minimal 12 karakter saat seed admin pertama")
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("hash password admin: %w", err)
		}
		seedAuth := models.Autentikasi{IDAutentikasi: 1, NamaPengguna: cfg.AdminUsername, HashKataSandi: string(hash)}
		if err := db.Create(&seedAuth).Error; err != nil {
			return fmt.Errorf("seed admin: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("baca admin: %w", err)
	}

	var content models.KontenWebsite
	if err := db.First(&content, 1).Error; err == gorm.ErrRecordNotFound {
		content = models.KontenWebsite{
			IDKonten: 1, NamaTampilan: "Kelompok 1", LabelProfesi: "KELOMPOK 1",
			JudulUtama:      "Belajar bareng, berkarya bareng.",
			TeksUtamaHero:   "Kami dari Kelompok 1, lagi belajar dan bikin proyek bersama.",
			TeksPendukung:   "Di sini kami kumpulkan tugas, proyek, dan pengalaman selama belajar.",
			TeksHeroLainnya: "",
			Deskripsi:       "Selamat datang! Lihat proyek dan perjalanan kami di sini.",
			TeksTombolCTA:   "Lihat proyek kami", URLTujuanCTA: "#proyek",
			DeskripsiFooter: "Terima kasih sudah berkunjung.",
			TeksHakCipta:    "Hak cipta Nama Anda", TeksTambahanFooter: "Dibuat untuk belajar.",
		}
		if err := db.Create(&content).Error; err != nil {
			return fmt.Errorf("seed website content: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("baca website content: %w", err)
	} else if isDefaultHeroContent(content) || isLegacyGroupHeroContent(content) {
		content.NamaTampilan = "Kelompok 1"
		content.LabelProfesi = "KELOMPOK 1"
		content.JudulUtama = "Belajar bareng, berkarya bareng."
		content.TeksUtamaHero = "Kami dari Kelompok 1, lagi belajar dan bikin proyek bersama."
		content.TeksPendukung = "Di sini kami kumpulkan tugas, proyek, dan pengalaman selama belajar."
		content.TeksHeroLainnya = ""
		content.Deskripsi = "Selamat datang! Lihat proyek dan perjalanan kami di sini."
		content.TeksTombolCTA = "Lihat proyek kami"
		content.URLTujuanCTA = "#proyek"
		if err := db.Save(&content).Error; err != nil {
			return fmt.Errorf("perbarui konten hero default: %w", err)
		}
	}

	var profile models.Profil
	if err := db.First(&profile, 1).Error; err == gorm.ErrRecordNotFound {
		profile = models.Profil{
			IDProfil: 1, NamaLengkap: "Nama Anda", JabatanProfesional: "Web Developer",
			Lokasi: "Indonesia", Email: "halo@example.com", Biografi: "Saya sedang belajar pengembangan web.",
			DeskripsiTentangSaya: "Saya tertarik membangun aplikasi web dan terus belajar teknologi baru.",
		}
		if err := db.Create(&profile).Error; err != nil {
			return fmt.Errorf("seed profil: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("baca profil: %w", err)
	}

	ikonContoh := []models.Ikon{
		{IDIkon: 1, NamaIkon: "GitHub", URLIkon: "https://cdn.simpleicons.org/github", Kategori: "media_sosial"},
		{IDIkon: 2, NamaIkon: "Go", URLIkon: "https://cdn.simpleicons.org/go", Kategori: "teknologi"},
		{IDIkon: 3, NamaIkon: "React", URLIkon: "https://cdn.simpleicons.org/react", Kategori: "teknologi"},
	}
	for _, ikon := range ikonContoh {
		var data models.Ikon
		if err := db.First(&data, ikon.IDIkon).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&ikon).Error; err != nil {
				return fmt.Errorf("seed ikon: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("baca ikon: %w", err)
		}
	}

	teknologiContoh := []models.Teknologi{
		{IDTeknologi: 1, NamaTeknologi: "Go", IDIkon: uintPointer(2)},
		{IDTeknologi: 2, NamaTeknologi: "React", IDIkon: uintPointer(3)},
	}
	for _, teknologi := range teknologiContoh {
		var data models.Teknologi
		if err := db.First(&data, teknologi.IDTeknologi).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&teknologi).Error; err != nil {
				return fmt.Errorf("seed teknologi: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("baca teknologi: %w", err)
		}
	}

	perjalananContoh := []models.Perjalanan{
		{IDPerjalanan: 1, Institusi: "Universitas Contoh", Kategori: "pendidikan", TanggalMulai: time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC), InformasiTambahan: "Teknik Informatika", Deskripsi: "Belajar dasar pengembangan perangkat lunak.", UrutanTampil: 1},
		{IDPerjalanan: 2, Institusi: "BEM Fakultas", Kategori: "pekerjaan", TanggalMulai: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), InformasiTambahan: "Staf bidang teknologi", Deskripsi: "Berkolaborasi membuat program kerja digital.", UrutanTampil: 2},
	}
	for _, item := range perjalananContoh {
		var data models.Perjalanan
		if err := db.First(&data, item.IDPerjalanan).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&item).Error; err != nil {
				return fmt.Errorf("seed perjalanan: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("baca perjalanan: %w", err)
		}
	}

	proyekContoh := []models.Proyek{
		{IDProyek: 1, Judul: "Portfolio Pribadi", Kategori: "Website", Deskripsi: "Website untuk menampilkan profil dan proyek.", Tanggal: datePointer(2025, 3, 1), URLDemo: "https://example.com", URLGitHub: "https://github.com/example/portfolio", Status: "finished", Prioritas: 1},
		{IDProyek: 2, Judul: "API Portfolio", Kategori: "Website", Deskripsi: "REST API sederhana menggunakan Go, Fiber, GORM, dan MySQL.", Tanggal: datePointer(2025, 5, 1), URLGitHub: "https://github.com/example/portfolio-api", Status: "on-going", Prioritas: 2},
	}
	for _, item := range proyekContoh {
		var data models.Proyek
		if err := db.First(&data, item.IDProyek).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&item).Error; err != nil {
				return fmt.Errorf("seed proyek: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("baca proyek: %w", err)
		}
	}

	if err := db.Where(
		"nama_sertifikat = ? AND penerbit = ? AND nomor_sertifikat = ? AND deskripsi = ?",
		"Dasar Pemrograman Go", "Belajar Mandiri", "GO-001", "Contoh data sertifikasi.",
	).Delete(&models.Sertifikasi{}).Error; err != nil {
		return fmt.Errorf("hapus sertifikat contoh: %w", err)
	}

	prestasi := []models.Sertifikasi{
		{
			NamaSertifikat:  "Gold Medal — Nusantara Creative Competition (NCC) 3",
			Penerbit:        "Lembaga Setara Prisma Nusantara (Nusantara Muda) & Politeknik TEDC Bandung",
			TanggalTerbit:   datePointer(2026, time.April, 11),
			NomorSertifikat: "001/TEKN/NCC3/LSPN-POLTEK.TEDC/IV/2026",
			URLGambar:       "/images/Juara%203%20NCC.png",
			Deskripsi:       "Bidang Teknologi. Waktu: 11–12 April 2026.",
			UrutanTampil:    3,
		},
		{
			NamaSertifikat:  "Best Presentation — Nusantara Creative Competition (NCC) 3",
			Penerbit:        "Lembaga Setara Prisma Nusantara (Nusantara Muda) & Politeknik TEDC Bandung",
			TanggalTerbit:   datePointer(2026, time.April, 11),
			NomorSertifikat: "007/BEST-FAV/NCC3/LSPN-POLTEK.TEDC/IV/2026",
			URLGambar:       "/images/Sertifikat%20Best%20Presentation.png",
			Deskripsi:       "Waktu: 11–12 April 2026.",
			UrutanTampil:    4,
		},
		{
			NamaSertifikat:  "Silver Medal — Nusantara Creative Competition (NCC) 3",
			Penerbit:        "Lembaga Setara Prisma Nusantara (Nusantara Muda) & Politeknik TEDC Bandung",
			TanggalTerbit:   datePointer(2026, time.April, 11),
			NomorSertifikat: "004/PTN/NCC3/LSPN-POLTEK.TEDC/IV/2026",
			URLGambar:       "/images/Silver%20Medal%20Pertanian.jpeg",
			Deskripsi:       "Bidang Pertanian. Waktu: 11–12 April 2026.",
			UrutanTampil:    5,
		},
		{
			NamaSertifikat:  "Peringkat Umum I — Semester Ganjil 2024/2025",
			Penerbit:        "SMA Xaverius 3 Palembang",
			TanggalTerbit:   datePointer(2024, time.December, 20),
			NomorSertifikat: "448/SMA Xav./3.S.6/2024",
			URLGambar:       "/images/Sertifikat%20Peringkat%20Semester%20Ganjil.jpg",
			Deskripsi:       "Fhabian Rafael meraih peringkat umum I dengan jumlah nilai 1238 dan rerata 95,23 di kelas F (XII).",
			UrutanTampil:    1,
		},
		{
			NamaSertifikat:  "Peringkat Umum I — Semester Genap 2023/2024",
			Penerbit:        "SMA Xaverius 3 Palembang",
			TanggalTerbit:   datePointer(2024, time.June, 14),
			NomorSertifikat: "132/SMA Xav./3.S.6/2024",
			URLGambar:       "/images/Sertifikat%20Peringkat%20Semester%20Genap.jpg",
			Deskripsi:       "Fhabian Rafael meraih peringkat umum I dengan rerata nilai 93,71 di kelas F (Rumpun Ilmu Teknik dan Kesehatan).",
			UrutanTampil:    2,
		},
	}
	for _, item := range prestasi {
		var existing models.Sertifikasi
		err := db.Where("nomor_sertifikat = ?", item.NomorSertifikat).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&item).Error; err != nil {
				return fmt.Errorf("seed sertifikat prestasi %q: %w", item.NamaSertifikat, err)
			}
			continue
		}
		if err != nil {
			return fmt.Errorf("baca sertifikat prestasi %q: %w", item.NamaSertifikat, err)
		}
		if existing.UrutanTampil != item.UrutanTampil {
			if err := db.Model(&existing).Update("sort_order", item.UrutanTampil).Error; err != nil {
				return fmt.Errorf("perbarui urutan sertifikat %q: %w", item.NamaSertifikat, err)
			}
		}
	}

	mediaSosialContoh := models.MediaSosial{IDMediaSosial: 1, NamaPlatform: "GitHub", NamaAkun: "username", IDIkon: uintPointer(1), URLTujuan: "https://github.com/username", UrutanTampil: 1}
	var mediaSosial models.MediaSosial
	if err := db.First(&mediaSosial, 1).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(&mediaSosialContoh).Error; err != nil {
			return fmt.Errorf("seed media sosial: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("baca media sosial: %w", err)
	}

	teknologiProfilContoh := models.TeknologiProfil{IDTeknologiProfil: 1, IDTeknologi: 1, Prioritas: 1}
	var teknologiProfil models.TeknologiProfil
	if err := db.First(&teknologiProfil, 1).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(&teknologiProfilContoh).Error; err != nil {
			return fmt.Errorf("seed teknologi profil: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("baca teknologi profil: %w", err)
	}

	teknologiProyekContoh := []models.TeknologiProyek{
		{IDTeknologiProyek: 1, IDProyek: 2, IDTeknologi: 1, Kategori: "Backend"},
		{IDTeknologiProyek: 2, IDProyek: 2, IDTeknologi: 2, Kategori: "Frontend"},
	}
	for _, item := range teknologiProyekContoh {
		var data models.TeknologiProyek
		if err := db.First(&data, item.IDTeknologiProyek).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&item).Error; err != nil {
				return fmt.Errorf("seed teknologi proyek: %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("baca teknologi proyek: %w", err)
		}
	}
	return nil
}

func uintPointer(value uint) *uint { return &value }

func isDefaultHeroContent(content models.KontenWebsite) bool {
	return content.NamaTampilan == "Nama Anda" &&
		content.LabelProfesi == "PORTFOLIO" &&
		content.JudulUtama == "Halo, saya Nama Anda" &&
		content.TeksUtamaHero == "Saya seorang web developer" &&
		content.TeksPendukung == "Saya senang membangun aplikasi yang bermanfaat." &&
		content.TeksHeroLainnya == "" &&
		content.Deskripsi == "Selamat datang di portfolio pribadi saya." &&
		content.TeksTombolCTA == "Hubungi saya" &&
		content.URLTujuanCTA == "mailto:halo@example.com"
}

func isLegacyGroupHeroContent(content models.KontenWebsite) bool {
	return content.NamaTampilan == "Kelompok 1" &&
		content.LabelProfesi == "KELOMPOK 1" &&
		content.JudulUtama == "Belajar bareng, berkarya bareng." &&
		content.TeksUtamaHero == "Kami dari Kelompok 1, lagi belajar dan bikin proyek bersama." &&
		content.TeksPendukung == "Di sini kami kumpulkan tugas, proyek, dan pengalaman selama belajar." &&
		content.TeksHeroLainnya == "Akademik × PTI" &&
		content.Deskripsi == "Selamat datang! Lihat proyek dan perjalanan kami di sini." &&
		content.TeksTombolCTA == "Lihat proyek kami" &&
		content.URLTujuanCTA == "#proyek"
}

func datePointer(year int, month time.Month, day int) *time.Time {
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return &date
}

package handlers

import (
	"testing"

	"bem-portfolio/backend/models"
)

func TestValidLink(t *testing.T) {
	diterima := []string{"", "https://example.com", "http://localhost:5173", "#proyek", "/cv.pdf", "mailto:halo@example.com"}
	ditolak := []string{"javascript:alert(1)", "//evil.com", "data:text/html,x", "ftp://example.com", "#", "/../etc/passwd", "mailto:bukan email", "contoh.com"}
	for _, v := range diterima {
		if !validLink(v) {
			t.Errorf("harus diterima: %q", v)
		}
	}
	for _, v := range ditolak {
		if validLink(v) {
			t.Errorf("harus ditolak: %q", v)
		}
	}
}

func TestCertificateFromInputKeepsImageURL(t *testing.T) {
	const imageURL = "/images/certificate.png"
	certificate, err := certificateFromInput(models.SertifikasiRequest{
		NamaSertifikat: "Sertifikat contoh",
		Penerbit:       "Penerbit contoh",
		URLGambar:      imageURL,
	})
	if err != nil {
		t.Fatalf("certificateFromInput() error = %v", err)
	}
	if certificate.URLGambar != imageURL {
		t.Errorf("URLGambar = %q, want %q", certificate.URLGambar, imageURL)
	}

	_, err = certificateFromInput(models.SertifikasiRequest{
		NamaSertifikat: "Sertifikat contoh",
		Penerbit:       "Penerbit contoh",
		URLGambar:      "javascript:alert(1)",
	})
	if err == nil {
		t.Fatal("certificateFromInput() accepted an invalid image URL")
	}
}

package handlers

import (
	"strings"

	"bem-portfolio/backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AlamatFromInput(in models.AlamattRequest) (models.Alamatt, error) {
	in.NamaAlamat = strings.TrimSpace(in.NamaAlamat)
	in.KodePos = strings.TrimSpace(in.KodePos)
	if in.NamaAlamat == "" || in.KodePos == "" {
		return models.Alamatt{}, fiber.NewError(400, "Nama alamat dan kode pos wajib diisi")
	}
	return models.Alamatt{NamaAlamat: in.NamaAlamat, KodePos: in.KodePos}, nil
}

func (h *Handler) ListAlamats(c *fiber.Ctx) error {
	page, limit, err := pagination(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var items []models.Alamatt
	var total int64
	if err := h.DB.Model(&models.Alamatt{}).Count(&total).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	if err := h.DB.Order("id_alamat ASC").Offset((page - 1) * limit).Limit(limit).Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	return c.JSON(pageResult(items, page, limit, total))
}

func (h *Handler) CreateAlamat(c *fiber.Ctx) error {
	var input models.AlamattRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Body JSON tidak valid",
		})
	}
	item, err := AlamatFromInput(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := h.DB.Create(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

func (h *Handler) GetAlamat(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var item models.Alamatt
	if err := h.DB.First(&item, "id_alamat = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Alamat tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	return c.JSON(item)
}

func (h *Handler) UpdateAlamatt(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var input models.AlamattRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Body JSON tidak valid",
		})
	}
	item, err := AlamatFromInput(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var alamat models.Alamatt
	if err := h.DB.First(&alamat, "id_alamat = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	alamat.NamaAlamat = item.NamaAlamat
	alamat.KodePos = item.KodePos
	if err := h.DB.Save(&alamat).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	return h.GetAlamat(c)
}

func (h *Handler) DeleteAlamat(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	result := h.DB.Delete(&models.Alamatt{}, "id_alamat = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Alamat tidak ditemukan",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

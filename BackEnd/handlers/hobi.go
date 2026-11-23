package handlers

import (
	"strings"

	"bem-portfolio/backend/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HobiFromInput(in models.HobiRequest) (models.Hobi, error) {
	in.NamaHobi = strings.TrimSpace(in.NamaHobi)
	in.Kategori = strings.TrimSpace(strings.ToLower(in.Kategori))
	if in.NamaHobi == "" || in.Kategori == "" {
		return models.Hobi{}, fiber.NewError(400, "Nama hobi dan kategori wajib diisi")
	}
	switch in.Kategori {
	case "makan", "tidur", "bermain", "mancing":
	default:
		return models.Hobi{}, fiber.NewError(400, "Kategori harus makan, tidur, bermain, atau mancing")
	}
	return models.Hobi{NamaHobi: in.NamaHobi, Kategori: in.Kategori}, nil
}

func (h *Handler) ListHobis(c *fiber.Ctx) error {
	page, limit, err := pagination(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var items []models.Hobi
	var total int64
	if err := h.DB.Model(&models.Hobi{}).Count(&total).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	if err := h.DB.Order("id_hobi ASC").Offset((page - 1) * limit).Limit(limit).Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	return c.JSON(pageResult(items, page, limit, total))
}

func (h *Handler) CreateHobi(c *fiber.Ctx) error {
	var input models.HobiRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Body JSON tidak valid",
		})
	}
	item, err := HobiFromInput(input)
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

func (h *Handler) GetHobi(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var item models.Hobi
	if err := h.DB.First(&item, "id_hobi = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Hobi tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	return c.JSON(item)
}

func (h *Handler) UpdateHobi(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var input models.HobiRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Body JSON tidak valid",
		})
	}
	item, err := HobiFromInput(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var hobi models.Hobi
	if err := h.DB.First(&hobi, "id_hobi = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Hobi tidak ditemukan"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	hobi.NamaHobi = item.NamaHobi
	hobi.Kategori = item.Kategori
	if err := h.DB.Save(&hobi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	return h.GetHobi(c)
}

func (h *Handler) DeleteHobi(c *fiber.Ctx) error {
	id, err := parseID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	result := h.DB.Delete(&models.Hobi{}, "id_hobi = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Terjadi kesalahan pada server",
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Hobi tidak ditemukan",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

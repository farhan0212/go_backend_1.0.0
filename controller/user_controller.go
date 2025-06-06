package controllers

import (
	database "backend/config"
	"backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 5 // Fixed limit per page

	offset := (page - 1) * limit

	var total int64
	database.DB.Model(&models.User{}).Count(&total)
	database.DB.Model(&models.User{}).Limit(limit).Offset(offset).Find(&users)

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	// Calculate pagination window (max 6 pages shown)
	var pages []int
	startPage := page - 2
	if startPage < 1 {
		startPage = 1
	}
	endPage := startPage + 5
	if endPage > totalPages {
		endPage = totalPages
		startPage = endPage - 5
		if startPage < 1 {
			startPage = 1
		}
	}
	for i := startPage; i <= endPage; i++ {
		pages = append(pages, i)
	}

	return c.JSON(fiber.Map{
		"data":       users,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": totalPages,
		"pages":      pages, // pages to show in pagination
		"hasPrev":    page > 1,
		"hasNext":    page < totalPages,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.First(&user, id)
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	database.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
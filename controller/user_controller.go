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

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	
	var total int64
	database.DB.Model(&models.User{}).Count(&total)
	
	database.DB.Model(&models.User{}).Limit(limit).Offset(offset).Find(&users)
	
	totalPages := int((total + int64(limit) -1) / int64(limit))

	return c.JSON(fiber.Map{
		"data": users,
		"page": page,
		"limit": limit,
		"total": total,
		"totalPages": totalPages,
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
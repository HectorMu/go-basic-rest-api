package controllers

import (
	"github.com/HectorMu/go-rest-api/types"
	"github.com/HectorMu/go-rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var users = []types.User{
	{
		Id:   "23243",
		Name: "Pedro",
	},
	{
		Id:   "2q323",
		Name: "Juan",
	},
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func SaveUser(c *fiber.Ctx) error {

	var request types.User

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "JSON malformed",
		})
	}

	newUser := types.User{Name: request.Name, Id: uuid.New().String()}

	if err := util.ValidateUser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": "Missing object properties",
		})
	}

	users = append(users, newUser)

	return c.JSON(&fiber.Map{
		"message": "User created",
	})
}

func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")

	index := util.FindIndex[types.User](users, func(u types.User) bool {
		return u.Id == id
	})

	if index == -1 {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "User not found",
		})
	}

	users = append(users[:index], users[index+1:]...)

	return c.JSON(&fiber.Map{
		"message": "User deleted",
	})
}

func HandleEvery(c *fiber.Ctx) error {

	everyNameIsPedro := util.EverySlice[types.User](users, func(u types.User) bool {

		return u.Name == "Pedro"
	})

	if everyNameIsPedro {
		return c.JSON(everyNameIsPedro)
	}

	return c.JSON(everyNameIsPedro)
}

func HandleMap(c *fiber.Ctx) error {

	mappedNames := util.MapSlice[types.User](users, func(u types.User) string {
		return u.Id
	})

	return c.JSON(mappedNames)
}

func HandleFilter(c *fiber.Ctx) error {

	filteredUsers := util.FilterSlice[types.User](users, func(u types.User) bool {

		return u.Id == "23243"
	})

	return c.JSON(filteredUsers)
}

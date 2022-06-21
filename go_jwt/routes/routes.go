package routes

import (
	"time"

	"github.com/Shaviaditya/BasicGo/go_jwt/database"
	"github.com/Shaviaditya/BasicGo/go_jwt/models"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/jinzhu/gorm/dialects/mysql"
);

func AuthRequired() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx,err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":"Unauthorized",
			})
		},
		SigningKey: []byte("secret"),
	})
}

func Signup(c *fiber.Ctx) error{
	db := database.DBConn
	details := new(models.Details)
	if err := c.BodyParser(details); err!=nil {
		return c.Status(503).Send([]byte(err.Error()))
	}
	db.Create(&details)
	return c.Redirect("/login")
	// return c.JSON(details)	
}
func Login(c *fiber.Ctx) error{
	db := database.DBConn
	var details models.Details
	if err := c.BodyParser(&details); err!=nil {
		c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":"Cannot Parse JSON Data",
		})
	}
	var data models.Details
	db.Where(&models.Details{Email: details.Email,Password: details.Password}).Find(&data); 
	if len(data.Email)==0 && len(data.Password)==0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":"Bad Creds",
		})
	}
	token:= jwt.New(jwt.SigningMethodHS256)
	claims:= token.Claims.(jwt.MapClaims)
	claims["sub"] = "LoggedIn"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	_,err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	/* return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
		"user": details.Email,
	}) */
	// fmt.Println(t)
	return c.Redirect("/users",fiber.StatusOK)
}
func LoginGet(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title":"Hello World",
	})
}
func SignupGet(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{
		"Title":"Hello World",
	})
}
func GetUsers(c *fiber.Ctx) error{
	db := database.DBConn
	var details []models.Details
	db.Find(&details)
	return c.JSON(details)
}
func GetUserId(c *fiber.Ctx) error{
	db := database.DBConn
	var details models.Details
	db.Find(&details,c.Params("id"))
	return c.JSON(details)
}
func DeleteUserId(c *fiber.Ctx) error{
	db := database.DBConn
	var details models.Details
	db.Find(&details,c.Params("id"))
	if details.Username==""{
		return c.SendStatus(fiber.StatusBadRequest);
	}
	db.Delete(&details)
	return c.SendStatus(fiber.StatusAccepted)
}
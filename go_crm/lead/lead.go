package lead

import (
	"github.com/jinzhu/gorm"
	"github.com/Shaviaditya/BasicGo/go_crm/database"
	"github.com/gofiber/fiber/v2"
	_"github.com/jinzhu/gorm/dialects/mysql"
);
type Lead struct {
	gorm.Model
	Name	string	`json:"name"`
	Company string	`json:"company"`
	Email	string	`json:"email"`
	Phone	int		`json:"phone"`
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var leads Lead
	db.Find(&leads,id)
	return c.JSON(leads)
}
func GetLeads(c *fiber.Ctx)error {
	db := database.DBConn
	var lead []Lead
	db.Find(&lead)
	return c.JSON(lead)
}
func NewLead(c *fiber.Ctx)error { 
	db := database.DBConn
	lead := new(Lead)
	if err:= c.BodyParser(lead); err!=nil {
		return c.Status(503).Send([]byte(err.Error()))
	}
	db.Create(&lead)
	return c.JSON(lead)
}
func DeleteLead(c *fiber.Ctx)error {
	db := database.DBConn
	id := c.Params("id")
	var lead Lead
	db.First(&lead,id)
	if lead.Name == "" {
		return c.Status(500).SendString("No such data exists")
	}
	db.Delete(&lead)
	return c.SendString("Deleted Used successfully")
}


package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"fholl.net/go-pg-sample/models"
)

type Routes struct {
	DB *gorm.DB
}

type GetByIDConfig struct {
	model    interface{}
	preloads []string
}

func (r *Routes) Setup(e *echo.Echo) {

	// '/booking' group setup
	booking_group := e.Group("/booking")

	booking_group.GET("/:id", func(c echo.Context) error {
		return r.GetModelByID(c, GetByIDConfig{model: &models.Booking{}, preloads: []string{"Client.Contact", "Contractor.Contact"}})
	})
	booking_group.POST("/create", func(c echo.Context) error {
		return r.CreateModel(c, &models.Booking{})
	})
	booking_group.DELETE("/delete/:id", func(c echo.Context) error {
		return r.DeleteModel(c, &models.Booking{})
	})

	// '/client' group setup
	client_group := e.Group("/client")

	client_group.GET("/:id", func(c echo.Context) error {
		return r.GetModelByID(c, GetByIDConfig{model: &models.Client{}, preloads: []string{"Contact"}})
	})
	client_group.POST("/create", func(c echo.Context) error {
		return r.CreateModel(c, &models.Client{})
	})
	client_group.DELETE("/delete/:id", func(c echo.Context) error {
		return r.DeleteModel(c, &models.Client{})
	})

	// '/contact' group setup
	contact_group := e.Group("/contact")

	contact_group.GET("/:id", func(c echo.Context) error {
		return r.GetModelByID(c, GetByIDConfig{model: &models.Contact{}})
	})
	contact_group.POST("/create", func(c echo.Context) error {
		return r.CreateModel(c, &models.Contact{})
	})
	contact_group.DELETE("/delete/:id", func(c echo.Context) error {
		return r.DeleteModel(c, &models.Contact{})
	})

	// '/contractor' group setup
	contractor_group := e.Group("/contractor")

	contractor_group.GET("/:id", func(c echo.Context) error {
		return r.GetModelByID(c, GetByIDConfig{model: &models.Contractor{}, preloads: []string{"Contact"}})
	})
	contractor_group.POST("/create", func(c echo.Context) error {
		return r.CreateModel(c, &models.Contractor{})
	})
	contractor_group.DELETE("/delete/:id", func(c echo.Context) error {
		return r.DeleteModel(c, &models.Contractor{})
	})

	// 	booking_group.GET("/:id", bc.GetBookingByID)
	// 	booking_group.POST("/create", bc.CreateBooking)
	//	booking_group.DELETE("/delete/:id", func(c echo.Context) error {
	//		return bc.DeleteModel(c, &models.Booking{})
	//	})
}

// Create '/booking' routes
func (r *Routes) GetBookingByID(c echo.Context) error {

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "invalid int ID entered"})
	}

	booking := models.Booking{}
	err = r.DB.Preload("Client.Contact").Preload("Contractor.Contact").First(&booking, id).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not download booking",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, booking)
}

func (r *Routes) CreateBooking(c echo.Context) error {

	booking := models.Booking{}

	err := c.Bind(&booking)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "could not create a new booking",
			"error":   err,
		})
	}

	// 	var contractors []models.Contractor
	// 	if err := r.DB.Where("id in ?", booking.ContractorIDs).Find(&contractors).Error; err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 			"message": "Could not load contractors",
	// 			"error":   err,
	// 		})
	// 	}
	//
	// 	booking.Contractors = contractors

	err = r.DB.Create(&booking).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not upload new booking to database",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, booking)
}

func (r *Routes) DeleteBooking(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "invalid int ID entered"})
	}

	booking := models.Booking{}

	err = r.DB.Delete(booking, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not delete booking",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "booking deleted successfully",
		"id":      id,
	})
}

// func (r *Routes) GetModelByID(c echo.Context, model interface{}, preloads []string) error {
func (r *Routes) GetModelByID(c echo.Context, cfg GetByIDConfig) error {

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "invalid int ID entered"})
	}

	var q *gorm.DB = r.DB
	for _, pl := range cfg.preloads {
		q = q.Preload(pl)

	}
	err = q.Where("id = ?", id).First(cfg.model).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not download model",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, cfg.model)
}

func (r *Routes) CreateModel(c echo.Context, model interface{}) error {

	err := c.Bind(model)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"message": "could not create a new booking",
			"error":   err,
		})
	}

	err = r.DB.Create(model).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not upload new booking to database",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, model)
}

func (r *Routes) DeleteModel(c echo.Context, model interface{}) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid int ID entered",
		})
	}

	result := r.DB.Delete(model, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "could not delete item",
			"error":   result.Error,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "item deleted successfully",
		"id":      id,
	})
}

package routes

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"fholl.net/go-pg-sample/models"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Routes struct {
	DB *gorm.DB
}

type GetByIDConfig struct {
	model    interface{}
	preloads []string
	template string
}

func (r *Routes) Setup(e *echo.Echo) {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*html")),
	}

	e.Renderer = t

	// '/booking' group setup
	booking_group := e.Group("/booking")

	booking_group.GET("/:id", func(c echo.Context) error {
		return r.GetModelByID(c, GetByIDConfig{model: &models.Booking{}, preloads: []string{"Client.Contact", "Contractor.Contact"}, template: "booking"})
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
}

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

	hxRequest := c.Request().Header.Get("HX-Request")
	fmt.Println("Template: ", cfg.template)

	fmt.Println("HX-Request: ", hxRequest)

	if hxRequest == "true" {
		return c.Render(http.StatusOK, cfg.template, cfg.model)
	} else {
		return c.JSON(http.StatusOK, cfg.model)
	}

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

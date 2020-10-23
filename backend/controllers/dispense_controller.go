package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wipha26/app/ent"
	"github.com/wipha26/app/ent/drug"
	"github.com/wipha26/app/ent/patient"
	"github.com/wipha26/app/ent/user"
)

// DispenseController defines the struct for the dispense controller
type DispenseController struct {
	client *ent.Client
	router gin.IRouter
}

type Dispense struct {
	Patient int
	User    int
	Drug    int
	Note    string
}

// CreateDispenses handles POST requests for adding dispense entities
// @Summary Create dispense
// @Description Create dispense
// @ID create-dispense
// @Accept   json
// @Produce  json
// @Param dispense body ent.Dispense true "Dispense entity"
// @Success 200 {object} ent.Dispense
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dispenses [post]
func (ctl *DispenseController) CreateDispense(c *gin.Context) {
	obj := Dispense{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dispense binding failed",
		})
		return
	}

	d, err := ctl.client.Drug.
		Query().
		Where(drug.IDEQ(int(obj.Drug))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Drug not found",
		})
		return
	}
	// d, err := ctl.client.Drug.
	// 	Query().
	// 	Where(drug.IDEQ(int(obj.Drug))).
	// 	Only(context.Background())

	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": "Drug not found",
	// 	})
	// 	return
	// }

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "User not found",
		})
		return
	}

	di, err := ctl.client.Dispense.
		Create().
		SetDrug(d).
		SetPatient(p).
		SetUser(u).
		SetNote(obj.Note).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, di)
}

// ListDispense handles request to get a list of dispense entities
// @Summary List dispense entities
// @Description list dispense entities
// @ID list-dispense
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Dispense
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dispenses [get]
func (ctl *DispenseController) ListDispense(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	dispenses, err := ctl.client.Dispense.
		Query().
		WithDrug().
		WithPatient().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dispenses)
}

// NewDispenseController creates and registers handles for the dispense controller
func NewDispenseController(router gin.IRouter, client *ent.Client) *DispenseController {
	dic := &DispenseController{
		client: client,
		router: router,
	}
	dic.register()
	return dic
}

// InitDispenseController registers routes to the main engine
func (ctl *DispenseController) register() {
	dispenses := ctl.router.Group("/dispenses")

	dispenses.GET("", ctl.ListDispense)
	dispenses.POST("", ctl.CreateDispense)

}

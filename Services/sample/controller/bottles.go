package controller

import (
	"goprueba/Services/sample/repository"
	"goprueba/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowBottle godoc
// @Summary Show a bottle
// @Description get string by ID
// @ID get-string-by-int
// @Tags bottles
// @Accept  json
// @Produce  json
// @Param  id path int true "Bottle ID"
// @Success 200 {object} model.Bottle
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /bottles/{id} [get]
func (c *Controller) ShowBottle(ctx *gin.Context) {
	id := ctx.Param("id")
	bid, err := strconv.Atoi(id)
	if err != nil {
		utils.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	bottle, err := repository.BottleOne(bid)
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, bottle)
}

// ListBottles godoc
// @Summary List bottles
// @Description get bottles
// @Tags bottles
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Bottle
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /bottles [get]
func (c *Controller) ListBottles(ctx *gin.Context) {
	bottles, err := repository.BottlesAll()
	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, bottles)
}

package controller

import (
	"goprueba/Services/sample/repository"
	"goprueba/infrastructure"
	"goprueba/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListBras godoc
// @Summary List bras
// @Description get bras
// @Tags bras
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Bra
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /bras [get]
func (c *Controller) ListBras(ctx *gin.Context) {
	infrastructure.LoadEnv()

	results, err := repository.BrasAll()

	if err != nil {
		utils.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, results)
}

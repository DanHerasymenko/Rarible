package handler

import (
	"RaribleAPI/internal/model"
	"RaribleAPI/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RaribleHandler struct {
	service *service.RaribleService
}

func NewRaribleHandler(service *service.RaribleService) *RaribleHandler {
	return &RaribleHandler{service: service}
}

func (h *RaribleHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/rarible")
	{
		api.GET("/ownerships/:id", h.GetNFTOwnerships)
		api.POST("/traits/rarity", h.GetTraitRarities)
	}
}

// GetNFTOwnerships godoc
// @Summary      Get NFT Ownership by ID
// @Description  Returns NFT ownership details by ownership ID
// @Tags         ownership
// @Param        id   path      string  true  "Ownership ID"
// @Success      200  {object}  model.OwnershipResponse
// @Failure      500  {object}  gin.H
// @Router       /api/rarible/ownerships/{id} [get]
func (h *RaribleHandler) GetNFTOwnerships(c *gin.Context) {
	id := c.Param("id")
	result, err := h.service.GetNFTOwnerships(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetTraitRarities godoc
// @Summary      Get NFT Traits Rarity
// @Description  Returns rarity for NFT traits (recommended)
// @Tags         rarity
// @Accept       json
// @Produce      json
// @Param        body  body      model.RarityRequest  true  "Rarity request"
// @Success      200   {object}  model.RarityResponse
// @Failure      400   {object}  gin.H
// @Failure      500   {object}  gin.H
// @Router       /api/rarible/traits/rarity [post]
func (h *RaribleHandler) GetTraitRarities(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	result, err := h.service.GetTraitRarities(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Для swagger: явно використовуємо типи, щоб swag їх побачив
var (
	_ = model.OwnershipResponse{}
	_ = model.RarityRequest{}
	_ = model.RarityResponse{}
)

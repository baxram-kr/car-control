package handler

import (
	"app/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetList petrol_history godoc
// @ID get_list_petrol_history
// @Router /petrol_history [GET]
// @Summary Get List PetrolHistory
// @Description Get List PetrolHistory
// @Tags PetrolHistory
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListPetrolHistory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list PetrolHistory offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list PetrolHistory limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.PetrolHistory().GetList(c.Request.Context(), &models.PetrolHistoryGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})

	if err != nil {
		h.handlerResponse(c, "storage.PetrolHistory.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list PetrolHistory resposne", http.StatusOK, resp)
}

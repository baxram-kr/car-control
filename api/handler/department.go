package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create department godoc
// @ID create_department
// @Router /department [POST]
// @Summary Create Department
// @Description Create Department
// @Tags Department
// @Accept json
// @Procedure json
// @Param Department body models.CreateDepartment true "CreateDepartmentRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateDepartment(c *gin.Context) {

	var createDepartment models.CreateDepartment
	err := c.ShouldBindJSON(&createDepartment)
	if err != nil {
		h.handlerResponse(c, "error Department should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Department().Create(c.Request.Context(), &createDepartment)
	if err != nil {
		h.handlerResponse(c, "storage.Department.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Department().GetByID(c.Request.Context(), &models.DepartmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Department.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Department resposne", http.StatusCreated, resp)
}

// GetByID department godoc
// @ID get_by_id_department
// @Router /department/{id} [GET]
// @Summary Get By ID Department
// @Description Get By ID Department
// @Tags Department
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdDepartment(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// ToolData := val.(helper.TokenInfo)
	// if len(ToolData) > 0 {
	// 	id = ToolData.ToolID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Department().GetByID(c.Request.Context(), &models.DepartmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Department.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Department resposne", http.StatusOK, resp)
}

// GetList department godoc
// @ID get_list_department
// @Router /department [GET]
// @Summary Get List Department
// @Description Get List Department
// @Tags Department
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListDepartment(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Department offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Department limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Department().GetList(c.Request.Context(), &models.DepartmentGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})

	if err != nil {
		h.handlerResponse(c, "storage.Department.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Department resposne", http.StatusOK, resp)
}

// Update department godoc
// @ID update_department
// @Router /department/{id} [PUT]
// @Summary Update Department
// @Description Update Department
// @Tags Department
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Department body models.UpdateDepartment true "UpdateDepartmentRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateDepartment(c *gin.Context) {

	var (
		id               string = c.Param("id")
		updateDepartment models.UpdateDepartment
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateDepartment)
	if err != nil {
		h.handlerResponse(c, "error Department should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateDepartment.Id = id
	rowsAffected, err := h.strg.Department().Update(c.Request.Context(), &updateDepartment)
	if err != nil {
		h.handlerResponse(c, "storage.Department.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Department.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Department().GetByID(c.Request.Context(), &models.DepartmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Department.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Department resposne", http.StatusAccepted, resp)
}

// Delete department godoc
// @ID delete_department
// @Router /department/{id} [DELETE]
// @Summary Delete Department
// @Description Delete Department
// @Tags Department
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteDepartment(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Department().Delete(c.Request.Context(), &models.DepartmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Department.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Department resposne", http.StatusNoContent, nil)
}

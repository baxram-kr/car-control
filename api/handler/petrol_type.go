package handler

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"app/api/models"
// 	"app/pkg/helper"
// )

// // Create petrol_type godoc
// // @ID create_petrol_type
// // @Router /petrol_type [POST]
// // @Summary Create PetrolType
// // @Description Create PetrolType
// // @Tags PetrolType
// // @Accept json
// // @Procedure json
// // @Param PetrolType body models.CreatePetrolType true "CreatePetrolTypeRequest"
// // @Success 200 {object} Response{data=string} "Success Request"
// // @Response 400 {object} Response{data=string} "Bad Request"
// // @Failure 500 {object} Response{data=string} "Server error"
// func (h *handler) CreatePetrolType(c *gin.Context) {

// 	var createPetrolType models.CreatePetrolType
// 	err := c.ShouldBindJSON(&createPetrolType)
// 	if err != nil {
// 		h.handlerResponse(c, "error PetrolType should bind json", http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	id, err := h.strg.PetrolType().Create(c.Request.Context(), &createPetrolType)
// 	if err != nil {
// 		h.handlerResponse(c, "storage.PetrolType.create", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	resp, err := h.strg.PetrolType().GetByID(c.Request.Context(), &models.PetrolTypePrimaryKey{Id: id})
// 	if err != nil {
// 		h.handlerResponse(c, "storage.PetrolType.getById", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	h.handlerResponse(c, "create PetrolType resposne", http.StatusCreated, resp)
// }

// // GetByID petrol_type godoc
// // @ID get_by_id_petrol_type
// // @Router /petrol_type/{id} [GET]
// // @Summary Get By ID PetrolType
// // @Description Get By ID PetrolType
// // @Tags PetrolType
// // @Accept json
// // @Procedure json
// // @Param id path string false "id"
// // @Success 200 {object} Response{data=string} "Success Request"
// // @Response 400 {object} Response{data=string} "Bad Request"
// // @Failure 500 {object} Response{data=string} "Server error"
// func (h *handler) GetByIdPetrolType(c *gin.Context) {
// 	var id string = c.Param("id")

// 	// Here We Check id from Token
// 	// val, exist := c.Get("Auth")
// 	// if !exist {
// 	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
// 	// 	return
// 	// }

// 	// ToolData := val.(helper.TokenInfo)
// 	// if len(ToolData) > 0 {
// 	// 	id = ToolData.ToolID
// 	// } else {
// 	// 	id = c.Param("id")
// 	// }

// 	if !helper.IsValidUUID(id) {
// 		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
// 		return
// 	}

// 	resp, err := h.strg.PetrolType().GetByID(c.Request.Context(), &models.PetrolTypePrimaryKey{Id: id})
// 	if err != nil {
// 		h.handlerResponse(c, "storage.PetrolType.getById", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	h.handlerResponse(c, "get by id PetrolType resposne", http.StatusOK, resp)
// }

// // GetList petrol_type godoc
// // @ID get_list_petrol_type
// // @Router /petrol_type [GET]
// // @Summary Get List PetrolType
// // @Description Get List PetrolType
// // @Tags PetrolType
// // @Accept json
// // @Procedure json
// // @Param offset query string false "offset"
// // @Param limit query string false "limit"
// // @Param search query string false "search"
// // @Success 200 {object} Response{data=string} "Success Request"
// // @Response 400 {object} Response{data=string} "Bad Request"
// // @Failure 500 {object} Response{data=string} "Server error"
// func (h *handler) GetListPetrolType(c *gin.Context) {

// 	offset, err := h.getOffsetQuery(c.Query("offset"))
// 	if err != nil {
// 		h.handlerResponse(c, "get list PetrolType offset", http.StatusBadRequest, "invalid offset")
// 		return
// 	}

// 	limit, err := h.getLimitQuery(c.Query("limit"))
// 	if err != nil {
// 		h.handlerResponse(c, "get list PetrolType limit", http.StatusBadRequest, "invalid limit")
// 		return
// 	}

// 	resp, err := h.strg.PetrolType().GetList(c.Request.Context(), &models.PetrolTypeGetListRequest{
// 		Offset: offset,
// 		Limit:  limit,
// 		Search: c.Query("search"),
// 	})
// 	if err != nil {
// 		h.handlerResponse(c, "storage.PetrolType.get_list", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	h.handlerResponse(c, "get list PetrolType resposne", http.StatusOK, resp)
// }

// // Update petrol_type godoc
// // @ID update_petrol_type
// // @Router /petrol_type/{id} [PUT]
// // @Summary Update PetrolType
// // @Description Update PetrolType
// // @Tags PetrolType
// // @Accept json
// // @Procedure json
// // @Param id path string true "id"
// // @Param PetrolType body models.UpdatePetrolType true "UpdatePetrolTypeRequest"
// // @Success 200 {object} Response{data=string} "Success Request"
// // @Response 400 {object} Response{data=string} "Bad Request"
// // @Failure 500 {object} Response{data=string} "Server error"
// func (h *handler) UpdatePetrolType(c *gin.Context) {

// 	var (
// 		id               string = c.Param("id")
// 		updatePetrolType models.UpdatePetrolType
// 	)

// 	if !helper.IsValidUUID(id) {
// 		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
// 		return
// 	}

// 	err := c.ShouldBindJSON(&updatePetrolType)
// 	if err != nil {
// 		h.handlerResponse(c, "error PetrolType should bind json", http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	updatePetrolType.Id = id
// 	rowsAffected, err := h.strg.PetrolType().Update(c.Request.Context(), &updatePetrolType)
// 	if err != nil {
// 		h.handlerResponse(c, "storage.PetrolType.update", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	if rowsAffected <= 0 {
// 		h.handlerResponse(c, "storage.PetrolType.update", http.StatusBadRequest, "now rows affected")
// 		return
// 	}

// 	resp, err := h.strg.PetrolType().GetByID(c.Request.Context(), &models.PetrolTypePrimaryKey{Id: id})
// 	if err != nil {
// 		h.handlerResponse(c, "storage.PetrolType.getById", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	h.handlerResponse(c, "create PetrolType resposne", http.StatusAccepted, resp)
// }

// // Delete petrol_type godoc
// // @ID delete_petrol_type
// // @Router /petrol_type/{id} [DELETE]
// // @Summary Delete PetrolType
// // @Description Delete PetrolType
// // @Tags PetrolType
// // @Accept json
// // @Procedure json
// // @Param id path string true "id"
// // @Success 200 {object} Response{data=string} "Success Request"
// // @Response 400 {object} Response{data=string} "Bad Request"
// // @Failure 500 {object} Response{data=string} "Server error"
// func (h *handler) DeletePetrolType(c *gin.Context) {

// 	var id string = c.Param("id")

// 	if !helper.IsValidUUID(id) {
// 		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
// 		return
// 	}

// 	err := h.strg.PetrolType().Delete(c.Request.Context(), &models.PetrolTypePrimaryKey{Id: id})
// 	if err != nil {
// 		h.handlerResponse(c, "storage.PetrolType.delete", http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	h.handlerResponse(c, "create PetrolType resposne", http.StatusNoContent, nil)
// }

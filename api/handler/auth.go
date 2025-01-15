package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Procedure json
// @Param login body models.LoginInfo true "LoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Login(c *gin.Context) {
	var login models.LoginInfo
	var role string
	var user_id string
	var passw string

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	doctor, err := h.strg.Admin().GetByID(context.Background(), &models.AdminPrimaryKey{Email: login.Email})
	if err == nil {
		user_id = doctor.Id
		role = "admin"
		passw = doctor.Password
		// hashedPassword = doctor.Password
	} else {
		// Check in HeadNurse table
		headNurse, err := h.strg.Department().GetByID(context.Background(), &models.DepartmentPrimaryKey{Email: login.Email})
		if err == nil {
			user_id = headNurse.Id
			role = "department"
			passw = headNurse.Password
		} else {
			// No matching user found
			if err.Error() == "no rows in result set" {
				h.handlerResponse(c, "User does not exist", http.StatusBadRequest, "User does not exist")
				return
			}
			h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
			return
		}
	}

	// resp, err := h.strg.Admin().GetByID(context.Background(), &models.AdminPrimaryKey{Email: login.Email})
	// if err == nil {
	// 	user_type = "admin"
	// 	user_id = resp.Id
	// 	login.Password = resp.Password
	// } else if err.Error() == "no rows in result set" {
	// 	// If not found in Admin, try the Department storage
	// 	resp2, err := h.strg.Department().GetByID(context.Background(), &models.DepartmentPrimaryKey{Email: login.Email})
	// 	if err == nil {
	// 		user_type = "department"
	// 		user_id = resp2.Id
	// 		login.Password = resp2.Password
	// 	} else if err.Error() == "no rows in result set" {
	// 		// User not found in both Admin and Department
	// 		h.handlerResponse(c, "User does not exist", http.StatusBadRequest, "User does not exist")
	// 		return
	// 	} else {
	// 		// Unexpected error in Department storage
	// 		h.handlerResponse(c, "storage.department.getByID", http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// } else {
	// 	// Unexpected error in Admin storage
	// 	h.handlerResponse(c, "storage.admin.getByID", http.StatusInternalServerError, err.Error())
	// 	return
	// }

	if login.Password != passw {
		h.handlerResponse(c, "Wrong password", http.StatusBadRequest, "Wrong password")
		return
	}

	token, err := helper.GenerateJWT(map[string]interface{}{
		"user_id": user_id,
		"role":    role,
	}, time.Hour*24, h.cfg.SecretKey)

	h.handlerResponseLogin(c, "token", http.StatusCreated, token, role)

}

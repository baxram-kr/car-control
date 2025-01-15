package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	r.Use(customCORSMiddleware())

	// r.POST("/login", handler.Login)

	//SisAmin
	r.POST("/admin", handler.CreateAdmin)
	r.GET("/admin/:id", handler.GetByIdAdmin)
	r.GET("/admin", handler.GetListAdmin)
	r.PUT("/admin/:id", handler.UpdateAdmin)
	r.DELETE("/admin/:id", handler.DeleteAdmin)
	///////////////////////////////////////////////////
	//Region
	// r.POST("/region", handler.CreateRegion)
	// r.GET("/region/:id", handler.GetByIdRegion)
	// r.GET("/region", handler.GetListRegion)
	// r.PUT("/region/:id", handler.UpdateRegion)
	// r.DELETE("/region/:id", handler.DeleteRegion)
	//PetrolTtype
	// r.POST("/petrol_type", handler.CreatePetrolType)
	// r.GET("/petrol_type/:id", handler.GetByIdPetrolType)
	// r.GET("/petrol_type", handler.GetListPetrolType)
	// r.PUT("/petrol_type/:id", handler.UpdatePetrolType)
	// r.DELETE("/petrol_type/:id", handler.DeletePetrolType)
	//DEpartment
	r.POST("/department", handler.CreateDepartment)
	r.GET("/department/:id", handler.GetByIdDepartment)
	r.GET("/department", handler.GetListDepartment)
	r.PUT("/department/:id", handler.UpdateDepartment)
	r.DELETE("/department/:id", handler.DeleteDepartment)
	//Car
	r.POST("/car", handler.CreateCar)
	r.GET("/car/:id", handler.GetByIdCar)
	r.GET("/car", handler.GetListCar)
	r.PUT("/car/:id", handler.UpdateCar)
	r.DELETE("/car/:id", handler.DeleteCar)

	r.POST("/login", handler.Login)

	r.GET("/petrol_history", handler.GetListPetrolHistory)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

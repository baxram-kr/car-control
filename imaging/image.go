package imaging

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"path/filepath"

// 	firebase "firebase.google.com/go"
// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/api/option"
// )

// // FirebaseApp holds the Firebase app instance
// var FirebaseApp *firebase.App

// // InitializeFirebase initializes the Firebase app
// func InitializeFirebase() {
// 	ctx := context.Background()
// 	opt := option.WithCredentialsFile("firebase-service-account.json")

// 	app, err := firebase.NewApp(ctx, nil, opt)
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to initialize Firebase: %v", err))
// 	}

// 	FirebaseApp = app
// }

// // UploadImageToFirebase uploads an image to Firebase Storage
// func UploadImageToFirebase(c *gin.Context) {
// 	ctx := context.Background()
// 	client, err := FirebaseApp.Storage(ctx)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Firebase Storage"})
// 		return
// 	}

// 	// Get the uploaded file
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image"})
// 		return
// 	}

// 	// Open the file
// 	fileContent, err := file.Open()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
// 		return
// 	}
// 	defer fileContent.Close()

// 	// Get the storage bucket
// 	bucket, err := client.Bucket("<your-bucket-name>.appspot.com")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to access Firebase bucket"})
// 		return
// 	}

// 	// Create a reference for the file
// 	object := bucket.Object(filepath.Base(file.Filename))

// 	// Upload the file
// 	writer := object.NewWriter(ctx)
// 	if _, err := writer.Write([]byte(fileContent)); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
// 		return
// 	}
// 	writer.Close()

// 	// Generate a download URL
// 	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket.Name(), object.ObjectName())
// 	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": url})
// }

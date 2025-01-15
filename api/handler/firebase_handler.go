package handler

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"cloud.google.com/go/storage"
// 	firebase "firebase.google.com/go"
// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/api/option"
// )

// var FirebaseApp *firebase.App

// // InitializeFirebase initializes Firebase using the service account
// func InitializeFirebase() {
// 	ctx := context.Background()
// 	opt := option.WithCredentialsFile("srm-system-4e98b-firebase-adminsdk-54jil-5f6d8f9d9b.json")
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

// 	// Access the storage bucket
// 	bucket, err := client.Bucket("srm-system-4e98b.appspot.com")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to access Firebase bucket"})
// 		return
// 	}

// 	// Create a reference for the file
// 	object := bucket.Object(file.Filename)

// 	// Upload the file
// 	writer := object.NewWriter(ctx)
// 	defer writer.Close()

// 	if _, err := writer.Write([]byte(fileContent)); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
// 		return
// 	}

// 	// Generate a public download URL
// 	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket.Name(), object.ObjectName())
// 	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": url})
// }

// // GetImageFromFirebase retrieves a signed URL for an image
// func GetImageFromFirebase(c *gin.Context) {
// 	imageName := c.Param("imageName")

// 	ctx := context.Background()
// 	client, err := FirebaseApp.Storage(ctx)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Firebase Storage"})
// 		return
// 	}

// 	// Access the storage bucket
// 	bucket, err := client.Bucket("srm-system-4e98b.appspot.com")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to access Firebase bucket"})
// 		return
// 	}

// 	// Get the object reference
// 	object := bucket.Object(imageName)

// 	// Generate a signed URL for temporary access (valid for 1 hour)
// 	url, err := object.SignedURL(ctx, &storage.SignedURLOptions{
// 		Method:  "GET",
// 		Expires: time.Now().Add(1 * time.Hour), // URL valid for 1 hour
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate download URL"})
// 		return
// 	}

// 	// Return the signed URL
// 	c.JSON(http.StatusOK, gin.H{"url": url})
// }

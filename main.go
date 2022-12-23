package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// Retrieve variables from .env file
	cloudName := os.Getenv("CLOUD_NAME")
	cloudinaryApiKey := os.Getenv("CLOUDINARY_KEY")
	cloudinarySecret := os.Getenv("CLOUDINARY_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, cloudinaryApiKey, cloudinarySecret)
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	var ctx = context.Background()

	// Uploading
	uploadResult, err := cld.Upload.Upload(
		ctx,
		"<YOUR_FILE>",
		uploader.UploadParams{PublicID: "logo"})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	log.Println(uploadResult.SecureURL)
}

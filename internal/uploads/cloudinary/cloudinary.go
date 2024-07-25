package cloudinary

import (
    "context"
    "github.com/cloudinary/cloudinary-go"
    "github.com/cloudinary/cloudinary-go/api/uploader"
    "log"
)

var cld *cloudinary.Cloudinary

func Init(cloudName, apiKey, apiSecret string) {
    var err error
    cld, err = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
    if err != nil {
        log.Fatalf("Failed to initialize Cloudinary: %v", err)
    }
}

func UploadImage(filePath string, folder string) (string, error) {
    resp, err := cld.Upload.Upload(context.TODO(), filePath, uploader.UploadParams{Folder: folder})
    if err != nil {
        return "", err
    }
    return resp.SecureURL, nil
}

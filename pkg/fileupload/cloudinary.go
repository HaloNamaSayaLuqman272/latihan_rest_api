package fileupload

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CludinaryConfig struct {
	CloudinaryURL string
}

type Uploader interface {
	UploadFile(ctx context.Context, file any) (string, error)
}

type CloudinaryUploader struct {
	Cld     *cloudinary.Cloudinary
	File    any
	Options uploader.UploadParams
}

func (c *CludinaryConfig) InitCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL(c.CloudinaryURL)
	if err != nil {
		log.Fatalf("error when connecting to the cloudinary: %s\n", err)
	}

	log.Println("cloudinary initialization succed")
	return cld
}

func (c *CloudinaryUploader) UploadFile(ctx context.Context, file any) (string, error) {
	resp, err := c.Cld.Upload.Upload(ctx, file, c.Options)
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}

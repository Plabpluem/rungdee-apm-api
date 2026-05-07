package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"rungdee-apm-api/internal/entities"
	usecases "rungdee-apm-api/internal/usecases/storage"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

type StorageServiceRepository struct {
	Bucket string
}

func NewStorageRepository() usecases.StorageRepository {
	return &StorageServiceRepository{
		Bucket: "upload_rungdee",
	}
}

func (r *StorageServiceRepository) newClient(ctx context.Context) (*storage.Client, error) {
	return storage.NewClient(ctx)
}

func (r *StorageServiceRepository) Save(file io.Reader, filename, contentType string) (*entities.StorageResponse, error) {
	ctx := context.Background()
	client, err := r.newClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage new client: %w", err)
	}
	defer client.Close()

	// file, err := fileHeader.Open()
	// if err != nil {
	// 	return nil, fmt.Errorf("fileHeader.open: %w", err)
	// }
	// defer file.Close()

	// ext := filepath.Ext(fileHeader.Filename)
	ext := filepath.Ext(filename)
	file_name := uuid.New().String() + ext
	subFolder := "pdf"
	objectName := fmt.Sprintf("%s/%s", subFolder, file_name)

	wc := client.Bucket(r.Bucket).Object(objectName).NewWriter(ctx)
	wc.ContentType = contentType

	if _, err = io.Copy(wc, file); err != nil {
		return nil, fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return nil, fmt.Errorf("Writer.Close: %w", err)
	}

	url, err := r.GetUrl(file_name, subFolder)
	if err != nil {
		return nil, fmt.Errorf("No Found this picture %w", err)
	}

	return &entities.StorageResponse{
		Url:     url,
		BaseUrl: file_name,
	}, nil
}

func (r *StorageServiceRepository) GetUrl(fileName string, subFolder string) (string, error) {
	ctx := context.Background()
	client, err := r.newClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage new client: %w", err)
	}
	defer client.Close()

	filepath := fmt.Sprintf("%s/%s", subFolder, fileName)
	_, err = client.Bucket(r.Bucket).Object(filepath).Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("No Found this picture")
	}

	privateKey, clientEmail, err := r.loadServiceAccount(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		return "", fmt.Errorf("loadServiceAccount: %w", err)
	}

	url, err := storage.SignedURL(r.Bucket, filepath, &storage.SignedURLOptions{
		Expires:        time.Now().Add(24 * time.Hour),
		Method:         "GET",
		PrivateKey:     []byte(privateKey),
		GoogleAccessID: clientEmail,
	})
	return url, nil
}

func (r *StorageServiceRepository) loadServiceAccount(path string) (privateKey string, clientEmail string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

	var creds struct {
		PrivateKey  string `json:"private_key"`
		ClientEmail string `json:"client_email"`
	}
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", "", err
	}

	return creds.PrivateKey, creds.ClientEmail, nil
}

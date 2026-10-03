package service

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
)

type CreateEventSrv struct {
	CEr *repo.DbCreateEventRepo
}

func ProviderCreateEventService(cer *repo.DbCreateEventRepo) *CreateEventSrv {
	return &CreateEventSrv{
		CEr: cer,
	}
}

func (pr *CreateEventSrv) CreateEventService(data dto.CreateEvent) (string, error) {
	splitData := strings.Split(data.Images, ",")
	if len(splitData) != 2 {
		return "", fmt.Errorf("invalid base64 image")
	}

	imageEvent, err := base64.StdEncoding.DecodeString(splitData[1])
	if err != nil {
		return "", err
	}

	uploadPath := "public/uploads/event"
	fileName := fmt.Sprintf("%s-%d.jpg", strings.ReplaceAll(data.Title, " ", "-"), time.Now().Unix())
	filePath := filepath.Join(uploadPath, fileName)

	if e := os.WriteFile(filePath, imageEvent, 0755); e != nil {
		return "", e
	}

	eventPath := "public/uploads/event/" + fileName

	data.Images = fileName

	if e := pr.CEr.CreateEventRepo(data); e != nil {
		return "", e
	}

	return eventPath, nil
}

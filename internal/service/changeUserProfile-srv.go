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

type ChangeUserProfileSrvS struct {
	CUPr *repo.DbChangeUserProfileRepoS
}

func ChangeUserProfileService(cupr *repo.DbChangeUserProfileRepoS) *ChangeUserProfileSrvS {
	return &ChangeUserProfileSrvS{
		CUPr: cupr,
	}
}

func (s *ChangeUserProfileSrvS) ChangeUserProfileSrv(userID int, data dto.ChangeUserProfile) (string, error) {

	profilePath := ""

	if data.Profile != "" {
		parts := strings.Split(data.Profile, ",")

		if len(parts) != 2 {
			return "", fmt.Errorf("invalid base64 image")
		}

		imageData, err := base64.StdEncoding.DecodeString(parts[1])

		if err != nil {
			return "", err
		}

		uploadPath := "public/uploads/profile"

		fileName := fmt.Sprintf(
			"user-%d-%s.jpg", userID, time.Now(),
		)

		filePath := filepath.Join(uploadPath, fileName)

		if err := os.WriteFile(filePath, imageData, 0755); err != nil {
			return "", err
		}

		profilePath = "/uploads/profile/" + fileName
	}

	err := s.CUPr.ChangeUserProfileRpo(userID, data, profilePath)

	if err != nil {
		return "", err
	}

	return profilePath, nil
}

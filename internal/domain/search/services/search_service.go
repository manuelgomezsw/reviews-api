package services

import (
	"reviews-api/internal/domain"
	"reviews-api/internal/domain/search/repository"
)

func GetByID(reviewID int64) (domain.Review, error) {
	review, err := repository.GetByID(reviewID)
	if err != nil {
		return domain.Review{}, err
	}

	return review, nil
}

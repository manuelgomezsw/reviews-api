package services

import (
	"reviews-api/internal/domain"
	"reviews-api/internal/domain/registry/repository"
	"strings"
)

func Create(review *domain.Review) error {
	formatReview(review)

	if err := repository.Create(review); err != nil {
		return err
	}

	return nil
}

func Update(currentReview *domain.Review) error {
	formatReview(currentReview)

	if err := repository.Update(currentReview); err != nil {
		return err
	}

	return nil
}

func Delete(reviewID int64) error {
	if err := repository.Delete(reviewID); err != nil {
		return err
	}

	return nil
}

func formatReview(review *domain.Review) {
	review.Title = strings.TrimSpace(review.Title)
	review.Review = strings.TrimSpace(review.Review)
}

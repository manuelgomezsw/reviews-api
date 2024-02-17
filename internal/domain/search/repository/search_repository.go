package repository

import (
	"reviews-api/internal/domain"
	"reviews-api/internal/util/mysql"
)

func GetByID(reviewID int64) (domain.Review, error) {
	resultReview, err := mysql.ClientDB.Query(
		"SELECT review_id, title, review, date_created FROM quotes.reviews WHERE review_id = ?", reviewID)
	if err != nil {
		return domain.Review{}, err
	}

	var review domain.Review
	for resultReview.Next() {
		err = resultReview.Scan(&review.ReviewID, &review.Title, &review.Review, &review.DateCreated)
		if err != nil {
			return domain.Review{}, err
		}
	}

	return review, nil
}

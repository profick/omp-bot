package review

import (
	"errors"
	"fmt"
	"sort"

	"github.com/ozonmp/omp-bot/internal/model/communication"
)

type ReviewService interface {
	Describe(reviewID uint64) (*communication.Review, error)
	List(cursor uint64, limit uint64) ([]communication.Review, error)
	Create(communication.Review) (uint64, error)
	Update(reviewID uint64, review communication.Review) error
	Remove(reviewID uint64) (bool, error)
}

type CommunicationReviewService struct {
	reviews map[uint64]communication.Review
	keys    []uint64
	lastID  uint64
}

func NewCommunicationReviewService() *CommunicationReviewService {
	return &CommunicationReviewService{
		reviews: make(map[uint64]communication.Review),
	}
}

func (crs *CommunicationReviewService) Describe(reviewID uint64) (*communication.Review, error) {
	review, ok := crs.reviews[reviewID]
	if !ok {
		return nil, fmt.Errorf("review (id: %d) not found", reviewID)
	}
	return &review, nil
}

func (crs *CommunicationReviewService) List(cursor uint64, limit uint64) ([]communication.Review, error) {
	var res []communication.Review
	for ; cursor < uint64(len(crs.keys)) && uint64(len(res)) < limit; cursor++ {
		review, ok := crs.reviews[crs.keys[cursor]]
		if !ok {
			return nil, errors.New("incompatible keys slice")
		}
		res = append(res, review)
	}
	return res, nil
}

func (crs *CommunicationReviewService) Create(review communication.Review) (uint64, error) {
	crs.lastID++
	reviewID := crs.lastID
	crs.keys = append(crs.keys, reviewID)

	review.ReviewID = reviewID
	crs.reviews[reviewID] = review
	return reviewID, nil
}

func (crs *CommunicationReviewService) Update(reviewID uint64, review communication.Review) error {
	_, ok := crs.reviews[reviewID]
	if !ok {
		return fmt.Errorf("review (id: %d) not found", reviewID)
	}

	review.ReviewID = reviewID
	crs.reviews[reviewID] = review
	return nil
}

func (crs *CommunicationReviewService) Remove(reviewID uint64) (bool, error) {
	if _, ok := crs.reviews[reviewID]; ok {
		delete(crs.reviews, reviewID)
		idx := sort.Search(len(crs.keys), func(i int) bool {
			return crs.keys[i] >= reviewID
		})
		if idx == len(crs.keys) || crs.keys[idx] != reviewID {
			return false, errors.New("failed to remove key")
		}
		crs.keys = append(crs.keys[:idx], crs.keys[idx+1:]...)
		return true, nil
	}
	return false, nil
}

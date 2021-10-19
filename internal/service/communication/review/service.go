package review

import (
	"errors"
	"fmt"
	"sort"

	"github.com/ozonmp/omp-bot/internal/model/communication"
)

type Service struct {
	reviews map[uint64]communication.Review
	keys    []uint64
	lastID  uint64
}

func NewService() *Service {
	return &Service{
		reviews: make(map[uint64]communication.Review),
	}
}

func (s *Service) Describe(reviewID uint64) (*communication.Review, error) {
	review, ok := s.reviews[reviewID]
	if !ok {
		return nil, fmt.Errorf("review (id: %d) not found", reviewID)
	}
	return &review, nil
}

func (s *Service) List(cursor uint64, limit uint64) ([]communication.Review, error) {
	var res []communication.Review
	for ; cursor < uint64(len(s.keys)) && uint64(len(res)) < limit; cursor++ {
		review, ok := s.reviews[s.keys[cursor]]
		if !ok {
			return nil, errors.New("incompatible keys slice")
		}
		res = append(res, review)
	}
	return res, nil
}

func (s *Service) Create(review communication.Review) (uint64, error) {
	s.lastID++
	reviewID := s.lastID
	s.keys = append(s.keys, reviewID)

	review.ReviewID = reviewID
	s.reviews[reviewID] = review
	return reviewID, nil
}

func (s *Service) Update(reviewID uint64, review communication.Review) error {
	_, ok := s.reviews[reviewID]
	if !ok {
		return fmt.Errorf("review (id: %d) not found", reviewID)
	}

	review.ReviewID = reviewID
	s.reviews[reviewID] = review
	return nil
}

func (s *Service) Remove(reviewID uint64) (bool, error) {
	if _, ok := s.reviews[reviewID]; ok {
		delete(s.reviews, reviewID)
		idx := sort.Search(len(s.keys), func(i int) bool {
			return s.keys[i] >= reviewID
		})
		if idx == len(s.keys) || s.keys[idx] != reviewID {
			return false, errors.New("failed to remove key")
		}
		s.keys = append(s.keys[:idx], s.keys[idx+1:]...)
		return true, nil
	}
	return false, nil
}

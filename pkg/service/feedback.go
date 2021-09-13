package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type FeedbackService struct {
	repo repository.Feedback
}

func (s *FeedbackService) AddFeedback(ctx context.Context, feedback models.Feedback, advertId string) error {
	err := s.repo.AddFeedback(ctx, feedback, advertId)

	if err != nil {
		return err
	}

	err = s.repo.UpdateRating(ctx, advertId)
	if err != nil {
		return err
	}

	return nil
}

func NewFeedbackService(repository *repository.Repository) *FeedbackService {
	return &FeedbackService{repo: repository}
}

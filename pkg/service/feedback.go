package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
	"time"
)

type FeedbackService struct {
	repo repository.Feedback
}

func (s *FeedbackService) UpdateFeedback(ctx context.Context,  feedbackId string, feedback models.Feedback) error {
	return s.UpdateFeedback(ctx, feedbackId, feedback)
}

func (s *FeedbackService) DeleteFeedback(ctx context.Context, feedbackId string) error {
	return s.DeleteFeedback(ctx, feedbackId)
}

func (s *FeedbackService) GetFeedbackByUserId(ctx context.Context, feedbackId string) (*models.Feedback, error) {
	panic("implement me")
}

func (s *FeedbackService) AddFeedback(ctx context.Context, feedback models.Feedback, advertId string) error {
	feedback.CreatedAt = time.Now()
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

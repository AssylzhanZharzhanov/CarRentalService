package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type BookmarkService struct {
	repo repository.Bookmarks
}

func (s *BookmarkService) RemoveUserBookmark(ctx context.Context, userId string, advertId string) error {
	return s.repo.RemoveUserBookmark(ctx, userId, advertId)
}

func (s *BookmarkService) AddUserBookmark(ctx context.Context, userId string, advertId string) error {
	return s.repo.AddUserBookmark(ctx, userId, advertId)
}

func (s *BookmarkService) GetUserBookmarks(ctx context.Context, userId string) ([]models.Advert, error) {
	return s.repo.GetUserBookmarks(ctx, userId)
}

func NewBookmarkService(repository *repository.Repository) *BookmarkService {
	return &BookmarkService{
		repo: repository.Bookmarks,
	}
}
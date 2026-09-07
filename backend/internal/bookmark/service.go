package bookmark

import "github.com/google/uuid"

type BookmarkService struct{}

func (b *BookmarkService) GetAll(userId uuid.UUID) ([]BookmarkResponseModel, error) {
	return nil, nil
}

func (b *BookmarkService) Create(userId uuid.UUID) (BookmarkResponseModel, error) {
	return BookmarkResponseModel{}, nil
}

func (b *BookmarkService) Get(userId uuid.UUID, id uuid.UUID) (BookmarkResponseModel, error) {
	return BookmarkResponseModel{}, nil
}

func (b *BookmarkService) Update(userId uuid.UUID, id uuid.UUID) (BookmarkResponseModel, error) {
	return BookmarkResponseModel{}, nil
}

func (b *BookmarkService) Delete(userId uuid.UUID, id uuid.UUID) error {
	return nil
}

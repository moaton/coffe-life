package admin

import (
	"coffe-life/internal/dto"
	"coffe-life/internal/interfaces"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type translates struct {
	repo interfaces.Repository
}

func newTranslates(deps Dependencies) *translates {
	return &translates{
		repo: deps.Repository,
	}
}

var _ interfaces.TranslatesUsecase = (*translates)(nil)

func (u *translates) GetTranslates(ctx context.Context) ([]*dto.Translate, error) {
	translates, err := u.repo.Admin().Translates().GetTranslates(u.repo.Conn().WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to get translates: %w", err)
	}

	return convertTranslatesToDto(translates), nil
}

func (u *translates) GetTranslateById(ctx context.Context, id string) (*dto.Translate, error) {
	translate, err := u.repo.Admin().Translates().GetTranslateById(u.repo.Conn().WithContext(ctx), id)
	if err != nil {
		return nil, fmt.Errorf("failed to get translates: %w", err)
	}

	return convertTranslateToDto(translate), nil
}

func (u *translates) CreateTranslate(ctx context.Context, translate dto.Translate) error {
	err := u.repo.Admin().Translates().CreateTranslate(u.repo.Conn().WithContext(ctx), convertTranslateToEntity(translate))
	if err != nil {
		return fmt.Errorf("failed to create translate: %w", err)
	}
	return nil
}

func (u *translates) UpdateTranslate(ctx context.Context, id string, translate dto.Translate) error {
	trs := convertTranslateToEntity(translate)

	ID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("failed to parse uuid: %w", err)
	}
	trs.ID = ID
	err = u.repo.Admin().Translates().UpdateTranslate(u.repo.Conn().WithContext(ctx), trs)
	if err != nil {
		return fmt.Errorf("failed to create translate: %w", err)
	}
	return nil
}

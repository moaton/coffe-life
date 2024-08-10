package admin

import (
	"coffe-life/internal/dto"
	"coffe-life/internal/entity"
	"encoding/json"
	"fmt"
)

func convertCategoryToDto(r entity.Category) dto.Category {
	return dto.Category{
		ID:          r.ID.String(),
		Name:        r.Name,
		Description: r.Description,
	}
}

func convertCategoriesToDto(r entity.Categories) dto.Categories {
	out := make(dto.Categories, 0, len(r))
	for _, v := range r {
		out = append(out, convertCategoryToDto(v))
	}
	return out
}

func convertCategoryToEntity(r dto.Category) *entity.Category {
	return &entity.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}

func convertFoodToDto(r entity.Food) dto.Food {
	var composition []dto.Composition
	out := dto.Food{
		ID:          r.ID.String(),
		Name:        r.Name,
		Type:        r.Type,
		Category:    r.Category,
		Price:       r.Price,
		IsNew:       r.IsNew,
		IsSpicy:     r.IsSpicy,
		Description: r.Description,
		Composition: composition,
	}

	err := json.Unmarshal(r.Composition, &composition)
	if err != nil {
		out.ErrorReason = "failed to parse compositions"
		return out
	}
	out.Composition = composition
	return out
}

func convertFoodsToDto(r entity.Foods) dto.Foods {
	out := make(dto.Foods, 0, len(r))
	for _, v := range r {
		out = append(out, convertFoodToDto(v))
	}
	return out
}

func convertFoodToEntity(r dto.Food) (*entity.Food, error) {
	composition, err := json.Marshal(r.Composition)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}
	return &entity.Food{
		Name:        r.Name,
		Type:        r.Type,
		Category:    r.Category,
		Price:       r.Price,
		IsNew:       r.IsNew,
		IsSpicy:     r.IsSpicy,
		Description: r.Description,
		Composition: composition,
	}, nil
}

func convertCreateUserRequestToEntity(req dto.CreateUserRequest) *entity.User {
	return &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Username:  req.Username,
		IsFirst:   true,
	}
}

func convertUserToDto(user *entity.User) *dto.User {
	return &dto.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		IsFirst:   user.IsFirst,
	}
}

func convertUsersToDto(users []*entity.User) []*dto.User {
	out := make([]*dto.User, 0, len(users))

	for _, user := range users {
		out = append(out, convertUserToDto(user))
	}
	return out
}

func convertGetUsersRequestToEntity(req dto.GetUsersRequest) entity.GetUsersRequest {
	return entity.GetUsersRequest{
		Pagination: entity.Pagination{
			Limit:  int(req.Size),
			Offset: int((req.Page - 1) * req.Size),
		},
		Search: req.Search,
	}
}

func convertUserToEntity(user dto.User) entity.User {
	return entity.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		IsFirst:   user.IsFirst,
	}
}

func convertTranslateToDto(translate entity.Translate) dto.Translate {
	return dto.Translate{
		ID:  translate.ID,
		RU:  translate.RU,
		KZ:  translate.KZ,
		ENG: translate.ENG,
	}
}

func convertTranslatesToDto(translates []entity.Translate) []dto.Translate {
	out := make([]dto.Translate, 0, len(translates))

	for _, translate := range translates {
		out = append(out, convertTranslateToDto(translate))
	}

	return out
}

func convertTranslateToEntity(translate dto.Translate) entity.Translate {
	return entity.Translate{
		RU:  translate.RU,
		KZ:  translate.KZ,
		ENG: translate.ENG,
	}
}

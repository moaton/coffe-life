package admin

import (
	"coffe-life/internal/dto"
	"coffe-life/internal/entity"
	"encoding/json"
	"fmt"
	"log"
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

func convertCategoryRequestToEntity(r dto.CategoryRequest) *entity.Category {
	return &entity.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}

func convertCompostionToDto(r entity.Composition) dto.Composition {
	return dto.Composition{
		Name:   r.Name,
		Weight: r.Weight,
		Unit:   r.Unit,
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

func convertFoodRequestToEntity(r dto.FoodRequest) (*entity.Food, error) {
	composition, err := json.Marshal(r.Composition)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}
	log.Println("cmp", composition, r.Composition)
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

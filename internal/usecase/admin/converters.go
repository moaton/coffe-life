package admin

import (
	"coffe-life/internal/dto"
	"coffe-life/internal/entity"
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

func convertCategoryRequestToEntity(r dto.CategoryRequest) entity.Category {
	return entity.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}

// mapper.go
package postgres

import "catalog/internal/entity"


func toEntityBrand(model CarBrandModel) entity.CarBrand {
	return entity.CarBrand{
		ID:    model.ID,
		Name: model.Name,
		IsVisible:     model.IsVisible,
	}
}

func toModelBrand(brand entity.CarBrand) CarBrandModel {
	return CarBrandModel{
		ID:    brand.ID,
		Name: brand.Name,
		IsVisible: brand.IsVisible,
	}
}

func toEntityType(model CarTypeModel) entity.CarType {
	return entity.CarType{
		ID:   model.ID,
		Name: model.Name,
		IsVisible: model.IsVisible,
	}
}

func toModelType(t entity.CarType) CarTypeModel {
	return CarTypeModel{
		ID:   t.ID,
		Name: t.Name,
		IsVisible: t.IsVisible,
	}
}

func toEntity(model CarModel) entity.Car {
	return entity.Car{
		ID:            model.ID,
		Brand:         toEntityBrand(model.Brand),
		Type:          toEntityType(model.Type),
		Model:         model.ModelCar,
		Year:          model.Year,
		License_plate: model.License_plate,
		Status:        model.Status,
		Price:         model.Price,
		IsVisible:     model.IsVisible,
	}
}

func toModel(car entity.Car) CarModel {
	return CarModel{
		ID:            car.ID,
		BrandID:       car.Brand.ID,
		TypeID:        car.Type.ID,
		ModelCar:      car.Model,
		Year:          car.Year,
		License_plate: car.License_plate,
		Status:        car.Status,
		Price:         car.Price,
		IsVisible:     car.IsVisible,
	}
}

func toEntityList(models []CarModel) []entity.Car {
	result := make([]entity.Car, 0, len(models))
	for _, m := range models {
		result = append(result, toEntity(m))
	}
	return result
}

func toEntityBrandList(models []CarBrandModel) []entity.CarBrand {
	result := make([]entity.CarBrand, 0, len(models))
	for _, m := range models {
		result = append(result, toEntityBrand(m))
	}
	return result
}

func toEntityTypeList(models []CarTypeModel) []entity.CarType {
	result := make([]entity.CarType, 0, len(models))
	for _, m := range models {
		result = append(result, toEntityType(m))
	}
	return result
}


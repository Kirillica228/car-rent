package usecase

import (
	"catalog/internal/domain/entity"
	"catalog/internal/interfaces/http/dto"
	"mime/multipart"
	"strings"
)

type CarUseCase struct {
	carRepo  entity.CarRepository
	fileRepo entity.FileRepository
}

func NewCarUseCase(car entity.CarRepository , file entity.FileRepository) entity.CarUseCase {
	return &CarUseCase{
		carRepo:  car,
		fileRepo: file,
	}
}

func (uc *CarUseCase) GetByID(id uint) (entity.Car, error) {
	return uc.carRepo.GetByID(id)
}

func (uc *CarUseCase) List(params dto.ListCarsParams, public bool) ([]entity.Car, error) {
	return uc.carRepo.List(params, public)
}

func (uc *CarUseCase) Create(car *entity.Car) error {
	return uc.carRepo.Save(car)
}

func (uc *CarUseCase) Update(id uint, car entity.Car) error {
	return uc.carRepo.Update(id, car)
}

func (uc *CarUseCase) Delete(ids []uint) error {
	return uc.carRepo.Delete(ids)
}

func (uc *CarUseCase) UploadPhoto(carID uint, file multipart.File, header *multipart.FileHeader) (string, error) {
	url, err := uc.fileRepo.Save(file, header, carID)
	if err != nil {
		return "", err
	}
	if err := uc.carRepo.AddPhoto(carID, url); err != nil {
		return "", err
	}
	return url, nil
}

func (uc *CarUseCase) DeletePhotos(photoURLs []string) error {
    for _, url := range photoURLs {
        parts := strings.Split(url, "/")
        fileName := parts[len(parts)-1]

        if err := uc.fileRepo.Delete(fileName); err != nil {
            return err
        }
        if err := uc.carRepo.DeletePhoto(url); err != nil {
            return err
        }
    }
    return nil
}

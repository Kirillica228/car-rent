package catalog_service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"rent/internal/client/catalog_service/dto"
	"rent/internal/entity"
)

type CarsService interface {
	GetCarByID(id uint) (entity.Car, error)
}

type carsService struct {
	baseURL string
}

func NewCarService(baseURL string) CarsService {
	return &carsService{baseURL: baseURL}
}

func (c *carsService) GetCarByID(id uint) (entity.Car, error) {
    resp, err := http.Get(fmt.Sprintf("%s/cars/%d", c.baseURL, id))
    if err != nil {
        return entity.Car{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return entity.Car{}, fmt.Errorf("error: %s", resp.Status)
    }

    var car dto.Car
    if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
        return entity.Car{}, err
    }
	
	if len(car.Image) == 0 {
    log.Printf("Машина %d (%s) не имеет картинок", car.ID, car.Model)
	}
    // Безопасно достаём картинку
    var image string
    if len(car.Image) > 0 {
        image = car.Image[0].URL
    }

    return entity.Car{
        ID:    car.ID,
        Model: car.Model,
        Price: car.Price,
        Brand: car.Brand.Name,
        Type:  car.Type.Name,
        Image: image, // пустая строка если нет картинок
    }, nil
}

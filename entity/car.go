package entity

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/heroku/cmanager/service"
)

type Car struct {
	Model string
	Year  int
	Plate string
	Brand string
}

func GetAllCars() *[] Car {



	result := [] Car{}
	iter := service.GetCollection("car").Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func GetCarByPlate(plate string) *Car {
	result := Car{}
	service.GetCollection("car").Find(bson.M{"plate": plate}).One(&result)
	return &result
}

func UpsertCar(model string, year int, plate string, brand string) {

	carParam := &Car{model, year, plate, brand}
	carInDb := GetCarByPlate(plate)

	if carInDb.Brand != "" {
		service.GetCollection("car").Update(bson.M{"plate": plate}, carParam)
	} else {
		service.GetCollection("car").Insert(carParam)
	}
}

func DeleteCar(plate string) {
	err := service.GetCollection("car").Remove(bson.M{"plate": plate})
	if err != nil {
		panic(err)
	}
}
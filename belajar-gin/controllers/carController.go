package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string 	`json:"car_id"`
	Brand string  	`json:"brand"`
	Model  string  	`json:"model"`
	Price int		`json:"price"`
}


var CarDatas[] Car{

}

func main() {
	//logika

	
}
//CREATE CAR
func CreateCar (ctx *gin.Context){
	var newCar Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

//UPDATE CAR
func UpdateCar (ctx *gin.Context){
	carID := ctx.Param("carID")
	condition := flase
	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updateCar
			CarDatas[i].carID
			break
		}
	}

	if !condition{
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found ", carID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		ctx.JSON (http.StatusOK, gin.H{
			",message": fmt.Sprintf("car with id %v has been successfully update", carID),
		}),
	})
}

//GET CAR
func GetCar (ctx *gin.Context){
	carID := ctx.Param ("carID")
	condition := false
	var carData Car

	for i, car := range CarDatas {
		if carID == car.CarID{
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status:" : "Data Not Found",
			"error+message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

//DELETE CAR
func DeleteCar (ctx *gin.Context) {
carID := ctx.Param ("carID")
	condition := false
	var carIndex int

	for i, car := range CarDatas {
		if carID == car.CarID{
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status:" : "Data Not Found",
			"error+message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex + 1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas [:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),

	})

}

package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosEduardoAD/Go-ing_on/domain"
	"github.com/CarlosEduardoAD/Go-ing_on/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Routes(route *gin.Engine) {
	rts := route.Group("/api")
	{
		rts.POST("/create", Create)
		rts.GET("/find", Find)
		rts.GET("/findWithEmail/:UserEmail", FindWithEmail)
		rts.PUT("/edit/:UserEmail", Update)
		rts.DELETE("/delete/:UserEmail", Delete)
	}
}

func Create(g *gin.Context) {
	NewUser := &domain.Customer{}

	if err := g.ShouldBindJSON(&NewUser); err != nil {
		log.Println(err)
		return
	}

	res, err := services.InsertUserService(NewUser)

	if err != nil {
		log.Println("err[app] service failed =>", err)
		g.JSON(http.StatusBadRequest, gin.H{"error": "400 - Bad Request"})
	}

	formatedString := fmt.Sprintf("%d - User Created", res)
	g.JSON(http.StatusAccepted, gin.H{"sucess": formatedString})
}

func FindWithEmail(g *gin.Context) {
	filter := &domain.Customer{
		Email: g.Param("UserEmail"),
	}

	res, err := services.FindDataByEmail(filter)

	if err != nil {
		log.Panic(err)
	}

	g.JSON(http.StatusAccepted, gin.H{"sucess": res})
}

func Find(g *gin.Context) {
	res, err := services.FindData()

	if err != nil {
		log.Panic(err)
	}

	g.JSON(http.StatusAccepted, gin.H{"sucess": res})
}

func Update(g *gin.Context) {
	UpdateUser := &domain.Customer{}
	filter := bson.D{{"email", g.Param("UserEmail")}}

	if err := g.ShouldBindJSON(&UpdateUser); err != nil {
		log.Println(err)
		return
	}

	res, err := services.UpdateOneRecordByEmail(UpdateUser, filter)

	if err != nil {
		log.Panic(err)
	}

	g.JSON(http.StatusAccepted, gin.H{"sucess": res})
}

func Delete(g *gin.Context) {
	DeleteUser := &domain.Customer{
		Email: g.Param("UserEmail"),
	}

	if err := g.ShouldBindJSON(&DeleteUser); err != nil {
		log.Println(err)
		return
	}

	res, err := services.DeleteOneRecordByEmail(DeleteUser)

	if err != nil {
		log.Panic(err)
	}

	g.JSON(http.StatusAccepted, gin.H{"sucess": res})
}

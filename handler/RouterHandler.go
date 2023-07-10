package handler
import (
	"strconv"
	"github.com/liujiage/restapi/service"
	"github.com/gin-gonic/gin"
)

//api address
const(
	ADD_USER = "/user/add"
	UPDATE_USER = "/user/update"
	DELETE_USER = "/user/delete"
	QUERY_USER = "/user/:id"
)


/****
Start api http server
****/
func Start(){
	router := gin.Default()
	router.POST(ADD_USER, UserAdd)
	router.DELETE(DELETE_USER, UserDeleteById)
	router.PUT(UPDATE_USER, UserUpdateById)
	router.GET(QUERY_USER, UserQueryById)
    //get app config
	config := service.GetConfig()
	router.Run(config.Ip + ":" +strconv.Itoa(config.Port))
}
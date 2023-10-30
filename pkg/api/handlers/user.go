package handlers

import (
	"context"
	"net/http"
	"strconv"

	client "github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/client/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/models"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/pb"
)

type UserHandler struct {
	Client client.UserClient
}

func NewUserHandler(client client.UserClient) UserHandler {
	return UserHandler{
		Client: client,
	}
}

// LIST USER BY ID
// @Summary API FOR LISTING USER BASED ON ID
// @ID LIST-USER
// @Description LISTING USER BASED ON ID
// @Tags USER
// @Accept json
// @Produce json
// @Param userId path string true "Enter the user id"
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user/get/byid/{userId} [get]
func (cr *UserHandler) GetUserData(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("userId"), 10, 32)
	userId := uint32(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse the id" + err.Error(),
		})
		return
	}
	if userId <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Check the Id",
		})
		return
	}

	response, err := cr.Client.GetUserData(context.Background(), &pb.GetUserRequest{UserId: int32(userId)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Error from handler-client side" + err.Error(),
		})
		return
	}
	c.JSON(int(response.Status), gin.H{
		"Success": response.Result,
	})
}

// UPDATE USER
// @Summary API FOR UPDATING USER
// @ID UPDATE-USER
// @Description UPDATING USER DETAILS WITH ID
// @Tags USER
// @Accept json
// @Produce json
// @Param user_details body models.RequestUserIDList true "Enter the user id's"
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /user/get/bylist [get]
func (cr *UserHandler) GetUserDataList(c *gin.Context) {
	var inputData models.RequestUserIDList
	if err := c.ShouldBindJSON(&inputData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Binding error" + err.Error(),
		})
		return
	}

	if len(inputData.UserId) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Check the Id once again",
		})
		return
	}
	response, err := cr.Client.GetUserDataList(context.Background(), &pb.GetUserDataListRequest{UserIdList: inputData.UserId})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Error from handler-client side " + err.Error(),
		})
		return
	}

	c.JSON(int(response.Status), gin.H{
		"Response":        &response.Response,
		"Users found":     &response.Result,
		"Users not found": &response.NotFound,
	})
}

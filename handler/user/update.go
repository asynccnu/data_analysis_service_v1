package user

import (
	"strconv"

	. "github.com/asynccnu/data_analysis_service_v1/handler"
	"github.com/asynccnu/data_analysis_service_v1/model"
	"github.com/asynccnu/data_analysis_service_v1/pkg/errno"
	"github.com/asynccnu/data_analysis_service_v1/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	var (
		postData model.UserModel
	)

	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	userID, _ := strconv.Atoi(c.Param("id"))

	if user, err := model.GetUserById(uint64(userID)); err != nil {
		SendError(c, errno.ErrUserNotFound, nil, err.Error())
		return
	} else {
		if err := c.Bind(&postData); err != nil {
			SendBadRequest(c, errno.ErrBind, nil, err.Error())
			return
		}

		// 更新数据
		user.Username = postData.Username
		user.Password = postData.Password

		// Validate the data.
		if err := user.Validate(); err != nil {
			SendError(c, errno.ErrValidation, nil, err.Error())
			return
		}

		// Encrypt the user password.
		if err := user.Encrypt(); err != nil {
			SendError(c, errno.ErrEncrypt, nil, err.Error())
			return
		}

		// Save changed fields.
		if err := user.Update(); err != nil {
			SendError(c, errno.ErrDatabase, nil, err.Error())
			return
		}

		SendResponse(c, nil, nil)
	}

}

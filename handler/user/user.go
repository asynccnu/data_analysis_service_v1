package user

import (
	"github.com/asynccnu/data_analysis_service_v1/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

type LoginResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    model.Token `json:"data"`
}

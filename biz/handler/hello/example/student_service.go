// Code generated by hertz generator.

package example

import (
	example "classbackend/biz/model/hello/example"
	"classbackend/db"
	"classbackend/enum"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"
	"log"
	"time"
)

// Hello .
// @router /hello [GET]
func Hello(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(example.HelloResp)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /login [GET]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.LoginReq
	err = c.BindAndValidate(&req)
	log.Println(req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := example.LoginResp{
		Msg:   enum.Success,
		Code:  enum.OK,
		Token: "",
	}
	res := db.DB.Table("login").Where("number = ? and password = ?", req.Number, req.Password).Find(&example.User{})
	if res.RowsAffected == 0 {
		zap.S().Error(enum.NoUserID)
		c.JSON(consts.StatusOK, example.LoginResp{
			Msg:   enum.Fail,
			Code:  enum.Error,
			Token: "",
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// Register .student
// @router /register [GET]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.RegisterReq
	//TODO 还需要根据时间线来生成学号
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	lastTenDigits := fmt.Sprintf("%010d", timestamp)
	stringLastTenDigits := lastTenDigits[len(lastTenDigits)-10:]
	fmt.Println(stringLastTenDigits)
	err = c.BindAndValidate(&req)
	user := example.User{UserName: req.UserName, Password: req.Password, Number: stringLastTenDigits, Role: 1}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(example.RegisterResp)
	res := db.DB.Table("login").Create(&user)
	if res.RowsAffected == 0 {
		zap.S().Error(enum.SystemError)
		c.JSON(consts.StatusOK, example.RegisterResp{
			Msg:   enum.SystemError,
			Code:  enum.Error,
			Token: "",
		})
		return
	}
	resp.Msg = enum.Success
	resp.Code = enum.OK
	c.JSON(consts.StatusOK, resp)
}

// QueryStudents .
// @router /querystudents [GET] TODO 查询学生信息 需要校验的是老师和管理员的权限。还有分页的操作
func QueryStudents(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.QueryAllStudentsReq

	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(example.QueryAllStudentsResp)
	res := db.DB.Table("login").Where("role = ?", 1).Find(&resp.Students)
	if res.RowsAffected == 0 {
		zap.S().Error(enum.SystemError)
		c.JSON(consts.StatusOK, example.QueryPersonDetailResp{
			Msg:  enum.SystemError,
			Code: enum.Error,
			User: nil,
		})
		return
	}
	//success
	resp.Msg = enum.Success
	resp.Code = enum.OK
	c.JSON(consts.StatusOK, resp)
}

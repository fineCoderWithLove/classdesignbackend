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
	"sync"
	"time"
)

// QueryPersonDetail .
// @router /person/details [GET]
func QueryPersonDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.QueryPersonDetailReq
	err = c.BindAndValidate(&req)
	//参数校验
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := example.QueryPersonDetailResp{
		Msg:  enum.Success,
		Code: enum.OK,
		User: nil,
	}
	res := db.DB.Table("login").Where("user_id = ?", req.UserID).Find(&resp.User)
	if res.RowsAffected == 0 {
		zap.S().Error(enum.NoUserID)
		c.JSON(consts.StatusOK, example.QueryPersonDetailResp{
			Msg:  enum.Fail,
			Code: enum.Error,
			User: nil,
		})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

/*
删除老师或者学生只需要通过权限来判断
*/
var (
	mutex sync.Mutex
)

// AddPerson .
// @router /person/add [GET]
func AddPerson(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.AddStudentReq
	err = c.BindAndValidate(&req)
	// 加锁，确保同一时间只有一个goroutine可以生成时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	lastTenDigits := fmt.Sprintf("%010d", timestamp)
	stringLastTenDigits := lastTenDigits[len(lastTenDigits)-10:]
	err = c.BindAndValidate(&req)
	user := example.Person{}
	if req.Role == "1" {
		user = example.Person{UserName: req.Person.UserName, Password: req.Person.Password, Number: stringLastTenDigits, Role: 1, Gender: req.Person.Gender, FromWhere: req.Person.FromWhere, Email: req.Person.Email, Tel: req.Person.Tel}
	}
	if req.Role == "2" {
		user = example.Person{UserName: req.Person.UserName, Password: req.Person.Password, Number: stringLastTenDigits, Role: 2, Gender: req.Person.Gender, FromWhere: req.Person.FromWhere, Email: req.Person.Email, Tel: req.Person.Tel}
	}
	log.Println(user)
	if err != nil {
		zap.S().Error(err.Error())
		return
	}
	res := db.DB.Table("login").Create(&user)
	if res.RowsAffected == 0 {
		zap.S().Error(enum.SystemError)
		c.JSON(consts.StatusOK, example.AddStudentResp{
			Msg:  enum.SystemError,
			Code: enum.Error,
		})
		return
	}
	resp := new(example.AddStudentResp)
	resp.Msg = enum.Success
	resp.Code = enum.OK
	c.JSON(consts.StatusOK, resp)
}

// DelPerson .
// @router /person/del [GET]
func DelPerson(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.DelStudentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	res := db.DB.Table("login").Delete(&example.Person{}, req.UserID)
	if res.RowsAffected == 0 {
		zap.S().Error(enum.NoUserID)
		c.JSON(consts.StatusOK, example.DelStudentResp{
			Msg:  enum.Fail,
			Code: enum.Error,
		})
		return
	}
	resp := new(example.DelStudentResp)
	resp.Msg = enum.Success
	resp.Code = enum.OK
	c.JSON(consts.StatusOK, resp)
}

// UpdatePerson .
// @router /person/update [GET]
func UpdatePerson(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.UpdateStudentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	res := db.DB.Table("login").Where("user_id = ?", req.Person.UserID).Save(&req.Person)
	if res.RowsAffected == 0 {
		zap.S().Error(enum.NoUserID)
		c.JSON(consts.StatusOK, example.UpdateStudentResp{
			Msg:  enum.Fail,
			Code: enum.Error,
		})
		return
	}
	resp := new(example.UpdateStudentResp)
	resp.Msg = enum.Success
	resp.Code = enum.OK
	c.JSON(consts.StatusOK, resp)
}

// SearchForPersonReq .
// @router /person/search [GET]
func SearchForPersonReq(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.SearchForPersonReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(example.SearchForPersonResp)

	if req.Role == "1" {
		//学生查询
		res := db.DB.Table("login").Where("user_name like ? and role = 1", req.UserName+"%").Find(&resp.Person)
		if res.RowsAffected == 0 {
			zap.S().Error(enum.NoUserID)
			c.JSON(consts.StatusOK, example.UpdateStudentResp{
				Msg:  enum.Fail,
				Code: enum.Error,
			})
			return
		}
		resp := new(example.SearchForPersonResp)
		resp.Msg = enum.Success
		resp.Code = enum.OK
	} else if req.Role == "2" {
		//教师查询
		res := db.DB.Table("login").Where("user_name like ? and role = 2", req.UserName+"%").Find(&resp.Person)
		if res.RowsAffected == 0 {
			zap.S().Error(enum.NoUserID)
			c.JSON(consts.StatusOK, example.UpdateStudentResp{
				Msg:  enum.Fail,
				Code: enum.Error,
			})
			return
		}
		resp := new(example.SearchForPersonResp)
		resp.Msg = enum.Success
		resp.Code = enum.OK
	}
	c.JSON(consts.StatusOK, resp)
}

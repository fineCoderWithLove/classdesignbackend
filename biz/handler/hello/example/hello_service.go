// Code generated by hertz generator.

package example

import (
	"context"

	example "classbackend/biz/model/hello/example"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
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
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(example.LoginResp)

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /register [GET]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(example.RegisterResp)

	c.JSON(consts.StatusOK, resp)
}

// QueryStudents .
// @router /querystudents [GET]
func QueryStudents(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.QueryAllStudentsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(example.QueryAllStudentsResp)

	c.JSON(consts.StatusOK, resp)
}

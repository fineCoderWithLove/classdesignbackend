// Code generated by hertz generator. DO NOT EDIT.

package example

import (
	example "classbackend/biz/handler/hello/example"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/hello", append(_helloMw(), example.Hello)...)
	root.GET("/login", append(_loginMw(), example.Login)...)
	root.GET("/querystudents", append(_querystudentsMw(), example.QueryStudents)...)
	root.GET("/register", append(_registerMw(), example.Register)...)
	{
		_person := root.Group("/person", _personMw()...)
		_person.POST("/add", append(_addpersonMw(), example.AddPerson)...)
		_person.GET("/del", append(_delpersonMw(), example.DelPerson)...)
		_person.GET("/details", append(_querypersondetailMw(), example.QueryPersonDetail)...)
		_person.GET("/search", append(_searchforpersonreqMw(), example.SearchForPersonReq)...)
		_person.POST("/update", append(_updatepersonMw(), example.UpdatePerson)...)
	}
	{
		_tech := root.Group("/tech", _techMw()...)
		_tech.GET("/classstu", append(_selectclassstuMw(), example.SelectClassStu)...)
		_tech.GET("/mycourse", append(_selectmytechcourseMw(), example.SelectMyTechCourse)...)
		_tech.GET("/queryclass", append(_selectclassbycourseidMw(), example.SelectClassByCourseId)...)
		_tech.POST("/score", append(_ratescoreMw(), example.RateScore)...)
	}
}

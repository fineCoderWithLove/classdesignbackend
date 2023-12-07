// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct HelloResp {
    1: string RespBody;
}

//1.登录的请求
struct LoginReq {
    1: string number (api.query="number"); // 添加 api 注解为方便进行参数绑定
    2: string password (api.query="password"); // 添加 api 注解为方便进行参数绑定
}

struct LoginResp {
    1: string msg;
    2: i64 code;
    3: string token;
}
//2注册的请求
struct RegisterReq {
    1: string user_name (api.query="userName"); // 添加 api 注解为方便进行参数绑定
    2: string password (api.query="password"); // 添加 api 注解为方便进行参数绑定
}

struct RegisterResp {
    1: string msg;
    2: i64 code;
    3: string token;
}
//3.查询所有学生
struct QueryAllStudentsReq {
    1: string token (api.query="token"); // 添加 api 注解为方便进行参数绑定
}

struct QueryAllStudentsResp {
    1: string msg;
    2: i64 code;
    3: list<Student> students;
}
//=======================================>1.学生的服务的服务信息<=============================================================
service StudentService {
    HelloResp Hello(1: HelloReq request) (api.get="/hello");
    LoginResp  Login(1: LoginReq request) (api.get="/login");
    RegisterResp  Register(1: RegisterReq request) (api.get="/register");
    QueryAllStudentsResp  QueryStudents(1: QueryAllStudentsReq request) (api.get="/querystudents");
}
//=======================================>2.老师的服务信息<=============================================================
//1.update某学生评分
struct RateScoreReq {
    1: string token (api.body="token"); // 添加 api 注解为方便进行参数绑定
    2: RateItem rateItem (api.body="rateItem");
}

struct RateScoreResp {
    1: string msg;
    2: i64 code;
}
//2.查询自己所带的课程
struct SelectMyTechCourseReq {
    1: string token (api.get="token"); // 添加 api 注解为方便进行参数绑定
    2: i64 user_id (api.get="userId");
}

struct SelectMyTechCourseResp {
    1: string msg;
    2: i64 code;
    3: list<Course> course;
}
//3.根据课程id查询，课程内的班级
struct SelectClassByCourseIdReq {
    1: string token (api.get="token"); // 添加 api 注解为方便进行参数绑定
    2: i64 course_id (api.get="courseId");
}

struct SelectClassByCourseIdResp {
    1: string msg;
    2: i64 code;
    3: list<string> from_where;
}
//4.根据班级来查询学生的信息tips:注意要查询的是该课程内的学生
struct SelectClassStuReq {
    1: string token (api.get="token"); // 添加 api 注解为方便进行参数绑定
    2: string from_where (api.get="fromWhere");
}

struct SelectClassStuResp {
    1: string msg;
    2: i64 code;
    3: list<RateItem> rateItem;
}
service TeacherService {
    RateScoreResp RateScore(1: RateScoreReq request) (api.post="/tech/score");
    SelectMyTechCourseResp SelectMyTechCourse(1: SelectMyTechCourseReq request) (api.get="/tech/mycourse");
    SelectClassByCourseIdResp SelectClassByCourseId(1: SelectClassByCourseIdReq request) (api.get="/tech/queryclass");
    SelectClassStuResp SelectClassStu(1: SelectClassStuReq request) (api.get="/tech/classstu");
}
//=======================================>3.管理员的服务信息<=============================================================
//1.添加学生
//2.添加老师
//3.查询学生详情
//4.查询老师详情
//5.删除学生
//6.删除老师
 //=======================================>4.个人服务信息的服务信息<=============================================================
 //1.查询个人详情
 struct QueryPersonDetailReq {
     1: string token (api.body="token"); // 添加 api 注解为方便进行参数绑定
     2: i64 user_id (api.body="userId");
 }

 struct QueryPersonDetailResp {
     1: string msg;
     2: i64 code;
     3: User user;
 }
 //2.添加学生/老师
  struct AddStudentReq {
      1: string token (api.query="token"); // 添加 api 注解为方便进行参数绑定
      2: Person person (api.query="person");
  }

  struct AddStudentResp {
      1: string msg;
      2: i64 code;
  }
 //3.删除学生/老师
  struct DelStudentReq {
      1: string token (api.query="token"); // 添加 api 注解为方便进行参数绑定
        2: i64 user_id (api.query="userId");
  }

  struct DelStudentResp {
      1: string msg;
      2: i64 code;
  }
   //3.更新学生/老师
    struct UpdateStudentReq {
        1: string token (api.body="token"); // 添加 api 注解为方便进行参数绑定
        2: Person person (api.body="person");
    }

    struct UpdateStudentResp {
        1: string msg;
        2: i64 code;
    }
   //4.搜索的信息老师/学生
    struct SearchForPersonReq {
        1: string token (api.query="token"); // 添加 api 注解为方便进行参数绑定
        2: string user_name (api.query="userName");
        3: string role (api.query="role");
    }
    struct SearchForPersonResp {
        1: string msg;
        2: i64 code;
        3: list<Person> person;
    }
 service AdminService {
        QueryPersonDetailResp QueryPersonDetail(1: QueryPersonDetailReq request) (api.get="/person/details");
        AddStudentResp AddPerson(1: AddStudentReq request) (api.post="/person/add");
        DelStudentResp DelPerson(1: DelStudentReq request) (api.get="/person/del");
        UpdateStudentResp UpdatePerson(1: UpdateStudentReq request) (api.post="/person/update");
        SearchForPersonResp SearchForPersonReq(1: SearchForPersonReq request) (api.get="/person/search");
 }
 struct Course {
       1:i64 course_id;
       2:string course_name;
 }

 // 仅仅供评分的一个item
  struct RateItem {
      1:i64 user_id;
      2:string  course_name;
      3:string  user_name;
      4:string course_total_score;
      5:string course_test;
      6:string course_normal;
      7:i64 course_id;
  }
  struct Person {
      1:i64 user_id;
      2:string  user_name;
      3:string  password;
      4:string avatar;
      5:i64 role;
      6:string number;
      7:string email;
      8:string gender;
      9:string from_where;
  }
 struct User {
     1:i64 user_id;
     2:string  user_name;
     3:string  password;
     4:string avatar;
     5:i64 role;
     6:string number;
     7:string email;
     8:string gender;
     9:string from_where;
 }
struct Student {
    1:i64 user_id;
    2:string  user_name;
    3:string  password;
    4:string avatar;
    5:i64 role;
    6:string number;
    7:string email;
    8:string gender;
    9:string from_where;
}
struct Teacher {
    1:i64 user_id;
    2:string  user_name;
    3:string  password;
    4:string avatar;
    5:i64 role;
    6:string number;
    7:string email;
    8:string gender;
    9:string from_where;
}
struct Admin {
    1:i64 user_id;
    2:string  user_name;
    3:string  password;
    4:string avatar;
    5:i64 role;
    6:string number;
    7:string email;
    8:string gender;
    9:string from_where;
}

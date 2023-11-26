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
    1: string number (api.query="number"); // 添加 api 注解为方便进行参数绑定
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
service HelloService {
    HelloResp Hello(1: HelloReq request) (api.get="/hello");
    LoginResp  Login(1: LoginReq request) (api.get="/login");
    RegisterResp  Register(1: RegisterReq request) (api.get="/register");
    QueryAllStudentsResp  QueryStudents(1: QueryAllStudentsReq request) (api.get="/querystudents");
}
struct Student {
    1:i64 user_id;
    2:string  user_name;
    3:string  password;
    4:string avatar;
}

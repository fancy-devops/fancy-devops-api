package controller

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/fancy-devops/fancy-devops-api/model/codes"
	"github.com/fancy-devops/fancy-devops-api/model/db"
	"github.com/fancy-devops/fancy-devops-api/model/requests"
	"github.com/fancy-devops/fancy-devops-api/model/responses"
	"github.com/fancy-devops/fancy-devops-api/service"
	"github.com/fancy-devops/fancy-devops-api/utils/common"
	"github.com/fancy-devops/fancy-devops-api/utils/redis"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

// @Tags User
// @Summary 发送认证邮件
// @Accept json
// @Produce json
// @Param "" body requests.SendVerifyEmailReq true "发送认证邮件"
// @Success 200 {object} responses.Result "成功"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/sendverifyemail [POST]
func (obj *User) SendVerifyEmail(c *gin.Context) {
	// 发送认证邮件
	var req requests.SendVerifyEmailReq
	err := c.BindJSON(&req)
	if err != nil {
		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, "参数错误")
		return
	}
	valid := validation.Validation{}
	valid.Email(req.Email, "Email").Message("邮箱格式错误；")
	if valid.HasErrors() {
		var errStr string
		for _, err := range valid.Errors {
			errStr += err.Message
		}

		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, errStr)
		return
	}
	n := common.NewCommon().GetRandomNumber(999999)

	redis.RedisClient.Set(redis.Prefix+req.Email, n, 5*60*1000*1000*1000)
	common.NewEmail().SendEmail(req.Email, "【fancy-devops】验证码", fmt.Sprintf("您的验证码为：%d", n))
	common.NewGinRes().SuccessResponse(c, "")
}

// @Tags User
// @Summary 重置密码
// @Accept json
// @Produce json
// @Param "" body requests.UserResetPwdReq true "重置密码"
// @Success 200 {object} responses.Result "成功"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/resetpwd [POST]
func (obj *User) ResetPwd(c *gin.Context) {
	// 重置密码，并将新密码发送给对应邮箱
	var req requests.UserResetPwdReq
	err := c.BindJSON(&req)
	if err != nil {
		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, "参数错误")
		return
	}
	valid := validation.Validation{}
	valid.Email(req.Email, "Email").Message("邮箱格式错误；")
	valid.Required(req.Code, "Code").Message("验证码不能为空；")
	if !valid.HasErrors() {
		authCode := redis.RedisClient.Get(redis.Prefix + req.Email).Val()
		if authCode != req.Code {
			valid.AddError("code", "验证码错误；")
		}
	}
	if valid.HasErrors() {
		var errStr string
		for _, err := range valid.Errors {
			errStr += err.Message
		}

		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, errStr)
		return
	}
	user := service.NewUser().GetUserByEmail(req.Email)
	if user.ID <= 0 {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_NotExist)
		return
	}
	newPwd := common.NewCommon().GetRandomString(12)
	common.NewEmail().SendEmail(req.Email, "【fancy-devops】重置密码", "您的新密码为："+newPwd)
	service.NewUser().UpdateUserPwd(req.Email, newPwd)
	common.NewGinRes().SuccessResponse(c, "")
}

// @Tags User
// @Summary 修改密码
// @Accept json
// @Produce json
// @Param "" body requests.UserChangePwdReq true "修改密码"
// @Param token header string true "token"
// @Success 200 {object} responses.Result "成功"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/changepwd [POST]
func (obj *User) ChangePwd(c *gin.Context) {
	// 修改密码
	var req requests.UserChangePwdReq
	err := c.BindJSON(&req)
	if err != nil {
		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, "参数错误")
		return
	}
	claims, err := common.NewJwt().GetClaims(c)
	if err != nil {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_AuthTokenParseError)
		return
	}

	valid := validation.Validation{}
	valid.Required(req.OldPwd, "OldPwd").Message("旧密码不能为空；")
	valid.Required(req.NewPwd, "NewPwd").Message("新密码不能为空；")
	valid.Length(req.NewPwd, 8, "NewPwd").Message("新密码长度不能低于8位；")
	if valid.HasErrors() {
		var errStr string
		for _, err := range valid.Errors {
			errStr += err.Message
		}

		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, errStr)
		return
	}

	user := service.NewUser().GetUserByPassword(claims.Email, req.OldPwd)
	if user.ID <= 0 {
		common.NewGinRes().ErrorResponse2(c, codes.Error_User_WrongPassword, "旧密码错误")
		return
	}

	service.NewUser().UpdateUserPwd(claims.Email, req.NewPwd)
	common.NewGinRes().SuccessResponse(c, "")
}

// @Tags User
// @Summary 注册
// @Accept json
// @Produce json
// @Param "" body requests.UserRegisterReq true "注册"
// @Success 200 {object} responses.IdResult "成功"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/register [POST]
func (obj *User) Register(c *gin.Context) {
	// 注册（要校验邮箱验证码）
	var req requests.UserRegisterReq
	err := c.BindJSON(&req)
	if err != nil {
		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, "参数错误")
		return
	}
	valid := validation.Validation{}
	valid.Email(req.Email, "Email").Message("邮箱格式错误；")
	valid.Required(req.Name, "Code").Message("用户名不能为空；")
	valid.Required(req.Code, "Code").Message("验证码不能为空；")
	valid.Length(req.Pwd, 8, "Pwd").Message("密码长度不能低于8位；")
	if !valid.HasErrors() {
		authCode := redis.RedisClient.Get(redis.Prefix + req.Email).String()
		if authCode != req.Code {
			valid.AddError("code", "验证码错误；")
		}
	}
	if valid.HasErrors() {
		var errStr string
		for _, err := range valid.Errors {
			errStr += err.Message
		}

		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, errStr)
		return
	}
	user := service.NewUser().GetUserByEmail(req.Email)
	if user.ID > 0 {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_Exist)
		return
	}
	newId := service.NewUser().CreateUser(req.Name, req.Email, req.Pwd, 0)

	service.NewUser().SendSecretEmail(req.Email)

	common.NewGinRes().SuccessResponse(c, newId)
}

// @Tags User
// @Summary 登录
// @Accept json
// @Produce json
// @Param "" body requests.UserLoginReq true "登录"
// @Success 200 {object} responses.UserLoginRes "成功"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/login [POST]
func (obj *User) Login(c *gin.Context) {
	// 登录
	var req requests.UserLoginReq
	err := c.BindJSON(&req)
	if err != nil {
		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, "参数错误")
		return
	}
	valid := validation.Validation{}
	valid.Required(req.Email, "Email").Message("邮箱不能为空；")
	valid.Required(req.Pwd, "Pwd").Message("密码不能为空；")
	if valid.HasErrors() {
		var errStr string
		for _, err := range valid.Errors {
			errStr += err.Message
		}

		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, errStr)
		return
	}

	user := service.NewUser().GetUserByPassword(req.Email, req.Pwd)
	if user.ID <= 0 {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_WrongPassword)
		return
	}
	if user.Role == 0 {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_Guest)
		return
	}

	if user.Secret == "" {
		service.NewUser().SendSecretEmail(req.Email)
	}

	token, err := common.NewJwt().GenerateToken(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_AuthTokenGenFailed)
		return
	}

	data := responses.UserLoginModel{Token: token}
	common.NewGinRes().SuccessResponse(c, data)
}

// @Tags User
// @Summary 登录
// @Accept json
// @Produce json
// @Param "" body requests.UserLoginReq true "登录"
// @Success 200 {object} responses.UserLoginRes "成功"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/glogin [POST]
func (obj *User) LoginByGoogle(c *gin.Context) {
	// 登录
	var req requests.UserLoginReq
	err := c.BindJSON(&req)
	if err != nil {
		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, "参数错误")
		return
	}
	valid := validation.Validation{}
	valid.Required(req.Email, "Email").Message("邮箱不能为空；")
	valid.Required(req.Pwd, "Pwd").Message("密码不能为空；")
	if valid.HasErrors() {
		var errStr string
		for _, err := range valid.Errors {
			errStr += err.Message
		}

		common.NewGinRes().ErrorResponse2(c, codes.Common_BadRequest, errStr)
		return
	}

	user := service.NewUser().GetUserByEmail(req.Email)
	if user.ID <= 0 {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_NotExist)
		return
	}

	verifyRes, err := common.NewGoogleAuth().VerifyCode(user.Secret, req.Pwd)
	if err != nil {
		common.NewGinRes().ErrorResponse2(c, codes.Error_Failed, "code校验异常")
		return
	}
	if !verifyRes {
		common.NewGinRes().ErrorResponse2(c, codes.Error_User_WrongCode, "code校验不通过")
		return
	}

	if user.Role == 0 {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_Guest)
		return
	}

	token, err := common.NewJwt().GenerateToken(user.ID, user.Name, user.Email, user.Role)
	if err != nil {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_AuthTokenGenFailed)
		return
	}

	data := responses.UserLoginModel{Token: token}
	common.NewGinRes().SuccessResponse(c, data)
}

// @Tags User
// @Summary 获取用户详情
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param token header string true "token"
// @Success 200 {object} responses.UserDetailModel "获取用户详情"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/detail/{id} [GET]
func (obj *User) GetUserDetail(c *gin.Context) {
	// 获取用户详情
	id := com.StrTo(c.Param("id")).MustInt()

	var data responses.UserDetailModel

	user := service.NewUser().GetUser(id)
	if user.ID > 0 {
		data.ID = user.ID
		data.Name = user.Name
		data.Email = user.Email
		data.Role = user.Role
	}

	common.NewGinRes().SuccessResponse(c, data)
}

// @Tags User
// @Summary 获取用户列表
// @Accept json
// @Produce json
// @Param "" query requests.GetUserListReq false "请求参数"
// @Param token header string true "token"
// @Success 200 {object} responses.PagedUserList "获取用户列表"
// @Failure 400 {object} responses.Result "错误的请求"
// @Failure 401 {object} responses.Result "未登录"
// @Failure 403 {object} responses.Result "无权限"
// @Failure 500 {object} responses.Result "系统异常"
// @Router /api/user/list [GET]
func (obj *User) GetUserList(c *gin.Context) {
	// 获取用户列表
	// 校验权限
	claims, err := common.NewJwt().GetClaims(c)
	if err != nil {
		common.NewGinRes().ErrorResponse(c, codes.Error_User_AuthTokenParseError)
		return
	}
	if claims.Role != db.Role_Admin {
		common.NewGinRes().ErrorResponse(c, codes.Common_Forbidden)
		return
	}

	role := c.Query("role")

	maps := make(map[string]interface{})

	if role != "" {
		maps["role"] = role
	}

	var data responses.PagedUserList
	skip, size := common.NewPagination().GetPager(c)
	userList := service.NewUser().GetUsers(skip, size, maps)

	for _, u := range userList {
		data.List = append(data.List, responses.UserDetailModel{ID: u.ID, Name: u.Name, Email: u.Email, Role: u.Role})
	}
	data.Total = service.NewUser().GetUserTotal(maps)

	common.NewGinRes().SuccessResponse(c, data)
}

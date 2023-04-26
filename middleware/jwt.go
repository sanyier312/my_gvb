package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"my_gvb/utils"
	"my_gvb/utils/errmsg"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte // 密钥
}

func NewJWT() *JWT {
	return &JWT{ // 生成密钥
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct { // 自定义声明
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录") // errors.New() 生成一个错误
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)

// CreateToken 生成token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token     传入token  返回自定义声明和错误
func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return j.JwtKey, nil
		})
	//  判断err
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok { // 类型断言  判断err是否是ValidationError类型  如果是则赋值给ve  ok为true
			if ve.Errors&jwt.ValidationErrorMalformed != 0 { // 判断错误类型
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	//  判断token
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc { // 返回一个gin.HandlerFunc类型的函数
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization") // 获取请求头中的Authorization字段
		if tokenHeader == "" {                               // 判断token是否存在
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort() // 中断请求
			return
		}

		checkToken := strings.Split(tokenHeader, " ") // 以空格分割token
		if len(checkToken) == 0 {                     // 判断token是否正确
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" { //   Bearer是jwt的认证头部  一般都是这样写的
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJWT() // 实例化jwt
		// 解析token
		claims, err := j.ParserToken(checkToken[1]) // 传入token  返回自定义声明和错误
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  errmsg.ERROR,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims) //  将claims存入gin的上下文中 以便后续使用
		c.Next()                  // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

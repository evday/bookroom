package token

import (
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"bookroom/pkg/errno"
	"bookroom/util"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token.
type Context struct {
	ID int64
	Username string
	Isadmin bool
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc{
	return func(token *jwt.Token) (interface{},error){
		if _,ok := token.Method.(*jwt.SigningMethodHMAC);!ok {
			return nil,jwt.ErrSignatureInvalid
		}
		return []byte(secret),nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString,secret string)(*Context,error){
	ctx := &Context{}
	conn := util.Get()
	defer conn.Close()
	//Parse the token
	token,err := jwt.Parse(tokenString,secretFunc(secret))
	//Parse the error
	if err != nil {
		return ctx,err
	//Read the token if it's valid
	}else if claims,ok := token.Claims.(jwt.MapClaims);ok && token.Valid{
		ctx.ID = int64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		redisToken,err := redis.String(conn.Do("GET","token:"+fmt.Sprintf("%d",ctx.ID)))
		if err != nil || redisToken != tokenString {
			return ctx,errno.ErrTokenInvilid
		}
		return ctx,nil
	//Other errors
	}else {
		return ctx,err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context)(*Context,error){
	header := c.Request.Header.Get("Authorization")
	//Load the jwt secret from config
	secret := viper.GetString("jwt_secret")
	if len(header) == 0 {
		return &Context{},ErrMissingHeader
	}
	return Parse(header,secret)
}

//Sign signs the context with the specified secret
func Sign(ctx *gin.Context,c Context,secret string)(tokenString,username string,isAdmin bool,err error){
	//Load the jwt secret from ght Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	//The token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id":c.ID,
		"username":c.Username,
		"nbf":time.Now().Unix(),
		"iat":time.Now().Unix(),
	})

	tokenString,err = token.SignedString([]byte(secret))
	return tokenString,c.Username,c.Isadmin,err
}
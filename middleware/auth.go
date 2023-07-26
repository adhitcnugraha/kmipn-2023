package middleware

import (
	"kmipn-2023/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		url := ctx.Request.Host + ctx.Request.URL.Path
		jwtClaims := &model.Claims{}
		cookie, err := ctx.Cookie("session_token")

		if err != nil {
			if url != "/" {
				ctx.String(http.StatusUnauthorized, "Cookie doesn't Exist")
				ctx.Abort()
				return
			}
			if ctx.Request.Header.Get("Content-Type") == "application/json" {
				ctx.String(http.StatusUnauthorized, "Unauthorized")
				ctx.Abort()
				return
			}
			ctx.Redirect(http.StatusSeeOther, "/client/login")
			ctx.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(cookie, jwtClaims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil || !token.Valid {
			ctx.String(http.StatusBadRequest, "Bad Request")
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			ctx.Set("email", claims.Email)
		} else {
			ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("Unauthorized"))
			ctx.Abort()
			return
		}

		ctx.Next()
	})
}

package api

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	zxcvbn "github.com/trustelem/zxcvbn"
)

var API_KEY string

func init() {
	API_KEY = os.Getenv("GO_ENV")
	if len(API_KEY) < 8 {
		panic("bad API_KEY")
	}
	strength := zxcvbn.PasswordStrength(API_KEY, []string{})
	if strength.Score < 3 {
		panic("low strength API_KEY - please use a more secure api key")
	}
	spew.Dump(strength)
}

func CheckApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_key, ok := c.Request.Header.Get("X-Api-Key")
		if !ok || api_key == "" {
			c.JSON(422, gin.H{
				"message": "no api key",
			})
			c.Abort()
			return
		}
		if api_key != API_KEY {
			c.JSON(422, gin.H{
				"message": "incorrect api key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

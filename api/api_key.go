package api

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs-pro/sshkeymanager/config"
	zxcvbn "github.com/trustelem/zxcvbn"
)

var API_KEY string

const MIN_LENGTH = 10
const MIN_STRENGTH = 3

func init() {
	API_KEY = config.Config.ApiKey
	if len(API_KEY) < MIN_LENGTH {
		panic("bad API_KEY - minimum " + strconv.Itoa(MIN_LENGTH) + " characters")
	}

	strength := zxcvbn.PasswordStrength(API_KEY, []string{"go", "docker", "exec", "create", "your", "api", "key", "create-your-key"})
	if strength.Score < MIN_STRENGTH {
		log.Println("key strength score (from zxcvbn):", strength.Score)
		log.Println("minimum strength: ", strconv.Itoa(MIN_STRENGTH))
		panic("low strength API_KEY - please use a more secure api key")
	}
}

func CheckApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_key := c.Request.Header.Get("X-Api-Key")
		if api_key == "" {
			c.JSON(410, gin.H{
				"message": "no api key",
			})
			c.Abort()
			return
		}
		if api_key != API_KEY {
			// Avoid returning bad api key too fast, so it's less easy to brute force
			time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
			c.JSON(410, gin.H{
				"message": "incorrect api key",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

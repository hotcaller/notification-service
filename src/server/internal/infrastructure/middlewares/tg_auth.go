package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/url"
	"service/internal/infrastructure/config"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UnifiedAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.GetTelegram()
		secret := config.GetJWT().Secret
		if isValidJWT(c, secret) {
			c.Next()
			return
		}

		// Попытка аутентификации через Telegram Login Widget
		if isValidTelegramUser(c, cfg.BotToken) {
			c.Next()
			return
		}

		// Попытка аутентификации бота
		if isValidBotRequest(c, cfg.BotSecret) {
			c.Next()
			return
		}

		// Если обе проверки не прошли
		c.AbortWithStatusJSON(401, gin.H{"error": "Неавторизованный доступ"})
	}
}

func isValidTelegramUser(c *gin.Context, botToken string) bool {
	// Извлекаем данные аутентификации из заголовка
	authHeader := c.GetHeader("X-Telegram-Auth")
	if authHeader == "" {
		return false
	}

	// Парсим данные
	data, err := url.ParseQuery(authHeader)
	if err != nil {
		return false
	}

	// Проверяем подпись данных
	if !validateTelegramAuthData(data, botToken) {
		return false
	}

	// Проверяем срок действия auth_date
	authDateStr := data.Get("auth_date")
	authDateInt, err := strconv.ParseInt(authDateStr, 10, 64)
	if err != nil {
		return false
	}
	authDate := time.Unix(authDateInt, 0)
	if time.Since(authDate) > 24*time.Hour {
		return false
	}

	// Устанавливаем идентификатор пользователя в контекст
	userID := data.Get("id")
	c.Set("telegram_user_id", userID)

	return true
}

func isValidJWT(c *gin.Context, jwtSecret string) bool {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return false
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return false
	}

	tokenString := parts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи токена")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID, ok := claims["sub"].(string)
		if !ok {
			return false
		}
		c.Set("telegram_user_id", userID)
		return true
	} else {
		return false
	}
}

func validateTelegramAuthData(data url.Values, token string) bool {
	checkHash := data.Get("hash")
	if checkHash == "" {
		return false
	}
	data.Del("hash")
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var dataStrings []string
	for _, key := range keys {
		value := data.Get(key)
		dataStrings = append(dataStrings, key+"="+value)
	}
	dataCheckString := strings.Join(dataStrings, "\n")

	secretKey := sha256.Sum256([]byte(token))
	h := hmac.New(sha256.New, secretKey[:])
	h.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(h.Sum(nil))

	return calculatedHash == checkHash
}

func isValidBotRequest(c *gin.Context, botSecret string) bool {
	// Проверяем секретный токен бота
	token := c.GetHeader("X-Bot-Token")
	if token != botSecret {
		return false
	}

	// Получаем идентификатор пользователя из заголовка
	userID := c.GetHeader("X-Telegram-User-ID")
	if userID == "" {
		return false
	}

	// Устанавливаем идентификатор пользователя в контекст
	c.Set("telegram_user_id", userID)

	return true
}

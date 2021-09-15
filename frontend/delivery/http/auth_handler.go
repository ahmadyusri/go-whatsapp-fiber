package http

import (
	"database/sql"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ahmadyusri/go-whatsapp-fiber/domain"
	"github.com/ahmadyusri/go-whatsapp-fiber/utils"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthHandler struct {
	Validate *validator.Validate
}

func NewAuthHandler(rPublic fiber.Router) {
	handler := &WhatsappHandler{
		Validate: utils.NewValidator(),
	}

	rAuth := rPublic.Group("/")
	rAuth.Post("/auth", handler.Auth)
}

// Auth func for get JWT Token.
// @Summary get JWT Token
// @Description Get JWT Token.
// @Tags Auth
// @Produce json
// @Param X-SERVER-KEY header string true "Server Key Whatsapp API."
// @Param Authorization header string true "Token Super Admin User."
// @Success 200 {object} domain.JSONResult{data=object,message=string} "Description"
// @Failure 422 {object} []domain.HTTPErrorValidation
// @Failure 400 {object} domain.HTTPError
// @Failure 401 {object} domain.JSONResult{message=string}
// @Failure 404 {object} domain.HTTPError
// @Failure 500 {object} domain.HTTPError
// @Router /v1/auth [post]
func (w *WhatsappHandler) Auth(c *fiber.Ctx) error {

	SERVER_KEY := c.Get("X-SERVER-KEY")
	if SERVER_KEY != os.Getenv("SERVER_KEY") {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.JSONResult{
			Code:    400,
			Message: "Anda tidak memiliki akses",
		})
	}
	Authorization := strings.Split(c.Get("Authorization"), " ")
	if Authorization[0] != "Bearer" || len(Authorization) < 1 {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.JSONResult{
			Code:    400,
			Message: "Format Authorization tidak valid",
		})
	}
	api_token := Authorization[1]

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	// if there is an error opening the connection, handle it
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.JSONResult{
			Code:    500,
			Message: err.Error(),
		})
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var tag Tag
	// perform a db.Query select
	err = db.QueryRow("select "+os.Getenv("DB_TABLE_FIELD_ID")+", "+os.Getenv("DB_TABLE_FIELD_NAME")+", "+os.Getenv("DB_TABLE_FIELD_TYPE")+" from "+os.Getenv("DB_TABLE_NAME")+" where "+os.Getenv("DB_TABLE_FIELD_TOKEN")+" = '"+api_token+"'").Scan(&tag.ID, &tag.Nama, &tag.Tipe)

	// if there is an error selecting, handle it
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(domain.JSONResult{
			Code:    401,
			Message: "Unauthorized.",
		})
	}

	tipe := tag.Tipe
	if Strstr(tipe, os.Getenv("DB_TABLE_TYPE_ALLOW")) == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(domain.JSONResult{
			Code:    401,
			Message: "Anda tidak memiliki Akses.",
		})
	}
	nama := tag.Nama

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = nama

	JWTExpiredTime, err := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES"))
	if err != nil {
		JWTExpiredTime = 60
	}
	AddedTimeJWT := time.Minute * time.Duration(JWTExpiredTime)
	claims["exp"] = time.Now().Add(AddedTimeJWT).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return domain.NewHttpError(c, fiber.StatusInternalServerError, err)
	}

	return c.JSON(domain.JSONResult{
		Data: fiber.Map{
			"token":        t,
			"expired_in":   claims["exp"],
			"expired_time": JWTExpiredTime,
		},
		Message: "Token Generate",
	})
}

/*
 * Tag... - a very simple struct
 */
type Tag struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
	Tipe string `json:"tipe"`
}

func Strstr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle))-1:]
}

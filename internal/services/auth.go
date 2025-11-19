package services

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/dto"
	"github.com/Alwin18/sistem-jadwal-pelajaran-sekolah/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewAuthService(db *gorm.DB, log *logrus.Logger) *AuthService {
	return &AuthService{
		db:  db,
		log: log,
	}
}

func (s *AuthService) Login(ctx *fiber.Ctx, username, password string) (dto.LoginResponse, error) {
	var result dto.LoginResponse
	var user models.User

	if err := s.db.Where("username = ?", username).
		Find(&user).Error; err != nil {
		s.log.Error(err.Error())
		ctx.Status(fiber.StatusInternalServerError)
		return result, errors.New("terjadi kesalahan sistem")
	}

	if user.ID == 0 {
		s.log.Error("user tidak ditemukan")
		ctx.Status(fiber.StatusNotFound)
		return result, errors.New("user tidak ditemukan")
	}

	fixedHash := strings.Replace(user.Password, "$2y$", "$2a$", 1)
	if err := bcrypt.CompareHashAndPassword([]byte(fixedHash), []byte(password)); err != nil {
		s.log.Error(err.Error())
		ctx.Status(fiber.StatusUnauthorized)
		return result, errors.New("password yang anda masukan salah")
	}

	namaPengguna := "ADMIN YAYASAN"
	nik := ""
	switch user.Role {
	case "TEACHER":
		var teacher models.Teacher
		if err := s.db.Where("user_id = ?", user.ID).First(&teacher).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				s.log.Error("akun tidak ditemukan")
				ctx.Status(fiber.StatusInternalServerError)
				return result, errors.New("akun tidak ditemukan")
			}
			s.log.Error(err.Error())
			ctx.Status(fiber.StatusInternalServerError)
			return result, errors.New("terjadi kesalahan sistem")
		}
		namaPengguna = teacher.Name
		nik = teacher.NIK
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"nama":     namaPengguna,
		"nik":      nik,
		"exp":      time.Now().Add(time.Hour * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		s.log.Error(err.Error())
		ctx.Status(fiber.StatusInternalServerError)
		return result, errors.New("terjadi kesalahan sistem")
	}

	result.User = dto.UserLogin{
		ID:           user.ID,
		NamaPengguna: namaPengguna,
		Username:     username,
		Role:         user.Role,
		Nik:          nik,
	}
	result.Token = tokenString
	return result, nil
}

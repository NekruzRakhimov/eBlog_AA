package service

import (
	"crypto/sha256"
	"eBlog/internal/errs"
	"eBlog/internal/models"
	"encoding/hex"
	"errors"
	"fmt"
)

func (s *Service) CreateUser(u models.User) error {
	// 1. проверить существует ли такой пользователь
	dbUser, err := s.repository.GetUserByUsername(u.Username)
	if err != nil && !errors.Is(err, errs.ErrNotFound) {
		return err
	}

	if dbUser.ID != 0 {
		return errs.ErrUsernameAlreadyExists
	}

	// 2. захещировать пароль
	u.Password, err = s.GenerateHash(u.Password)
	if err != nil {
		return err
	}

	// 3. создать пользователя
	return s.repository.CreateUser(u)
}

func (s *Service) AuthenticateUser(username, password string) (accessToken string, refreshToken string, err error) {
	// хешируем входной пароль
	hashedPassword, err := s.GenerateHash(password)
	if err != nil {
		return "", "", err
	}

	// получаем из бд данные
	u, err := s.repository.GetUserByUsernameAndPassword(username, hashedPassword)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return "", "", errs.ErrIncorrectUsernameOrPassword
		}

		return "", "", err
	}

	// генерация токена
	accessToken, err = s.GenerateToken(u.ID, false)
	if err != nil {
		return "", "", fmt.Errorf(" GenerateToken(u.ID): %v\n", err)
	}

	refreshToken, err = s.GenerateToken(u.ID, true)
	if err != nil {
		return "", "", fmt.Errorf(" GenerateToken(u.ID): %v\n", err)
	}

	return accessToken, refreshToken, nil
}

func (s *Service) GenerateHash(input string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(input)); err != nil {
		return "", err
	}

	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes), nil
}

package service

import (
	"crypto/sha256"
	"database/sql"
	"eBlog/internal/models"
	"eBlog/internal/repository"
	"encoding/hex"
	"errors"
	"fmt"
)

func CreateUser(u models.User) error {
	// 1. проверить существует ли такой пользователь
	dbUser, err := repository.GetUserByUsername(u.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if dbUser.ID != 0 {
		return errors.New("username already exists")
	}

	// 2. захещировать пароль
	u.Password, err = GenerateHash(u.Password)
	if err != nil {
		return err
	}

	// 3. создать пользователя
	return repository.CreateUser(u)
}

func AuthenticateUser(username, password string) (accessToken string, refreshToken string, err error) {
	// хешируем входной пароль
	hashedPassword, err := GenerateHash(password)
	if err != nil {
		return "", "", err
	}

	// получаем из бд данные
	u, err := repository.GetUserByUsernameAndPassword(username, hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", errors.New("неправильный логин или пароль")
		}

		return "", "", err
	}

	// генерация токена
	accessToken, err = GenerateToken(u.ID, false)
	if err != nil {
		return "", "", fmt.Errorf(" GenerateToken(u.ID): %v\n", err)
	}

	refreshToken, err = GenerateToken(u.ID, true)
	if err != nil {
		return "", "", fmt.Errorf(" GenerateToken(u.ID): %v\n", err)
	}

	return accessToken, refreshToken, nil
}

func GenerateHash(input string) (string, error) {
	hash := sha256.New()

	if _, err := hash.Write([]byte(input)); err != nil {
		return "", err
	}

	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes), nil
}

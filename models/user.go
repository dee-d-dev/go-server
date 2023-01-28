package models

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/go-redis/redis"
	"errors"
)

var (
	ErrNotFound = errors.New("user not found")
	ErrInvalidLogin = errors.New("Invalid login")
)

func RegisterUser(username, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return err
	}

	return client.Set("user:"+username, hash, 0).Err()


}


func AuthenticateUser(username, password string) error {
	hash, err := client.Get("user:" + username).Bytes()

	if err == redis.Nil {
		return ErrNotFound
	} else if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))

	if err != nil {
		return ErrInvalidLogin
	}

	return nil
}
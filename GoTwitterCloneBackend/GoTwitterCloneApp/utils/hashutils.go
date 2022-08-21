package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type IHashUtility interface {
	GetHashData(string) string
	GetIsCompareHashValid(string, string) bool
}

type HashUtility struct {
	HashStrategy IHashStrategy
}

func (s *HashUtility) GetHashData(data string) string {
	return s.HashStrategy.ExecuteGetHashData(data)
}

func (s *HashUtility) GetIsCompareHashValid(hashedData, data string) bool {
	return s.HashStrategy.ExecuteGetIsCompareHashValid(hashedData, data)
}

type IHashStrategy interface {
	ExecuteGetHashData(string) string
	ExecuteGetIsCompareHashValid(string, string) bool
}

type HashUserPasswordStrategy struct {
}

const hashUserPasswordSalt = "UHbCCPrHtxBK4w4bm8wcGkLdcbE0nxcD5MHUPKtK"

func (s *HashUserPasswordStrategy) ExecuteGetHashData(data string) string {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(data +hashUserPasswordSalt), 10)
	return string(hashBytes)
}

func (s *HashUserPasswordStrategy) ExecuteGetIsCompareHashValid(hashedData, data string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data +hashUserPasswordSalt))
	return err == nil
}

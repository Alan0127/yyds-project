package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

//
//  HashAndSalt
//  @Description: 密码加密
//  @param pwd
//  @return string
//
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//
//  comparePasswords
//  @Description: 密码验证
//  @param hashedPwd
//  @param plainPwd
//  @return bool
//
func ComparePwd(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println("err is: ", err)
		return false
	}
	return true
}

/*
 * @Author: your name
 * @Date: 2021-03-22 11:08:24
 * @LastEditTime: 2021-03-22 11:24:21
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \exchange\internal\pkg\password\password.go
 */
package password

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptionPassword(cleartextPassword string) (string, error) {
	if b, err := bcrypt.GenerateFromPassword([]byte(cleartextPassword), bcrypt.DefaultCost); err == nil {
		return string(b), err
	}
	return "", nil
}

func VerifyPassword(encryptionPassword, cleartextPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptionPassword), []byte(cleartextPassword)); err != nil {
		return err
	}
	return nil
}

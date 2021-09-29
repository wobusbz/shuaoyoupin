/*
 * @Author: your name
 * @Date: 2021-03-22 11:12:50
 * @LastEditTime: 2021-03-22 11:24:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \exchange\internal\pkg\password\password_test.go
 */
package password

import "testing"

func TestEncryptionPassword(t *testing.T) {
	var password = "123456"
	passwords, _ := EncryptionPassword(password)
	t.Log(passwords)
	t.Log(VerifyPassword(passwords, password))
}

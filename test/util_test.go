package test

import (
	"fmt"
	"go-gin-duty-master/e"
	"log"
	"testing"

	"go-gin-duty-master/util"
)

func TestGenerateToken(t *testing.T) {
	token, err := util.GenerateToken("yaojunyi", "1007227432")
	if err != nil {
		log.Fatalf("util.TestGenerateToken err: %v", err)
	} else {
		fmt.Print(token)
	}
}

func TestParseToken(t *testing.T) {
	//auth := &models.DutyAuth{}
	//token,_ :=util.GenerateToken("yaojunyi","1007227432")
	//claim,_ := util.ParseToken(token)
	//username,_ := util.Decrypt(claim.Username,[]byte(e.KEY))
	//err := models.Db.Select("name").Where(models.DutyAuth{Username: username}).First(&auth).Error
	//fmt.Print(err)
	//fmt.Printf(username)

}

func TestEncrypt(t *testing.T) {
	username := "yaojunyi"
	username, _ = util.Encrypt(username, []byte(e.KEY))
	username, _ = util.Decrypt(username, []byte(e.KEY))
	fmt.Print(username)
}

package util

import "go-gin-duty-master/e"

func DecrpytToken(token string) (string, error) {
	var (
		username string
		err      error
	)

	claim, err := ParseToken(token)

	if err != nil {
		return "", err
	}

	username, err = Decrypt(claim.Username, []byte(e.KEY))

	if err != nil {
		return "", err
	}

	return username, nil
}

package black_kiwi_auth_structs

import (
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)


const tokenExpirationTime = 3600 // 1 hour

type AuthTokenStruct struct {
	Token string
	Role int8
	LastUsed int64
}
type AuthToken *AuthTokenStruct

var tokenArr *map[string]AuthToken

func InitTokenArr() {
	log.Debug("Initializing tokenArr")
	tokenArr =  &map[string]AuthToken{}
}

func AddToken(user *User)(authToken AuthToken) {
	authToken = &AuthTokenStruct{
		Token: uuid.New().String(),
		Role: (*user).Role,
		LastUsed: time.Now().Unix(),
	}
	log.WithFields(log.Fields{"user": *user, "token": *authToken}).Debug("Generating new token")
	(*user).Token = (*authToken).Token
	(*tokenArr)[(*authToken).Token] = authToken
	return
}

func (authToken AuthTokenStruct) IsExpired() (expired bool) {
	expired = time.Now().Unix() - authToken.LastUsed > tokenExpirationTime
	log.WithFields(log.Fields{"token": authToken, "expired": expired}).Debug("Checking if token is expired")
	return 
}

func GetToken(authTokenStr string) (authToken AuthToken) {
	log.WithFields(log.Fields{"authTokenStr": authTokenStr, "token": (*tokenArr)[authTokenStr]}).Debug("GetToken called")
	if authTokenStr == "" {
		return nil
	}

	// If no struct is found, Go initializes the fields with zero values
	if (*tokenArr)[authTokenStr] == nil {
		return nil
	}

	authToken = (*tokenArr)[authTokenStr]

	// If token is not expired, update LastUsed
	if !(*authToken).IsExpired() {
		log.WithFields(log.Fields{"token": *authToken}).Debug("Updating token LastUsed")
		(*authToken).LastUsed = time.Now().Unix()
	}

	return
}

func (authToken AuthTokenStruct) DeleteToken() {
	delete(*tokenArr, authToken.Token)
}

func (authToken AuthTokenStruct) GetRole() int8 {
	return authToken.Role
}

func (authToken AuthTokenStruct) IsAdmin() bool {
	return authToken.Role == 2
}

func (authToken AuthTokenStruct) IsUser() bool {
	return authToken.Role == 1
}
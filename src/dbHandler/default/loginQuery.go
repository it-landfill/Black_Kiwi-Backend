package black_kiwi_default_queries

import (
	"context"

	log "github.com/sirupsen/logrus"

	black_kiwi_db_utils "ITLandfill/Black-Kiwi/dbHandler/utils"
	"ITLandfill/Black-Kiwi/structs/auth_structs"
)

//https://app.swaggerhub.com/apis/ITLandfill/Black-Kiwi/1.0.2

/*
SELECT username, role
FROM "black-kiwi_authentication".Users
WHERE username = 'testUser' and password = 'testPassword';
*/
func GetUser(username string, password string) (result bool ,user *black_kiwi_auth_structs.User) {

	var user_name string
	var role int8
	err := black_kiwi_db_utils.ConnPool.QueryRow(context.Background(),  "SELECT username, role FROM \"black-kiwi_authentication\".Users WHERE username=$1 and password=$2;", username, password).Scan(&user_name, &role)
	if err != nil {
		if (err.Error() == "no rows in result set") {
			log.WithFields(log.Fields{"username": username, "error":err}).Info("No user found in db.")
			return true, nil
		}
		
		log.WithFields(log.Fields{"username": username, "error":err}).Error("QueryRow failed while retrieving user.")
		return false, nil
	}

	log.WithFields(log.Fields{"username":user_name, "role":role}).Debug("QueryRow succeeded while retrieving user.")

	user = new(black_kiwi_auth_structs.User)
	(*user).Username = user_name
	(*user).Role = role

	return true, user
}
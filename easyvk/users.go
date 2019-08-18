package easyvk

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Users struct {
	vk *VK
}

type UserData struct {
	ID         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type UsersGetInfoResponse struct {
	Response []UserData `json:"response"`
}

func (u Users) Get() (UsersGetInfoResponse, error) {
	resp, err := u.vk.Request("users.get", nil)
	if err != nil {
		return UsersGetInfoResponse{}, err
	}
	logrus.Warn("easyvk.Users.Get: ", string(resp))
	var info UsersGetInfoResponse
	info.Response = make([]UserData, 0)
	err = json.Unmarshal(resp, &info.Response)
	if err != nil {
		return UsersGetInfoResponse{}, err
	}
	return info, nil
}

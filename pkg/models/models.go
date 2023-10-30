package models

import "github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/pb"

type RequestUserID struct {
	UserId int32 `json:"userId"`
}

type RequestUserIDList struct {
	UserId []int32 `json:"userId"`
}

type ResponseUserData struct {
	UserId    int32
	FirstName string
	City      string
	Phone     string
	Height    float32
	Married   bool
}

type ResponseUserDataList struct {
	Users         []*pb.UserData
	UsersNotFound []int32
}

type Response struct {
	Response string
}

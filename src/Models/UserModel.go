package Models

import (
	"../Sys"
	"../DBStruct"
)

func GetUserLogin(userName string, Password string) []DBStruct.Messages  {
	defer Sys.DB().Close()

	Messages := make([]DBStruct.Messages, 0, 0)
	Sys.DB().
		Preload("MessagesLang").
		Find(&Messages)


	return Messages
}



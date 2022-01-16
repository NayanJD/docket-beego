package main

import (
	"flag"
	"fmt"

	"docket-beego/auth"
	"docket-beego/utils"

	pg "github.com/vgarvardt/go-oauth2-pg/v4"
)

type Oauth2ClientInfo struct {
	pg.ClientStoreItem
}

func (ci Oauth2ClientInfo) GetID() string {
	fmt.Println(ci.ID)
	return ci.ID
}

func (ci Oauth2ClientInfo) GetSecret() string {
	fmt.Println(ci.Secret)
	return ci.Secret
}

func (ci Oauth2ClientInfo) GetDomain() string {
	fmt.Println(ci.Domain)
	return ci.Domain
}

func (ci Oauth2ClientInfo) GetUserID() string {
	return ""
}

func CreateOauthClient(flagSet *flag.FlagSet, commands []string) {
	// newClient := pg.ClientStoreItem{
	// 	ID:     utils.GenerateSecureToken(8),
	// 	Secret: utils.GenerateSecureToken(16),
	// }

	newClientInfo := Oauth2ClientInfo{
		pg.ClientStoreItem{
			ID:     utils.GenerateSecureToken(8),
			Secret: utils.GenerateSecureToken(16),
			Domain: "",
		},
	}

	auth.ClientStore.Create(newClientInfo)

	fmt.Println("client id: ", newClientInfo.ID)
	fmt.Println("client secret: ", newClientInfo.Secret)
}

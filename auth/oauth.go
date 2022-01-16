package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"

	"docket-beego/models"
	"docket-beego/utils"
)

var Srv *server.Server
var ClientStore *pg.ClientStore

func IsStatusSuccess(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices
}

func init() {
	pgxConn, _ := pgx.Connect(context.TODO(), models.SqlConnString)

	manager := manage.NewDefaultManager()

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()

	ClientStore, _ = pg.NewClientStore(adapter)

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(ClientStore)

	Srv = server.NewDefaultServer(manager)

	Srv.SetClientInfoHandler(server.ClientFormHandler)

	Srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		user := &models.User{}

		// err = models.GetDB().First(user, "username = ?", username).Error

		err = models.Orm.QueryTable("users").Filter("username", username).One(user)

		if err == nil && user.ComparePassword(&password) {
			userID = user.Id
		}

		return userID, nil
	})

	Srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		utils.Log.Error(fmt.Sprintf("Internal Error: %v", err.Error()))
		return
	})

	Srv.SetResponseErrorHandler(func(re *errors.Response) {
		utils.Log.Error(fmt.Sprintf("Response Error: %v", re.Error.Error()))
	})

	Srv.SetResponseTokenHandler(
		func(w http.ResponseWriter, data map[string]interface{}, header http.Header, statusCode ...int) error {
			body := map[string]interface{}{
				"data":      data,
				"errors":    nil,
				"isSuccess": len(statusCode) > 0 && IsStatusSuccess(statusCode[0]),
				"meta":      nil,
			}

			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.Header().Set("Cache-Control", "no-store")
			w.Header().Set("Pragma", "no-cache")

			for key := range header {
				w.Header().Set(key, header.Get(key))
			}

			status := http.StatusOK
			if len(statusCode) > 0 && statusCode[0] > 0 {
				status = statusCode[0]
			}

			w.WriteHeader(status)
			return json.NewEncoder(w).Encode(body)
		},
	)
}

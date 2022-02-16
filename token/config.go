package token

import (
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"infradev-practice/Wade/OAuth2.0-server/models"
)

var (
	manager = manage.NewDefaultManager()
	srv     = server.NewServer(server.NewConfig(), manager)
	cs      = store.NewClientStore()

	authInfo = new(models.AuthInfo)
)

func init() {
	manager.SetClientTokenCfg(manage.DefaultClientTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(cs)
}

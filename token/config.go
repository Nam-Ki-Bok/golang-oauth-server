package token

import (
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"infradev-practice/Wade/OAuth2.0-server/models"
)

var (
	mgr = manage.NewDefaultManager()
	srv = server.NewServer(server.NewConfig(), mgr)
	cs  = store.NewClientStore()

	authInfo = new(models.AuthInfo)
)

func init() {
	mgr.SetClientTokenCfg(manage.DefaultClientTokenCfg)
	mgr.MustTokenStorage(store.NewMemoryTokenStore())
	mgr.MapAccessGenerate(generates.NewAccessGenerate())
	mgr.MapClientStorage(cs)
}

package token

import (
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

var (
	Mgr = manage.NewDefaultManager()
	Srv = server.NewServer(server.NewConfig(), Mgr)
	Cs  = store.NewClientStore()
)

func init() {
	// You can use this if you want to change TTL.
	//Mgr.SetClientTokenCfg(&manage.Config{
	//	 AccessTokenExp: time.Duration(time.Minute * 15),
	//})

	Mgr.SetClientTokenCfg(manage.DefaultClientTokenCfg)
	Mgr.MustTokenStorage(store.NewMemoryTokenStore())
	Mgr.MapAccessGenerate(generates.NewAccessGenerate())
	Mgr.MapClientStorage(Cs)
}

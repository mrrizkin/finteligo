package system

import (
	"github.com/mrrizkin/finteligo/routes"

	"github.com/mrrizkin/finteligo/app/config"
	"github.com/mrrizkin/finteligo/app/models"
	"github.com/mrrizkin/finteligo/system/database"
	"github.com/mrrizkin/finteligo/system/server"
	"github.com/mrrizkin/finteligo/system/session"
	"github.com/mrrizkin/finteligo/system/types"
	"github.com/mrrizkin/finteligo/system/validator"
	"github.com/mrrizkin/finteligo/third_party/logger"
)

func Run() {
	conf, err := config.New()
	if err != nil {
		panic(err)
	}
	log, err := logger.New(conf)
	if err != nil {
		panic(err)
	}
	sess, err := session.New(conf)
	if err != nil {
		panic(err)
	}
	defer sess.Stop()
	model := models.New()
	db, err := database.New(conf, model, log)
	if err != nil {
		panic(err)
	}
	defer db.Stop()
	err = db.Start()
	if err != nil {
		panic(err)
	}

	valid := validator.New()
	serv := server.New(conf, log, sess)

	routes.Setup(&types.App{
		App: serv.App,
		System: &types.System{
			Logger:    log,
			Database:  db,
			Config:    conf,
			Session:   sess,
			Validator: valid,
		},
		Library: &types.Library{},
	})

	log.Info().Msgf("Server is running on port %d", conf.PORT)

	if err := serv.Serve(); err != nil {
		panic(err)
	}
}

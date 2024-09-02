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
	"github.com/mrrizkin/finteligo/third_party/argon2"
	"github.com/mrrizkin/finteligo/third_party/langchain"
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
	argon2 := argon2.New(
		uint32(conf.HASH_MEMORY),
		uint32(conf.HASH_ITERATIONS),
		uint32(conf.HASH_KEY_LEN),
		uint32(conf.HASH_SALT_LEN),
		uint8(conf.HASH_PARALLELISM),
	)

	model := models.New(conf, argon2)
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
	serv := server.New(conf, log)

	lc := langchain.New(log, db)
	lc.InitializeLLMs()

	routes.Setup(&types.App{
		App: serv.App,
		System: &types.System{
			Logger:    log,
			Database:  db,
			Config:    conf,
			Session:   sess,
			Validator: valid,
		},
		Library: &types.Library{
			LangChain: lc,
			Argon2:    argon2,
		},
	}, sess)

	log.Info().Msgf("Server is running on port %d", conf.PORT)

	if err := serv.Serve(); err != nil {
		panic(err)
	}
}

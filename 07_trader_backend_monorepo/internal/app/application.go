package app

import (
	"context"
	"fmt"

	strategiesHTTP "trader-backend_monorepo/internal/adapters/http/strategies"
	mongoAdapters "trader-backend_monorepo/internal/adapters/mongo"
	"trader-backend_monorepo/internal/config"
	strategiesUC "trader-backend_monorepo/internal/usecases/strategies"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Application interface {
	Run(ctx context.Context) error
}

type application struct {
	conf       *config.Config
	httpRouter *gin.Engine
}

func NewApplication(ctx context.Context) (Application, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	app := &application{
		conf:       conf,
		httpRouter: gin.Default(),
	}

	// Configure and connect to MongoDB:
	mongoOpts := options.Client().
		SetAppName(conf.MongoConfig.AppName).
		ApplyURI(conf.MongoConfig.GetConnectionURI()).
		SetMinPoolSize(uint64(conf.MongoConfig.MinPoolSize)).
		SetMaxPoolSize(uint64(conf.MongoConfig.MaxPoolSize))

	mongoClient, err := mongo.Connect(mongoOpts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo: %w", err)
	}

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("error getting ping from mongo: %w", err)
	}

	// Connect to 'traderdb' and get its collections:
	mongoDB := mongoClient.Database(conf.MongoConfig.Database)
	strategiesColl := mongoDB.Collection("strategies")

	// Create the Strategies adapter:
	strategiesRepo := mongoAdapters.NewStrategiesRepository(strategiesColl)
	// Create strategies service:
	strategiesSvc := strategiesUC.NewService(strategiesRepo)
	// Create strategies HTTP handlers:
	handlers := strategiesHTTP.NewHandlers(strategiesSvc)
	// Register strategies handlers:
	strategiesHTTP.RegisterRoutes(app.httpRouter, handlers)

	return app, nil
}

func (a *application) Run(ctx context.Context) error {
	return a.httpRouter.Run(fmt.Sprintf(":%d", a.conf.HTTPConfig.Port))
}

package app

import (
	"github.com/hudayberdipolat/golang-auth-jwtToken/pkg/config"
	"github.com/hudayberdipolat/golang-auth-jwtToken/pkg/database/dbConfig"
	customHttp "github.com/hudayberdipolat/golang-auth-jwtToken/pkg/http"
	"gorm.io/gorm"
	"net/http"
)

type Dependencies struct {
	Config     *config.Config
	DB         *gorm.DB
	HttpClient *http.Client
}

func GetAppDependencies() (*Dependencies, error) {
	// get config
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	// database config
	newDbConfig := dbConfig.NewDBConnection(*getConfig)
	db, errDbConnection := newDbConfig.DBConnection()
	if errDbConnection != nil {
		return nil, errDbConnection
	}
	// httpClient config
	httpClient := customHttp.NewHttpClient()
	return &Dependencies{
		Config:     getConfig,
		DB:         db,
		HttpClient: httpClient,
	}, nil
}

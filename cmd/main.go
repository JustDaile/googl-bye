package main

import (
	"context"
	"log"

	"github.com/google/go-github/v39/github"
	"github.com/jwtly10/googl-bye/internal/common"
	"github.com/jwtly10/googl-bye/internal/models"
	"github.com/jwtly10/googl-bye/internal/parser"
	"github.com/jwtly10/googl-bye/internal/repository"
	"github.com/jwtly10/googl-bye/internal/search"
	"go.uber.org/zap/zapcore"
)

func main() {
	config, err := common.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config")
	}

	useJsonLogger := false
	logger := common.NewLogger(useJsonLogger, zapcore.InfoLevel)

	db, err := common.ConnectDB(config)
	if err != nil {
		logger.Errorf("Error connecting to DB: %v", err)
	}

	repoRepo := repository.NewRepoRepository(db)
	stateRepo := repository.NewParserStateRepository(db)
	linkRepo := repository.NewParserLinkRepository(db)

	searchParams := &models.SearchParams{
		Query: "stars:>10000",
		Opts: &github.SearchOptions{
			Sort:  "stars",
			Order: "desc",
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		},
		Page: 2,
	}

	_ = search.NewRepoSearch(searchParams, config, logger, repoRepo)

	search := search.NewRepoSearch(searchParams, config, logger, repoRepo)
	search.StartSearch(context.Background())

	parser := parser.NewParser(logger, repoRepo, stateRepo, linkRepo)

	parser.StartParser(context.Background(), 5)
}

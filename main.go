package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/GriffyHome/go-skeleton/pkg/config"
	"github.com/GriffyHome/go-skeleton/pkg/db"
	"github.com/GriffyHome/go-skeleton/pkg/logger"
	"github.com/GriffyHome/go-skeleton/pkg/seq"

	"github.com/rs/zerolog/log"
)

func main() {
	logger.InitLogger()
	environment := flag.String("e", "production", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	if *environment == "production" {
		seq.InitSeqLogger()
	}

	dbConnection, err := db.NewSQLDB()
	if err != nil {
		log.Error().Msg("Can't initialize db " + err.Error())
		panic("Can't initialize db " + err.Error())
	}

	log.Printf("Connected to sql database")
	redisConnection, err := db.NewRedisConnection()
	if err != nil {
		log.Error().Msg("Can't initialize redis " + err.Error())
		panic("Can't initialize redis " + err.Error())
	}
	log.Printf("Connected to redis database")

	cassandraConnection, err := db.NewCassandraDB()
	if err != nil {
		log.Error().Msg("Can't initialize Cassandra: " + err.Error())
		panic("Can't initialize Cassandra: " + err.Error())
	}
	log.Printf("Connected to cassandra database")
	
	//remove below line before running
	fmt.Println(redisConnection,cassandraConnection,dbConnection)

}

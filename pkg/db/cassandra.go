package db

import (
	"github.com/GriffyHome/go-skeleton/pkg/config"

	"github.com/gocql/gocql"
	"github.com/rs/zerolog/log"
)

type cassandraConnection struct {
	session *gocql.Session
}

func NewCassandraDB() (*cassandraConnection, error) {
	dbConfig := config.CassandraConfig()

	log.Info().Msg("Initializing Cassandra cluster connection...")
	cluster := gocql.NewCluster(dbConfig.Hosts...)
	cluster.Keyspace = dbConfig.Keyspace
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Error().Msg("Failed to create Cassandra session: " + err.Error())
		return nil, err
	}

	conn := &cassandraConnection{

		session: session,
	}

	if err := conn.AutoMigrateTables(); err != nil {
		log.Error().Msg("Failed to auto-migrate tables: " + err.Error())
		session.Close()
		return nil, err
	}

	log.Info().Msg("Connected to Cassandra successfully")
	return conn, nil
}

func (d *cassandraConnection) AutoMigrateTables() error {
	return nil
}

func (d *cassandraConnection) GetSession() *gocql.Session {
	return d.session
}

func (d cassandraConnection) Close() {
	if d.session != nil {
		d.session.Close()
	}
}

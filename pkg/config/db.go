package config

type DB struct {
	URL      string
	DBType   string
	Database string
	Username string
	Password string
}


type Cassandra struct {
	Hosts    []string
	Keyspace string
	Username string
	Password string
}

type DBType string

const (
	RDS    DBType = "rds"
	Aurora DBType = "aurora"
)

func (d DBType) String() string {
	switch d {
	case Aurora:
		return "aurora"
	case RDS:
		return "rds"
	}
	return "invalid"
}

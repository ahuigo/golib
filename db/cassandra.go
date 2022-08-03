package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	conf := &Config{
		Hosts:    "host1,host2",
		Port:     1042,
		Keyspace: "cadence_dev",
		User:     "cassandra",
		Password: "123456",
	}
	_, err := NewQueueManager(conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("connect cassandra successfully!\n")
}

const (
	cassandraProtoVersion = 4
	defaultSessionTimeout = 10 * time.Second
)

type (
	// QueueManager -
	QueueManager struct {
		session *gocql.Session
	}
)

func parseHosts(input string) []string {
	var hosts = make([]string, 0)
	for _, h := range strings.Split(input, ",") {
		if host := strings.TrimSpace(h); len(host) > 0 {
			hosts = append(hosts, host)
		}
	}
	return hosts
}

type (
	// Config -
	Config struct {
		Hosts    string
		Port     int
		Keyspace string
		User     string
		Password string
	}
)

// NewQueueManager -
func NewQueueManager(cfg *Config) (*QueueManager, error) {
	cluster := NewCassandraCluster(cfg)
	cluster.ProtoVersion = cassandraProtoVersion
	cluster.Consistency = gocql.LocalQuorum
	cluster.SerialConsistency = gocql.LocalSerial
	cluster.Timeout = defaultSessionTimeout
	session, err := cluster.CreateSession()
	if err != nil {
		err = fmt.Errorf("cassandra create session:%v", err)
		return nil, err
	}
	mgr := &QueueManager{
		session: session,
	}
	return mgr, nil
}

// NewCassandraCluster creates a cassandra cluster from a given configuration
func NewCassandraCluster(cfg *Config) *gocql.ClusterConfig {
	hosts := parseHosts(cfg.Hosts)
	cluster := gocql.NewCluster(hosts...)
	cluster.ProtoVersion = 4
	if cfg.Port > 0 {
		cluster.Port = cfg.Port
	}
	if cfg.Keyspace != "" {
		cluster.Keyspace = cfg.Keyspace
	}
	if cfg.User != "" && cfg.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: cfg.User,
			Password: cfg.Password,
		}
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}

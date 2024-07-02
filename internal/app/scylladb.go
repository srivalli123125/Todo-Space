package scylladb

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"
)

func SetupScyllaDB() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.23.0.2")
	cluster.Keyspace = "todo"

	session, err := cluster.CreateSession()
	if err == nil {
		return session, nil
	}

	if !strings.Contains(err.Error(), "Keyspace 'todo' does not exist") {
		return nil, fmt.Errorf("failed to create session for todo keyspace: %v", err)
	}

	cluster.Keyspace = "system"
	session, err = cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session for system keyspace: %v", err)
	}

	err = session.Query(`CREATE KEYSPACE todo WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec()
	if err != nil {
		return nil, fmt.Errorf("failed to create keyspace todo: %v", err)
	}

	cluster.Keyspace = "todo"
	session, err = cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session for todo keyspace: %v", err)
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS todos (
			id UUID,
			user_id UUID,
			title TEXT,
			description TEXT,
			status TEXT,
			created TEXT,
			updated TEXT,
			PRIMARY KEY (id, user_id)
		)
	`).Exec()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	fmt.Println("Keyspace 'todo' and table 'todos' created successfully")
	return session, nil
}

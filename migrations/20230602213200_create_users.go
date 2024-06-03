package migrations

import (
	"gofr.dev/pkg/gofr/migration"
	"log"
)

const createUsers = `CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password VARCHAR(100) NOT NULL,
                        name VARCHAR(100) NOT NULL,
                        address VARCHAR(255) NOT NULL,
                        latitude DOUBLE PRECISION,
    					longitude DOUBLE PRECISION,
                        phone_number VARCHAR(50) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

func createUsersTable() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createUsers)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		},
	}
}

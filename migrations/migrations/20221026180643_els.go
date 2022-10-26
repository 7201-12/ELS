package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upEls, downEls)
}

var migrate = `
CREATE TABLE IF NOT EXISTS questions(
    id uuid,
    q_type int,
    value text,
    variants uuid[],
    answer uuid,
    time real,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS variants(
    id uuid,
    value text,
    PRIMARY KEY(id)
);

INSERT INTO questions(id, q_type, value, variants, answer, time) VALUES('ab169a28-124f-454d-8c84-99e110f3b013', 1, 'What is your name?', '{"dae7a8aa-6168-485d-96d6-188a5a9fbdaf", "a11366b7-6e7f-4d93-ba8d-4d68e423da31", "7744edaa-b34d-47cb-96cf-a66136406cd9"}', '7744edaa-b34d-47cb-96cf-a66136406cd9', 0.5);
INSERT INTO variants(id, value) VALUES('dae7a8aa-6168-485d-96d6-188a5a9fbdaf', 'My name is Sir Lancelot of Camelot');
INSERT INTO variants(id, value) VALUES('a11366b7-6e7f-4d93-ba8d-4d68e423da31', 'Sir Robin of Camelot');
INSERT INTO variants(id, value) VALUES('7744edaa-b34d-47cb-96cf-a66136406cd9', 'It is Arthur, King of the Britons');
`

func upEls(tx *sql.Tx) error {
	_, err := tx.Exec(migrate)
	if err != nil {
		return err
	}
	return nil
}

func downEls(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}

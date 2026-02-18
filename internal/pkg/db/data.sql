CREATE TABLE IF NOT EXISTS scheduler (
	id INTEGER primary key autoincrement,
	date text NOT NULL,
	title text NOT NULL,
	comment text NOT NULL,
	repeat varchar(250) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);

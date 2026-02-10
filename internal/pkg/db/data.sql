CREATE TABLE scheduler (
	id integer primary key autoincrement,
	date date NOT NULL,
	title text NOT NULL,
	comment text NOT NULL,
	repeat varchar(250) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);

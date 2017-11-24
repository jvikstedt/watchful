package model

var schema = `
PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS jobs (
	id integer PRIMARY KEY,
	name text,
	active integer DEFAULT 0,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	deleted_at timestamp
);
CREATE TABLE IF NOT EXISTS tasks (
	id integer PRIMARY KEY,
	job_id integer,
	executable text,
	seq integer,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	deleted_at timestamp,
	FOREIGN KEY(job_id) REFERENCES jobs(id),
	UNIQUE(job_id, seq)
);
CREATE TABLE IF NOT EXISTS inputs (
	id integer PRIMARY KEY,
	task_id integer,
	name text,
	value text,
	type integer DEFAULT 0,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	deleted_at timestamp,
	FOREIGN KEY(task_id) REFERENCES tasks(id)
);
CREATE TABLE IF NOT EXISTS results (
	id integer PRIMARY KEY,
	uuid text,
	test_run integer DEFAULT 0,
	job_id integer,
	status text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	FOREIGN KEY(job_id) REFERENCES jobs(id)
);
CREATE TABLE IF NOT EXISTS result_items (
	id integer PRIMARY KEY,
	result_id integer,
	task_id integer,
	output text,
	error text,
	status text,
	created_at timestamp,
	updated_at timestamp DEFAULT current_timestamp,
	FOREIGN KEY(result_id) REFERENCES results(id),
	FOREIGN KEY(task_id) REFERENCES tasks(id)
);
`

package logger

import (
	"database/sql"
	"fmt"
)

type Logger struct {
	db *sql.DB
}

func NewLogger(db *sql.DB) Logger {
	return Logger{db: db}
}

func (l *Logger) latestHelper(stmt, prefix string) ([]string, error) {
	var err error
	var rows *sql.Rows
	if prefix != "" {
		rows, err = l.db.Query(stmt, prefix)
	} else {
		rows, err = l.db.Query(stmt)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []string
	for rows.Next() {
		log := &parsedLog{}
		err = rows.Scan(&log.Prefix, &log.LogTime, &log.File, &log.Payload)
		if err != nil {
			return nil, err
		}
		logs = append(logs, fmt.Sprintf("%v", log))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func (l *Logger) Latest1Day() ([]string, error) {
	stmt := `select prefix, log_time, file, payload from log where log_time >= now() - '1 day'::interval`
	return l.latestHelper(stmt, "")
}

func (l *Logger) Latest1Week() ([]string, error) {
	stmt := `select prefix, log_time, file, payload from log where log_time >= now() - '1 week'::interval`
	return l.latestHelper(stmt, "")
}

func (l *Logger) Latest1DayWithPrefix(prefix string) ([]string, error) {
	stmt := `
		select prefix, log_time, file, payload
		from log
		where log_time >= now() - '24 hours'::interval and prefix = $1`
	return l.latestHelper(stmt, prefix)
}

func (l *Logger) Latest1WeekWithPrefix(prefix string) ([]string, error) {
	stmt := `
		select prefix, log_time, file, payload
		from log
		where log_time >= now() - '1 Week'::interval and prefix = $1`
	return l.latestHelper(stmt, prefix)
}

func (l *Logger) ClearLogs() error {
	_, err := l.db.Exec("truncate log")
	return err
}

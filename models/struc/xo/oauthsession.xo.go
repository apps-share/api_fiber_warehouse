package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// OauthSession represents a row from 'spc_holding.oauth_sessions'.
type OauthSession struct {
	SessionID           string         `json:"session_id"`            // session_id
	SessionStart        sql.NullInt64  `json:"session_start"`         // session_start
	SessionLastActivity sql.NullInt64  `json:"session_last_activity"` // session_last_activity
	SessionIPAddress    sql.NullString `json:"session_ip_address"`    // session_ip_address
	SessionUserAgent    sql.NullString `json:"session_user_agent"`    // session_user_agent
	SessionData         sql.NullString `json:"session_data"`          // session_data
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the OauthSession exists in the database.
func (os *OauthSession) Exists() bool {
	return os._exists
}

// Deleted returns true when the OauthSession has been marked for deletion from
// the database.
func (os *OauthSession) Deleted() bool {
	return os._deleted
}

// Insert inserts the OauthSession to the database.
func (os *OauthSession) Insert(ctx context.Context, db DB) error {
	switch {
	case os._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case os._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO spc_holding.oauth_sessions (` +
		`session_id, session_start, session_last_activity, session_ip_address, session_user_agent, session_data` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, os.SessionID, os.SessionStart, os.SessionLastActivity, os.SessionIPAddress, os.SessionUserAgent, os.SessionData)
	if _, err := db.ExecContext(ctx, sqlstr, os.SessionID, os.SessionStart, os.SessionLastActivity, os.SessionIPAddress, os.SessionUserAgent, os.SessionData); err != nil {
		return logerror(err)
	}
	// set exists
	os._exists = true
	return nil
}

// Update updates a OauthSession in the database.
func (os *OauthSession) Update(ctx context.Context, db DB) error {
	switch {
	case !os._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case os._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.oauth_sessions SET ` +
		`session_start = ?, session_last_activity = ?, session_ip_address = ?, session_user_agent = ?, session_data = ? ` +
		`WHERE session_id = ?`
	// run
	logf(sqlstr, os.SessionStart, os.SessionLastActivity, os.SessionIPAddress, os.SessionUserAgent, os.SessionData, os.SessionID)
	if _, err := db.ExecContext(ctx, sqlstr, os.SessionStart, os.SessionLastActivity, os.SessionIPAddress, os.SessionUserAgent, os.SessionData, os.SessionID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the OauthSession to the database.
func (os *OauthSession) Save(ctx context.Context, db DB) error {
	if os.Exists() {
		return os.Update(ctx, db)
	}
	return os.Insert(ctx, db)
}

// Upsert performs an upsert for OauthSession.
func (os *OauthSession) Upsert(ctx context.Context, db DB) error {
	switch {
	case os._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.oauth_sessions (` +
		`session_id, session_start, session_last_activity, session_ip_address, session_user_agent, session_data` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`session_id = VALUES(session_id), session_start = VALUES(session_start), session_last_activity = VALUES(session_last_activity), session_ip_address = VALUES(session_ip_address), session_user_agent = VALUES(session_user_agent), session_data = VALUES(session_data)`
	// run
	logf(sqlstr, os.SessionID, os.SessionStart, os.SessionLastActivity, os.SessionIPAddress, os.SessionUserAgent, os.SessionData)
	if _, err := db.ExecContext(ctx, sqlstr, os.SessionID, os.SessionStart, os.SessionLastActivity, os.SessionIPAddress, os.SessionUserAgent, os.SessionData); err != nil {
		return logerror(err)
	}
	// set exists
	os._exists = true
	return nil
}

// Delete deletes the OauthSession from the database.
func (os *OauthSession) Delete(ctx context.Context, db DB) error {
	switch {
	case !os._exists: // doesn't exist
		return nil
	case os._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.oauth_sessions ` +
		`WHERE session_id = ?`
	// run
	logf(sqlstr, os.SessionID)
	if _, err := db.ExecContext(ctx, sqlstr, os.SessionID); err != nil {
		return logerror(err)
	}
	// set deleted
	os._deleted = true
	return nil
}

// OauthSessionBySessionID retrieves a row from 'spc_holding.oauth_sessions' as a OauthSession.
//
// Generated from index 'oauth_sessions_session_id_pkey'.
func OauthSessionBySessionID(ctx context.Context, db DB, sessionID string) (*OauthSession, error) {
	// query
	const sqlstr = `SELECT ` +
		`session_id, session_start, session_last_activity, session_ip_address, session_user_agent, session_data ` +
		`FROM spc_holding.oauth_sessions ` +
		`WHERE session_id = ?`
	// run
	logf(sqlstr, sessionID)
	os := OauthSession{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, sessionID).Scan(&os.SessionID, &os.SessionStart, &os.SessionLastActivity, &os.SessionIPAddress, &os.SessionUserAgent, &os.SessionData); err != nil {
		return nil, logerror(err)
	}
	return &os, nil
}

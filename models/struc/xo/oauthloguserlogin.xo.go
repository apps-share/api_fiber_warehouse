package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// OauthLogUserLogin represents a row from 'spc_holding.oauth_log_user_login'.
type OauthLogUserLogin struct {
	IDLogUserLogin      uint           `json:"id_log_user_login"`      // id_log_user_login
	SessionID           sql.NullString `json:"session_id"`             // session_id
	IDUser              sql.NullInt64  `json:"id_user"`                // id_user
	IDBranch            sql.NullInt64  `json:"id_branch"`              // id_branch
	EmailLogin          sql.NullString `json:"email_login"`            // email_login
	IsSuccess1Fail2Flag sql.NullBool   `json:"is_success1_fail2_flag"` // is_success1_fail2_flag
	LoginTime           sql.NullInt64  `json:"login_time"`             // login_time
	LogoutTime          sql.NullInt64  `json:"logout_time"`            // logout_time
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the OauthLogUserLogin exists in the database.
func (olul *OauthLogUserLogin) Exists() bool {
	return olul._exists
}

// Deleted returns true when the OauthLogUserLogin has been marked for deletion from
// the database.
func (olul *OauthLogUserLogin) Deleted() bool {
	return olul._deleted
}

// Insert inserts the OauthLogUserLogin to the database.
func (olul *OauthLogUserLogin) Insert(ctx context.Context, db DB) error {
	switch {
	case olul._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case olul._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO spc_holding.oauth_log_user_login (` +
		`session_id, id_user, id_branch, email_login, is_success1_fail2_flag, login_time, logout_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, olul.SessionID, olul.IDUser, olul.IDBranch, olul.EmailLogin, olul.IsSuccess1Fail2Flag, olul.LoginTime, olul.LogoutTime)
	res, err := db.ExecContext(ctx, sqlstr, olul.SessionID, olul.IDUser, olul.IDBranch, olul.EmailLogin, olul.IsSuccess1Fail2Flag, olul.LoginTime, olul.LogoutTime)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	olul.IDLogUserLogin = uint(id)
	// set exists
	olul._exists = true
	return nil
}

// Update updates a OauthLogUserLogin in the database.
func (olul *OauthLogUserLogin) Update(ctx context.Context, db DB) error {
	switch {
	case !olul._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case olul._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.oauth_log_user_login SET ` +
		`session_id = ?, id_user = ?, id_branch = ?, email_login = ?, is_success1_fail2_flag = ?, login_time = ?, logout_time = ? ` +
		`WHERE id_log_user_login = ?`
	// run
	logf(sqlstr, olul.SessionID, olul.IDUser, olul.IDBranch, olul.EmailLogin, olul.IsSuccess1Fail2Flag, olul.LoginTime, olul.LogoutTime, olul.IDLogUserLogin)
	if _, err := db.ExecContext(ctx, sqlstr, olul.SessionID, olul.IDUser, olul.IDBranch, olul.EmailLogin, olul.IsSuccess1Fail2Flag, olul.LoginTime, olul.LogoutTime, olul.IDLogUserLogin); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the OauthLogUserLogin to the database.
func (olul *OauthLogUserLogin) Save(ctx context.Context, db DB) error {
	if olul.Exists() {
		return olul.Update(ctx, db)
	}
	return olul.Insert(ctx, db)
}

// Upsert performs an upsert for OauthLogUserLogin.
func (olul *OauthLogUserLogin) Upsert(ctx context.Context, db DB) error {
	switch {
	case olul._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.oauth_log_user_login (` +
		`id_log_user_login, session_id, id_user, id_branch, email_login, is_success1_fail2_flag, login_time, logout_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`session_id = VALUES(session_id), id_user = VALUES(id_user), id_branch = VALUES(id_branch), email_login = VALUES(email_login), is_success1_fail2_flag = VALUES(is_success1_fail2_flag), login_time = VALUES(login_time), logout_time = VALUES(logout_time)`
	// run
	logf(sqlstr, olul.IDLogUserLogin, olul.SessionID, olul.IDUser, olul.IDBranch, olul.EmailLogin, olul.IsSuccess1Fail2Flag, olul.LoginTime, olul.LogoutTime)
	if _, err := db.ExecContext(ctx, sqlstr, olul.IDLogUserLogin, olul.SessionID, olul.IDUser, olul.IDBranch, olul.EmailLogin, olul.IsSuccess1Fail2Flag, olul.LoginTime, olul.LogoutTime); err != nil {
		return logerror(err)
	}
	// set exists
	olul._exists = true
	return nil
}

// Delete deletes the OauthLogUserLogin from the database.
func (olul *OauthLogUserLogin) Delete(ctx context.Context, db DB) error {
	switch {
	case !olul._exists: // doesn't exist
		return nil
	case olul._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.oauth_log_user_login ` +
		`WHERE id_log_user_login = ?`
	// run
	logf(sqlstr, olul.IDLogUserLogin)
	if _, err := db.ExecContext(ctx, sqlstr, olul.IDLogUserLogin); err != nil {
		return logerror(err)
	}
	// set deleted
	olul._deleted = true
	return nil
}

// OauthLogUserLoginByIDLogUserLogin retrieves a row from 'spc_holding.oauth_log_user_login' as a OauthLogUserLogin.
//
// Generated from index 'oauth_log_user_login_id_log_user_login_pkey'.
func OauthLogUserLoginByIDLogUserLogin(ctx context.Context, db DB, idLogUserLogin uint) (*OauthLogUserLogin, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_log_user_login, session_id, id_user, id_branch, email_login, is_success1_fail2_flag, login_time, logout_time ` +
		`FROM spc_holding.oauth_log_user_login ` +
		`WHERE id_log_user_login = ?`
	// run
	logf(sqlstr, idLogUserLogin)
	olul := OauthLogUserLogin{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, idLogUserLogin).Scan(&olul.IDLogUserLogin, &olul.SessionID, &olul.IDUser, &olul.IDBranch, &olul.EmailLogin, &olul.IsSuccess1Fail2Flag, &olul.LoginTime, &olul.LogoutTime); err != nil {
		return nil, logerror(err)
	}
	return &olul, nil
}

package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// JobTemplateType represents a row from 'spc_holding.job_template_type'.
type JobTemplateType struct {
	IDTemplateType uint          `json:"id_template_type"` // id_template_type
	TemplTypeName  sql.NullInt64 `json:"templ_type_name"`  // templ_type_name
	Show356        sql.NullBool  `json:"show_356"`         // show_356
	TemplTypeRoder sql.NullInt64 `json:"templ_type_roder"` // templ_type_roder
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the JobTemplateType exists in the database.
func (jtt *JobTemplateType) Exists() bool {
	return jtt._exists
}

// Deleted returns true when the JobTemplateType has been marked for deletion from
// the database.
func (jtt *JobTemplateType) Deleted() bool {
	return jtt._deleted
}

// Insert inserts the JobTemplateType to the database.
func (jtt *JobTemplateType) Insert(ctx context.Context, db DB) error {
	switch {
	case jtt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case jtt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO spc_holding.job_template_type (` +
		`templ_type_name, show_356, templ_type_roder` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`
	// run
	logf(sqlstr, jtt.TemplTypeName, jtt.Show356, jtt.TemplTypeRoder)
	res, err := db.ExecContext(ctx, sqlstr, jtt.TemplTypeName, jtt.Show356, jtt.TemplTypeRoder)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	jtt.IDTemplateType = uint(id)
	// set exists
	jtt._exists = true
	return nil
}

// Update updates a JobTemplateType in the database.
func (jtt *JobTemplateType) Update(ctx context.Context, db DB) error {
	switch {
	case !jtt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case jtt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.job_template_type SET ` +
		`templ_type_name = ?, show_356 = ?, templ_type_roder = ? ` +
		`WHERE id_template_type = ?`
	// run
	logf(sqlstr, jtt.TemplTypeName, jtt.Show356, jtt.TemplTypeRoder, jtt.IDTemplateType)
	if _, err := db.ExecContext(ctx, sqlstr, jtt.TemplTypeName, jtt.Show356, jtt.TemplTypeRoder, jtt.IDTemplateType); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the JobTemplateType to the database.
func (jtt *JobTemplateType) Save(ctx context.Context, db DB) error {
	if jtt.Exists() {
		return jtt.Update(ctx, db)
	}
	return jtt.Insert(ctx, db)
}

// Upsert performs an upsert for JobTemplateType.
func (jtt *JobTemplateType) Upsert(ctx context.Context, db DB) error {
	switch {
	case jtt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.job_template_type (` +
		`id_template_type, templ_type_name, show_356, templ_type_roder` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`templ_type_name = VALUES(templ_type_name), show_356 = VALUES(show_356), templ_type_roder = VALUES(templ_type_roder)`
	// run
	logf(sqlstr, jtt.IDTemplateType, jtt.TemplTypeName, jtt.Show356, jtt.TemplTypeRoder)
	if _, err := db.ExecContext(ctx, sqlstr, jtt.IDTemplateType, jtt.TemplTypeName, jtt.Show356, jtt.TemplTypeRoder); err != nil {
		return logerror(err)
	}
	// set exists
	jtt._exists = true
	return nil
}

// Delete deletes the JobTemplateType from the database.
func (jtt *JobTemplateType) Delete(ctx context.Context, db DB) error {
	switch {
	case !jtt._exists: // doesn't exist
		return nil
	case jtt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.job_template_type ` +
		`WHERE id_template_type = ?`
	// run
	logf(sqlstr, jtt.IDTemplateType)
	if _, err := db.ExecContext(ctx, sqlstr, jtt.IDTemplateType); err != nil {
		return logerror(err)
	}
	// set deleted
	jtt._deleted = true
	return nil
}

// JobTemplateTypeByIDTemplateType retrieves a row from 'spc_holding.job_template_type' as a JobTemplateType.
//
// Generated from index 'job_template_type_id_template_type_pkey'.
func JobTemplateTypeByIDTemplateType(ctx context.Context, db DB, idTemplateType uint) (*JobTemplateType, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_template_type, templ_type_name, show_356, templ_type_roder ` +
		`FROM spc_holding.job_template_type ` +
		`WHERE id_template_type = ?`
	// run
	logf(sqlstr, idTemplateType)
	jtt := JobTemplateType{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, idTemplateType).Scan(&jtt.IDTemplateType, &jtt.TemplTypeName, &jtt.Show356, &jtt.TemplTypeRoder); err != nil {
		return nil, logerror(err)
	}
	return &jtt, nil
}

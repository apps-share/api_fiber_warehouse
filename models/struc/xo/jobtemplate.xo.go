package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// JobTemplate represents a row from 'spc_holding.job_template'.
type JobTemplate struct {
	IDJobTemplate  uint          `json:"id_job_template"`  // id_job_template
	TemplName      sql.NullInt64 `json:"templ_name"`       // templ_name
	CostPrice      sql.NullInt64 `json:"cost_price"`       // cost_price
	IDTemplateType sql.NullBool  `json:"id_template_type"` // id_template_type
	TemplOrder     sql.NullInt64 `json:"templ_order"`      // templ_order
	Active356      sql.NullBool  `json:"active_356"`       // active_356
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the JobTemplate exists in the database.
func (jt *JobTemplate) Exists() bool {
	return jt._exists
}

// Deleted returns true when the JobTemplate has been marked for deletion from
// the database.
func (jt *JobTemplate) Deleted() bool {
	return jt._deleted
}

// Insert inserts the JobTemplate to the database.
func (jt *JobTemplate) Insert(ctx context.Context, db DB) error {
	switch {
	case jt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case jt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO spc_holding.job_template (` +
		`templ_name, cost_price, id_template_type, templ_order, active_356` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, jt.TemplName, jt.CostPrice, jt.IDTemplateType, jt.TemplOrder, jt.Active356)
	res, err := db.ExecContext(ctx, sqlstr, jt.TemplName, jt.CostPrice, jt.IDTemplateType, jt.TemplOrder, jt.Active356)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	jt.IDJobTemplate = uint(id)
	// set exists
	jt._exists = true
	return nil
}

// Update updates a JobTemplate in the database.
func (jt *JobTemplate) Update(ctx context.Context, db DB) error {
	switch {
	case !jt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case jt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.job_template SET ` +
		`templ_name = ?, cost_price = ?, id_template_type = ?, templ_order = ?, active_356 = ? ` +
		`WHERE id_job_template = ?`
	// run
	logf(sqlstr, jt.TemplName, jt.CostPrice, jt.IDTemplateType, jt.TemplOrder, jt.Active356, jt.IDJobTemplate)
	if _, err := db.ExecContext(ctx, sqlstr, jt.TemplName, jt.CostPrice, jt.IDTemplateType, jt.TemplOrder, jt.Active356, jt.IDJobTemplate); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the JobTemplate to the database.
func (jt *JobTemplate) Save(ctx context.Context, db DB) error {
	if jt.Exists() {
		return jt.Update(ctx, db)
	}
	return jt.Insert(ctx, db)
}

// Upsert performs an upsert for JobTemplate.
func (jt *JobTemplate) Upsert(ctx context.Context, db DB) error {
	switch {
	case jt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.job_template (` +
		`id_job_template, templ_name, cost_price, id_template_type, templ_order, active_356` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`templ_name = VALUES(templ_name), cost_price = VALUES(cost_price), id_template_type = VALUES(id_template_type), templ_order = VALUES(templ_order), active_356 = VALUES(active_356)`
	// run
	logf(sqlstr, jt.IDJobTemplate, jt.TemplName, jt.CostPrice, jt.IDTemplateType, jt.TemplOrder, jt.Active356)
	if _, err := db.ExecContext(ctx, sqlstr, jt.IDJobTemplate, jt.TemplName, jt.CostPrice, jt.IDTemplateType, jt.TemplOrder, jt.Active356); err != nil {
		return logerror(err)
	}
	// set exists
	jt._exists = true
	return nil
}

// Delete deletes the JobTemplate from the database.
func (jt *JobTemplate) Delete(ctx context.Context, db DB) error {
	switch {
	case !jt._exists: // doesn't exist
		return nil
	case jt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.job_template ` +
		`WHERE id_job_template = ?`
	// run
	logf(sqlstr, jt.IDJobTemplate)
	if _, err := db.ExecContext(ctx, sqlstr, jt.IDJobTemplate); err != nil {
		return logerror(err)
	}
	// set deleted
	jt._deleted = true
	return nil
}

// JobTemplateByIDJobTemplate retrieves a row from 'spc_holding.job_template' as a JobTemplate.
//
// Generated from index 'job_template_id_job_template_pkey'.
func JobTemplateByIDJobTemplate(ctx context.Context, db DB, idJobTemplate uint) (*JobTemplate, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_job_template, templ_name, cost_price, id_template_type, templ_order, active_356 ` +
		`FROM spc_holding.job_template ` +
		`WHERE id_job_template = ?`
	// run
	logf(sqlstr, idJobTemplate)
	jt := JobTemplate{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, idJobTemplate).Scan(&jt.IDJobTemplate, &jt.TemplName, &jt.CostPrice, &jt.IDTemplateType, &jt.TemplOrder, &jt.Active356); err != nil {
		return nil, logerror(err)
	}
	return &jt, nil
}

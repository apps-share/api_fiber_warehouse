package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// ServiceSchedule represents a row from 'spc_holding.service_schedule'.
type ServiceSchedule struct {
	IDServiceSchedule   uint           `json:"id_service schedule"`   // id_service schedule
	ServiceDateDdmm20yy sql.NullInt64  `json:"service_date_ddmm20yy"` // service_date_ddmm20yy
	ServiceMm           sql.NullString `json:"service_mm"`            // service_mm
	ServiceYear         sql.NullString `json:"service_year"`          // service_year
	IDServiceCore       sql.NullInt64  `json:"id_service_core"`       // id_service_core
	IDServiceTrans      sql.NullInt64  `json:"id_service_trans"`      // id_service_trans
	IDStaffCore         sql.NullInt64  `json:"id_staff_core"`         // id_staff_core
	ServiceDone356      sql.NullBool   `json:"service_done_356"`      // service_done_356
	RecordCreateTime    sql.NullInt64  `json:"record_create_time"`    // record_create_time
	RecordCreateByID    sql.NullInt64  `json:"record_create_by_id"`   // record_create_by_id
	RecordCreateByName  sql.NullInt64  `json:"record_create_by_name"` // record_create_by_name
	RecordUpdateTime    sql.NullInt64  `json:"record_update_time"`    // record_update_time
	RecordUpdateByID    sql.NullInt64  `json:"record_update_by_id"`   // record_update_by_id
	RecordUpdateByName  sql.NullString `json:"record_update_by_name"` // record_update_by_name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the ServiceSchedule exists in the database.
func (ss *ServiceSchedule) Exists() bool {
	return ss._exists
}

// Deleted returns true when the ServiceSchedule has been marked for deletion from
// the database.
func (ss *ServiceSchedule) Deleted() bool {
	return ss._deleted
}

// Insert inserts the ServiceSchedule to the database.
func (ss *ServiceSchedule) Insert(ctx context.Context, db DB) error {
	switch {
	case ss._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ss._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO spc_holding.service_schedule (` +
		`service_date_ddmm20yy, service_mm, service_year, id_service_core, id_service_trans, id_staff_core, service_done_356, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, ss.ServiceDateDdmm20yy, ss.ServiceMm, ss.ServiceYear, ss.IDServiceCore, ss.IDServiceTrans, ss.IDStaffCore, ss.ServiceDone356, ss.RecordCreateTime, ss.RecordCreateByID, ss.RecordCreateByName, ss.RecordUpdateTime, ss.RecordUpdateByID, ss.RecordUpdateByName)
	res, err := db.ExecContext(ctx, sqlstr, ss.ServiceDateDdmm20yy, ss.ServiceMm, ss.ServiceYear, ss.IDServiceCore, ss.IDServiceTrans, ss.IDStaffCore, ss.ServiceDone356, ss.RecordCreateTime, ss.RecordCreateByID, ss.RecordCreateByName, ss.RecordUpdateTime, ss.RecordUpdateByID, ss.RecordUpdateByName)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	ss.IDServiceSchedule = uint(id)
	// set exists
	ss._exists = true
	return nil
}

// Update updates a ServiceSchedule in the database.
func (ss *ServiceSchedule) Update(ctx context.Context, db DB) error {
	switch {
	case !ss._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ss._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.service_schedule SET ` +
		`service_date_ddmm20yy = ?, service_mm = ?, service_year = ?, id_service_core = ?, id_service_trans = ?, id_staff_core = ?, service_done_356 = ?, record_create_time = ?, record_create_by_id = ?, record_create_by_name = ?, record_update_time = ?, record_update_by_id = ?, record_update_by_name = ? ` +
		`WHERE id_service schedule = ?`
	// run
	logf(sqlstr, ss.ServiceDateDdmm20yy, ss.ServiceMm, ss.ServiceYear, ss.IDServiceCore, ss.IDServiceTrans, ss.IDStaffCore, ss.ServiceDone356, ss.RecordCreateTime, ss.RecordCreateByID, ss.RecordCreateByName, ss.RecordUpdateTime, ss.RecordUpdateByID, ss.RecordUpdateByName, ss.IDServiceSchedule)
	if _, err := db.ExecContext(ctx, sqlstr, ss.ServiceDateDdmm20yy, ss.ServiceMm, ss.ServiceYear, ss.IDServiceCore, ss.IDServiceTrans, ss.IDStaffCore, ss.ServiceDone356, ss.RecordCreateTime, ss.RecordCreateByID, ss.RecordCreateByName, ss.RecordUpdateTime, ss.RecordUpdateByID, ss.RecordUpdateByName, ss.IDServiceSchedule); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the ServiceSchedule to the database.
func (ss *ServiceSchedule) Save(ctx context.Context, db DB) error {
	if ss.Exists() {
		return ss.Update(ctx, db)
	}
	return ss.Insert(ctx, db)
}

// Upsert performs an upsert for ServiceSchedule.
func (ss *ServiceSchedule) Upsert(ctx context.Context, db DB) error {
	switch {
	case ss._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.service_schedule (` +
		`id_service schedule, service_date_ddmm20yy, service_mm, service_year, id_service_core, id_service_trans, id_staff_core, service_done_356, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`service_date_ddmm20yy = VALUES(service_date_ddmm20yy), service_mm = VALUES(service_mm), service_year = VALUES(service_year), id_service_core = VALUES(id_service_core), id_service_trans = VALUES(id_service_trans), id_staff_core = VALUES(id_staff_core), service_done_356 = VALUES(service_done_356), record_create_time = VALUES(record_create_time), record_create_by_id = VALUES(record_create_by_id), record_create_by_name = VALUES(record_create_by_name), record_update_time = VALUES(record_update_time), record_update_by_id = VALUES(record_update_by_id), record_update_by_name = VALUES(record_update_by_name)`
	// run
	logf(sqlstr, ss.IDServiceSchedule, ss.ServiceDateDdmm20yy, ss.ServiceMm, ss.ServiceYear, ss.IDServiceCore, ss.IDServiceTrans, ss.IDStaffCore, ss.ServiceDone356, ss.RecordCreateTime, ss.RecordCreateByID, ss.RecordCreateByName, ss.RecordUpdateTime, ss.RecordUpdateByID, ss.RecordUpdateByName)
	if _, err := db.ExecContext(ctx, sqlstr, ss.IDServiceSchedule, ss.ServiceDateDdmm20yy, ss.ServiceMm, ss.ServiceYear, ss.IDServiceCore, ss.IDServiceTrans, ss.IDStaffCore, ss.ServiceDone356, ss.RecordCreateTime, ss.RecordCreateByID, ss.RecordCreateByName, ss.RecordUpdateTime, ss.RecordUpdateByID, ss.RecordUpdateByName); err != nil {
		return logerror(err)
	}
	// set exists
	ss._exists = true
	return nil
}

// Delete deletes the ServiceSchedule from the database.
func (ss *ServiceSchedule) Delete(ctx context.Context, db DB) error {
	switch {
	case !ss._exists: // doesn't exist
		return nil
	case ss._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.service_schedule ` +
		`WHERE id_service schedule = ?`
	// run
	logf(sqlstr, ss.IDServiceSchedule)
	if _, err := db.ExecContext(ctx, sqlstr, ss.IDServiceSchedule); err != nil {
		return logerror(err)
	}
	// set deleted
	ss._deleted = true
	return nil
}

// ServiceScheduleByIDServiceSchedule retrieves a row from 'spc_holding.service_schedule' as a ServiceSchedule.
//
// Generated from index 'service_schedule_id_service schedule_pkey'.
func ServiceScheduleByIDServiceSchedule(ctx context.Context, db DB, idserviceSchedule uint) (*ServiceSchedule, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_service schedule, service_date_ddmm20yy, service_mm, service_year, id_service_core, id_service_trans, id_staff_core, service_done_356, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name ` +
		`FROM spc_holding.service_schedule ` +
		`WHERE id_service schedule = ?`
	// run
	logf(sqlstr, idserviceSchedule)
	ss := ServiceSchedule{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, idserviceSchedule).Scan(&ss.IDServiceSchedule, &ss.ServiceDateDdmm20yy, &ss.ServiceMm, &ss.ServiceYear, &ss.IDServiceCore, &ss.IDServiceTrans, &ss.IDStaffCore, &ss.ServiceDone356, &ss.RecordCreateTime, &ss.RecordCreateByID, &ss.RecordCreateByName, &ss.RecordUpdateTime, &ss.RecordUpdateByID, &ss.RecordUpdateByName); err != nil {
		return nil, logerror(err)
	}
	return &ss, nil
}

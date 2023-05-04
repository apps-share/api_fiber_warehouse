package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DocPayTran represents a row from 'spc_holding.doc_pay_trans'.
type DocPayTran struct {
	IDDocPayTrans      uint            `json:"id_doc_pay_trans"`      // id_doc_pay_trans
	IDDocPayCore       sql.NullInt64   `json:"id_doc_pay_core"`       // id_doc_pay_core
	IDDocCore          sql.NullInt64   `json:"id_doc_core"`           // id_doc_core
	Paid01Ddmm20yy     sql.NullString  `json:"paid_01_ddmm20yy"`      // paid_01_ddmm20yy
	Paid01TimeInt      sql.NullInt64   `json:"paid_01_time_int"`      // paid_01_time_int
	Paid01Amt          sql.NullFloat64 `json:"paid_01_amt"`           // paid_01_amt
	Paid01Desc         sql.NullString  `json:"paid_01_desc"`          // paid_01_desc
	Paid02Ddmm20yy     sql.NullString  `json:"paid_02_ddmm20yy"`      // paid_02_ddmm20yy
	Paid02TimeInt      sql.NullInt64   `json:"paid_02_time_int"`      // paid_02_time_int
	Paid02Amt          sql.NullFloat64 `json:"paid_02_amt"`           // paid_02_amt
	Paid02Desc         sql.NullString  `json:"paid_02_desc"`          // paid_02_desc
	Paid03Ddmm20yy     sql.NullString  `json:"paid_03_ddmm20yy"`      // paid_03_ddmm20yy
	Paid03TimeInt      sql.NullInt64   `json:"paid_03_time_int"`      // paid_03_time_int
	Paid03Amt          sql.NullFloat64 `json:"paid_03_amt"`           // paid_03_amt
	Paid03Desc         sql.NullString  `json:"paid_03_desc"`          // paid_03_desc
	Paid04Ddmm20yy     sql.NullString  `json:"paid_04_ddmm20yy"`      // paid_04_ddmm20yy
	Paid04TimeInt      sql.NullInt64   `json:"paid_04_time_int"`      // paid_04_time_int
	Paid04Amt          sql.NullFloat64 `json:"paid_04_amt"`           // paid_04_amt
	Paid04Desc         sql.NullString  `json:"paid_04_desc"`          // paid_04_desc
	Paid05Ddmm20yy     sql.NullString  `json:"paid_05_ddmm20yy"`      // paid_05_ddmm20yy
	Paid05TimeInt      sql.NullInt64   `json:"paid_05_time_int"`      // paid_05_time_int
	Paid05Amt          sql.NullFloat64 `json:"paid_05_amt"`           // paid_05_amt
	Paid05Desc         sql.NullString  `json:"paid_05_desc"`          // paid_05_desc
	Paid06Ddmm20yy     sql.NullString  `json:"paid_06_ddmm20yy"`      // paid_06_ddmm20yy
	Paid06TimeInt      sql.NullInt64   `json:"paid_06_time_int"`      // paid_06_time_int
	Paid06Amt          sql.NullFloat64 `json:"paid_06_amt"`           // paid_06_amt
	Paid06Desc         sql.NullString  `json:"paid_06_desc"`          // paid_06_desc
	Paid07Ddmm20yy     sql.NullString  `json:"paid_07_ddmm20yy"`      // paid_07_ddmm20yy
	Paid07TimeInt      sql.NullInt64   `json:"paid_07_time_int"`      // paid_07_time_int
	Paid07Amt          sql.NullFloat64 `json:"paid_07_amt"`           // paid_07_amt
	Paid07Desc         sql.NullString  `json:"paid_07_desc"`          // paid_07_desc
	Paid08Ddmm20yy     sql.NullString  `json:"paid_08_ddmm20yy"`      // paid_08_ddmm20yy
	Paid08TimeInt      sql.NullInt64   `json:"paid_08_time_int"`      // paid_08_time_int
	Paid08Amt          sql.NullFloat64 `json:"paid_08_amt"`           // paid_08_amt
	Paid08Desc         sql.NullString  `json:"paid_08_desc"`          // paid_08_desc
	Paid09Ddmm20yy     sql.NullString  `json:"paid_09_ddmm20yy"`      // paid_09_ddmm20yy
	Paid09TimeInt      sql.NullInt64   `json:"paid_09_time_int"`      // paid_09_time_int
	Paid09Amt          sql.NullFloat64 `json:"paid_09_amt"`           // paid_09_amt
	Paid09Desc         sql.NullString  `json:"paid_09_desc"`          // paid_09_desc
	Paid10Ddmm20yy     sql.NullString  `json:"paid_10_ddmm20yy"`      // paid_10_ddmm20yy
	Paid10TimeInt      sql.NullInt64   `json:"paid_10_time_int"`      // paid_10_time_int
	Paid10Amt          sql.NullFloat64 `json:"paid_10_amt"`           // paid_10_amt
	Paid10Desc         sql.NullString  `json:"paid_10_desc"`          // paid_10_desc
	Paid11Ddmm20yy     sql.NullString  `json:"paid_11_ddmm20yy"`      // paid_11_ddmm20yy
	Paid11TimeInt      sql.NullInt64   `json:"paid_11_time_int"`      // paid_11_time_int
	Paid11Amt          sql.NullFloat64 `json:"paid_11_amt"`           // paid_11_amt
	Paid11Desc         sql.NullString  `json:"paid_11_desc"`          // paid_11_desc
	Paid12Ddmm20yy     sql.NullString  `json:"paid_12_ddmm20yy"`      // paid_12_ddmm20yy
	Paid12TimeInt      sql.NullInt64   `json:"paid_12_time_int"`      // paid_12_time_int
	Paid12Amt          sql.NullFloat64 `json:"paid_12_amt"`           // paid_12_amt
	Paid12Desc         sql.NullString  `json:"paid_12_desc"`          // paid_12_desc
	IsPrepaid356       sql.NullBool    `json:"is_prepaid_356"`        // is_prepaid_356
	PayTransDesc       sql.NullString  `json:"pay_trans_desc"`        // pay_trans_desc
	IsActivegFlag      sql.NullBool    `json:"is_activeg_flag"`       // is_activeg_flag
	RecordCreateTime   sql.NullInt64   `json:"record_create_time"`    // record_create_time
	RecordCreateByID   sql.NullInt64   `json:"record_create_by_id"`   // record_create_by_id
	RecordCreateByName sql.NullString  `json:"record_create_by_name"` // record_create_by_name
	RecordUpdateTime   sql.NullInt64   `json:"record_update_time"`    // record_update_time
	RecordUpdateByID   sql.NullInt64   `json:"record_update_by_id"`   // record_update_by_id
	RecordUpdateByName sql.NullString  `json:"record_update_by_name"` // record_update_by_name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DocPayTran exists in the database.
func (dpt *DocPayTran) Exists() bool {
	return dpt._exists
}

// Deleted returns true when the DocPayTran has been marked for deletion from
// the database.
func (dpt *DocPayTran) Deleted() bool {
	return dpt._deleted
}

// Insert inserts the DocPayTran to the database.
func (dpt *DocPayTran) Insert(ctx context.Context, db DB) error {
	switch {
	case dpt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dpt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO spc_holding.doc_pay_trans (` +
		`id_doc_pay_core, id_doc_core, paid_01_ddmm20yy, paid_01_time_int, paid_01_amt, paid_01_desc, paid_02_ddmm20yy, paid_02_time_int, paid_02_amt, paid_02_desc, paid_03_ddmm20yy, paid_03_time_int, paid_03_amt, paid_03_desc, paid_04_ddmm20yy, paid_04_time_int, paid_04_amt, paid_04_desc, paid_05_ddmm20yy, paid_05_time_int, paid_05_amt, paid_05_desc, paid_06_ddmm20yy, paid_06_time_int, paid_06_amt, paid_06_desc, paid_07_ddmm20yy, paid_07_time_int, paid_07_amt, paid_07_desc, paid_08_ddmm20yy, paid_08_time_int, paid_08_amt, paid_08_desc, paid_09_ddmm20yy, paid_09_time_int, paid_09_amt, paid_09_desc, paid_10_ddmm20yy, paid_10_time_int, paid_10_amt, paid_10_desc, paid_11_ddmm20yy, paid_11_time_int, paid_11_amt, paid_11_desc, paid_12_ddmm20yy, paid_12_time_int, paid_12_amt, paid_12_desc, is_prepaid_356, pay_trans_desc, is_activeg_flag, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, dpt.IDDocPayCore, dpt.IDDocCore, dpt.Paid01Ddmm20yy, dpt.Paid01TimeInt, dpt.Paid01Amt, dpt.Paid01Desc, dpt.Paid02Ddmm20yy, dpt.Paid02TimeInt, dpt.Paid02Amt, dpt.Paid02Desc, dpt.Paid03Ddmm20yy, dpt.Paid03TimeInt, dpt.Paid03Amt, dpt.Paid03Desc, dpt.Paid04Ddmm20yy, dpt.Paid04TimeInt, dpt.Paid04Amt, dpt.Paid04Desc, dpt.Paid05Ddmm20yy, dpt.Paid05TimeInt, dpt.Paid05Amt, dpt.Paid05Desc, dpt.Paid06Ddmm20yy, dpt.Paid06TimeInt, dpt.Paid06Amt, dpt.Paid06Desc, dpt.Paid07Ddmm20yy, dpt.Paid07TimeInt, dpt.Paid07Amt, dpt.Paid07Desc, dpt.Paid08Ddmm20yy, dpt.Paid08TimeInt, dpt.Paid08Amt, dpt.Paid08Desc, dpt.Paid09Ddmm20yy, dpt.Paid09TimeInt, dpt.Paid09Amt, dpt.Paid09Desc, dpt.Paid10Ddmm20yy, dpt.Paid10TimeInt, dpt.Paid10Amt, dpt.Paid10Desc, dpt.Paid11Ddmm20yy, dpt.Paid11TimeInt, dpt.Paid11Amt, dpt.Paid11Desc, dpt.Paid12Ddmm20yy, dpt.Paid12TimeInt, dpt.Paid12Amt, dpt.Paid12Desc, dpt.IsPrepaid356, dpt.PayTransDesc, dpt.IsActivegFlag, dpt.RecordCreateTime, dpt.RecordCreateByID, dpt.RecordCreateByName, dpt.RecordUpdateTime, dpt.RecordUpdateByID, dpt.RecordUpdateByName)
	res, err := db.ExecContext(ctx, sqlstr, dpt.IDDocPayCore, dpt.IDDocCore, dpt.Paid01Ddmm20yy, dpt.Paid01TimeInt, dpt.Paid01Amt, dpt.Paid01Desc, dpt.Paid02Ddmm20yy, dpt.Paid02TimeInt, dpt.Paid02Amt, dpt.Paid02Desc, dpt.Paid03Ddmm20yy, dpt.Paid03TimeInt, dpt.Paid03Amt, dpt.Paid03Desc, dpt.Paid04Ddmm20yy, dpt.Paid04TimeInt, dpt.Paid04Amt, dpt.Paid04Desc, dpt.Paid05Ddmm20yy, dpt.Paid05TimeInt, dpt.Paid05Amt, dpt.Paid05Desc, dpt.Paid06Ddmm20yy, dpt.Paid06TimeInt, dpt.Paid06Amt, dpt.Paid06Desc, dpt.Paid07Ddmm20yy, dpt.Paid07TimeInt, dpt.Paid07Amt, dpt.Paid07Desc, dpt.Paid08Ddmm20yy, dpt.Paid08TimeInt, dpt.Paid08Amt, dpt.Paid08Desc, dpt.Paid09Ddmm20yy, dpt.Paid09TimeInt, dpt.Paid09Amt, dpt.Paid09Desc, dpt.Paid10Ddmm20yy, dpt.Paid10TimeInt, dpt.Paid10Amt, dpt.Paid10Desc, dpt.Paid11Ddmm20yy, dpt.Paid11TimeInt, dpt.Paid11Amt, dpt.Paid11Desc, dpt.Paid12Ddmm20yy, dpt.Paid12TimeInt, dpt.Paid12Amt, dpt.Paid12Desc, dpt.IsPrepaid356, dpt.PayTransDesc, dpt.IsActivegFlag, dpt.RecordCreateTime, dpt.RecordCreateByID, dpt.RecordCreateByName, dpt.RecordUpdateTime, dpt.RecordUpdateByID, dpt.RecordUpdateByName)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	dpt.IDDocPayTrans = uint(id)
	// set exists
	dpt._exists = true
	return nil
}

// Update updates a DocPayTran in the database.
func (dpt *DocPayTran) Update(ctx context.Context, db DB) error {
	switch {
	case !dpt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dpt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.doc_pay_trans SET ` +
		`id_doc_pay_core = ?, id_doc_core = ?, paid_01_ddmm20yy = ?, paid_01_time_int = ?, paid_01_amt = ?, paid_01_desc = ?, paid_02_ddmm20yy = ?, paid_02_time_int = ?, paid_02_amt = ?, paid_02_desc = ?, paid_03_ddmm20yy = ?, paid_03_time_int = ?, paid_03_amt = ?, paid_03_desc = ?, paid_04_ddmm20yy = ?, paid_04_time_int = ?, paid_04_amt = ?, paid_04_desc = ?, paid_05_ddmm20yy = ?, paid_05_time_int = ?, paid_05_amt = ?, paid_05_desc = ?, paid_06_ddmm20yy = ?, paid_06_time_int = ?, paid_06_amt = ?, paid_06_desc = ?, paid_07_ddmm20yy = ?, paid_07_time_int = ?, paid_07_amt = ?, paid_07_desc = ?, paid_08_ddmm20yy = ?, paid_08_time_int = ?, paid_08_amt = ?, paid_08_desc = ?, paid_09_ddmm20yy = ?, paid_09_time_int = ?, paid_09_amt = ?, paid_09_desc = ?, paid_10_ddmm20yy = ?, paid_10_time_int = ?, paid_10_amt = ?, paid_10_desc = ?, paid_11_ddmm20yy = ?, paid_11_time_int = ?, paid_11_amt = ?, paid_11_desc = ?, paid_12_ddmm20yy = ?, paid_12_time_int = ?, paid_12_amt = ?, paid_12_desc = ?, is_prepaid_356 = ?, pay_trans_desc = ?, is_activeg_flag = ?, record_create_time = ?, record_create_by_id = ?, record_create_by_name = ?, record_update_time = ?, record_update_by_id = ?, record_update_by_name = ? ` +
		`WHERE id_doc_pay_trans = ?`
	// run
	logf(sqlstr, dpt.IDDocPayCore, dpt.IDDocCore, dpt.Paid01Ddmm20yy, dpt.Paid01TimeInt, dpt.Paid01Amt, dpt.Paid01Desc, dpt.Paid02Ddmm20yy, dpt.Paid02TimeInt, dpt.Paid02Amt, dpt.Paid02Desc, dpt.Paid03Ddmm20yy, dpt.Paid03TimeInt, dpt.Paid03Amt, dpt.Paid03Desc, dpt.Paid04Ddmm20yy, dpt.Paid04TimeInt, dpt.Paid04Amt, dpt.Paid04Desc, dpt.Paid05Ddmm20yy, dpt.Paid05TimeInt, dpt.Paid05Amt, dpt.Paid05Desc, dpt.Paid06Ddmm20yy, dpt.Paid06TimeInt, dpt.Paid06Amt, dpt.Paid06Desc, dpt.Paid07Ddmm20yy, dpt.Paid07TimeInt, dpt.Paid07Amt, dpt.Paid07Desc, dpt.Paid08Ddmm20yy, dpt.Paid08TimeInt, dpt.Paid08Amt, dpt.Paid08Desc, dpt.Paid09Ddmm20yy, dpt.Paid09TimeInt, dpt.Paid09Amt, dpt.Paid09Desc, dpt.Paid10Ddmm20yy, dpt.Paid10TimeInt, dpt.Paid10Amt, dpt.Paid10Desc, dpt.Paid11Ddmm20yy, dpt.Paid11TimeInt, dpt.Paid11Amt, dpt.Paid11Desc, dpt.Paid12Ddmm20yy, dpt.Paid12TimeInt, dpt.Paid12Amt, dpt.Paid12Desc, dpt.IsPrepaid356, dpt.PayTransDesc, dpt.IsActivegFlag, dpt.RecordCreateTime, dpt.RecordCreateByID, dpt.RecordCreateByName, dpt.RecordUpdateTime, dpt.RecordUpdateByID, dpt.RecordUpdateByName, dpt.IDDocPayTrans)
	if _, err := db.ExecContext(ctx, sqlstr, dpt.IDDocPayCore, dpt.IDDocCore, dpt.Paid01Ddmm20yy, dpt.Paid01TimeInt, dpt.Paid01Amt, dpt.Paid01Desc, dpt.Paid02Ddmm20yy, dpt.Paid02TimeInt, dpt.Paid02Amt, dpt.Paid02Desc, dpt.Paid03Ddmm20yy, dpt.Paid03TimeInt, dpt.Paid03Amt, dpt.Paid03Desc, dpt.Paid04Ddmm20yy, dpt.Paid04TimeInt, dpt.Paid04Amt, dpt.Paid04Desc, dpt.Paid05Ddmm20yy, dpt.Paid05TimeInt, dpt.Paid05Amt, dpt.Paid05Desc, dpt.Paid06Ddmm20yy, dpt.Paid06TimeInt, dpt.Paid06Amt, dpt.Paid06Desc, dpt.Paid07Ddmm20yy, dpt.Paid07TimeInt, dpt.Paid07Amt, dpt.Paid07Desc, dpt.Paid08Ddmm20yy, dpt.Paid08TimeInt, dpt.Paid08Amt, dpt.Paid08Desc, dpt.Paid09Ddmm20yy, dpt.Paid09TimeInt, dpt.Paid09Amt, dpt.Paid09Desc, dpt.Paid10Ddmm20yy, dpt.Paid10TimeInt, dpt.Paid10Amt, dpt.Paid10Desc, dpt.Paid11Ddmm20yy, dpt.Paid11TimeInt, dpt.Paid11Amt, dpt.Paid11Desc, dpt.Paid12Ddmm20yy, dpt.Paid12TimeInt, dpt.Paid12Amt, dpt.Paid12Desc, dpt.IsPrepaid356, dpt.PayTransDesc, dpt.IsActivegFlag, dpt.RecordCreateTime, dpt.RecordCreateByID, dpt.RecordCreateByName, dpt.RecordUpdateTime, dpt.RecordUpdateByID, dpt.RecordUpdateByName, dpt.IDDocPayTrans); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DocPayTran to the database.
func (dpt *DocPayTran) Save(ctx context.Context, db DB) error {
	if dpt.Exists() {
		return dpt.Update(ctx, db)
	}
	return dpt.Insert(ctx, db)
}

// Upsert performs an upsert for DocPayTran.
func (dpt *DocPayTran) Upsert(ctx context.Context, db DB) error {
	switch {
	case dpt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.doc_pay_trans (` +
		`id_doc_pay_trans, id_doc_pay_core, id_doc_core, paid_01_ddmm20yy, paid_01_time_int, paid_01_amt, paid_01_desc, paid_02_ddmm20yy, paid_02_time_int, paid_02_amt, paid_02_desc, paid_03_ddmm20yy, paid_03_time_int, paid_03_amt, paid_03_desc, paid_04_ddmm20yy, paid_04_time_int, paid_04_amt, paid_04_desc, paid_05_ddmm20yy, paid_05_time_int, paid_05_amt, paid_05_desc, paid_06_ddmm20yy, paid_06_time_int, paid_06_amt, paid_06_desc, paid_07_ddmm20yy, paid_07_time_int, paid_07_amt, paid_07_desc, paid_08_ddmm20yy, paid_08_time_int, paid_08_amt, paid_08_desc, paid_09_ddmm20yy, paid_09_time_int, paid_09_amt, paid_09_desc, paid_10_ddmm20yy, paid_10_time_int, paid_10_amt, paid_10_desc, paid_11_ddmm20yy, paid_11_time_int, paid_11_amt, paid_11_desc, paid_12_ddmm20yy, paid_12_time_int, paid_12_amt, paid_12_desc, is_prepaid_356, pay_trans_desc, is_activeg_flag, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id_doc_pay_core = VALUES(id_doc_pay_core), id_doc_core = VALUES(id_doc_core), paid_01_ddmm20yy = VALUES(paid_01_ddmm20yy), paid_01_time_int = VALUES(paid_01_time_int), paid_01_amt = VALUES(paid_01_amt), paid_01_desc = VALUES(paid_01_desc), paid_02_ddmm20yy = VALUES(paid_02_ddmm20yy), paid_02_time_int = VALUES(paid_02_time_int), paid_02_amt = VALUES(paid_02_amt), paid_02_desc = VALUES(paid_02_desc), paid_03_ddmm20yy = VALUES(paid_03_ddmm20yy), paid_03_time_int = VALUES(paid_03_time_int), paid_03_amt = VALUES(paid_03_amt), paid_03_desc = VALUES(paid_03_desc), paid_04_ddmm20yy = VALUES(paid_04_ddmm20yy), paid_04_time_int = VALUES(paid_04_time_int), paid_04_amt = VALUES(paid_04_amt), paid_04_desc = VALUES(paid_04_desc), paid_05_ddmm20yy = VALUES(paid_05_ddmm20yy), paid_05_time_int = VALUES(paid_05_time_int), paid_05_amt = VALUES(paid_05_amt), paid_05_desc = VALUES(paid_05_desc), paid_06_ddmm20yy = VALUES(paid_06_ddmm20yy), paid_06_time_int = VALUES(paid_06_time_int), paid_06_amt = VALUES(paid_06_amt), paid_06_desc = VALUES(paid_06_desc), paid_07_ddmm20yy = VALUES(paid_07_ddmm20yy), paid_07_time_int = VALUES(paid_07_time_int), paid_07_amt = VALUES(paid_07_amt), paid_07_desc = VALUES(paid_07_desc), paid_08_ddmm20yy = VALUES(paid_08_ddmm20yy), paid_08_time_int = VALUES(paid_08_time_int), paid_08_amt = VALUES(paid_08_amt), paid_08_desc = VALUES(paid_08_desc), paid_09_ddmm20yy = VALUES(paid_09_ddmm20yy), paid_09_time_int = VALUES(paid_09_time_int), paid_09_amt = VALUES(paid_09_amt), paid_09_desc = VALUES(paid_09_desc), paid_10_ddmm20yy = VALUES(paid_10_ddmm20yy), paid_10_time_int = VALUES(paid_10_time_int), paid_10_amt = VALUES(paid_10_amt), paid_10_desc = VALUES(paid_10_desc), paid_11_ddmm20yy = VALUES(paid_11_ddmm20yy), paid_11_time_int = VALUES(paid_11_time_int), paid_11_amt = VALUES(paid_11_amt), paid_11_desc = VALUES(paid_11_desc), paid_12_ddmm20yy = VALUES(paid_12_ddmm20yy), paid_12_time_int = VALUES(paid_12_time_int), paid_12_amt = VALUES(paid_12_amt), paid_12_desc = VALUES(paid_12_desc), is_prepaid_356 = VALUES(is_prepaid_356), pay_trans_desc = VALUES(pay_trans_desc), is_activeg_flag = VALUES(is_activeg_flag), record_create_time = VALUES(record_create_time), record_create_by_id = VALUES(record_create_by_id), record_create_by_name = VALUES(record_create_by_name), record_update_time = VALUES(record_update_time), record_update_by_id = VALUES(record_update_by_id), record_update_by_name = VALUES(record_update_by_name)`
	// run
	logf(sqlstr, dpt.IDDocPayTrans, dpt.IDDocPayCore, dpt.IDDocCore, dpt.Paid01Ddmm20yy, dpt.Paid01TimeInt, dpt.Paid01Amt, dpt.Paid01Desc, dpt.Paid02Ddmm20yy, dpt.Paid02TimeInt, dpt.Paid02Amt, dpt.Paid02Desc, dpt.Paid03Ddmm20yy, dpt.Paid03TimeInt, dpt.Paid03Amt, dpt.Paid03Desc, dpt.Paid04Ddmm20yy, dpt.Paid04TimeInt, dpt.Paid04Amt, dpt.Paid04Desc, dpt.Paid05Ddmm20yy, dpt.Paid05TimeInt, dpt.Paid05Amt, dpt.Paid05Desc, dpt.Paid06Ddmm20yy, dpt.Paid06TimeInt, dpt.Paid06Amt, dpt.Paid06Desc, dpt.Paid07Ddmm20yy, dpt.Paid07TimeInt, dpt.Paid07Amt, dpt.Paid07Desc, dpt.Paid08Ddmm20yy, dpt.Paid08TimeInt, dpt.Paid08Amt, dpt.Paid08Desc, dpt.Paid09Ddmm20yy, dpt.Paid09TimeInt, dpt.Paid09Amt, dpt.Paid09Desc, dpt.Paid10Ddmm20yy, dpt.Paid10TimeInt, dpt.Paid10Amt, dpt.Paid10Desc, dpt.Paid11Ddmm20yy, dpt.Paid11TimeInt, dpt.Paid11Amt, dpt.Paid11Desc, dpt.Paid12Ddmm20yy, dpt.Paid12TimeInt, dpt.Paid12Amt, dpt.Paid12Desc, dpt.IsPrepaid356, dpt.PayTransDesc, dpt.IsActivegFlag, dpt.RecordCreateTime, dpt.RecordCreateByID, dpt.RecordCreateByName, dpt.RecordUpdateTime, dpt.RecordUpdateByID, dpt.RecordUpdateByName)
	if _, err := db.ExecContext(ctx, sqlstr, dpt.IDDocPayTrans, dpt.IDDocPayCore, dpt.IDDocCore, dpt.Paid01Ddmm20yy, dpt.Paid01TimeInt, dpt.Paid01Amt, dpt.Paid01Desc, dpt.Paid02Ddmm20yy, dpt.Paid02TimeInt, dpt.Paid02Amt, dpt.Paid02Desc, dpt.Paid03Ddmm20yy, dpt.Paid03TimeInt, dpt.Paid03Amt, dpt.Paid03Desc, dpt.Paid04Ddmm20yy, dpt.Paid04TimeInt, dpt.Paid04Amt, dpt.Paid04Desc, dpt.Paid05Ddmm20yy, dpt.Paid05TimeInt, dpt.Paid05Amt, dpt.Paid05Desc, dpt.Paid06Ddmm20yy, dpt.Paid06TimeInt, dpt.Paid06Amt, dpt.Paid06Desc, dpt.Paid07Ddmm20yy, dpt.Paid07TimeInt, dpt.Paid07Amt, dpt.Paid07Desc, dpt.Paid08Ddmm20yy, dpt.Paid08TimeInt, dpt.Paid08Amt, dpt.Paid08Desc, dpt.Paid09Ddmm20yy, dpt.Paid09TimeInt, dpt.Paid09Amt, dpt.Paid09Desc, dpt.Paid10Ddmm20yy, dpt.Paid10TimeInt, dpt.Paid10Amt, dpt.Paid10Desc, dpt.Paid11Ddmm20yy, dpt.Paid11TimeInt, dpt.Paid11Amt, dpt.Paid11Desc, dpt.Paid12Ddmm20yy, dpt.Paid12TimeInt, dpt.Paid12Amt, dpt.Paid12Desc, dpt.IsPrepaid356, dpt.PayTransDesc, dpt.IsActivegFlag, dpt.RecordCreateTime, dpt.RecordCreateByID, dpt.RecordCreateByName, dpt.RecordUpdateTime, dpt.RecordUpdateByID, dpt.RecordUpdateByName); err != nil {
		return logerror(err)
	}
	// set exists
	dpt._exists = true
	return nil
}

// Delete deletes the DocPayTran from the database.
func (dpt *DocPayTran) Delete(ctx context.Context, db DB) error {
	switch {
	case !dpt._exists: // doesn't exist
		return nil
	case dpt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.doc_pay_trans ` +
		`WHERE id_doc_pay_trans = ?`
	// run
	logf(sqlstr, dpt.IDDocPayTrans)
	if _, err := db.ExecContext(ctx, sqlstr, dpt.IDDocPayTrans); err != nil {
		return logerror(err)
	}
	// set deleted
	dpt._deleted = true
	return nil
}

// DocPayTranByIDDocPayTrans retrieves a row from 'spc_holding.doc_pay_trans' as a DocPayTran.
//
// Generated from index 'doc_pay_trans_id_doc_pay_trans_pkey'.
func DocPayTranByIDDocPayTrans(ctx context.Context, db DB, idDocPayTrans uint) (*DocPayTran, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_doc_pay_trans, id_doc_pay_core, id_doc_core, paid_01_ddmm20yy, paid_01_time_int, paid_01_amt, paid_01_desc, paid_02_ddmm20yy, paid_02_time_int, paid_02_amt, paid_02_desc, paid_03_ddmm20yy, paid_03_time_int, paid_03_amt, paid_03_desc, paid_04_ddmm20yy, paid_04_time_int, paid_04_amt, paid_04_desc, paid_05_ddmm20yy, paid_05_time_int, paid_05_amt, paid_05_desc, paid_06_ddmm20yy, paid_06_time_int, paid_06_amt, paid_06_desc, paid_07_ddmm20yy, paid_07_time_int, paid_07_amt, paid_07_desc, paid_08_ddmm20yy, paid_08_time_int, paid_08_amt, paid_08_desc, paid_09_ddmm20yy, paid_09_time_int, paid_09_amt, paid_09_desc, paid_10_ddmm20yy, paid_10_time_int, paid_10_amt, paid_10_desc, paid_11_ddmm20yy, paid_11_time_int, paid_11_amt, paid_11_desc, paid_12_ddmm20yy, paid_12_time_int, paid_12_amt, paid_12_desc, is_prepaid_356, pay_trans_desc, is_activeg_flag, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name ` +
		`FROM spc_holding.doc_pay_trans ` +
		`WHERE id_doc_pay_trans = ?`
	// run
	logf(sqlstr, idDocPayTrans)
	dpt := DocPayTran{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, idDocPayTrans).Scan(&dpt.IDDocPayTrans, &dpt.IDDocPayCore, &dpt.IDDocCore, &dpt.Paid01Ddmm20yy, &dpt.Paid01TimeInt, &dpt.Paid01Amt, &dpt.Paid01Desc, &dpt.Paid02Ddmm20yy, &dpt.Paid02TimeInt, &dpt.Paid02Amt, &dpt.Paid02Desc, &dpt.Paid03Ddmm20yy, &dpt.Paid03TimeInt, &dpt.Paid03Amt, &dpt.Paid03Desc, &dpt.Paid04Ddmm20yy, &dpt.Paid04TimeInt, &dpt.Paid04Amt, &dpt.Paid04Desc, &dpt.Paid05Ddmm20yy, &dpt.Paid05TimeInt, &dpt.Paid05Amt, &dpt.Paid05Desc, &dpt.Paid06Ddmm20yy, &dpt.Paid06TimeInt, &dpt.Paid06Amt, &dpt.Paid06Desc, &dpt.Paid07Ddmm20yy, &dpt.Paid07TimeInt, &dpt.Paid07Amt, &dpt.Paid07Desc, &dpt.Paid08Ddmm20yy, &dpt.Paid08TimeInt, &dpt.Paid08Amt, &dpt.Paid08Desc, &dpt.Paid09Ddmm20yy, &dpt.Paid09TimeInt, &dpt.Paid09Amt, &dpt.Paid09Desc, &dpt.Paid10Ddmm20yy, &dpt.Paid10TimeInt, &dpt.Paid10Amt, &dpt.Paid10Desc, &dpt.Paid11Ddmm20yy, &dpt.Paid11TimeInt, &dpt.Paid11Amt, &dpt.Paid11Desc, &dpt.Paid12Ddmm20yy, &dpt.Paid12TimeInt, &dpt.Paid12Amt, &dpt.Paid12Desc, &dpt.IsPrepaid356, &dpt.PayTransDesc, &dpt.IsActivegFlag, &dpt.RecordCreateTime, &dpt.RecordCreateByID, &dpt.RecordCreateByName, &dpt.RecordUpdateTime, &dpt.RecordUpdateByID, &dpt.RecordUpdateByName); err != nil {
		return nil, logerror(err)
	}
	return &dpt, nil
}

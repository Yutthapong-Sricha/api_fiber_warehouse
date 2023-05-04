// Package models contains generated code for schema 'spc_holding'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// ActionLog represents a row from 'spc_holding.action_log'.
type ActionLog struct {
	IDActionLog        int            `json:"id_action_log"`         // id_action_log
	IDActionName       int            `json:"id_action_name"`        // id_action_name
	TmpActorName       sql.NullString `json:"tmp_actor_name"`        // tmp_actor_name
	TmpActionName      sql.NullString `json:"tmp_action_name"`       // tmp_action_name
	IDDoceCore         sql.NullInt64  `json:"id_doce_core"`          // id_doce_core
	ActionNote         sql.NullString `json:"action_note"`           // action_note
	ActionJSON         sql.NullString `json:"action_json"`           // action_json
	RecordUpdateTime   sql.NullInt64  `json:"record_update_time"`    // record_update_time
	RecordUpdateID     sql.NullInt64  `json:"record_update_id"`      // record_update_id
	RecordUpdateByName sql.NullString `json:"record_update_by_name"` // record_update_by_name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the ActionLog exists in the database.
func (al *ActionLog) Exists() bool {
	return al._exists
}

// Deleted returns true when the ActionLog has been marked for deletion from
// the database.
func (al *ActionLog) Deleted() bool {
	return al._deleted
}

// Insert inserts the ActionLog to the database.
func (al *ActionLog) Insert(ctx context.Context, db DB) error {
	switch {
	case al._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case al._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO spc_holding.action_log (` +
		`id_action_name, tmp_actor_name, tmp_action_name, id_doce_core, action_note, action_json, record_update_time, record_update_id, record_update_by_name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, al.IDActionName, al.TmpActorName, al.TmpActionName, al.IDDoceCore, al.ActionNote, al.ActionJSON, al.RecordUpdateTime, al.RecordUpdateID, al.RecordUpdateByName)
	res, err := db.ExecContext(ctx, sqlstr, al.IDActionName, al.TmpActorName, al.TmpActionName, al.IDDoceCore, al.ActionNote, al.ActionJSON, al.RecordUpdateTime, al.RecordUpdateID, al.RecordUpdateByName)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	al.IDActionLog = int(id)
	// set exists
	al._exists = true
	return nil
}

// Update updates a ActionLog in the database.
func (al *ActionLog) Update(ctx context.Context, db DB) error {
	switch {
	case !al._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case al._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.action_log SET ` +
		`id_action_name = ?, tmp_actor_name = ?, tmp_action_name = ?, id_doce_core = ?, action_note = ?, action_json = ?, record_update_time = ?, record_update_id = ?, record_update_by_name = ? ` +
		`WHERE id_action_log = ?`
	// run
	logf(sqlstr, al.IDActionName, al.TmpActorName, al.TmpActionName, al.IDDoceCore, al.ActionNote, al.ActionJSON, al.RecordUpdateTime, al.RecordUpdateID, al.RecordUpdateByName, al.IDActionLog)
	if _, err := db.ExecContext(ctx, sqlstr, al.IDActionName, al.TmpActorName, al.TmpActionName, al.IDDoceCore, al.ActionNote, al.ActionJSON, al.RecordUpdateTime, al.RecordUpdateID, al.RecordUpdateByName, al.IDActionLog); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the ActionLog to the database.
func (al *ActionLog) Save(ctx context.Context, db DB) error {
	if al.Exists() {
		return al.Update(ctx, db)
	}
	return al.Insert(ctx, db)
}

// Upsert performs an upsert for ActionLog.
func (al *ActionLog) Upsert(ctx context.Context, db DB) error {
	switch {
	case al._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.action_log (` +
		`id_action_log, id_action_name, tmp_actor_name, tmp_action_name, id_doce_core, action_note, action_json, record_update_time, record_update_id, record_update_by_name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`id_action_name = VALUES(id_action_name), tmp_actor_name = VALUES(tmp_actor_name), tmp_action_name = VALUES(tmp_action_name), id_doce_core = VALUES(id_doce_core), action_note = VALUES(action_note), action_json = VALUES(action_json), record_update_time = VALUES(record_update_time), record_update_id = VALUES(record_update_id), record_update_by_name = VALUES(record_update_by_name)`
	// run
	logf(sqlstr, al.IDActionLog, al.IDActionName, al.TmpActorName, al.TmpActionName, al.IDDoceCore, al.ActionNote, al.ActionJSON, al.RecordUpdateTime, al.RecordUpdateID, al.RecordUpdateByName)
	if _, err := db.ExecContext(ctx, sqlstr, al.IDActionLog, al.IDActionName, al.TmpActorName, al.TmpActionName, al.IDDoceCore, al.ActionNote, al.ActionJSON, al.RecordUpdateTime, al.RecordUpdateID, al.RecordUpdateByName); err != nil {
		return logerror(err)
	}
	// set exists
	al._exists = true
	return nil
}

// Delete deletes the ActionLog from the database.
func (al *ActionLog) Delete(ctx context.Context, db DB) error {
	switch {
	case !al._exists: // doesn't exist
		return nil
	case al._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.action_log ` +
		`WHERE id_action_log = ?`
	// run
	logf(sqlstr, al.IDActionLog)
	if _, err := db.ExecContext(ctx, sqlstr, al.IDActionLog); err != nil {
		return logerror(err)
	}
	// set deleted
	al._deleted = true
	return nil
}

// ActionLogByIDActionName retrieves a row from 'spc_holding.action_log' as a ActionLog.
//
// Generated from index 'acl_id_action_name'.
func ActionLogByIDActionName(ctx context.Context, db DB, idActionName int) ([]*ActionLog, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_action_log, id_action_name, tmp_actor_name, tmp_action_name, id_doce_core, action_note, action_json, record_update_time, record_update_id, record_update_by_name ` +
		`FROM spc_holding.action_log ` +
		`WHERE id_action_name = ?`
	// run
	logf(sqlstr, idActionName)
	rows, err := db.QueryContext(ctx, sqlstr, idActionName)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ActionLog
	for rows.Next() {
		al := ActionLog{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&al.IDActionLog, &al.IDActionName, &al.TmpActorName, &al.TmpActionName, &al.IDDoceCore, &al.ActionNote, &al.ActionJSON, &al.RecordUpdateTime, &al.RecordUpdateID, &al.RecordUpdateByName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &al)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ActionLogByIDDoceCore retrieves a row from 'spc_holding.action_log' as a ActionLog.
//
// Generated from index 'acl_id_doce_core'.
func ActionLogByIDDoceCore(ctx context.Context, db DB, idDoceCore sql.NullInt64) ([]*ActionLog, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_action_log, id_action_name, tmp_actor_name, tmp_action_name, id_doce_core, action_note, action_json, record_update_time, record_update_id, record_update_by_name ` +
		`FROM spc_holding.action_log ` +
		`WHERE id_doce_core = ?`
	// run
	logf(sqlstr, idDoceCore)
	rows, err := db.QueryContext(ctx, sqlstr, idDoceCore)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ActionLog
	for rows.Next() {
		al := ActionLog{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&al.IDActionLog, &al.IDActionName, &al.TmpActorName, &al.TmpActionName, &al.IDDoceCore, &al.ActionNote, &al.ActionJSON, &al.RecordUpdateTime, &al.RecordUpdateID, &al.RecordUpdateByName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &al)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ActionLogByIDActionLog retrieves a row from 'spc_holding.action_log' as a ActionLog.
//
// Generated from index 'action_log_id_action_log_pkey'.
func ActionLogByIDActionLog(ctx context.Context, db DB, idActionLog int) (*ActionLog, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_action_log, id_action_name, tmp_actor_name, tmp_action_name, id_doce_core, action_note, action_json, record_update_time, record_update_id, record_update_by_name ` +
		`FROM spc_holding.action_log ` +
		`WHERE id_action_log = ?`
	// run
	logf(sqlstr, idActionLog)
	al := ActionLog{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, idActionLog).Scan(&al.IDActionLog, &al.IDActionName, &al.TmpActorName, &al.TmpActionName, &al.IDDoceCore, &al.ActionNote, &al.ActionJSON, &al.RecordUpdateTime, &al.RecordUpdateID, &al.RecordUpdateByName); err != nil {
		return nil, logerror(err)
	}
	return &al, nil
}
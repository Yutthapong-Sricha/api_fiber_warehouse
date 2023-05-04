package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Product represents a row from 'spc_holding.product'.
type Product struct {
	IDProduct          uint            `json:"id_product"`            // id_product
	ProductName        sql.NullString  `json:"product_name"`          // product_name
	ProductBrand       sql.NullString  `json:"product_brand"`         // product_brand
	ProductModel       sql.NullString  `json:"product_model"`         // product_model
	ProductDetail      sql.NullString  `json:"product_detail"`        // product_detail
	ProductYear        sql.NullInt64   `json:"product_year"`          // product_year
	IsFifo356          sql.NullBool    `json:"is_fifo_356"`           // is_fifo_356
	IDProdCategory     sql.NullInt64   `json:"id_prod_category"`      // id_prod_category
	QtyPerPack         sql.NullInt64   `json:"qty_per_pack"`          // qty_per_pack
	UnitName           sql.NullString  `json:"unit_name"`             // unit_name
	UnitNamePack       sql.NullString  `json:"unit_name_pack"`        // unit_name_pack
	IsMain356          sql.NullBool    `json:"is_main_356"`           // is_main_356
	IsVolumeFlag       sql.NullInt64   `json:"is_volume_flag"`        // is_volume_flag
	AvgPriceVol        sql.NullFloat64 `json:"avg_price_vol"`         // avg_price_vol
	SalePricePerUnit   sql.NullFloat64 `json:"sale_price_per_unit"`   // sale_price_per_unit
	IsActiveFlag       sql.NullBool    `json:"is_active_flag"`        // is_active_flag
	RecordCreateTime   sql.NullInt64   `json:"record_create_time"`    // record_create_time
	RecordCreateByID   sql.NullInt64   `json:"record_create_by_id"`   // record_create_by_id
	RecordCreateByName sql.NullString  `json:"record_create_by_name"` // record_create_by_name
	RecordUpdateTime   sql.NullInt64   `json:"record_update_time"`    // record_update_time
	RecordUpdateByID   sql.NullInt64   `json:"record_update_by_id"`   // record_update_by_id
	RecordUpdateByName sql.NullString  `json:"record_update_by_name"` // record_update_by_name
	KeySearch          sql.NullString  `json:"key_search"`            // key_search
	ProductBarcode     sql.NullString  `json:"product_barcode"`       // product_barcode
	IncludeVatFlag     sql.NullInt64   `json:"include_vat_flag"`      // include_vat_flag
	LastCostPrice      sql.NullFloat64 `json:"last_cost_price"`       // last_cost_price
	LastCostTimeInt    sql.NullInt64   `json:"last_cost_time_int"`    // last_cost_time_int
	IDSupplier         sql.NullInt64   `json:"id_supplier"`           // id_supplier
	IsUseCar356        sql.NullBool    `json:"is_use_car_356"`        // is_use_car_356
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Product exists in the database.
func (p *Product) Exists() bool {
	return p._exists
}

// Deleted returns true when the Product has been marked for deletion from
// the database.
func (p *Product) Deleted() bool {
	return p._deleted
}

// Insert inserts the Product to the database.
func (p *Product) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO spc_holding.product (` +
		`product_name, product_brand, product_model, product_detail, product_year, is_fifo_356, id_prod_category, qty_per_pack, unit_name, unit_name_pack, is_main_356, is_volume_flag, avg_price_vol, sale_price_per_unit, is_active_flag, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name, key_search, product_barcode, include_vat_flag, last_cost_price, last_cost_time_int, id_supplier, is_use_car_356` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, p.ProductName, p.ProductBrand, p.ProductModel, p.ProductDetail, p.ProductYear, p.IsFifo356, p.IDProdCategory, p.QtyPerPack, p.UnitName, p.UnitNamePack, p.IsMain356, p.IsVolumeFlag, p.AvgPriceVol, p.SalePricePerUnit, p.IsActiveFlag, p.RecordCreateTime, p.RecordCreateByID, p.RecordCreateByName, p.RecordUpdateTime, p.RecordUpdateByID, p.RecordUpdateByName, p.KeySearch, p.ProductBarcode, p.IncludeVatFlag, p.LastCostPrice, p.LastCostTimeInt, p.IDSupplier, p.IsUseCar356)
	res, err := db.ExecContext(ctx, sqlstr, p.ProductName, p.ProductBrand, p.ProductModel, p.ProductDetail, p.ProductYear, p.IsFifo356, p.IDProdCategory, p.QtyPerPack, p.UnitName, p.UnitNamePack, p.IsMain356, p.IsVolumeFlag, p.AvgPriceVol, p.SalePricePerUnit, p.IsActiveFlag, p.RecordCreateTime, p.RecordCreateByID, p.RecordCreateByName, p.RecordUpdateTime, p.RecordUpdateByID, p.RecordUpdateByName, p.KeySearch, p.ProductBarcode, p.IncludeVatFlag, p.LastCostPrice, p.LastCostTimeInt, p.IDSupplier, p.IsUseCar356)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	p.IDProduct = uint(id)
	// set exists
	p._exists = true
	return nil
}

// Update updates a Product in the database.
func (p *Product) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE spc_holding.product SET ` +
		`product_name = ?, product_brand = ?, product_model = ?, product_detail = ?, product_year = ?, is_fifo_356 = ?, id_prod_category = ?, qty_per_pack = ?, unit_name = ?, unit_name_pack = ?, is_main_356 = ?, is_volume_flag = ?, avg_price_vol = ?, sale_price_per_unit = ?, is_active_flag = ?, record_create_time = ?, record_create_by_id = ?, record_create_by_name = ?, record_update_time = ?, record_update_by_id = ?, record_update_by_name = ?, key_search = ?, product_barcode = ?, include_vat_flag = ?, last_cost_price = ?, last_cost_time_int = ?, id_supplier = ?, is_use_car_356 = ? ` +
		`WHERE id_product = ?`
	// run
	logf(sqlstr, p.ProductName, p.ProductBrand, p.ProductModel, p.ProductDetail, p.ProductYear, p.IsFifo356, p.IDProdCategory, p.QtyPerPack, p.UnitName, p.UnitNamePack, p.IsMain356, p.IsVolumeFlag, p.AvgPriceVol, p.SalePricePerUnit, p.IsActiveFlag, p.RecordCreateTime, p.RecordCreateByID, p.RecordCreateByName, p.RecordUpdateTime, p.RecordUpdateByID, p.RecordUpdateByName, p.KeySearch, p.ProductBarcode, p.IncludeVatFlag, p.LastCostPrice, p.LastCostTimeInt, p.IDSupplier, p.IsUseCar356, p.IDProduct)
	if _, err := db.ExecContext(ctx, sqlstr, p.ProductName, p.ProductBrand, p.ProductModel, p.ProductDetail, p.ProductYear, p.IsFifo356, p.IDProdCategory, p.QtyPerPack, p.UnitName, p.UnitNamePack, p.IsMain356, p.IsVolumeFlag, p.AvgPriceVol, p.SalePricePerUnit, p.IsActiveFlag, p.RecordCreateTime, p.RecordCreateByID, p.RecordCreateByName, p.RecordUpdateTime, p.RecordUpdateByID, p.RecordUpdateByName, p.KeySearch, p.ProductBarcode, p.IncludeVatFlag, p.LastCostPrice, p.LastCostTimeInt, p.IDSupplier, p.IsUseCar356, p.IDProduct); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Product to the database.
func (p *Product) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for Product.
func (p *Product) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO spc_holding.product (` +
		`id_product, product_name, product_brand, product_model, product_detail, product_year, is_fifo_356, id_prod_category, qty_per_pack, unit_name, unit_name_pack, is_main_356, is_volume_flag, avg_price_vol, sale_price_per_unit, is_active_flag, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name, key_search, product_barcode, include_vat_flag, last_cost_price, last_cost_time_int, id_supplier, is_use_car_356` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`product_name = VALUES(product_name), product_brand = VALUES(product_brand), product_model = VALUES(product_model), product_detail = VALUES(product_detail), product_year = VALUES(product_year), is_fifo_356 = VALUES(is_fifo_356), id_prod_category = VALUES(id_prod_category), qty_per_pack = VALUES(qty_per_pack), unit_name = VALUES(unit_name), unit_name_pack = VALUES(unit_name_pack), is_main_356 = VALUES(is_main_356), is_volume_flag = VALUES(is_volume_flag), avg_price_vol = VALUES(avg_price_vol), sale_price_per_unit = VALUES(sale_price_per_unit), is_active_flag = VALUES(is_active_flag), record_create_time = VALUES(record_create_time), record_create_by_id = VALUES(record_create_by_id), record_create_by_name = VALUES(record_create_by_name), record_update_time = VALUES(record_update_time), record_update_by_id = VALUES(record_update_by_id), record_update_by_name = VALUES(record_update_by_name), key_search = VALUES(key_search), product_barcode = VALUES(product_barcode), include_vat_flag = VALUES(include_vat_flag), last_cost_price = VALUES(last_cost_price), last_cost_time_int = VALUES(last_cost_time_int), id_supplier = VALUES(id_supplier), is_use_car_356 = VALUES(is_use_car_356)`
	// run
	logf(sqlstr, p.IDProduct, p.ProductName, p.ProductBrand, p.ProductModel, p.ProductDetail, p.ProductYear, p.IsFifo356, p.IDProdCategory, p.QtyPerPack, p.UnitName, p.UnitNamePack, p.IsMain356, p.IsVolumeFlag, p.AvgPriceVol, p.SalePricePerUnit, p.IsActiveFlag, p.RecordCreateTime, p.RecordCreateByID, p.RecordCreateByName, p.RecordUpdateTime, p.RecordUpdateByID, p.RecordUpdateByName, p.KeySearch, p.ProductBarcode, p.IncludeVatFlag, p.LastCostPrice, p.LastCostTimeInt, p.IDSupplier, p.IsUseCar356)
	if _, err := db.ExecContext(ctx, sqlstr, p.IDProduct, p.ProductName, p.ProductBrand, p.ProductModel, p.ProductDetail, p.ProductYear, p.IsFifo356, p.IDProdCategory, p.QtyPerPack, p.UnitName, p.UnitNamePack, p.IsMain356, p.IsVolumeFlag, p.AvgPriceVol, p.SalePricePerUnit, p.IsActiveFlag, p.RecordCreateTime, p.RecordCreateByID, p.RecordCreateByName, p.RecordUpdateTime, p.RecordUpdateByID, p.RecordUpdateByName, p.KeySearch, p.ProductBarcode, p.IncludeVatFlag, p.LastCostPrice, p.LastCostTimeInt, p.IDSupplier, p.IsUseCar356); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the Product from the database.
func (p *Product) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM spc_holding.product ` +
		`WHERE id_product = ?`
	// run
	logf(sqlstr, p.IDProduct)
	if _, err := db.ExecContext(ctx, sqlstr, p.IDProduct); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// ProductByIDProduct retrieves a row from 'spc_holding.product' as a Product.
//
// Generated from index 'product_id_product_pkey'.
func ProductByIDProduct(ctx context.Context, db DB, idProduct uint) (*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`id_product, product_name, product_brand, product_model, product_detail, product_year, is_fifo_356, id_prod_category, qty_per_pack, unit_name, unit_name_pack, is_main_356, is_volume_flag, avg_price_vol, sale_price_per_unit, is_active_flag, record_create_time, record_create_by_id, record_create_by_name, record_update_time, record_update_by_id, record_update_by_name, key_search, product_barcode, include_vat_flag, last_cost_price, last_cost_time_int, id_supplier, is_use_car_356 ` +
		`FROM spc_holding.product ` +
		`WHERE id_product = ?`
	// run
	logf(sqlstr, idProduct)
	p := Product{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, idProduct).Scan(&p.IDProduct, &p.ProductName, &p.ProductBrand, &p.ProductModel, &p.ProductDetail, &p.ProductYear, &p.IsFifo356, &p.IDProdCategory, &p.QtyPerPack, &p.UnitName, &p.UnitNamePack, &p.IsMain356, &p.IsVolumeFlag, &p.AvgPriceVol, &p.SalePricePerUnit, &p.IsActiveFlag, &p.RecordCreateTime, &p.RecordCreateByID, &p.RecordCreateByName, &p.RecordUpdateTime, &p.RecordUpdateByID, &p.RecordUpdateByName, &p.KeySearch, &p.ProductBarcode, &p.IncludeVatFlag, &p.LastCostPrice, &p.LastCostTimeInt, &p.IDSupplier, &p.IsUseCar356); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}
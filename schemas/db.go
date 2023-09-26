/*
 * SPDX-License-Identifier: MPL-2.0
 *   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package schemas

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/url"
)

func openSqlite(dbFile string) (*gorm.Dialector, error) {
	dial := sqlite.Open(dbFile)
	return &dial, nil
}

func OpenDb(db_url string) (*gorm.DB, error) {
	var err error = nil

	urlObj, err := url.Parse(db_url)
	driver := urlObj.Scheme

	var dial *gorm.Dialector
	if err == nil {
		switch driver {
		case "sqlite3":
			dial, err = openSqlite(urlObj.Path)
		default:
			err = fmt.Errorf("Unsupported driver: %s", driver)
		}
	}
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(*dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&User{},
		&HealthEntry{},
		&Category{},
		&CategoryMultiChoice{},
		&CategorySingleChoiceGroup{},
		&CategorySingleChoiceItem{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Convertable[T interface{}] interface {
	ToModel(ctx context.Context, db *gorm.DB) (T, error)
	FromModel(ctx context.Context, db *gorm.DB, model T) error
	SetParentId(id uint)
}

func DbGetFromId[T interface{}, N int32 | int64 | uint](ctx context.Context, db *gorm.DB, id N) (T, error) {
	var obj T
	db.WithContext(ctx).Where("id = ?", uint(id)).First(&obj)
	if db.Error == nil {
		return obj, nil
	} else {
		return *new(T), db.Error
	}
}

func DbGetManyFromIds[T interface{}, N int32 | int64 | uint](ctx context.Context, db *gorm.DB, ids []N) ([]T, error) {
	objs := make([]T, len(ids))
	for i := range ids {
		obj, err := DbGetFromId[T](ctx, db, uint(ids[i]))
		if err != nil {
			return nil, err
		}
		// todo check permission
		objs[i] = obj
	}
	return objs, nil
}

func DbIdsOfAssoc[N int32 | int64 | uint](ctx context.Context, db *gorm.DB, model interface{}, column string) ([]N, error) {
	var array []gorm.Model
	if err := db.WithContext(ctx).Model(model).Association(column).Find(&array); err != nil {
		return nil, err
	}

	ids := make([]N, len(array))
	for i := range array {
		ids[i] = N(array[i].ID)
	}

	return ids, nil
}

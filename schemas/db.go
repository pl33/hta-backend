/*
 * SPDX-License-Identifier: MPL-2.0
 *   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package schemas

import (
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

	//err = db.AutoMigrate(&User{}, &HealthEntry{})
	err = db.AutoMigrate(&User{}, &User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

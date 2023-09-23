/*
 * SPDX-License-Identifier: MPL-2.0
 *   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	OidcIssuer string
	OidcSub    string
	Name       string
	FirstName  string
	//HealthEntries []HealthEntry
	//Categories    []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

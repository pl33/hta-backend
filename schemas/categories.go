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
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserID        uint
	Title         string
	MultiChoices  []CategoryMultiChoice       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SingleChoices []CategorySingleChoiceGroup `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (owner *User) AddCategory(ctx context.Context, db *gorm.DB, category *Category) error {
	category.UserID = owner.ID
	db.WithContext(ctx).Create(category)
	return db.Error
}

func (category *Category) AddChoice(ctx context.Context, db gorm.DB, choice *CategoryMultiChoice) error {
	choice.CategoryID = category.ID
	db.WithContext(ctx).Create(choice)
	return db.Error
}

func (category *Category) authIsWritable(loggedInUser User) bool {
	return category.UserID == loggedInUser.ID
}

func (category *Category) authIsReadable(loggedInUser User) bool {
	return category.UserID == loggedInUser.ID
}

type CategoryMultiChoice struct {
	gorm.Model
	CategoryID  uint
	Title       string
	Description string
	Entries     []*HealthEntry `gorm:"many2many:entry_multi_choices;"`
}

type CategorySingleChoiceGroup struct {
	gorm.Model
	CategoryID  uint
	Title       string
	Description string
	Choices     []CategorySingleChoiceItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (group *CategorySingleChoiceGroup) AddChoice(ctx context.Context, db gorm.DB, choice *CategorySingleChoiceItem) error {
	choice.CategorySingleChoiceGroupID = group.ID
	db.WithContext(ctx).Create(choice)
	return db.Error
}

type CategorySingleChoiceItem struct {
	gorm.Model
	CategorySingleChoiceGroupID uint
	Title                       string
	Description                 string
	Entries                     []*HealthEntry `gorm:"many2many:entry_single_choices;"`
}

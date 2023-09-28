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
	"gorm.io/gorm"
	"hta_backend_2/models"
)

type Category struct {
	gorm.Model
	UserID        uint
	Title         string
	MultiChoices  []CategoryMultiChoice       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SingleChoices []CategorySingleChoiceGroup `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (category Category) GetOwnerID(context.Context, *gorm.DB) uint {
	return category.UserID
}

func (category *Category) ToModel(context.Context, *gorm.DB) (models.Category, error) {
	model := models.Category{
		ID:     int32(category.ID),
		UserID: int32(category.UserID),
		Title:  &category.Title,
	}
	return model, nil
}

func (category *Category) FromModel(_ context.Context, _ *gorm.DB, model models.Category) error {
	if model.Title == nil {
		return fmt.Errorf("Title must not be nil")
	}

	category.Title = *model.Title

	return nil
}

func (category *Category) SetParentId(id uint) {
	category.UserID = id
}

type CategoryMultiChoice struct {
	gorm.Model
	CategoryID  uint
	Title       string
	Description string
	Entries     []*HealthEntry `gorm:"many2many:entry_multi_choices;"`
}

func (choice *CategoryMultiChoice) GetParent(ctx context.Context, db *gorm.DB) (Category, error) {
	category, err := DbGetFromId[Category](ctx, db, choice.CategoryID)
	return category, err
}

func (choice CategoryMultiChoice) GetOwnerID(ctx context.Context, db *gorm.DB) uint {
	category, err := choice.GetParent(ctx, db)
	if err == nil {
		return category.UserID
	} else {
		return 0 // Invalid user ID
	}
}

func (choice *CategoryMultiChoice) ToModel(context.Context, *gorm.DB) (models.CategoryMultiChoice, error) {
	model := models.CategoryMultiChoice{
		ID:          int32(choice.ID),
		CategoryID:  int32(choice.CategoryID),
		Title:       &choice.Title,
		Description: choice.Description,
	}
	return model, nil
}

func (choice *CategoryMultiChoice) FromModel(_ context.Context, _ *gorm.DB, model models.CategoryMultiChoice) error {
	choice.Title = *model.Title
	choice.Description = model.Description

	return nil
}

func (choice *CategoryMultiChoice) SetParentId(id uint) {
	choice.CategoryID = id
}

type CategorySingleChoiceGroup struct {
	gorm.Model
	CategoryID  uint
	Title       string
	Description string
	Choices     []CategorySingleChoiceItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (choice *CategorySingleChoiceGroup) GetParent(ctx context.Context, db *gorm.DB) (Category, error) {
	category, err := DbGetFromId[Category](ctx, db, choice.CategoryID)
	return category, err
}

func (choice CategorySingleChoiceGroup) GetOwnerID(ctx context.Context, db *gorm.DB) uint {
	category, err := choice.GetParent(ctx, db)
	if err == nil {
		return category.UserID
	} else {
		return 0 // Invalid user ID
	}
}

func (choice *CategorySingleChoiceGroup) ToModel(context.Context, *gorm.DB) (models.CategorySingleChoiceGroup, error) {
	model := models.CategorySingleChoiceGroup{
		ID:          int32(choice.ID),
		CategoryID:  int32(choice.CategoryID),
		Title:       &choice.Title,
		Description: choice.Description,
	}
	return model, nil
}

func (choice *CategorySingleChoiceGroup) FromModel(_ context.Context, _ *gorm.DB, model models.CategorySingleChoiceGroup) error {
	choice.Title = *model.Title
	choice.Description = model.Description

	return nil
}

func (choice *CategorySingleChoiceGroup) SetParentId(id uint) {
	choice.CategoryID = id
}

type CategorySingleChoiceItem struct {
	gorm.Model
	CategorySingleChoiceGroupID uint
	Title                       string
	Description                 string
	Entries                     []*HealthEntry `gorm:"many2many:entry_single_choices;"`
}

func (choice *CategorySingleChoiceItem) GetParent(ctx context.Context, db *gorm.DB) (CategorySingleChoiceGroup, error) {
	group, err := DbGetFromId[CategorySingleChoiceGroup](ctx, db, choice.CategorySingleChoiceGroupID)
	return group, err
}

func (choice CategorySingleChoiceItem) GetOwnerID(ctx context.Context, db *gorm.DB) uint {
	group, err1 := choice.GetParent(ctx, db)
	if err1 != nil {
		return 0 // Invalid user ID
	}
	category, err2 := group.GetParent(ctx, db)
	if err2 != nil {
		return 0 // Invalid user ID
	}
	return category.UserID
}

func (choice *CategorySingleChoiceItem) ToModel(context.Context, *gorm.DB) (models.CategorySingleChoice, error) {
	model := models.CategorySingleChoice{
		ID:          int32(choice.ID),
		GroupID:     int32(choice.CategorySingleChoiceGroupID),
		Title:       &choice.Title,
		Description: choice.Description,
	}
	return model, nil
}

func (choice *CategorySingleChoiceItem) FromModel(_ context.Context, _ *gorm.DB, model models.CategorySingleChoice) error {
	choice.Title = *model.Title
	choice.Description = model.Description

	return nil
}

func (choice *CategorySingleChoiceItem) SetParentId(id uint) {
	choice.CategorySingleChoiceGroupID = id
}

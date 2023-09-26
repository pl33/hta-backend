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
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
	"hta_backend_2/models"
	"time"
)

type HealthEntry struct {
	gorm.Model

	UserID uint `json:"user_id"`

	Remarks string `json:"remarks"`

	HaveBloodPressure bool    `json:"have_blood_pressure"`
	Systole           float32 `json:"systole"`
	Diastole          float32 `json:"diastole"`
	Pulse             float32 `json:"pulse"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	MultiChoices  []CategoryMultiChoice      `gorm:"many2many:entry_multi_choices;" json:"multi_choices"`
	SingleChoices []CategorySingleChoiceItem `gorm:"many2many:entry_single_choices;" json:"single_choices"`
}

func (owner *User) AddHealthEntry(ctx context.Context, db *gorm.DB, entry *HealthEntry) error {
	entry.UserID = owner.ID
	db.WithContext(ctx).Create(entry)
	return db.Error
}

func (owner *User) ListHealthEntries(ctx context.Context, db *gorm.DB, first *int32, limit *int32) ([]HealthEntry, error) {
	var entries []HealthEntry

	if err := db.WithContext(ctx).Model(owner).Association("HealthEntries").Find(&entries); err != nil {
		return entries, err
	}

	var start, end uint32
	if first != nil {
		start = uint32(*first)
	} else {
		start = 0
	}
	if limit != nil {
		end = start + uint32(*limit)
	} else {
		end = uint32(len(entries))
	}
	if start > uint32(len(entries)) {
		start = uint32(len(entries))
	}
	if end > uint32(len(entries)) {
		end = uint32(len(entries))
	}
	entries = entries[start:end]

	return entries, nil
}

func (entry *HealthEntry) authIsWritable(principalId int32) bool {
	return entry.UserID == uint(principalId)
}

func (entry *HealthEntry) authIsReadable(principalId int32) bool {
	return entry.UserID == uint(principalId)
}

func (entry *HealthEntry) ToModel(ctx context.Context, db *gorm.DB) (models.Entry, error) {
	var err error
	var multiIds []int64
	var singleIds []int64

	multiIds, err = DbIdsOfAssoc[int64](ctx, db, entry, "MultiChoices")
	if err != nil {
		return models.Entry{}, err
	}
	singleIds, err = DbIdsOfAssoc[int64](ctx, db, entry, "SingleChoices")
	if err != nil {
		return models.Entry{}, err
	}

	startTime := strfmt.DateTime(entry.StartTime)
	model := models.Entry{
		ID:                int32(entry.ID),
		UserID:            int32(entry.UserID),
		Remarks:           entry.Remarks,
		HaveBloodPressure: &entry.HaveBloodPressure,
		Systole:           entry.Systole,
		Diastole:          entry.Diastole,
		Pulse:             entry.Pulse,
		StartTime:         &startTime,
		EndTime:           strfmt.DateTime(entry.EndTime),
		MultiChoices:      multiIds,
		SingleChoices:     singleIds,
	}
	return model, nil
}

func (entry *HealthEntry) FromModel(ctx context.Context, db *gorm.DB, model models.Entry) error {
	var err error

	if model.StartTime == nil {
		return fmt.Errorf("StartTime must not be nil")
	}
	if model.HaveBloodPressure == nil {
		return fmt.Errorf("HaveBloodPressure must not be nil")
	}

	startTime := time.Time(*model.StartTime)
	entry.Remarks = model.Remarks
	entry.HaveBloodPressure = *model.HaveBloodPressure
	entry.Systole = model.Systole
	entry.Diastole = model.Diastole
	entry.Pulse = model.Pulse
	entry.StartTime = startTime
	entry.EndTime = time.Time(model.EndTime)
	entry.Pulse = model.Pulse
	entry.MultiChoices, err = DbGetManyFromIds[CategoryMultiChoice](ctx, db, model.MultiChoices)
	if err != nil {
		return err
	}
	entry.SingleChoices, err = DbGetManyFromIds[CategorySingleChoiceItem](ctx, db, model.SingleChoices)
	if err != nil {
		return err
	}

	return nil
}

func (entry *HealthEntry) SetParentId(id uint) {
	entry.UserID = id
}

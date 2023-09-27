/*
 * SPDX-License-Identifier: MPL-2.0
 *   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package restapi

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
	"hta_backend_2/models"
	"hta_backend_2/schemas"
	"net/http"
)

type Convertable[T interface{}] interface {
	ToModel(ctx context.Context, db *gorm.DB) (T, error)
	FromModel(ctx context.Context, db *gorm.DB, model T) error
	SetParentId(id uint)
}

func FromModelFunc[Model interface{}, SchemaPtr Convertable[Model]](schema SchemaPtr, ctx context.Context, db *gorm.DB, model Model) error {
	return schema.FromModel(ctx, db, model)
}

func SetParentIdFunc[Model interface{}, SchemaPtr Convertable[Model]](schema SchemaPtr, parentId uint) {
	schema.SetParentId(parentId)
}

func ToModelFunc[Model interface{}, SchemaPtr Convertable[Model]](schema SchemaPtr, ctx context.Context, db *gorm.DB) (Model, error) {
	return schema.ToModel(ctx, db)
}

func errorResponder(code int, err error) middleware.Responder {
	msg := err.Error()
	responder := middleware.Error(code, models.Error{
		Code:    int32(code),
		Message: &msg,
	})
	return responder
}

func ListHandler[Model interface{}, Schema interface{}, Parent interface{}](
	r *http.Request,
	db *gorm.DB,
	getParent func(ctx context.Context, db *gorm.DB) (Parent, error),
	listItems func(ctx context.Context, db *gorm.DB, parent *Parent) ([]Schema, error),
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(modelList []*Model) middleware.Responder,
) middleware.Responder {
	ctx := r.Context()

	parent, parentErr := getParent(ctx, db)
	if parentErr != nil {
		return errorResponder(http.StatusInternalServerError, parentErr)
	}

	schemaList, schemaErr := listItems(ctx, db, &parent)
	if schemaErr != nil {
		return errorResponder(http.StatusInternalServerError, schemaErr)
	}

	modelList := make([]*Model, len(schemaList))
	for idx := range schemaList {
		item, err := toModelFunc(&schemaList[idx], ctx, db)
		if err != nil {
			return errorResponder(http.StatusInternalServerError, err)
		}
		modelList[idx] = &item
	}

	return successFunc(modelList)
}

func GetHandler[Model interface{}, Schema interface{}, N int32 | int64 | uint](
	r *http.Request,
	db *gorm.DB,
	id N,
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(model *Model) middleware.Responder,
) middleware.Responder {
	var err error
	ctx := r.Context()

	var schema Schema
	schema, err = schemas.DbGetFromId[Schema](ctx, db, id)
	if err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	model, _ := toModelFunc(&schema, ctx, db)
	return successFunc(&model)
}

func PostHandler[Model interface{}, Schema interface{}, N int32 | int64 | uint](
	r *http.Request,
	db *gorm.DB,
	body *Model,
	parentId N,
	fromModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB, model Model) error,
	setParentIdFunc func(schema *Schema, parentId uint),
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(model *Model) middleware.Responder,
) middleware.Responder {
	ctx := r.Context()

	if body == nil {
		return errorResponder(http.StatusBadRequest, fmt.Errorf("body is missing"))
	}

	var schema Schema
	if err := fromModelFunc(&schema, ctx, db, *body); err != nil {
		return errorResponder(http.StatusBadRequest, err)
	}
	setParentIdFunc(&schema, uint(parentId))

	if err := db.WithContext(ctx).Create(&schema).Error; err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	model, _ := toModelFunc(&schema, ctx, db)
	return successFunc(&model)
}

func PutHandler[Model interface{}, Schema interface{}, N int32 | int64 | uint](
	r *http.Request,
	db *gorm.DB,
	body *Model,
	id N,
	fromModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB, model Model) error,
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(model *Model) middleware.Responder,
) middleware.Responder {
	var err error
	ctx := r.Context()

	if body == nil {
		return errorResponder(http.StatusBadRequest, fmt.Errorf("body is missing"))
	}

	var schema Schema
	schema, err = schemas.DbGetFromId[Schema](ctx, db, id)
	if err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	err = fromModelFunc(&schema, ctx, db, *body)
	if err != nil {
		return errorResponder(http.StatusBadRequest, err)
	}

	err = db.WithContext(ctx).Save(&schema).Error
	if err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	model, _ := toModelFunc(&schema, ctx, db)
	return successFunc(&model)
}

func DeleteHandler[Schema interface{}, N int32 | int64 | uint](r *http.Request, db *gorm.DB, id N, successFunc func() middleware.Responder) middleware.Responder {
	var err error
	ctx := r.Context()

	var schema Schema
	schema, err = schemas.DbGetFromId[Schema](ctx, db, id)
	if err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	err = db.WithContext(ctx).Delete(&schema).Error
	if err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	return successFunc()
}

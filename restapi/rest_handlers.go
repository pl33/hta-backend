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

type SchemaAuth interface {
	GetOwnerID(ctx context.Context, db *gorm.DB) uint
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

func ListHandler[Model interface{}, Schema SchemaAuth, Parent interface{}](
	r *http.Request,
	db *gorm.DB,
	principal *schemas.User,
	getParent func(ctx context.Context, db *gorm.DB) (Parent, int, error),
	listItems func(ctx context.Context, db *gorm.DB, parent *Parent) ([]Schema, error),
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(modelList []*Model) middleware.Responder,
) middleware.Responder {
	ctx := r.Context()

	if principal == nil {
		return errorResponder(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
	}

	parent, parentCode, parentErr := getParent(ctx, db)
	if parentErr != nil {
		return errorResponder(parentCode, parentErr)
	}

	schemaList, schemaErr := listItems(ctx, db, &parent)
	if schemaErr != nil {
		return errorResponder(http.StatusInternalServerError, schemaErr)
	}

	modelList := make([]*Model, len(schemaList))
	for idx := range schemaList {
		if principal.ReadAllowed(schemaList[idx].GetOwnerID(ctx, db)) {
			item, err := toModelFunc(&schemaList[idx], ctx, db)
			if err != nil {
				return errorResponder(http.StatusInternalServerError, err)
			}
			modelList[idx] = &item
		}
	}

	return successFunc(modelList)
}

func GetHandler[Model interface{}, Schema SchemaAuth, N int32 | int64 | uint](
	r *http.Request,
	db *gorm.DB,
	id N,
	principal *schemas.User,
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(model *Model) middleware.Responder,
) middleware.Responder {
	var err error
	ctx := r.Context()

	if principal == nil {
		return errorResponder(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
	}

	var schema Schema
	schema, err = schemas.DbGetFromId[Schema](ctx, db, id)
	if err != nil {
		return errorResponder(http.StatusNotFound, err)
	}

	if !principal.ReadAllowed(schema.GetOwnerID(ctx, db)) {
		return errorResponder(http.StatusForbidden, fmt.Errorf("action not permitted"))
	}

	model, _ := toModelFunc(&schema, ctx, db)
	return successFunc(&model)
}

func PostHandler[Model interface{}, Schema SchemaAuth, N int32 | int64 | uint](
	r *http.Request,
	db *gorm.DB,
	body *Model,
	parentId N,
	principal *schemas.User,
	fromModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB, model Model) error,
	modelAuthFunc func(ctx context.Context, db *gorm.DB, principal *schemas.User, schema *Schema) error,
	setParentIdFunc func(schema *Schema, parentId uint),
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(model *Model) middleware.Responder,
) middleware.Responder {
	ctx := r.Context()

	if principal == nil {
		return errorResponder(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
	}

	if body == nil {
		return errorResponder(http.StatusBadRequest, fmt.Errorf("body is missing"))
	}

	var schema Schema
	if err := fromModelFunc(&schema, ctx, db, *body); err != nil {
		return errorResponder(http.StatusBadRequest, err)
	}
	setParentIdFunc(&schema, uint(parentId))

	if !principal.CreateAllowed(schema.GetOwnerID(ctx, db)) {
		return errorResponder(http.StatusForbidden, fmt.Errorf("action not permitted"))
	}

	if modelAuthFunc != nil {
		if err := modelAuthFunc(ctx, db, principal, &schema); err != nil {
			return errorResponder(http.StatusForbidden, err)
		}
	}

	if err := db.WithContext(ctx).Create(&schema).Error; err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	model, _ := toModelFunc(&schema, ctx, db)
	return successFunc(&model)
}

func PutHandler[Model interface{}, Schema SchemaAuth, N int32 | int64 | uint](
	r *http.Request,
	db *gorm.DB,
	body *Model,
	id N,
	principal *schemas.User,
	fromModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB, model Model) error,
	modelAuthFunc func(ctx context.Context, db *gorm.DB, principal *schemas.User, schema *Schema) error,
	replaceAssociationsFunc func(schema *Schema, ctx context.Context, db *gorm.DB) error,
	toModelFunc func(schema *Schema, ctx context.Context, db *gorm.DB) (Model, error),
	successFunc func(model *Model) middleware.Responder,
) middleware.Responder {
	var err error
	ctx := r.Context()

	if principal == nil {
		return errorResponder(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
	}

	if body == nil {
		return errorResponder(http.StatusBadRequest, fmt.Errorf("body is missing"))
	}

	var schema Schema
	schema, err = schemas.DbGetFromId[Schema](ctx, db, id)
	if err != nil {
		return errorResponder(http.StatusNotFound, err)
	}

	if !principal.UpdateAllowed(schema.GetOwnerID(ctx, db)) {
		return errorResponder(http.StatusForbidden, fmt.Errorf("action not permitted"))
	}

	err = fromModelFunc(&schema, ctx, db, *body)
	if err != nil {
		return errorResponder(http.StatusBadRequest, err)
	}

	if modelAuthFunc != nil {
		if err := modelAuthFunc(ctx, db, principal, &schema); err != nil {
			return errorResponder(http.StatusForbidden, err)
		}
	}

	if replaceAssociationsFunc != nil {
		if err := replaceAssociationsFunc(&schema, ctx, db); err != nil {
			return errorResponder(http.StatusInternalServerError, err)
		}
	}

	err = db.WithContext(ctx).Save(&schema).Error
	if err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	model, _ := toModelFunc(&schema, ctx, db)
	return successFunc(&model)
}

func DeleteHandler[Schema SchemaAuth, N int32 | int64 | uint](r *http.Request, db *gorm.DB, id N, principal *schemas.User, successFunc func() middleware.Responder) middleware.Responder {
	var err error
	ctx := r.Context()

	if principal == nil {
		return errorResponder(http.StatusUnauthorized, fmt.Errorf("unauthorized"))
	}

	var schema Schema
	schema, err = schemas.DbGetFromId[Schema](ctx, db, id)
	if err != nil {
		return errorResponder(http.StatusNotFound, err)
	}

	if !principal.DeleteAllowed(schema.GetOwnerID(ctx, db)) {
		return errorResponder(http.StatusForbidden, fmt.Errorf("action not permitted"))
	}

	err = db.WithContext(ctx).Delete(&schema).Error
	if err != nil {
		return errorResponder(http.StatusInternalServerError, err)
	}

	return successFunc()
}

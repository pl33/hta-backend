// This file is safe to edit. Once it exists it will not be overwritten

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
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	"gorm.io/gorm"
	"hta_backend_2/schemas"
	"net/http"

	"hta_backend_2/models"
	"hta_backend_2/restapi/operations"
	"hta_backend_2/restapi/operations/category"
	"hta_backend_2/restapi/operations/entry"
	"hta_backend_2/restapi/operations/login"
)

//go:generate swagger generate server --target ../../hta_backend_2 --name Hta --spec ../swagger.yaml --principal schemas.User

func configureFlags(api *operations.HtaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HtaAPI) http.Handler {
	db, dbErr := schemas.OpenDb(GetEnvOrPanic("DB"))
	if dbErr != nil {
		panic(dbErr)
	}

	var auth Auth
	var err error
	auth, err = AuthSetup(
		db,
		GetEnvOrPanic("OIDC_ISSUER"),
		GetEnvOrPanic("OIDC_CLIENT_ID"),
		GetEnvOrPanic("OIDC_CLIENT_SECRET"),
		GetEnvOrPanic("OIDC_FRONTEND_CLIENT_ID"),
	)
	if err != nil {
		panic(err)
	}

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BearerTokenAuth = func(token string) (*schemas.User, error) {
		user, err := AuthGetUser(context.Background(), &auth, token)
		return &user, err
	}

	api.OauthSecurityAuth = func(token string, scopes []string) (*schemas.User, error) {
		user, err := AuthGetUser(context.Background(), &auth, token)
		return &user, err
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	/////////////////////////////////////////////////////////////////
	// Login
	api.LoginGetOidcInfoHandler = login.GetOidcInfoHandlerFunc(func(params login.GetOidcInfoParams) middleware.Responder {
		discoveryUrl := auth.issuer + "/.well-known/openid-configuration"
		model := models.OidcInfo{
			DiscoveryURL: &discoveryUrl,
			ClientID:     auth.frontendClientId,
			ClientSecret: "",
		}
		return login.NewGetOidcInfoOK().WithPayload(&model)
	})

	/////////////////////////////////////////////////////////////////
	// User
	api.LoginGetUserHandler = login.GetUserHandlerFunc(func(params login.GetUserParams, principal *schemas.User) middleware.Responder {
		model := models.User{
			ID:        int32(principal.ID),
			Name:      principal.Name,
			FirstName: principal.FirstName,
		}
		return login.NewGetUserOK().WithPayload(&model)
	})

	/////////////////////////////////////////////////////////////////
	// Entry
	api.EntryGetEntriesHandler = entry.GetEntriesHandlerFunc(func(params entry.GetEntriesParams, principal *schemas.User) middleware.Responder {
		return ListHandler[models.Entry, schemas.HealthEntry](
			params.HTTPRequest,
			db,
			principal,
			func(ctx context.Context, db *gorm.DB) (schemas.User, int, error) {
				item, err := schemas.LookupUser(ctx, db, principal.ID)
				if err != nil {
					return schemas.User{}, http.StatusNotFound, err
				}
				return item, http.StatusOK, nil
			},
			func(ctx context.Context, db *gorm.DB, owner *schemas.User) ([]schemas.HealthEntry, error) {
				return schemas.DbGetChildSlice[schemas.User, schemas.HealthEntry](ctx, db, owner, "HealthEntries", params.First, params.Limit, "start_time desc")
			},
			ToModelFunc[models.Entry, *schemas.HealthEntry],
			func(modelList []*models.Entry) middleware.Responder {
				return entry.NewGetEntriesOK().WithPayload(modelList)
			},
		)
	})
	api.EntryPostEntriesHandler = entry.PostEntriesHandlerFunc(func(params entry.PostEntriesParams, principal *schemas.User) middleware.Responder {
		return PostHandler[models.Entry, schemas.HealthEntry](
			params.HTTPRequest,
			db,
			params.Body,
			principal.ID,
			principal,
			FromModelFunc[models.Entry, *schemas.HealthEntry],
			schemas.ModelAuthHealthEntry,
			SetParentIdFunc[models.Entry, *schemas.HealthEntry],
			ToModelFunc[models.Entry, *schemas.HealthEntry],
			func(model *models.Entry) middleware.Responder {
				return entry.NewPostEntriesCreated().WithPayload(model)
			},
		)
	})
	api.EntryGetEntriesIDHandler = entry.GetEntriesIDHandlerFunc(func(params entry.GetEntriesIDParams, principal *schemas.User) middleware.Responder {
		return GetHandler[models.Entry, schemas.HealthEntry](
			params.HTTPRequest,
			db,
			params.ID,
			principal,
			ToModelFunc[models.Entry, *schemas.HealthEntry],
			func(model *models.Entry) middleware.Responder {
				return entry.NewGetEntriesIDOK().WithPayload(model)
			},
		)
	})
	api.EntryPutEntriesIDHandler = entry.PutEntriesIDHandlerFunc(func(params entry.PutEntriesIDParams, principal *schemas.User) middleware.Responder {
		return PutHandler[models.Entry, schemas.HealthEntry](
			params.HTTPRequest,
			db,
			params.Body,
			params.ID,
			principal,
			FromModelFunc[models.Entry, *schemas.HealthEntry],
			schemas.ModelAuthHealthEntry,
			func(schema *schemas.HealthEntry, ctx context.Context, db *gorm.DB) error {
				if err := db.WithContext(ctx).Model(schema).Association("MultiChoices").Replace(schema.MultiChoices); err != nil {
					return err
				}
				if err := db.WithContext(ctx).Model(schema).Association("SingleChoices").Replace(schema.SingleChoices); err != nil {
					return err
				}
				return nil
			},
			ToModelFunc[models.Entry, *schemas.HealthEntry],
			func(model *models.Entry) middleware.Responder {
				return entry.NewPutEntriesIDOK().WithPayload(model)
			},
		)
	})
	api.EntryDeleteEntriesIDHandler = entry.DeleteEntriesIDHandlerFunc(func(params entry.DeleteEntriesIDParams, principal *schemas.User) middleware.Responder {
		return DeleteHandler[schemas.HealthEntry](params.HTTPRequest, db, params.ID, principal, func() middleware.Responder {
			return entry.NewDeleteEntriesIDNoContent()
		})
	})

	/////////////////////////////////////////////////////////////////
	// Category
	api.CategoryGetCategoryHandler = category.GetCategoryHandlerFunc(func(params category.GetCategoryParams, principal *schemas.User) middleware.Responder {
		return ListHandler[models.Category, schemas.Category](
			params.HTTPRequest,
			db,
			principal,
			func(ctx context.Context, db *gorm.DB) (schemas.User, int, error) {
				item, err := schemas.LookupUser(ctx, db, principal.ID)
				if err != nil {
					return schemas.User{}, http.StatusNotFound, err
				}
				return item, http.StatusOK, nil
			},
			func(ctx context.Context, db *gorm.DB, owner *schemas.User) ([]schemas.Category, error) {
				return schemas.DbGetChildSlice[schemas.User, schemas.Category](ctx, db, owner, "Categories", params.First, params.Limit, "id asc")
			},
			ToModelFunc[models.Category, *schemas.Category],
			func(modelList []*models.Category) middleware.Responder {
				return category.NewGetCategoryOK().WithPayload(modelList)
			},
		)
	})
	api.CategoryPostCategoryHandler = category.PostCategoryHandlerFunc(func(params category.PostCategoryParams, principal *schemas.User) middleware.Responder {
		return PostHandler[models.Category, schemas.Category](
			params.HTTPRequest,
			db,
			params.Body,
			principal.ID,
			principal,
			FromModelFunc[models.Category, *schemas.Category],
			nil,
			SetParentIdFunc[models.Category, *schemas.Category],
			ToModelFunc[models.Category, *schemas.Category],
			func(model *models.Category) middleware.Responder {
				return category.NewPostCategoryCreated().WithPayload(model)
			},
		)
	})
	api.CategoryGetCategoryIDHandler = category.GetCategoryIDHandlerFunc(func(params category.GetCategoryIDParams, principal *schemas.User) middleware.Responder {
		return GetHandler[models.Category, schemas.Category](
			params.HTTPRequest,
			db,
			params.ID,
			principal,
			ToModelFunc[models.Category, *schemas.Category],
			func(model *models.Category) middleware.Responder {
				return category.NewGetCategoryIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryPutCategoryIDHandler = category.PutCategoryIDHandlerFunc(func(params category.PutCategoryIDParams, principal *schemas.User) middleware.Responder {
		return PutHandler[models.Category, schemas.Category](
			params.HTTPRequest,
			db,
			params.Body,
			params.ID,
			principal,
			FromModelFunc[models.Category, *schemas.Category],
			nil,
			nil,
			ToModelFunc[models.Category, *schemas.Category],
			func(model *models.Category) middleware.Responder {
				return category.NewPutCategoryIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryDeleteCategoryIDHandler = category.DeleteCategoryIDHandlerFunc(func(params category.DeleteCategoryIDParams, principal *schemas.User) middleware.Responder {
		return DeleteHandler[schemas.Category](params.HTTPRequest, db, params.ID, principal, func() middleware.Responder {
			return category.NewDeleteCategoryIDNoContent()
		})
	})

	/////////////////////////////////////////////////////////////////
	// Multi Choice
	api.CategoryGetCategoryCategoryIDMultiChoiceHandler = category.GetCategoryCategoryIDMultiChoiceHandlerFunc(func(params category.GetCategoryCategoryIDMultiChoiceParams, principal *schemas.User) middleware.Responder {
		return ListHandler[models.CategoryMultiChoice, schemas.CategoryMultiChoice](
			params.HTTPRequest,
			db,
			principal,
			func(ctx context.Context, db *gorm.DB) (schemas.Category, int, error) {
				item, err := schemas.DbGetFromId[schemas.Category](ctx, db, params.CategoryID)
				if err != nil {
					return schemas.Category{}, http.StatusNotFound, err
				}
				if !principal.ReadAllowed(item.GetOwnerID(ctx, db)) {
					return schemas.Category{}, http.StatusForbidden, fmt.Errorf("action not permitted")
				}
				return item, http.StatusOK, nil
			},
			func(ctx context.Context, db *gorm.DB, parent *schemas.Category) ([]schemas.CategoryMultiChoice, error) {
				return schemas.DbGetChildSlice[schemas.Category, schemas.CategoryMultiChoice](ctx, db, parent, "MultiChoices", params.First, params.Limit, "id asc")
			},
			ToModelFunc[models.CategoryMultiChoice, *schemas.CategoryMultiChoice],
			func(modelList []*models.CategoryMultiChoice) middleware.Responder {
				return category.NewGetCategoryCategoryIDMultiChoiceOK().WithPayload(modelList)
			},
		)
	})
	api.CategoryPostCategoryCategoryIDMultiChoiceHandler = category.PostCategoryCategoryIDMultiChoiceHandlerFunc(func(params category.PostCategoryCategoryIDMultiChoiceParams, principal *schemas.User) middleware.Responder {
		return PostHandler[models.CategoryMultiChoice, schemas.CategoryMultiChoice](
			params.HTTPRequest,
			db,
			params.Body,
			params.CategoryID,
			principal,
			FromModelFunc[models.CategoryMultiChoice, *schemas.CategoryMultiChoice],
			nil,
			SetParentIdFunc[models.CategoryMultiChoice, *schemas.CategoryMultiChoice],
			ToModelFunc[models.CategoryMultiChoice, *schemas.CategoryMultiChoice],
			func(model *models.CategoryMultiChoice) middleware.Responder {
				return category.NewPostCategoryCategoryIDMultiChoiceCreated().WithPayload(model)
			},
		)
	})
	api.CategoryGetMultiChoiceIDHandler = category.GetMultiChoiceIDHandlerFunc(func(params category.GetMultiChoiceIDParams, principal *schemas.User) middleware.Responder {
		return GetHandler[models.CategoryMultiChoice, schemas.CategoryMultiChoice](
			params.HTTPRequest,
			db,
			params.ID,
			principal,
			ToModelFunc[models.CategoryMultiChoice, *schemas.CategoryMultiChoice],
			func(model *models.CategoryMultiChoice) middleware.Responder {
				return category.NewGetMultiChoiceIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryPutMultiChoiceIDHandler = category.PutMultiChoiceIDHandlerFunc(func(params category.PutMultiChoiceIDParams, principal *schemas.User) middleware.Responder {
		return PutHandler[models.CategoryMultiChoice, schemas.CategoryMultiChoice](
			params.HTTPRequest,
			db,
			params.Body,
			params.ID,
			principal,
			FromModelFunc[models.CategoryMultiChoice, *schemas.CategoryMultiChoice],
			nil,
			nil,
			ToModelFunc[models.CategoryMultiChoice, *schemas.CategoryMultiChoice],
			func(model *models.CategoryMultiChoice) middleware.Responder {
				return category.NewPutMultiChoiceIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryDeleteMultiChoiceIDHandler = category.DeleteMultiChoiceIDHandlerFunc(func(params category.DeleteMultiChoiceIDParams, principal *schemas.User) middleware.Responder {
		return DeleteHandler[schemas.CategoryMultiChoice](params.HTTPRequest, db, params.ID, principal, func() middleware.Responder {
			return category.NewDeleteMultiChoiceIDNoContent()
		})
	})

	/////////////////////////////////////////////////////////////////
	// Single Choice Group
	api.CategoryGetCategoryCategoryIDSingleChoiceGroupHandler = category.GetCategoryCategoryIDSingleChoiceGroupHandlerFunc(func(params category.GetCategoryCategoryIDSingleChoiceGroupParams, principal *schemas.User) middleware.Responder {
		return ListHandler[models.CategorySingleChoiceGroup, schemas.CategorySingleChoiceGroup](
			params.HTTPRequest,
			db,
			principal,
			func(ctx context.Context, db *gorm.DB) (schemas.Category, int, error) {
				item, err := schemas.DbGetFromId[schemas.Category](ctx, db, params.CategoryID)
				if err != nil {
					return schemas.Category{}, http.StatusNotFound, err
				}
				if !principal.ReadAllowed(item.GetOwnerID(ctx, db)) {
					return schemas.Category{}, http.StatusForbidden, fmt.Errorf("action not permitted")
				}
				return item, http.StatusOK, nil
			},
			func(ctx context.Context, db *gorm.DB, parent *schemas.Category) ([]schemas.CategorySingleChoiceGroup, error) {
				return schemas.DbGetChildSlice[schemas.Category, schemas.CategorySingleChoiceGroup](ctx, db, parent, "SingleChoices", params.First, params.Limit, "id asc")
			},
			ToModelFunc[models.CategorySingleChoiceGroup, *schemas.CategorySingleChoiceGroup],
			func(modelList []*models.CategorySingleChoiceGroup) middleware.Responder {
				return category.NewGetCategoryCategoryIDSingleChoiceGroupOK().WithPayload(modelList)
			},
		)
	})
	api.CategoryPostCategoryCategoryIDSingleChoiceGroupHandler = category.PostCategoryCategoryIDSingleChoiceGroupHandlerFunc(func(params category.PostCategoryCategoryIDSingleChoiceGroupParams, principal *schemas.User) middleware.Responder {
		return PostHandler[models.CategorySingleChoiceGroup, schemas.CategorySingleChoiceGroup](
			params.HTTPRequest,
			db,
			params.Body,
			params.CategoryID,
			principal,
			FromModelFunc[models.CategorySingleChoiceGroup, *schemas.CategorySingleChoiceGroup],
			nil,
			SetParentIdFunc[models.CategorySingleChoiceGroup, *schemas.CategorySingleChoiceGroup],
			ToModelFunc[models.CategorySingleChoiceGroup, *schemas.CategorySingleChoiceGroup],
			func(model *models.CategorySingleChoiceGroup) middleware.Responder {
				return category.NewPostCategoryCategoryIDSingleChoiceGroupCreated().WithPayload(model)
			},
		)
	})
	api.CategoryGetSingleChoiceGroupIDHandler = category.GetSingleChoiceGroupIDHandlerFunc(func(params category.GetSingleChoiceGroupIDParams, principal *schemas.User) middleware.Responder {
		return GetHandler[models.CategorySingleChoiceGroup, schemas.CategorySingleChoiceGroup](
			params.HTTPRequest,
			db,
			params.ID,
			principal,
			ToModelFunc[models.CategorySingleChoiceGroup, *schemas.CategorySingleChoiceGroup],
			func(model *models.CategorySingleChoiceGroup) middleware.Responder {
				return category.NewGetSingleChoiceGroupIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryPutSingleChoiceGroupIDHandler = category.PutSingleChoiceGroupIDHandlerFunc(func(params category.PutSingleChoiceGroupIDParams, principal *schemas.User) middleware.Responder {
		return PutHandler[models.CategorySingleChoiceGroup, schemas.CategorySingleChoiceGroup](
			params.HTTPRequest,
			db,
			params.Body,
			params.ID,
			principal,
			FromModelFunc[models.CategorySingleChoiceGroup, *schemas.CategorySingleChoiceGroup],
			nil,
			nil,
			ToModelFunc[models.CategorySingleChoiceGroup, *schemas.CategorySingleChoiceGroup],
			func(model *models.CategorySingleChoiceGroup) middleware.Responder {
				return category.NewPutSingleChoiceGroupIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryDeleteSingleChoiceGroupIDHandler = category.DeleteSingleChoiceGroupIDHandlerFunc(func(params category.DeleteSingleChoiceGroupIDParams, principal *schemas.User) middleware.Responder {
		return DeleteHandler[schemas.CategorySingleChoiceGroup](params.HTTPRequest, db, params.ID, principal, func() middleware.Responder {
			return category.NewDeleteSingleChoiceGroupIDNoContent()
		})
	})

	/////////////////////////////////////////////////////////////////
	// Single Choice Item
	api.CategoryGetSingleChoiceGroupGroupIDSingleChoiceHandler = category.GetSingleChoiceGroupGroupIDSingleChoiceHandlerFunc(func(params category.GetSingleChoiceGroupGroupIDSingleChoiceParams, principal *schemas.User) middleware.Responder {
		return ListHandler[models.CategorySingleChoice, schemas.CategorySingleChoiceItem](
			params.HTTPRequest,
			db,
			principal,
			func(ctx context.Context, db *gorm.DB) (schemas.CategorySingleChoiceGroup, int, error) {
				item, err := schemas.DbGetFromId[schemas.CategorySingleChoiceGroup](ctx, db, params.GroupID)
				if err != nil {
					return schemas.CategorySingleChoiceGroup{}, http.StatusNotFound, err
				}
				if !principal.ReadAllowed(item.GetOwnerID(ctx, db)) {
					return schemas.CategorySingleChoiceGroup{}, http.StatusForbidden, fmt.Errorf("action not permitted")
				}
				return item, http.StatusOK, nil
			},
			func(ctx context.Context, db *gorm.DB, parent *schemas.CategorySingleChoiceGroup) ([]schemas.CategorySingleChoiceItem, error) {
				return schemas.DbGetChildSlice[schemas.CategorySingleChoiceGroup, schemas.CategorySingleChoiceItem](ctx, db, parent, "Choices", params.First, params.Limit, "id asc")
			},
			ToModelFunc[models.CategorySingleChoice, *schemas.CategorySingleChoiceItem],
			func(modelList []*models.CategorySingleChoice) middleware.Responder {
				return category.NewGetSingleChoiceGroupGroupIDSingleChoiceOK().WithPayload(modelList)
			},
		)
	})
	api.CategoryPostSingleChoiceGroupGroupIDSingleChoiceHandler = category.PostSingleChoiceGroupGroupIDSingleChoiceHandlerFunc(func(params category.PostSingleChoiceGroupGroupIDSingleChoiceParams, principal *schemas.User) middleware.Responder {
		return PostHandler[models.CategorySingleChoice, schemas.CategorySingleChoiceItem](
			params.HTTPRequest,
			db,
			params.Body,
			params.GroupID,
			principal,
			FromModelFunc[models.CategorySingleChoice, *schemas.CategorySingleChoiceItem],
			nil,
			SetParentIdFunc[models.CategorySingleChoice, *schemas.CategorySingleChoiceItem],
			ToModelFunc[models.CategorySingleChoice, *schemas.CategorySingleChoiceItem],
			func(model *models.CategorySingleChoice) middleware.Responder {
				return category.NewPostSingleChoiceGroupGroupIDSingleChoiceCreated().WithPayload(model)
			},
		)
	})
	api.CategoryGetSingleChoiceIDHandler = category.GetSingleChoiceIDHandlerFunc(func(params category.GetSingleChoiceIDParams, principal *schemas.User) middleware.Responder {
		return GetHandler[models.CategorySingleChoice, schemas.CategorySingleChoiceItem](
			params.HTTPRequest,
			db,
			params.ID,
			principal,
			ToModelFunc[models.CategorySingleChoice, *schemas.CategorySingleChoiceItem],
			func(model *models.CategorySingleChoice) middleware.Responder {
				return category.NewGetSingleChoiceIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryPutSingleChoiceIDHandler = category.PutSingleChoiceIDHandlerFunc(func(params category.PutSingleChoiceIDParams, principal *schemas.User) middleware.Responder {
		return PutHandler[models.CategorySingleChoice, schemas.CategorySingleChoiceItem](
			params.HTTPRequest,
			db,
			params.Body,
			params.ID,
			principal,
			FromModelFunc[models.CategorySingleChoice, *schemas.CategorySingleChoiceItem],
			nil,
			nil,
			ToModelFunc[models.CategorySingleChoice, *schemas.CategorySingleChoiceItem],
			func(model *models.CategorySingleChoice) middleware.Responder {
				return category.NewPutSingleChoiceIDOK().WithPayload(model)
			},
		)
	})
	api.CategoryDeleteSingleChoiceIDHandler = category.DeleteSingleChoiceIDHandlerFunc(func(params category.DeleteSingleChoiceIDParams, principal *schemas.User) middleware.Responder {
		return DeleteHandler[schemas.CategorySingleChoiceItem](params.HTTPRequest, db, params.ID, principal, func() middleware.Responder {
			return category.NewDeleteSingleChoiceIDNoContent()
		})
	})

	/////////////////////////////////////////////////////////////////

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.AllowAll().Handler(handler)

	return corsHandler
}

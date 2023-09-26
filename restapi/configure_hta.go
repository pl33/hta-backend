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
	"hta_backend_2/schemas"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"hta_backend_2/models"
	"hta_backend_2/restapi/operations"
	"hta_backend_2/restapi/operations/category"
	"hta_backend_2/restapi/operations/entry"
	"hta_backend_2/restapi/operations/login"
)

//go:generate swagger generate server --target ../../hta_backend_2 --name Hta --spec ../swagger.yaml --principal models.User

func errorResponder(code int, err error) middleware.Responder {
	msg := err.Error()
	responder := middleware.Error(code, models.Error{
		Code:    int32(code),
		Message: &msg,
	})
	return responder
}

func configureFlags(api *operations.HtaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HtaAPI) http.Handler {
	db, db_err := schemas.OpenDb(GetEnvOrPanic("DB"))
	if db_err != nil {
		panic(db_err)
	}

	auth, auth_err := AuthSetup(
		db,
		GetEnvOrPanic("OIDC_ISSUER"),
		GetEnvOrPanic("OIDC_CLIENT_ID"),
		GetEnvOrPanic("OIDC_CLIENT_SECRET"),
	)
	if auth_err != nil {
		panic(auth_err)
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

	api.BearerTokenAuth = func(token string) (*models.User, error) {
		user, err := AuthGetUser(context.Background(), &auth, token)
		principal := models.User{
			ID:        int32(user.ID),
			Name:      user.Name,
			FirstName: user.FirstName,
		}
		return &principal, err
	}

	api.OauthSecurityAuth = func(token string, scopes []string) (*models.User, error) {
		user, err := AuthGetUser(context.Background(), &auth, token)
		principal := models.User{
			ID:        int32(user.ID),
			Name:      user.Name,
			FirstName: user.FirstName,
		}
		return &principal, err
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.CategoryDeleteCategoryIDHandler == nil {
		api.CategoryDeleteCategoryIDHandler = category.DeleteCategoryIDHandlerFunc(func(params category.DeleteCategoryIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.DeleteCategoryID has not yet been implemented")
		})
	}
	api.EntryDeleteEntriesIDHandler = entry.DeleteEntriesIDHandlerFunc(func(params entry.DeleteEntriesIDParams, principal *models.User) middleware.Responder {
		var err error
		ctx := params.HTTPRequest.Context()

		var ent schemas.HealthEntry
		ent, err = schemas.DbGetFromId[schemas.HealthEntry](ctx, db, params.ID)
		if err != nil {
			return errorResponder(http.StatusInternalServerError, err)
		}

		err = db.WithContext(ctx).Delete(&ent).Error
		if err != nil {
			return errorResponder(http.StatusInternalServerError, err)
		}

		return entry.NewDeleteEntriesIDNoContent()
	})
	if api.CategoryDeleteMultiChoiceIDHandler == nil {
		api.CategoryDeleteMultiChoiceIDHandler = category.DeleteMultiChoiceIDHandlerFunc(func(params category.DeleteMultiChoiceIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.DeleteMultiChoiceID has not yet been implemented")
		})
	}
	if api.CategoryDeleteSingleChoiceGroupIDHandler == nil {
		api.CategoryDeleteSingleChoiceGroupIDHandler = category.DeleteSingleChoiceGroupIDHandlerFunc(func(params category.DeleteSingleChoiceGroupIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.DeleteSingleChoiceGroupID has not yet been implemented")
		})
	}
	if api.CategoryDeleteSingleChoiceIDHandler == nil {
		api.CategoryDeleteSingleChoiceIDHandler = category.DeleteSingleChoiceIDHandlerFunc(func(params category.DeleteSingleChoiceIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.DeleteSingleChoiceID has not yet been implemented")
		})
	}
	if api.CategoryGetCategoryHandler == nil {
		api.CategoryGetCategoryHandler = category.GetCategoryHandlerFunc(func(params category.GetCategoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.GetCategory has not yet been implemented")
		})
	}
	if api.CategoryGetCategoryCategoryIDMultiChoiceHandler == nil {
		api.CategoryGetCategoryCategoryIDMultiChoiceHandler = category.GetCategoryCategoryIDMultiChoiceHandlerFunc(func(params category.GetCategoryCategoryIDMultiChoiceParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.GetCategoryCategoryIDMultiChoice has not yet been implemented")
		})
	}
	if api.CategoryGetCategoryCategoryIDSingleChoiceGroupHandler == nil {
		api.CategoryGetCategoryCategoryIDSingleChoiceGroupHandler = category.GetCategoryCategoryIDSingleChoiceGroupHandlerFunc(func(params category.GetCategoryCategoryIDSingleChoiceGroupParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.GetCategoryCategoryIDSingleChoiceGroup has not yet been implemented")
		})
	}
	api.EntryGetEntriesHandler = entry.GetEntriesHandlerFunc(func(params entry.GetEntriesParams, principal *models.User) middleware.Responder {
		ctx := params.HTTPRequest.Context()

		user, userErr := schemas.LookupUser(ctx, db, principal.ID)
		if userErr != nil {
			return errorResponder(http.StatusInternalServerError, userErr)
		}

		entries, listErr := user.ListHealthEntries(ctx, db, params.First, params.Limit)
		if listErr != nil {
			return errorResponder(http.StatusInternalServerError, listErr)
		}

		entriesArray := make([]*models.Entry, len(entries))
		for idx := range entries {
			item, err := entries[idx].ToModel(ctx, db)
			if err != nil {
				return errorResponder(http.StatusInternalServerError, err)
			}
			entriesArray[idx] = &item
		}

		return entry.NewGetEntriesOK().WithPayload(entriesArray)
	})
	api.LoginGetLoginHandler = login.GetLoginHandlerFunc(func(params login.GetLoginParams) middleware.Responder {
		return AuthLogin(&auth, params.HTTPRequest)
	})
	api.LoginGetOidcCallbackHandler = login.GetOidcCallbackHandlerFunc(func(params login.GetOidcCallbackParams) middleware.Responder {
		return AuthCallback(&auth, params.HTTPRequest)
	})
	if api.CategoryGetSingleChoiceGroupGroupIDSingleChoiceHandler == nil {
		api.CategoryGetSingleChoiceGroupGroupIDSingleChoiceHandler = category.GetSingleChoiceGroupGroupIDSingleChoiceHandlerFunc(func(params category.GetSingleChoiceGroupGroupIDSingleChoiceParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.GetSingleChoiceGroupGroupIDSingleChoice has not yet been implemented")
		})
	}
	api.LoginGetUserHandler = login.GetUserHandlerFunc(func(params login.GetUserParams, principal *models.User) middleware.Responder {
		return login.NewGetUserOK().WithPayload(principal)
	})
	if api.CategoryPostCategoryHandler == nil {
		api.CategoryPostCategoryHandler = category.PostCategoryHandlerFunc(func(params category.PostCategoryParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PostCategory has not yet been implemented")
		})
	}
	if api.CategoryPostCategoryCategoryIDMultiChoiceHandler == nil {
		api.CategoryPostCategoryCategoryIDMultiChoiceHandler = category.PostCategoryCategoryIDMultiChoiceHandlerFunc(func(params category.PostCategoryCategoryIDMultiChoiceParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PostCategoryCategoryIDMultiChoice has not yet been implemented")
		})
	}
	if api.CategoryPostCategoryCategoryIDSingleChoiceGroupHandler == nil {
		api.CategoryPostCategoryCategoryIDSingleChoiceGroupHandler = category.PostCategoryCategoryIDSingleChoiceGroupHandlerFunc(func(params category.PostCategoryCategoryIDSingleChoiceGroupParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PostCategoryCategoryIDSingleChoiceGroup has not yet been implemented")
		})
	}
	if api.CategoryPostSingleChoiceGroupGroupIDSingleChoiceHandler == nil {
		api.CategoryPostSingleChoiceGroupGroupIDSingleChoiceHandler = category.PostSingleChoiceGroupGroupIDSingleChoiceHandlerFunc(func(params category.PostSingleChoiceGroupGroupIDSingleChoiceParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PostSingleChoiceGroupGroupIDSingleChoice has not yet been implemented")
		})
	}
	api.EntryPostEntriesHandler = entry.PostEntriesHandlerFunc(func(params entry.PostEntriesParams, principal *models.User) middleware.Responder {
		ctx := params.HTTPRequest.Context()

		if params.Body == nil {
			return errorResponder(http.StatusBadRequest, fmt.Errorf("body is missing"))
		}

		var ent schemas.HealthEntry
		if err := ent.FromModel(ctx, db, *params.Body); err != nil {
			return errorResponder(http.StatusBadRequest, err)
		}
		ent.SetParentId(uint(principal.ID))

		if err := db.WithContext(ctx).Create(&ent).Error; err != nil {
			return errorResponder(http.StatusInternalServerError, err)
		}

		model, _ := ent.ToModel(ctx, db)
		return entry.NewPostEntriesCreated().WithPayload(&model)
	})
	if api.CategoryPutCategoryIDHandler == nil {
		api.CategoryPutCategoryIDHandler = category.PutCategoryIDHandlerFunc(func(params category.PutCategoryIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PutCategoryID has not yet been implemented")
		})
	}
	api.EntryPutEntriesIDHandler = entry.PutEntriesIDHandlerFunc(func(params entry.PutEntriesIDParams, principal *models.User) middleware.Responder {
		var err error
		ctx := params.HTTPRequest.Context()

		if params.Body == nil {
			return errorResponder(http.StatusBadRequest, fmt.Errorf("body is missing"))
		}

		var ent schemas.HealthEntry
		ent, err = schemas.DbGetFromId[schemas.HealthEntry](ctx, db, params.ID)
		if err != nil {
			return errorResponder(http.StatusInternalServerError, err)
		}

		err = ent.FromModel(ctx, db, *params.Body)
		if err != nil {
			return errorResponder(http.StatusBadRequest, err)
		}

		err = db.WithContext(ctx).Save(&ent).Error
		if err != nil {
			return errorResponder(http.StatusInternalServerError, err)
		}

		model, _ := ent.ToModel(ctx, db)
		return entry.NewPutEntriesIDOK().WithPayload(&model)
	})
	if api.CategoryPutMultiChoiceIDHandler == nil {
		api.CategoryPutMultiChoiceIDHandler = category.PutMultiChoiceIDHandlerFunc(func(params category.PutMultiChoiceIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PutMultiChoiceID has not yet been implemented")
		})
	}
	if api.CategoryPutSingleChoiceGroupIDHandler == nil {
		api.CategoryPutSingleChoiceGroupIDHandler = category.PutSingleChoiceGroupIDHandlerFunc(func(params category.PutSingleChoiceGroupIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PutSingleChoiceGroupID has not yet been implemented")
		})
	}
	if api.CategoryPutSingleChoiceIDHandler == nil {
		api.CategoryPutSingleChoiceIDHandler = category.PutSingleChoiceIDHandlerFunc(func(params category.PutSingleChoiceIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation category.PutSingleChoiceID has not yet been implemented")
		})
	}

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
	return handler
}

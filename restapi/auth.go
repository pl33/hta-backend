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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/pkg/client/rp"
	"github.com/zitadel/oidc/pkg/client/rs"
	"github.com/zitadel/oidc/pkg/oidc"
	"gorm.io/gorm"
	"hta_backend_2/schemas"
	"net/http"
	"strings"
	"time"
)

type Auth struct {
	issuer           string
	clientId         string
	frontendClientId string
	oidcRS           rs.ResourceServer
	oidcRP           rp.RelyingParty
	db               *gorm.DB
}

func AuthSetup(db *gorm.DB, issuer string, clientId string, clientSecret string, frontendClientId string) (Auth, error) {
	var auth Auth

	auth.db = db
	auth.issuer = issuer
	auth.clientId = clientId
	auth.frontendClientId = frontendClientId

	var err error
	auth.oidcRS, err = rs.NewResourceServerClientCredentials(
		issuer,
		clientId,
		clientSecret,
	)
	if err != nil {
		return auth, err
	}

	auth.oidcRP, err = rp.NewRelyingPartyOIDC(
		issuer,
		clientId,
		clientSecret,
		"http://localhost:8080/oidc_callback",
		strings.Split("openid", " "),
	)
	if err != nil {
		return auth, err
	}

	return auth, nil
}

func AuthLogin(auth *Auth, r *http.Request) middleware.Responder {
	fn := rp.AuthURLHandler(func() string {
		return uuid.New().String()
	}, auth.oidcRP, rp.WithPromptURLParam("Welcome to HTA!"))
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			fn(w, r)
		})
}

func AuthCallback(auth *Auth, r *http.Request) middleware.Responder {
	fn := rp.CodeExchangeHandler(func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string, rp rp.RelyingParty) {
		fmt.Printf("Received ID Token %s", tokens.IDToken)
		fmt.Printf("Received Access Token %s", tokens.AccessToken)
		data, err := json.Marshal(tokens)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			w.Write(data)
		}
	}, auth.oidcRP)
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			fn(w, r)
		})
}

func getIntrospect(ctx context.Context, auth *Auth, token string) (oidc.IntrospectionResponse, error) {
	rctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	introspect, err := rs.Introspect(rctx, auth.oidcRS, token)
	if err != nil {
		msg := fmt.Sprintf("introspection error: %s", err.Error())
		return nil, errors.New(msg)
	}

	if !introspect.IsActive() {
		return nil, errors.New("token expired")
	}

	return introspect, nil
}

func getUser(ctx context.Context, auth *Auth, introspect oidc.IntrospectionResponse) (schemas.User, error) {
	rctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var user schemas.User
	auth.db.WithContext(rctx).FirstOrCreate(&user, schemas.User{
		OidcIssuer: introspect.GetIssuer(),
		OidcSub:    introspect.GetSubject(),
	})

	user.Name = introspect.GetName()
	user.FirstName = introspect.GetGivenName()
	auth.db.WithContext(rctx).Save(&user)

	return user, auth.db.Error
}

func AuthGetUser(ctx context.Context, auth *Auth, token string) (schemas.User, error) {
	introspect, err := getIntrospect(ctx, auth, token)
	if err != nil {
		return schemas.User{}, err
	}

	user, err := getUser(ctx, auth, introspect)
	if err != nil {
		return schemas.User{}, err
	}

	return user, nil
}

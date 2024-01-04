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
	"errors"
	"fmt"
	"github.com/pl33/hta-backend/schemas"
	"github.com/zitadel/oidc/pkg/client/rs"
	"github.com/zitadel/oidc/pkg/oidc"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	issuer           string
	clientId         string
	frontendClientId string
	oidcRS           rs.ResourceServer
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

	return auth, nil
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

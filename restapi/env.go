/*
 * SPDX-License-Identifier: MPL-2.0
 *   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package restapi

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	env, found := os.LookupEnv(key)
	if !found {
		return "", fmt.Errorf("%s environment variable is missing", key)
	}
	return env, nil
}

func GetEnvOrPanic(key string) string {
	env, err := GetEnv(key)
	if err != nil {
		panic(err)
	}
	return env
}

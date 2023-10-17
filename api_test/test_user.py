# SPDX-License-Identifier: MPL-2.0
#   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

import pydantic
import http

from utils import server, http_get


class User(pydantic.BaseModel):
    id: int
    first_name: str
    name: str


def test_get_user(server):
    code, obj = http_get(server, "/user", 1)
    assert code == http.HTTPStatus.OK
    user = User.model_validate(obj)
    assert user == User(id=1, first_name="Alice", name="Alice Wonderland")

    code, obj = http_get(server, "/user", 2)
    assert code == http.HTTPStatus.OK
    user = User.model_validate(obj)
    assert user == User(id=2, first_name="John", name="John Doe")

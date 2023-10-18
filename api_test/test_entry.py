# SPDX-License-Identifier: MPL-2.0
#   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
import datetime

import pydantic
import http
from typing import List, Optional

from utils import server, http_get, http_post, http_put, http_delete, Server
from category_common import (
    create_categories,
    create_multi_choice_items,
    create_single_choice_groups,
    create_single_choice_items,
)


class Entry(pydantic.BaseModel):
    have_blood_pressure: bool
    start_time: datetime.datetime
    multi_choices: List[int]
    single_choices: List[int]
    end_time: Optional[datetime.datetime] = None
    diastole: Optional[float] = None
    systole: Optional[float] = None
    pulse: Optional[float] = None
    remarks: Optional[str] = None
    id: Optional[int] = None
    user_id: Optional[int] = None


class Entries(pydantic.RootModel[List[Entry]]):
    pass


ENTRY_1 = Entry(
    have_blood_pressure=True,
    start_time="2023-10-18T23:34:12Z",
    multi_choices=[2],
    single_choices=[1],
    diastole=120,
    systole=80,
    pulse=60,
    end_time="0001-01-01T00:00:00.000Z",
    remarks=None,
    id=1,
    user_id=1,
)
ENTRY_2 = Entry(
    have_blood_pressure=True,
    start_time="2023-10-17T21:34:12Z",
    multi_choices=[],
    single_choices=[],
    diastole=130,
    systole=70,
    pulse=50,
    end_time="2023-10-17T23:45:55Z",
    remarks="Some",
    id=2,
    user_id=1,
)
ENTRY_2_CHANGED = Entry(
    have_blood_pressure=True,
    start_time="2023-10-17T21:34:12Z",
    multi_choices=[1],
    single_choices=[2],
    diastole=115,
    systole=75,
    pulse=55,
    end_time="2023-10-17T23:45:55Z",
    remarks="Other",
    id=2,
    user_id=1,
)
ENTRY_3 = Entry(
    have_blood_pressure=False,
    start_time="2023-10-16T23:34:12Z",
    multi_choices=[],
    single_choices=[],
    diastole=None,
    systole=None,
    pulse=None,
    end_time="0001-01-01T00:00:00.000Z",
    remarks=None,
    id=3,
    user_id=2,
)


def create_entries(server: Server):
    code, obj = http_post(
        server,
        "/entries",
        Entry(
            have_blood_pressure=True,
            start_time="2023-10-18T23:34:12Z",
            multi_choices=[2],
            single_choices=[1],
            diastole=120,
            systole=80,
            pulse=60,
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    category = Entry.model_validate(obj)
    assert category == ENTRY_1
    code, obj = http_post(
        server,
        "/entries",
        Entry(
            have_blood_pressure=True,
            start_time="2023-10-17T21:34:12Z",
            multi_choices=[],
            single_choices=[],
            diastole=130,
            systole=70,
            pulse=50,
            end_time="2023-10-17T23:45:55Z",
            remarks="Some",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    category = Entry.model_validate(obj)
    assert category == ENTRY_2

    code, obj = http_post(
        server,
        "/entries",
        Entry(
            have_blood_pressure=False,
            start_time="2023-10-16T23:34:12Z",
            multi_choices=[],
            single_choices=[],
        ),
        2,
    )
    assert code == http.HTTPStatus.CREATED
    category = Entry.model_validate(obj)
    assert category == ENTRY_3


def test_entry(server):
    create_categories(server)
    create_multi_choice_items(server)
    create_single_choice_groups(server)
    create_single_choice_items(server)

    # Check empty
    code, obj = http_get(server, "/entries", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    # Create
    create_entries(server)

    # Verify creation
    code, obj = http_get(server, "/entries", 2)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 1
    categories = Entries.model_validate(obj)
    assert ENTRY_3 in categories.root

    code, obj = http_get(server, "/entries", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 2
    categories = Entries.model_validate(obj)
    assert ENTRY_1 in categories.root
    assert ENTRY_2 in categories.root

    # Get
    code, obj = http_get(server, "/entries/1", 1)
    assert code == http.HTTPStatus.OK
    category = Entry.model_validate(obj)
    assert category == ENTRY_1

    code, obj = http_get(server, "/entries/2", 1)
    assert code == http.HTTPStatus.OK
    category = Entry.model_validate(obj)
    assert category == ENTRY_2

    # Update
    code, obj = http_put(
        server,
        "/entries/2",
        Entry(
            have_blood_pressure=True,
            start_time="2023-10-17T21:34:12Z",
            multi_choices=[1],
            single_choices=[2],
            diastole=115,
            systole=75,
            pulse=55,
            end_time="2023-10-17T23:45:55Z",
            remarks="Other",
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.OK
    category = Entry.model_validate(obj)
    assert category == ENTRY_2_CHANGED

    # Not found
    code, _ = http_get(server, "/entries/4", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_put(
        server,
        "/entries/4",
        Entry(
            have_blood_pressure=True,
            start_time="2023-10-17T21:34:12Z",
            multi_choices=[1],
            single_choices=[2],
            diastole=115,
            systole=75,
            pulse=55,
            end_time="2023-10-17T23:45:55Z",
            remarks="Other",
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.NOT_FOUND

    code = http_delete(server, "/entries/4", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    # Forbidden
    code, _ = http_get(server, "/entries/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/entries/2",
        Entry(
            have_blood_pressure=True,
            start_time="2023-10-17T21:34:12Z",
            multi_choices=[],
            single_choices=[],
            diastole=115,
            systole=75,
            pulse=55,
            end_time="2023-10-17T23:45:55Z",
            remarks="Other",
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        2,
    )

    code, _ = http_put(
        server,
        "/entries/3",
        Entry(
            have_blood_pressure=False,
            start_time="2023-10-16T23:34:12Z",
            multi_choices=[1],
            single_choices=[],
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/entries/3",
        Entry(
            have_blood_pressure=False,
            start_time="2023-10-16T23:34:12Z",
            multi_choices=[],
            single_choices=[1],
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code = http_delete(server, "/entries/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    # Delete
    code = http_delete(server, "/entries/1", 1)
    assert code == http.HTTPStatus.NO_CONTENT
    code, _ = http_get(server, "/entries/1", 1)
    assert code == http.HTTPStatus.NOT_FOUND
    code, obj = http_get(server, "/entries", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 1
    categories = Entries.model_validate(obj)
    assert ENTRY_2_CHANGED in categories.root

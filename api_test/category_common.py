# SPDX-License-Identifier: MPL-2.0
#   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

import pydantic
import http
from typing import List, Optional

from utils import http_post, Server


class Category(pydantic.BaseModel):
    title: str
    id: Optional[int] = None
    user_id: Optional[int] = None


class Categories(pydantic.RootModel[List[Category]]):
    pass


class CategoryMultiChoice(pydantic.BaseModel):
    title: str
    id: Optional[int] = None
    category_id: Optional[int] = None
    description: Optional[str] = None


class CategoryMultiChoices(pydantic.RootModel[List[CategoryMultiChoice]]):
    pass


class CategorySingleChoiceGroup(pydantic.BaseModel):
    title: str
    id: Optional[int] = None
    category_id: Optional[int] = None
    description: Optional[str] = None


class CategorySingleChoiceGroups(pydantic.RootModel[List[CategorySingleChoiceGroup]]):
    pass


class CategorySingleChoiceItem(pydantic.BaseModel):
    title: str
    id: Optional[int] = None
    group_id: Optional[int] = None
    description: Optional[str] = None


class CategorySingleChoiceItems(pydantic.RootModel[List[CategorySingleChoiceItem]]):
    pass


def create_categories(server: Server):
    code, obj = http_post(
        server,
        "/category",
        Category(
            title="Test 1",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    category = Category.model_validate(obj)
    assert category == Category(id=1, user_id=1, title="Test 1")
    code, obj = http_post(
        server,
        "/category",
        Category(
            title="Test 2",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    category = Category.model_validate(obj)
    assert category == Category(id=2, user_id=1, title="Test 2")

    code, obj = http_post(
        server,
        "/category",
        Category(
            title="User 2 Test 1",
        ),
        2,
    )
    assert code == http.HTTPStatus.CREATED
    category = Category.model_validate(obj)
    assert category == Category(id=3, user_id=2, title="User 2 Test 1")

def create_multi_choice_items(server: Server):
    code, obj = http_post(
        server,
        "/category/1/multi_choice",
        CategoryMultiChoice(
            title="Test 1",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    item = CategoryMultiChoice.model_validate(obj)
    assert item == CategoryMultiChoice(id=1, category_id=1, title="Test 1", description=None)
    code, obj = http_post(
        server,
        "/category/1/multi_choice",
        CategoryMultiChoice(
            title="Test 2",
            description="Some",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    item = CategoryMultiChoice.model_validate(obj)
    assert item == CategoryMultiChoice(id=2, category_id=1, title="Test 2", description="Some")


def create_single_choice_groups(server: Server):
    code, obj = http_post(
        server,
        "/category/1/single_choice_group",
        CategorySingleChoiceGroup(
            title="Test 1",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    item = CategorySingleChoiceGroup.model_validate(obj)
    assert item == CategorySingleChoiceGroup(id=1, category_id=1, title="Test 1", description=None)
    code, obj = http_post(
        server,
        "/category/1/single_choice_group",
        CategorySingleChoiceGroup(
            title="Test 2",
            description="Some",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    item = CategorySingleChoiceGroup.model_validate(obj)
    assert item == CategorySingleChoiceGroup(id=2, category_id=1, title="Test 2", description="Some")


def create_single_choice_items(server: Server):
    code, obj = http_post(
        server,
        "/single_choice_group/1/single_choice",
        CategorySingleChoiceItem(
            title="Test 1",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    item = CategorySingleChoiceItem.model_validate(obj)
    assert item == CategorySingleChoiceItem(id=1, group_id=1, title="Test 1", description=None)
    code, obj = http_post(
        server,
        "/single_choice_group/1/single_choice",
        CategorySingleChoiceItem(
            title="Test 2",
            description="Some",
        ),
        1,
    )
    assert code == http.HTTPStatus.CREATED
    item = CategorySingleChoiceItem.model_validate(obj)
    assert item == CategorySingleChoiceItem(id=2, group_id=1, title="Test 2", description="Some")

# SPDX-License-Identifier: MPL-2.0
#   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

import http

from utils import server, http_get, http_post, http_put, http_delete
from category_common import (
    create_categories,
    create_multi_choice_items,
    create_single_choice_groups,
    create_single_choice_items,
    Category,
    Categories,
    CategoryMultiChoice,
    CategoryMultiChoices,
    CategorySingleChoiceGroup,
    CategorySingleChoiceGroups,
    CategorySingleChoiceItem,
    CategorySingleChoiceItems,
)


def test_category(server):
    # Check empty
    code, obj = http_get(server, "/category", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    # Create
    create_categories(server)

    # Verify creation
    code, obj = http_get(server, "/category", 2)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 1
    categories = Categories.model_validate(obj)
    assert Category(id=3, user_id=2, title="User 2 Test 1") in categories.root

    code, obj = http_get(server, "/category", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 2
    categories = Categories.model_validate(obj)
    assert Category(id=1, user_id=1, title="Test 1") in categories.root
    assert Category(id=2, user_id=1, title="Test 2") in categories.root

    # Get
    code, obj = http_get(server, "/category/1", 1)
    assert code == http.HTTPStatus.OK
    category = Category.model_validate(obj)
    assert category == Category(id=1, user_id=1, title="Test 1")

    code, obj = http_get(server, "/category/2", 1)
    assert code == http.HTTPStatus.OK
    category = Category.model_validate(obj)
    assert category == Category(id=2, user_id=1, title="Test 2")

    # Update
    code, obj = http_put(
        server,
        "/category/2",
        Category(
            title="Test 3",
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.OK
    category = Category.model_validate(obj)
    assert category == Category(id=2, user_id=1, title="Test 3")

    # Not found
    code, _ = http_get(server, "/category/4", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_put(
        server,
        "/category/4",
        Category(
            title="Test 3",
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.NOT_FOUND

    code = http_delete(server, "/category/4", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    # Forbidden
    code, _ = http_get(server, "/category/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/category/2",
        Category(
            title="Test 3",
            id=0,       # Shall be ignored
            user_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code = http_delete(server, "/category/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    # Delete
    code = http_delete(server, "/category/1", 1)
    assert code == http.HTTPStatus.NO_CONTENT
    code, _ = http_get(server, "/category/1", 1)
    assert code == http.HTTPStatus.NOT_FOUND
    code, obj = http_get(server, "/category", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 1
    categories = Categories.model_validate(obj)
    assert Category(id=2, user_id=1, title="Test 3") in categories.root


def test_multi_choice(server):
    create_categories(server)

    # Check empty
    code, obj = http_get(server, "/category/1/multi_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    # Create
    create_multi_choice_items(server)

    # Verify creation
    code, obj = http_get(server, "/category/2/multi_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    code, obj = http_get(server, "/category/1/multi_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 2
    items = CategoryMultiChoices.model_validate(obj)
    assert CategoryMultiChoice(id=1, category_id=1, title="Test 1", description=None) in items.root
    assert CategoryMultiChoice(id=2, category_id=1, title="Test 2", description="Some") in items.root

    # Get
    code, obj = http_get(server, "/multi_choice/1", 1)
    assert code == http.HTTPStatus.OK
    item = CategoryMultiChoice.model_validate(obj)
    assert item == CategoryMultiChoice(id=1, category_id=1, title="Test 1", description=None)

    code, obj = http_get(server, "/multi_choice/2", 1)
    assert code == http.HTTPStatus.OK
    item = CategoryMultiChoice.model_validate(obj)
    assert item == CategoryMultiChoice(id=2, category_id=1, title="Test 2", description="Some")

    # Update
    code, obj = http_put(
        server,
        "/multi_choice/2",
        CategoryMultiChoice(
            title="Test 3",
            description="Other",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.OK
    item = CategoryMultiChoice.model_validate(obj)
    assert item == CategoryMultiChoice(id=2, category_id=1, title="Test 3", description="Other")

    # Not found
    code, _ = http_get(server, "/category/4/multi_choice", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_get(server, "/multi_choice/3", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_post(
        server,
        "/category/4/multi_choice",
        CategoryMultiChoice(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/multi_choice/3",
        CategoryMultiChoice(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.NOT_FOUND

    code = http_delete(server, "/multi_choice/3", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    # Forbidden
    code, _ = http_get(server, "/category/1/multi_choice", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_get(server, "/multi_choice/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_post(
        server,
        "/category/1/multi_choice",
        CategoryMultiChoice(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/multi_choice/2",
        CategoryMultiChoice(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code = http_delete(server, "/multi_choice/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    # Delete
    code = http_delete(server, "/multi_choice/1", 1)
    assert code == http.HTTPStatus.NO_CONTENT
    code, _ = http_get(server, "/multi_choice/1", 1)
    assert code == http.HTTPStatus.NOT_FOUND
    code, obj = http_get(server, "/category/1/multi_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 1
    items = CategoryMultiChoices.model_validate(obj)
    assert CategoryMultiChoice(id=2, category_id=1, title="Test 3", description="Other") in items.root


def test_single_choice_groups(server):
    create_categories(server)

    # Check empty
    code, obj = http_get(server, "/category/1/single_choice_group", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    # Create
    create_single_choice_groups(server)

    # Verify creation
    code, obj = http_get(server, "/category/2/single_choice_group", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    code, obj = http_get(server, "/category/1/single_choice_group", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 2
    items = CategorySingleChoiceGroups.model_validate(obj)
    assert CategorySingleChoiceGroup(id=1, category_id=1, title="Test 1", description=None) in items.root
    assert CategorySingleChoiceGroup(id=2, category_id=1, title="Test 2", description="Some") in items.root

    # Get
    code, obj = http_get(server, "/single_choice_group/1", 1)
    assert code == http.HTTPStatus.OK
    item = CategorySingleChoiceGroup.model_validate(obj)
    assert item == CategorySingleChoiceGroup(id=1, category_id=1, title="Test 1", description=None)

    code, obj = http_get(server, "/single_choice_group/2", 1)
    assert code == http.HTTPStatus.OK
    item = CategorySingleChoiceGroup.model_validate(obj)
    assert item == CategorySingleChoiceGroup(id=2, category_id=1, title="Test 2", description="Some")

    # Update
    code, obj = http_put(
        server,
        "/single_choice_group/2",
        CategorySingleChoiceGroup(
            title="Test 3",
            description="Other",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.OK
    item = CategorySingleChoiceGroup.model_validate(obj)
    assert item == CategorySingleChoiceGroup(id=2, category_id=1, title="Test 3", description="Other")

    # Not found
    code, _ = http_get(server, "/category/4/single_choice_group", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_get(server, "/single_choice_group/3", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_post(
        server,
        "/category/4/single_choice_group",
        CategorySingleChoiceGroup(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/single_choice_group/3",
        CategorySingleChoiceGroup(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.NOT_FOUND

    code = http_delete(server, "/single_choice_group/3", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    # Forbidden
    code, _ = http_get(server, "/category/1/single_choice_group", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_get(server, "/single_choice_group/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_post(
        server,
        "/category/1/single_choice_group",
        CategorySingleChoiceGroup(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/single_choice_group/2",
        CategorySingleChoiceGroup(
            title="Test 3",
            id=0,       # Shall be ignored
            category_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code = http_delete(server, "/single_choice_group/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    # Delete
    code = http_delete(server, "/single_choice_group/1", 1)
    assert code == http.HTTPStatus.NO_CONTENT
    code, _ = http_get(server, "/single_choice_group/1", 1)
    assert code == http.HTTPStatus.NOT_FOUND
    code, obj = http_get(server, "/category/1/single_choice_group", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 1
    items = CategorySingleChoiceGroups.model_validate(obj)
    assert CategorySingleChoiceGroup(id=2, category_id=1, title="Test 3", description="Other") in items.root


def test_single_choice_items(server):
    create_categories(server)
    create_single_choice_groups(server)

    # Check empty
    code, obj = http_get(server, "/single_choice_group/1/single_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    # Create
    create_single_choice_items(server)

    # Verify creation
    code, obj = http_get(server, "/single_choice_group/2/single_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 0

    code, obj = http_get(server, "/single_choice_group/1/single_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 2
    items = CategorySingleChoiceItems.model_validate(obj)
    assert CategorySingleChoiceItem(id=1, group_id=1, title="Test 1", description=None) in items.root
    assert CategorySingleChoiceItem(id=2, group_id=1, title="Test 2", description="Some") in items.root

    # Get
    code, obj = http_get(server, "/single_choice/1", 1)
    assert code == http.HTTPStatus.OK
    item = CategorySingleChoiceItem.model_validate(obj)
    assert item == CategorySingleChoiceItem(id=1, group_id=1, title="Test 1", description=None)

    code, obj = http_get(server, "/single_choice/2", 1)
    assert code == http.HTTPStatus.OK
    item = CategorySingleChoiceItem.model_validate(obj)
    assert item == CategorySingleChoiceItem(id=2, group_id=1, title="Test 2", description="Some")

    # Update
    code, obj = http_put(
        server,
        "/single_choice/2",
        CategorySingleChoiceItem(
            title="Test 3",
            description="Other",
            id=0,       # Shall be ignored
            group_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.OK
    item = CategorySingleChoiceItem.model_validate(obj)
    assert item == CategorySingleChoiceItem(id=2, group_id=1, title="Test 3", description="Other")

    # Not found
    code, _ = http_get(server, "/single_choice_group/4/single_choice", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_get(server, "/single_choice/3", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    code, _ = http_post(
        server,
        "/single_choice_group/4/single_choice",
        CategorySingleChoiceItem(
            title="Test 3",
            id=0,       # Shall be ignored
            group_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/single_choice/3",
        CategorySingleChoiceItem(
            title="Test 3",
            id=0,       # Shall be ignored
            group_id=0,  # Shall be ignored
        ),
        1,
    )
    assert code == http.HTTPStatus.NOT_FOUND

    code = http_delete(server, "/single_choice/3", 1)
    assert code == http.HTTPStatus.NOT_FOUND

    # Forbidden
    code, _ = http_get(server, "/single_choice_group/1/single_choice", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_get(server, "/single_choice/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_post(
        server,
        "/single_choice_group/1/single_choice",
        CategorySingleChoiceItem(
            title="Test 3",
            id=0,       # Shall be ignored
            group_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code, _ = http_put(
        server,
        "/single_choice/2",
        CategorySingleChoiceItem(
            title="Test 3",
            id=0,       # Shall be ignored
            group_id=0,  # Shall be ignored
        ),
        2,
    )
    assert code == http.HTTPStatus.FORBIDDEN

    code = http_delete(server, "/single_choice/2", 2)
    assert code == http.HTTPStatus.FORBIDDEN

    # Delete
    code = http_delete(server, "/single_choice/1", 1)
    assert code == http.HTTPStatus.NO_CONTENT
    code, _ = http_get(server, "/single_choice/1", 1)
    assert code == http.HTTPStatus.NOT_FOUND
    code, obj = http_get(server, "/single_choice_group/1/single_choice", 1)
    assert code == http.HTTPStatus.OK
    assert len(obj) == 1
    items = CategorySingleChoiceItems.model_validate(obj)
    assert CategorySingleChoiceItem(id=2, group_id=1, title="Test 3", description="Other") in items.root

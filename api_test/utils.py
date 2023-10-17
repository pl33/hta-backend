# SPDX-License-Identifier: MPL-2.0
#   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

import copy
import sqlite3
import subprocess
import select
import signal

from typing import Any, NamedTuple, Optional, Tuple
import json
import pytest
import os
from http.client import HTTPConnection, HTTPResponse
import pydantic

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))

DB_FILE = f"{SCRIPT_DIR}/test.db"

SUBPROCESS = True


class Server(NamedTuple):
    host: str
    db: sqlite3.Connection


def terminate(proc: subprocess.Popen):
    pgid = os.getpgid(proc.pid)
    os.killpg(pgid, signal.SIGTERM)
    code = proc.wait()
    print("Exit code: ", code)


@pytest.fixture
def server():
    if SUBPROCESS:
        env = copy.deepcopy(os.environ)
        env.update({
            "DEBUG_ENABLE": "cfe58f39-9d21-48ad-b0f8-84563532bc24",
            "DB": f"sqlite3://{DB_FILE}",
            "GOPATH": "/home/ple/var/go",
            "CGO_ENABLED": "1",
        })
        proc = subprocess.Popen(
            [
                "go",
                "run",
                "cmd/hta-server/main.go",
                "--scheme=http",
                "--port=8080",
            ],
            cwd=f"{SCRIPT_DIR}/../",
            env=env,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            preexec_fn=os.setsid,
        )

        out_fno = proc.stdout.fileno()
        err_fno = proc.stderr.fileno()
        for step in range(50):
            rlist, _, _ = select.select([err_fno], [], [], 0.1)
            if err_fno in rlist:
                line = str(proc.stderr.readline(), "utf-8")
                if "Serving hta at" in line:
                    break
                else:
                    print("Ignore STDERR: ", line)
        else:
            terminate(proc)
            raise TimeoutError

    db = sqlite3.connect(DB_FILE)
    db.execute("INSERT INTO users (id, created_at, name, first_name) VALUES (1, CURRENT_TIMESTAMP, 'Alice Wonderland', 'Alice')")
    db.execute("INSERT INTO users (id, created_at, name, first_name) VALUES (2, CURRENT_TIMESTAMP, 'John Doe', 'John')")
    db.commit()

    yield Server("localhost:8080", db)

    db.close()
    os.remove(DB_FILE)

    if SUBPROCESS:
        terminate(proc)

        polling = [out_fno, err_fno]
        while len(polling) > 0:
            rlist, _, _ = select.select(polling, [], [], 0.1)
            if len(rlist) == 0:
                polling = []
            if err_fno in rlist:
                line = proc.stderr.readline()
                if len(line) == 0:
                    polling.remove(err_fno)
                else:
                    print("STDERR: ", str(line, "utf-8"))
            if out_fno in rlist:
                line = proc.stdout.readline()
                if len(line) == 0:
                    polling.remove(out_fno)
                else:
                    print("STDOUT: ", str(line, "utf-8"))


def _make_request(server: Server, method: str, url: str, body: Optional[bytes] = None, user_id: int = 1) -> Tuple[int, HTTPResponse]:
    con = HTTPConnection(server.host)
    con.request(
        method,
        url,
        body=body,
        headers={
            "x-token": f"{user_id}",
            "content-type": "application/json",
        },
    )
    resp = con.getresponse()
    code = resp.getcode()
    return code, resp


def _make_body(obj: Any) -> bytes:
    if isinstance(obj, pydantic.BaseModel):
        json_str = obj.model_dump_json()
    else:
        json_str = json.dumps(obj)
    data = json_str.encode("utf-8")
    return data


def http_get(server: Server, url: str, user_id: int = 1) -> Tuple[int, Any]:
    code, resp = _make_request(server, "GET", url, None, user_id)
    body = resp.read()
    obj = json.loads(body)
    return code, obj


def http_post(server: Server, url: str, obj: Any, user_id: int = 1) -> Tuple[int, Any]:
    body = _make_body(obj)
    code, resp = _make_request(server, "POST", url, body, user_id)
    resp_body = resp.read()
    resp_obj = json.loads(resp_body)
    return code, resp_obj


def http_put(server: Server, url: str, obj: Any, user_id: int = 1) -> Tuple[int, Any]:
    body = _make_body(obj)
    code, resp = _make_request(server, "PUT", url, body, user_id)
    resp_body = resp.read()
    resp_obj = json.loads(resp_body)
    return code, resp_obj


def http_delete(server: Server, url: str, user_id: int = 1) -> int:
    code, resp = _make_request(server, "DELETE", url, None, user_id)
    return code


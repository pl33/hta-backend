# SPDX-License-Identifier: MPL-2.0
#   Copyright (c) 2023 Philipp Le <philipp@philipple.de>.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

import copy
import multiprocessing
import sqlite3
import subprocess
import select
import signal
import threading

from typing import Any, NamedTuple, Optional, Tuple
import json
import pytest
import os
from http.client import HTTPConnection, HTTPResponse
import pydantic
from flask import Flask, request

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))

DB_FILE = f"{SCRIPT_DIR}/test.db"

SUBPROCESS = True


ISSUER_HOST = "localhost"
ISSUER_PORT = 7500
ISSUER_URL = f"http://{ISSUER_HOST}:{ISSUER_PORT}"
USER_MAP = {
    "1": {
        "name": "Alice Wonderland",
        "given_name": "Alice"
    },
    "2": {
        "name": "John Doe",
        "given_name": "John"
    }
}
oidc_server = Flask(__name__)


@oidc_server.get("/.well-known/openid-configuration")
def oidc_discover():
    return "{" + f'''
  "issuer": "{ISSUER_URL}",
  "authorization_endpoint": "{ISSUER_URL}/protocol/openid-connect/auth",
  "token_endpoint": "{ISSUER_URL}/protocol/openid-connect/token",
  "introspection_endpoint": "{ISSUER_URL}/protocol/openid-connect/token/introspect",
  "userinfo_endpoint": "{ISSUER_URL}/protocol/openid-connect/userinfo",
  "end_session_endpoint": "{ISSUER_URL}/protocol/openid-connect/logout",
  "jwks_uri": "{ISSUER_URL}/protocol/openid-connect/certs",
  "check_session_iframe": "{ISSUER_URL}/protocol/openid-connect/login-status-iframe.html",
  "grant_types_supported": [
    "authorization_code",
    "implicit",
    "refresh_token",
    "password",
    "client_credentials"
  ],
  "response_types_supported": [
    "code",
    "none",
    "id_token",
    "token",
    "id_token token",
    "code id_token",
    "code token",
    "code id_token token"
  ],
  "subject_types_supported": [
    "public",
    "pairwise"
  ],
  "id_token_signing_alg_values_supported": [
    "PS384",
    "ES384",
    "RS384",
    "HS256",
    "HS512",
    "ES256",
    "RS256",
    "HS384",
    "ES512",
    "PS256",
    "PS512",
    "RS512"
  ],
  "id_token_encryption_alg_values_supported": [
    "RSA-OAEP",
    "RSA-OAEP-256",
    "RSA1_5"
  ],
  "id_token_encryption_enc_values_supported": [
    "A256GCM",
    "A192GCM",
    "A128GCM",
    "A128CBC-HS256",
    "A192CBC-HS384",
    "A256CBC-HS512"
  ],
  "userinfo_signing_alg_values_supported": [
    "PS384",
    "ES384",
    "RS384",
    "HS256",
    "HS512",
    "ES256",
    "RS256",
    "HS384",
    "ES512",
    "PS256",
    "PS512",
    "RS512",
    "none"
  ],
  "request_object_signing_alg_values_supported": [
    "PS384",
    "ES384",
    "RS384",
    "HS256",
    "HS512",
    "ES256",
    "RS256",
    "HS384",
    "ES512",
    "PS256",
    "PS512",
    "RS512",
    "none"
  ],
  "response_modes_supported": [
    "query",
    "fragment",
    "form_post"
  ],
  "registration_endpoint": "{ISSUER_URL}/clients-registrations/openid-connect",
  "token_endpoint_auth_methods_supported": [
    "private_key_jwt",
    "client_secret_basic",
    "client_secret_post",
    "tls_client_auth",
    "client_secret_jwt"
  ],
  "token_endpoint_auth_signing_alg_values_supported": [
    "PS384",
    "ES384",
    "RS384",
    "HS256",
    "HS512",
    "ES256",
    "RS256",
    "HS384",
    "ES512",
    "PS256",
    "PS512",
    "RS512"
  ],
  "claims_supported": [
    "aud",
    "sub",
    "iss",
    "auth_time",
    "name",
    "given_name",
    "family_name",
    "preferred_username",
    "email",
    "acr"
  ],
  "claim_types_supported": [
    "normal"
  ],
  "claims_parameter_supported": true,
  "scopes_supported": [
    "openid",
    "offline_access",
    "profile",
    "email",
    "address",
    "phone",
    "roles",
    "web-origins",
    "microprofile-jwt",
    "timeclock"
  ],
  "request_parameter_supported": true,
  "request_uri_parameter_supported": true,
  "code_challenge_methods_supported": [
    "plain",
    "S256"
  ],
  "tls_client_certificate_bound_access_tokens": true,
  "revocation_endpoint": "{ISSUER_URL}/protocol/openid-connect/revoke",
  "revocation_endpoint_auth_methods_supported": [
    "private_key_jwt",
    "client_secret_basic",
    "client_secret_post",
    "tls_client_auth",
    "client_secret_jwt"
  ],
  "revocation_endpoint_auth_signing_alg_values_supported": [
    "PS384",
    "ES384",
    "RS384",
    "HS256",
    "HS512",
    "ES256",
    "RS256",
    "HS384",
    "ES512",
    "PS256",
    "PS512",
    "RS512"
  ],
  "backchannel_logout_supported": true,
  "backchannel_logout_session_supported": true
    ''' + "}"


@oidc_server.post("/protocol/openid-connect/token/introspect")
def oidc_introspect():
    token = request.values.get("token")
    try:
        user = USER_MAP[token]
        name = user["name"]
        given_name = user["given_name"]
    except KeyError:
        return "User not found", 403
    return "{" + f'''
  "active": true,
  "iss": "{ISSUER_URL}",
  "sub": "{token}",
  "name": "{name}",
  "given_name": "{given_name}"
    ''' + "}"


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
    t = multiprocessing.Process(target=lambda: oidc_server.run(ISSUER_HOST, ISSUER_PORT, False))
    t.start()

    if SUBPROCESS:
        env = copy.deepcopy(os.environ)
        env.update({
            "DB": f"sqlite3://{DB_FILE}",
            "OIDC_ISSUER": f"{ISSUER_URL}",
            "OIDC_CLIENT_ID": "",
            "OIDC_CLIENT_SECRET": "",
            "OIDC_FRONTEND_CLIENT_ID": "",
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
        for step in range(5000):
            rlist, _, _ = select.select([err_fno], [], [], 0.1)
            if err_fno in rlist:
                line = str(proc.stderr.readline(), "utf-8")
                if "Serving hta at" in line:
                    break
                else:
                    print("Ignore STDERR: ", line)
        else:
            terminate(proc)
            t.terminate()
            raise TimeoutError

    try:
        db = sqlite3.connect(DB_FILE)
        db.commit()

        yield Server("localhost:8080", db)

        db.close()
        os.remove(DB_FILE)
    finally:
        t.terminate()

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


#!/bin/bash

mockgen -source="api/presenter.go" -destination="api/presenter_mock.go" -package api
mockgen -source="api/server.go" -destination="api/server_mock.go" -package api
mockgen -source="base/usecase.go" -destination="base/usecase_mock.go" -package base
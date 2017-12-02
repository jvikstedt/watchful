#!/bin/sh

npm run build --prefix frontend

statik -src=frontend/dist

go build cmd/watchful/main.go

mv main watchful

#!/bin/sh
gofmt -s -w .
git add .
git commit -m "$1"
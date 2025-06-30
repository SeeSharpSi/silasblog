#!/bin/sh
wgo -file=.go -file=.templ -xfile=_templ.go -file=.css templ generate :: go run . -db="./data.db"

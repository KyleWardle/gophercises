#! /bin/bash
rm -rf ./quizgame
go build
./quizgame "$@"

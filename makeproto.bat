@echo off

protoc --proto_path=./src/robot/dyj/pb --go_out=./src/robot/dyj/pbgo --go_opt=paths=source_relative ./src/robot/dyj/pb/*.proto
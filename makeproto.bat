@echo off

protoc --proto_path=./src/robot/dyj/pb --go_out=./src/robot/dyj/pbgo --go_opt=paths=source_relative ./src/robot/dyj/pb/*.proto

protoc --jsonschema_out=./src/robot/dyj/schema --proto_path=./src/robot/dyj/pb --jsonschema_opt=enforce_oneof --jsonschema_opt=enums_as_strings_only ./src/robot/dyj/pb/game_msg.proto

::python ./DelProtoJsonTag.py

rem grpc版协议
.\protoc-3.7.1-win64.exe --go_out=plugins=grpc:..\src\ .\proto\gs\gs.proto
rem 标准protobuf
rem .\protoc-3.7.1-win64.exe --go_out=..\src\ .\proto\gs\gs.proto
pause
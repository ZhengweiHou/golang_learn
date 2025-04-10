- `kitex -module kitex_demo hello.thrift`
- `mkdir -p rpc && cd rpc`
- `cd rpc`
- `kitex -module kitex_demo -service hzwapi -use kitex_demo/kitex_gen ../hello.thrift`
- 

- 多service多Handler生成 `kitex -tpl multiple_services -service your_service path/to/idl`


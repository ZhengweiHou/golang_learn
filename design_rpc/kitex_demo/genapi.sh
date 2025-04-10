#!/bin/bash

# -gen-path 指定生成代码的路径，默认kitex_gen
# -record 自动记录每次执行的 Kitex 命令并生成脚本文件参数用于生成记录文件 kitex-all.sh
#kitex -module kitex_demo -gen-path api -record hzwcmd.thrift
#kitex -module kitex_demo -gen-path api -record hzwquery.thrift

ls idl | while read tfile
do
    kitex -module kitex_demo -gen-path api/kitex -record idl/${tfile}; 
done

# mkdir -p rpc && cd rpc
# kitex -tpl multiple_services -service hzwservice -use kitex_demo/kitex_gen -gen-path rpc ../idl/hzw.thrift
# kitex -tpl multiple_services -service hzwservice -use kitex_demo/kitex_gen -gen-path rpc ../idl/hello.thrift
# cd -


# kitex -module wiredemo -gen-path api/kitex api/kitex/idl/hzw.thrift
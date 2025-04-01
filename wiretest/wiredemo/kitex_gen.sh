#!/bin/bash

kitex -module wiredemo -gen-path api/kitex api/kitex/idl/hzw.thrift

# SkipDecoder
#kitex -module wiredemo -gen-path api/kitex -thrift no_default_serdes api/kitex/idl/hzw.thrift

#cd internal/adapter/adapter_kitex
#kitex -tpl multiple_services -service hzw -use api/kitex -gen-path internal/adapter/adapter_kitex ../../../api/kitex/idl/hzw.thrift
#cd -



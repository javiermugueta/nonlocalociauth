#!/bin/sh
# jmu, april 2021
#
# packages
#
go get context fmt io/ioutil os github.com/oracle/oci-go-sdk/common github.com/oracle/oci-go-sdk/identity time
#
# build
#
go build nonlocalociauth.go
#
# run
#
./nonlocalociauth 
#
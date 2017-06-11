#######################################################################
##                     Copyright (C) 2017 wystan
##
##       filename: makefile
##    description:
##        created: 2017-06-10 12:10:46
##         author: wystan
##
#######################################################################
.PHONY: all test install db c q clean

db := bitmain
bin := bitmain
src := $(shell find . -name '*.go'|grep -v *_test.go)

all: db $(bin)

$(bin): $(src)
	@echo "===> building ..."
	@go build -o $(bin)
	@echo "===> done."

db:
	@echo "===> install database ..."
	@PGOPTIONS="--client-min-messages=warning" psql -q -f store/pg.sql
	@echo "===> done."

q:
	@psql $(db) -c "select * from students;"
	@psql $(db) -c "select * from classes;"

c:
	@psql $(db) -c "delete from students;"
	@psql $(db) -c "delete from classes;"


clean:
	@echo "===> cleaning ..."
	@psql -c "drop database if exists $(db);"
	@rm -f $(bin) $(bin).log
	@echo "===> done."

#######################################################################

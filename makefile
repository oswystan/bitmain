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
os := $(shell uname -s)
ifeq ($(os), Darwin)
	pgsql := PGOPTIONS="--client-min-messages=warning" psql
else
	pgsql := sudo -u postgres PGOPTIONS="--client-min-messages=warning" psql
endif

all: db $(bin)

$(bin): $(src)
	@echo "===> building ..."
	@go build -o $(bin)
	@echo "===> done."

db:
	@echo "===> install database ..."
	@$(pgsql) -q -f store/pg.sql
	@echo "===> done."

q:
	@$(pgsql) $(db) -c "select * from students;"
	@$(pgsql) $(db) -c "select * from classes;"

c:
	@$(pgsql) $(db) -c "delete from students;"
	@$(pgsql) $(db) -c "delete from classes;"


clean:
	@echo "===> cleaning ..."
	@$(pgsql) -c "drop database if exists $(db);"
	@rm -f $(bin) $(bin).log*
	@echo "===> done."

test:
	@go get github.com/smartystreets/goconvey/convey
	@go test -v

#######################################################################

#######################################################################
##                     Copyright (C) 2017 wystan
##
##       filename: makefile
##    description:
##        created: 2017-06-10 12:10:46
##         author: wystan
##
#######################################################################
.PHONY: all build test install db c q clean

db := "bitmain"
bin := "bitmain"

all: db build


build:
	@echo "===> building ..."
	go build -o $(bin)
	@echo "===> done."

db:
	@echo "===> install database ..."
	@psql -f store/pg.sql > /dev/null
	@echo "===> done."

q:
	@psql $(db) -c "select * from students;"
	@psql $(db) -c "select * from classes;"

c:
	@psql $(db) -c "delete from students;"
	@psql $(db) -c "delete from classes;"


clean:
	@echo "===> cleaning ..."
	@psql -c "drop database if exists bitmain;"
	@rm -f $(bin)
	@echo "===> done."

#######################################################################

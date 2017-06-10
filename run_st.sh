#!/usr/bin/env bash
###########################################################################
##                     Copyright (C) 2017 wystan
##
##       filename: run_st.sh
##    description:
##        created: 2017-06-10 14:08:42
##         author: wystan
##
###########################################################################

function log_start() {
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}]##########################################################"
    echo "[${strNow}] start program  : $0"
    echo "[${strNow}]##########################################################"
    echo ""
}

function logi() {
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}] INFO:$*"
}

function logw() {
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}] WARN:$*"
}

function loge() {
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}]ERROR:$*"
}
function log_end() {
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo ""
    echo "[${strNow}]##########################################################"
    echo "[${strNow}] finished $0"
    echo "[${strNow}]##########################################################"
}

function safe_exec() {
    if [ $# -eq 0 ]; then
        exit 1
    fi

    $*
    if [ $? -ne 0 ]; then
        loge "fail to do [$*]"
        exit 1
    fi
}

function do_student() {
    data_post='{
        "id":"10000",
        "classNumber":1,
        "score": 100
    }'
    curl -XPOST -d "${data_post}" "http://localhost:8000/register-student" -w "%{http_code}\n"
    data_post='{
        "id":"10002",
        "classNumber":2,
        "score": 20
    }'
    curl -XPOST -d "${data_post}" "http://localhost:8000/register-student" -w "%{http_code}\n"
    curl -XGET "http://localhost:8000/get-class-total-score/10000" -w "%{http_code}\n"
    curl -XGET "http://localhost:8000/get-class-total-score/20000" -w "%{http_code}\n"
}

function do_class() {
    data_post='{
        "classNumber":1,
        "teacher": "lucy"
    }'
    curl -XPOST -d "${data_post}" "http://localhost:8000/register-class" -w "%{http_code}\n"
    data_post='{
        "classNumber":2,
        "teacher": "lily"
    }'
    curl -XPOST -d "${data_post}" "http://localhost:8000/register-class" -w "%{http_code}\n"
    curl -XGET "http://localhost:8000/get-top-teacher" -w "%{http_code}\n"
}


function do_work() {
    log_start
    do_student
    do_class
    log_end
}


################################
## main
################################

if [ $# -eq 0 ]; then
    do_work
else
    $*
fi

###########################################################################

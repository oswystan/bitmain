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
    safe_exec curl -XPOST "http://localhost:8000/register-student"
}

function do_class() {
    safe_exec curl -XPOST "http://localhost:8000/register-class"
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

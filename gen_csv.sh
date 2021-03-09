#!/bin/bash

#while [ true ]
#do
    PREFIX=`date '+%Y%m%d'`
    SHOPID=-100
    POSID=1
    FILENAME=${PREFIX}_SHOP${SHOPID}_POS${POSID}.csv
    if [ `uname` != "Darwin" ]; then
        gen_csv/gen_csv.exe -shopid ${SHOPID} -posid ${POSID} > data/before/${FILENAME}
    else
        go run gen_csv/gen_csv.go -shopid ${SHOPID} -posid ${POSID} > data/before/${FILENAME}
    fi
    echo ${FILENAME}
    sleep 1
#done

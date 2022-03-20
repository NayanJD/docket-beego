#!/bin/bash

if [ "$1" = "migrate" ]
then
    bee migrate -driver=postgres -conn=$DB_CONN_STR
else
    ./docket-beego
fi
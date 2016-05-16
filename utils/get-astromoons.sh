#!/bin/bash

USAGE="Usage:
CLIENT_ID=appid CLIENT_SECRET=secretkey ./get-astromoons.sh FROM_YEAR TO_YEAR"

# Aeris Weather API - Registered Apps
# Free Developer API Subscription
# 750 hits/day, 10 hits/minute
# http://www.aerisweather.com/account/apps

if [ "$CLIENT_ID" == "" -o "$CLIENT_SECRET" == "" -o "$1" == "" -o "$2" == "" ]; then
    echo $USAGE
    exit 1
fi

# There are ~208 moon records per year. API allows max 250 per request.
limit=250

from_year=$1
to_year=$2
place="london,uk" # for UTC time

year=$from_year
lastBatchStart=`date +%s`
reqPerMin=0

while [ $year -le $to_year ]; do
    echo -n "$year ... "

    # Stay under 10 requests per minute
    if [ $reqPerMin -gt 10 ]; then
        d=$(expr `date +%s` - $lastBatchStart)
        # +1s for error
        while [ $d -lt 61 ]; do
            echo -ne "\r$year ... wait: "$(expr 60 - $d )" "
            sleep 1
            d=$(expr `date +%s` - $lastBatchStart)
        done
        lastBatchStart=`date +%s`
        reqPerMin=0
    fi

    curl -s "http://api.aerisapi.com/sunmoon/moonphases/$place?from=$year/01/01&to=$year/12/31&limit=$limit&client_id=$CLIENT_ID&client_secret=$CLIENT_SECRET" > astro-$year.json
    ((year += 1))
    ((reqPerMin += 1))

    echo "OK"
done


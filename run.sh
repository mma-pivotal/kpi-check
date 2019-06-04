#!/bin/bash

curl $kpi_url > latest.html
diff latest.html last.html > /dev/null 2>&1
error=$?
echo $error
if [ $error -eq 1 ]
then
   curl -X POST -H 'Content-type: application/json' --data '{"text":"Changes found on KPI page '"$kpi_url"'."}' $slack_url 
elif [ $error -eq 0 ]
then 
   curl -X POST -H 'Content-type: application/json' --data '{"text":"No changes found on KPI page '"$kpi_url"'."}' $slack_url
fi
rm last.html
mv latest.html last.html

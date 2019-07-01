#!/bin/bash

curl $kpi_url > latest.html
{ read var1; read var2; } <<< $(grep -w main latest.html -n | cut -f 1 -d ":")
awk "NR>=$var1&&NR<=$var2" latest.html > extracted.html
diff extracted.html last.html > diff.txt 
error=$?
echo $error
if [ $error -eq 1 ]
then
   curl -X POST -H 'Content-type: application/json' --data '{"text":"Changes found on KPI page '"$kpi_url"'. \n Diff can be found : https://kpi-check.cfapps.io/diff"}' $slack_url 
elif [ $error -eq 0 ]
then 
   curl -X POST -H 'Content-type: application/json' --data '{"text":"No changes found on KPI page '"$kpi_url"'."}' $slack_url
fi
rm last.html
rm latest.html
mv extracted.html last.html

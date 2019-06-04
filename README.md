# PCF KPI DOC Check

A simple app which check whether there is any KPI change in PCF's official website. 

## Frontend

A empty go webapp which only returns 200 at this moment.
I am planning to add more functionality so that the frontend can take user input to support more pages / frequency of job runs / send messages to multiple channels.

## Backend

A simple bash script which download the latest KPI page and `diff` it with the last saved page.
POST a message specified in the environment variable $slack_url.
For now this job runs at 9am every Monday.

## Environment variables.

`slack_url` - The incoming web hook for slack.  
`kpi_url` - The KPI doc that is being checked.

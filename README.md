# monitoring_alert_tool
An alert and monitoring tool for checking the functionality and availability of a system

#### This tool or repo features 4 components.
- First is the simple application A, which is being monitored
- Second is the monitoring or alert tool B that monitors application A and records the event based on A response.
- Third is the monitor of the monitoring tool. The tool will check the status of the monitoring tool at intervals and send an email. 
- Fourth component is an elastic search cloud instance on aws cloud service. It will contain records of monitoring events at intervals.

All other components of this system can be started with docker using the following command.

    docker-compose up -d

or if you're rebuilding the docker images, you can use the following command.

    docker-compose up -d --build

List of required environment variables required for the system to run. Copy the following environment variables to your .env file using the sample .env file provided.

    - MONITORING_ALERT_TOOL_AWS_ACCESS_KEY_ID
    - MONITORING_ALERT_TOOL_AWS_SECRET_ACCESS_KEY
    - MONITORING_ALERT_TOOL_AWS_REGION
    - MONITORING_ALERT_TOOL_AWS_ES_HOST
    - MONITORING_ALERT_TOOL_AWS_ES_PORT
    - MONITORING_ALERT_TOOL_AWS_ES_INDEX
    - MONITORING_ALERT_TOOL_AWS_ES_TYPE

Available endpoints on the watched system.
    - http://localhost:8000/ - This is the root url of the watched system.
    - http://localhost:8000/docs - to view the openapi documentation
    - http://localhost:8000/api/v1/active_status - to get the active status of the watched system.
    - http://localhost:8000/api/v1/inactive_status - to get the inactive status of the watched system (if the system is down, returns round error code between 4XX and 5XX).
    - http://localhost:8000/api/v1/server_error - to get the server error status of the watched system.

Starting the system in production mode
    -   OPENAPI_URL= uvicorn main:app or $env:OPENAPI_URL = uvicorn main:app
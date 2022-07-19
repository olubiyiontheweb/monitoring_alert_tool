# monitoring_alert_tool
An alert and monitoring tool for checking the functionality and availability of a system

####This tool or repo features 4 components.
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
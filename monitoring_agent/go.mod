module github.com/olubiyiontheweb/monitoring_alert_tool

go 1.18

require (
	github.com/aws/aws-sdk-go v1.44.61
	github.com/joho/godotenv v1.4.0
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

replace github.com/olubiyiontheweb/monitoring_alert_tool/pkgs/api_caller => ../pkgs/api_caller

replace github.com/olubiyiontheweb/monitoring_alert_tool/pkgs/alert_sender => ../pkgs/alert_sender

replace github.com/olubiyiontheweb/monitoring_alert_tool/pkgs/database => ../pkgs/database

from pydantic import BaseSettings, HttpUrl

class HiddenSettings(BaseSettings):
    AWS_ES_DOMAIN_ENDPOINT: HttpUrl = "" # elasticsearch domain endpoint
    ES_INDEX: str = "" # elasticsearch index
    AWS_REGION: str = "" # AWS region
    AWS_SERVICE: str = "es" # AWS ES service
    AWS_ES_ACCESS_KEY_ID: str = "" # AWS access key ID
    AWS_ES_SECRET_ACCESS_KEY: str = "" # AWS secret access key
    AUTH_TOKEN: str = "" # authentication token
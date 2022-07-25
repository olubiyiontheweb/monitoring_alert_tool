import logging
from env import HiddenSettings

# global settings for the application
class Settings(HiddenSettings):
    APP_NAME: str = "Watched System"
    DESCRIPTION: str = "Watched System is a simple application being monitored by the tool."
    API_V1_STR: str = "/api/v1"
    VERSION: str = "0.1.0"
    OPENAPI_URL: str = "/openapi.json"
    DOCS_URL: str = "/docs"
    DEBUG: bool = False
    LOG_LEVEL: int = logging.INFO
    ELASTICSEARCH_PAGE_SIZE: int = 30
    ELASTICSEARCH_TIMEOUT: int = 10
    
settings = Settings()
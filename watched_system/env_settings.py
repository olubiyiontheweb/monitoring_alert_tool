import logging
from pydantic.env_settings import BaseSettings

# global settings for the application
class Settings(BaseSettings):
    APP_NAME: str = "Watched System"
    DESCRIPTION: str = "Watched System is a simple application being monitored by the tool."
    API_V1_STR: str = "/api/v1"
    VERSION: str = "0.1.0"
    OPENAPI_URL: str = "/openapi.json"
    DOCS_URL: str = "/docs"
    DEBUG: bool = False
    LOG_LEVEL = logging.INFO
    
settings = Settings()
import logging

from fastapi import FastAPI

from env_settings import settings
from endpoints import routes, index_routes

# setting logging level for the application
logging.basicConfig(level=settings.LOG_LEVEL)
logger = logging.getLogger(__name__)

# fastapi app configuration
app = FastAPI(title=settings.APP_NAME,
              version=settings.VERSION, 
              openapi_url=settings.OPENAPI_URL,
              docs_url=settings.DOCS_URL,
              debug=settings.DEBUG)

# router for endpoints
app.include_router(index_routes)
app.include_router(routes, prefix=settings.API_V1_STR)

@app.on_event("startup")
async def startup():
    # start all dependency services
    logger.info("\nWatched System Started Successfully!!!\n")

@app.on_event("shutdown")
async def shutdown():
    # shutdown all dependency services
    logger.info("\nWatched System Shutdown Successfully!!!\n")
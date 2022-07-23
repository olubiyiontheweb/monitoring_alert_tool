import random
from typing import Union

from fastapi import APIRouter, Request, status
from fastapi.responses import JSONResponse

from env_settings import settings

# setting up routers
index_routes = APIRouter()
routes = APIRouter()

# endpoint for index page
@index_routes.get("/")
def read_root():
    return JSONResponse(status_code=status.HTTP_200_OK, 
                        content={"message": settings.APP_NAME + " v" + settings.VERSION + " is running"})

# other endpoints

# return active status of the server
@routes.get("/active_status")
def read_app_status():
    return JSONResponse(status_code=status.HTTP_200_OK,
                        content={"message": settings.APP_NAME + " v" + settings.VERSION + " is working fine"})

# return inactive status of the server (any status code from server or client http response)
@routes.get("/inactive_status")
def read_app_status_inactive():
    return JSONResponse(status_code=random.randint(status.HTTP_400_BAD_REQUEST, status.HTTP_511_NETWORK_AUTHENTICATION_REQUIRED),
                        content={"message": "Service Unavailable"})

# return bad request status of the server
@routes.get("/server_error")
def read_server_error(request: Request):
    return JSONResponse(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR, 
                        content={"message": "Internal Server Error"})
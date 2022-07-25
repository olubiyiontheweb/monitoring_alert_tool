import random
from typing import Union

from fastapi import APIRouter, Request, status, Depends
from fastapi.responses import JSONResponse
from fastapi.security import HTTPBearer, HTTPAuthorizationCredentials

from env_settings import settings
from elasticsearch_query import elastic_search_query

# setting up routers
index_routes = APIRouter()
routes = APIRouter()

# validate bearer token in the request header
def validate_token(HTTPBearer):
    if HTTPBearer.credentials == settings.AUTH_TOKEN:
        return True
    else:
        return JSONResponse(status_code=status.HTTP_401_UNAUTHORIZED, content={"message": "Unauthorized"})

# endpoint for index page
@index_routes.get("/")
def read_root(Authorization: HTTPBearer = Depends(HTTPBearer())):
    if validate_token(Authorization):
        return JSONResponse(status_code=status.HTTP_200_OK, 
                        content={"message": settings.APP_NAME + " v" + settings.VERSION + " is running"})

# other endpoints

# return active status of the server
@routes.get("/active_status")
def read_app_status(Authorization: HTTPBearer = Depends(HTTPBearer())):
    
    if validate_token(Authorization):
        return JSONResponse(status_code=status.HTTP_200_OK,
                        content={"message": settings.APP_NAME + " v" + settings.VERSION + " is working fine"})

# return inactive status of the server (any status code from server or client http response)
@routes.get("/inactive_status")
def read_app_status_inactive(Authorization: HTTPBearer = Depends(HTTPBearer())):
    
    if validate_token(Authorization):
        return JSONResponse(status_code=random.randint(status.HTTP_400_BAD_REQUEST, status.HTTP_511_NETWORK_AUTHENTICATION_REQUIRED),
                        content={"message": "Service Unavailable"})

# return bad request status of the server
@routes.get("/server_error")
def read_server_error(Authorization: HTTPBearer = Depends(HTTPBearer())):
    
    if validate_token(Authorization):
        return JSONResponse(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR, 
                        content={"message": "Internal Server Error"})
    
# return elasticsearch query results
@routes.get("/elasticsearch_query")
def read_elasticsearch_query(Authorization: HTTPBearer = Depends(HTTPBearer())):
    if validate_token(Authorization):
        response = elastic_search_query.get_records()
        return JSONResponse(status_code=status.HTTP_200_OK, content=response)
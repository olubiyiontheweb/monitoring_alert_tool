import logging
from fastapi.encoders import jsonable_encoder
from elasticsearch import Elasticsearch, RequestsHttpConnection
from requests_aws4auth import AWS4Auth

from env_settings import settings

class ElasticSearchQuery:
    """
     Initialize and manage connection to Elasticsearch.
    """

    def __init__(self):
        awsauth = AWS4Auth(settings.AWS_ES_ACCESS_KEY_ID, settings.AWS_ES_SECRET_ACCESS_KEY, settings.AWS_REGION, settings.AWS_SERVICE)
    
        self.client = Elasticsearch(
            hosts=[settings.AWS_ES_DOMAIN_ENDPOINT],
            http_auth=awsauth,
            use_ssl=True,
            verify_certs=True,
            connection_class=RequestsHttpConnection,            
            timeout=settings.ELASTICSEARCH_TIMEOUT)
        
    def get_records(self):
        response = self.client.search(size=settings.ELASTICSEARCH_PAGE_SIZE, index=settings.ES_INDEX, sort="_id:desc")
        return response


# initiallize elasticsearch connection settings
elastic_search_query = ElasticSearchQuery()
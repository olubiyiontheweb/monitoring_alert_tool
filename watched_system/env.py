from pydantic import BaseSettings, HttpUrl

class HiddenSettings(BaseSettings):
    AWS_ES_DOMAIN_ENDPOINT: HttpUrl = "https://search-quodity-lqotvuewhisjyunjtx27r7zvlu.eu-west-2.es.amazonaws.com/"
    ES_INDEX: str = "quodity_status_logs"
    AWS_REGION: str = "eu-west-2"
    AWS_SERVICE: str = "es"
    AWS_ES_ACCESS_KEY_ID: str = "AKIAYPOOJNHSC7N3YGWA"
    AWS_ES_SECRET_ACCESS_KEY: str = "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"
    AUTH_TOKEN: str = "@2hbr6TB*Lya%$3DNljZ^^c$!9XOm7s5siG1KUYNUbo0!w4FmR5G2mvttDlrxiAMpzPnaMtyorfQI@kqWl"
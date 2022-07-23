FROM python:3.8

ENV PYTHONUNBUFFERED=1
ENV SYSTEM_ENVIRONMENT='PRODUCTION'

WORKDIR /app

COPY ./requirements.txt /code/requirements.txt
COPY ./watched_system /app

RUN pip install --no-cache-dir -r requirements.txt
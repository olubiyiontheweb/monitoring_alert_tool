FROM python:3.8-alpine3.15

ENV PYTHONUNBUFFERED=1
ENV SYSTEM_ENVIRONMENT='PRODUCTION'

WORKDIR /app

COPY . /app

RUN pip install --no-cache-dir -r requirements.txt

WORKDIR /app/watched_system
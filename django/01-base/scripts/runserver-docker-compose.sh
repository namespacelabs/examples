#!/usr/bin/env bash
env
sleep 2
aws --endpoint-url=${AWS_S3_ENDPOINT_URL} s3api create-bucket --bucket=${AWS_STORAGE_BUCKET_NAME}
poetry run python manage.py migrate --no-input
poetry run python manage.py tailwind build --no-input
poetry run python manage.py collectstatic --no-input
# poetry run python manage.py runserver 0.0.0.0:8000
poetry run python -m uvicorn todo.asgi:application --host 0.0.0.0 --port 8000 
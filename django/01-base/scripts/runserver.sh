#!/usr/bin/env bash
env
poetry run python manage.py migrate --no-input
poetry run python manage.py tailwind build --no-input
poetry run python manage.py collectstatic --no-input
poetry run python -m uvicorn todo.asgi:application --host 0.0.0.0 --port 8000 
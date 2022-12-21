FROM docker.io/python:3.10.8-slim-bullseye
ARG TARGETPLATFORM
ARG BUILDPLATFORM

ENV PYTHONDONTWRITEBYTECODE="1" \
    PYTHONUNBUFFERED="1" \
    PYTHONFAULTHANDLER=1 \
    PYTHONHASHSEED=random \
    PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on \
    PIP_DEFAULT_TIMEOUT=100 \
    POETRY_VERSION=1.3.1 

RUN apt-get update -y && apt-get install -y \
    build-essential \
    curl \
    postgresql-client \
    less \
    groff \
    unzip 

# Node
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - \
    && apt-get install -y nodejs --no-install-recommends 

# AWS Cli2
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then URL="https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip"; else URL="https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"; fi; \
    curl ${URL} -o "awscliv2.zip" && unzip awscliv2.zip && ./aws/install

# Purge
RUN apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false; \
    rm -rf /var/lib/apt/lists/* /usr/share/doc /usr/share/man; 

# Install poetry
RUN pip install "poetry==$POETRY_VERSION"

# Install app deps
WORKDIR /app

# Install python dependencies
COPY poetry.lock pyproject.toml /app/
RUN poetry install --no-ansi --no-interaction --without=dev

# Install javascript dependencies
COPY theme/static_src/package.json theme/static_src/package-lock.json /app/theme/static_src/
RUN cd theme/static_src/; npm install

COPY . /app

EXPOSE 8000

CMD ["./runserver.sh"]
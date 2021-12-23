FROM mysql

ENV DOCKER_DEFAULT_PLATFORM=linux/amd64   
ADD ./volume/databaseFile.sql /docker-entrypoint-initdb.d
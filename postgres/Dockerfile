FROM postgis/postgis:14-master

RUN mkdir -p /docker-entrypoint-initdb.d
COPY docker-entrypoint-initdb /docker-entrypoint-initdb.d/

#docker run --name testgis -e  POSTGRES_USER="black-kiwi_administrator" -e POSTGRES_DB="black-kiwi" -e  POSTGRES_PASSWORD="6whuUYT#E&hy@A2#" -p 5432 -d postgis/postgis:14-master
#docker run --name testgis -e  POSTGRES_USER="black-kiwi_administrator" -e POSTGRES_DB="black-kiwi" -e  POSTGRES_PASSWORD="6whuUYT#E&hy@A2#" -p 127.0.0.1:5432:5432 -d test
#docker build -t aleben/itlandfill-cas-postgis-preloaded:latest -t aleben/itlandfill-cas-postgis-preloaded:1.4 .
#docker push aleben/itlandfill-cas-postgis-preloaded --all-tags
FROM postgis/postgis:14-master

RUN mkdir -p /docker-entrypoint-initdb.d
COPY docker-entrypoint-initdb /docker-entrypoint-initdb.d/

#docker build -t aleben/itlandfill-cas-postgis-preloaded:latest -t aleben/itlandfill-cas-postgis-preloaded:1.6 .
#docker push aleben/itlandfill-cas-postgis-preloaded --all-tags
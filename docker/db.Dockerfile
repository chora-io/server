FROM postgres:15

# set data directory
ENV PGDATA=/var/pgdata
ENV LOGDIR=/var/log/postgresql

# install dependencies
RUN apt-get update
RUN apt-get install jq -y

# create directory to store PostgreSQL data and logs
RUN mkdir -p ${PGDATA} ${LOGDIR}

# change owner of directory to postgres group and user
RUN chown -R postgres:postgres ${PGDATA} ${LOGDIR}

# change working directory
WORKDIR /home/db

# copy migrations
COPY db/migrations /home/db/migrations

# copy data seed file
COPY docker/data/db_data.json /home/db/data/

# copy init script
COPY docker/scripts/db_init.sh /home/db/scripts/

# copy start script
COPY docker/scripts/db_start.sh /home/db/scripts/

# make init script executable
RUN ["chmod", "+x", "/home/db/scripts/db_init.sh"]

# make start script executable
RUN ["chmod", "+x", "/home/db/scripts/db_start.sh"]

# set user to run init script and db container
USER postgres

# run init script
RUN /home/db/scripts/db_init.sh

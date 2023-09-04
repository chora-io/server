FROM postgres:15

# Set data directory
ENV PGDATA=/var/pgdata
ENV LOGDIR=/var/log/postgresql

# Create directory to store PostgreSQL data and logs
RUN mkdir -p ${PGDATA} ${LOGDIR}

# Change owner of directory to postgres group and user
RUN chown -R postgres:postgres ${PGDATA} ${LOGDIR}

# Change working directory
WORKDIR /home/db

# Copy migrations
COPY db/migrations /home/db/migrations

# Copy init script
COPY docker/scripts/db_init.sh /home/db/scripts/

# Copy seed script
COPY docker/scripts/db_seed.sh /home/db/scripts/

# Copy start script
COPY docker/scripts/db_start.sh /home/db/scripts/

# Make init script executable
RUN ["chmod", "+x", "/home/db/scripts/db_init.sh"]

# Make seed script executable
RUN ["chmod", "+x", "/home/db/scripts/db_seed.sh"]

# Make start script executable
RUN ["chmod", "+x", "/home/db/scripts/db_start.sh"]

# Set user to run init script and db container
USER postgres

# Run init script
RUN /home/db/scripts/db_init.sh

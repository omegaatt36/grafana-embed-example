#!/bin/sh

if [ "$(docker images -q liquibase/liquibase-mysql:latest 2> /dev/null)" = "" ]; then
  echo \
"FROM liquibase/liquibase
RUN lpm add mysql --global" | \
docker build \
    -t "liquibase/liquibase-mysql" --quiet -f - \
    . || exit 1
fi

docker run --rm --net host \
  -e MYSQL_DATABASE=grafana \
  -v "$(pwd)/res/changeset.mysql.sql:/liquibase/changelog/changeset.mysql.sql" \
  liquibase/liquibase-mysql \
  --searchPath=/liquibase/changelog \
  --changelog-file=changeset.mysql.sql \
  --url="jdbc:mysql://localhost:3306/grafana" \
  --username="root" \
  --password="root" \
  update
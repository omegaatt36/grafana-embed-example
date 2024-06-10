setup: setup-containers setup-migration

setup-containers:
	docker-compose up -d

setup-migration:
	. bin/migration.sh

remove: stop-containers

stop-containers:
	docker-compose stop
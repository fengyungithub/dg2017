#!/usr/bin/make -f

start: up deps install modules open

up:
	docker-compose up --build -d

deps:
	docker-compose exec site1 /bin/bash -c 'composer install --prefer-dist --no-progress && composer drupal-scaffold'
	docker-compose exec site1 /bin/bash -c 'composer install --prefer-dist --no-progress && composer drupal-scaffold'

install:
	docker-compose exec site1 /bin/bash -c 'drush site-install -y --site-name="Dept. of Awesome" --account-pass=password standard && chown -R www-data:www-data app'
	docker-compose exec site2 /bin/bash -c 'drush site-install -y --site-name="Dept. of Cool" --account-pass=password standard && chown -R www-data:www-data app'

modules:
	docker-compose exec site1 /bin/bash -c 'drush en -y remote_article'
	# docker-compose exec site2 /bin/bash -c 'drush en -y remote_article'

open:
	xdg-open http://localhost:9000
	xdg-open http://localhost:8081
	xdg-open http://localhost:8082

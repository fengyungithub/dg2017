#!/bin/bash

docker-compose up -d

docker-compose exec site1 /bin/bash -c 'drush site-install -y --site-name="Dept. of Awesome" --account-pass=password standard && chown -R www-data:www-data app'
docker-compose exec site2 /bin/bash -c 'drush site-install -y --site-name="Dept. of Cool" --account-pass=password standard && chown -R www-data:www-data app'

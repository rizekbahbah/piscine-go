curl -s https://adam-jerusalem.nd.edu/assets/superhero/all.json | jq ".[] | select (.id == $HERO_ID) | .connections.relatives" |tr -d "\""


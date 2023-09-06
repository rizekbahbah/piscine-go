#! /bin/bash

curl https://adam-jerusalem.nd.edu/assets/superhero/all.json | jq '.[] | select (.id == 170)| /name
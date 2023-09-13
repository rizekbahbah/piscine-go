#! /bin/bash

name=$(curl -s https://adam-jerusalem.nd.edu/assets/superhero/all.json | jq " .[] | select(.id == 70) | .name")
echo  "$name" 
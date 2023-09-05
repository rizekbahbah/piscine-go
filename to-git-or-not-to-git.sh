#! /bin/bash

json_file="superhero.json"
superhero_id=170
result-$(jq -r ".[] | select(.id == $superhero_id) | .name, ,power, .gender" "$json_file")
if [[ -n "$result" ]];then
    printf "$result" | tr '\n''' 
    printf
else
    printf "Superhero with ID $superhero_id not found." >&2
    if

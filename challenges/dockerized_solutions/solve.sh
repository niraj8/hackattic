#!/bin/bash 

docker run -d -p 5001:5000 --name registry registry:2.7

tags=($(go run main.go | grep Tag | jq -r ".aux.Tag" | tr "\n" " "))

ignition_key=$(head -n 1 IGNITION_KEY.txt)

echo "$ignition_key"

for tag in "${tags[@]}"
do
    docker pull localhost:5001/hack:"$tag"
    output=$(docker run -e IGNITION_KEY="$ignition_key" --name hack-"$tag" localhost:5001/hack:"$tag")
    if [ "$output" != "oops wrong image" ]; then
        echo "correct image"
        curl --data "{\"secret\": \"$output\"}" --header 'Content-Type: application/json' "https://hackattic.com/challenges/dockerized_solutions/solve?access_token=$HACKATTIC_ACCESS_TOKEN"
    fi  
done

docker rm -f registry

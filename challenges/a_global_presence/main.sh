#!/usr/bin/bash

# 
presence_token=$(curl -s https://hackattic.com/challenges/a_global_presence/problem?access_token="${HACKATTIC_ACCESS_TOKEN}" | jq -r .presence_token)

country_codes=("US" "CA" "GB" "DE" "FR" "JP" "AU" "BR" "IN" "CN" "RU" "MX" "IT" "ES" "ZA")

geotargetly_url_part="https://geotargetly.com/geo-browse?url=https://hackattic.com/_/presence/$presence_token&country_code="

echo https://hackattic.com/_/presence/"$presence_token"

for country_code in "${country_codes[@]}"
do
    echo "Processing country code: $country_code"
    complete_url="$geotargetly_url_part$country_code"
    open "${complete_url}"
done

sleep 28
curl --location "https://hackattic.com/challenges/a_global_presence/solve?access_token=${HACKATTIC_ACCESS_TOKEN}" --header 'Content-Type: application/json' --data '{}'

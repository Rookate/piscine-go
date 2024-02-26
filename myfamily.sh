curl --silent "https://zone01normandie.org/assets/superhero/all.json" | jq -r --argjson hero_id $HERO_ID'.[] | select(.id == $hero_id) | .connections.relatives | gsub(("\n"; "\\n")'

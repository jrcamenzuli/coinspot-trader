docker run --rm --name coinspot-trader -p 8080:8080 -e COINSPOT_KEY=$env:COINSPOT_KEY -e COINSPOT_SECRET=$env:COINSPOT_SECRET -d coinspot-trader ./app -mode publisher

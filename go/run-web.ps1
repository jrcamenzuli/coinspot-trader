docker login -u jrcamenzuli -p $env:DOCKER_TOKEN_GITHUB ghcr.io ; docker run --restart always --pull always --name coinspot-trader -p 10000:8081 -e COINSPOT_KEY=$env:COINSPOT_KEY -e COINSPOT_SECRET=$env:COINSPOT_SECRET -d ghcr.io/jrcamenzuli/coinspot-trader ./app -mode web
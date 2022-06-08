#/bin/bash

#docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t tassilobalbo/planthub-client --push client/.
#docker buildx build --platform linux/arm64,linux/arm/v7 -t tassilobalbo/planthub-server --push server/.

# Local test
cd ..
docker buildx build -t orexinmaster/planthub-client --push client/.
docker buildx build -t orexinmaster/planthub-server --push test/.
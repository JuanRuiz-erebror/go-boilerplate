docker rm main_dev --force
docker rm redis --force
docker rm nginx-proxy --force
docker-compose -f docker-compose-dev.yml up -d
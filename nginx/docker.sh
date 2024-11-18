sudo docker run -d --name nginx \
  -p 80:80 \
  -v $(pwd)/nginx.conf:/etc/nginx/conf.d/default.conf \
  --link nb-frontend:frontend \
  --link nb-backend:backend \
  nginx
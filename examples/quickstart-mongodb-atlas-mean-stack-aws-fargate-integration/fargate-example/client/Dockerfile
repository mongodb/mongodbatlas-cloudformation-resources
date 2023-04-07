FROM --platform=linux/x86_64 node:17-slim AS build

WORKDIR /usr/src/app
COPY package.json package-lock.json ./

# Install dependencies and copy them to the container
RUN npm install
COPY . .

# Build the Angular application for production
RUN npm run build --prod

# Configure the nginx web server
FROM --platform=linux/x86_64 nginx:1.17.1-alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build /usr/src/app/dist/client /usr/share/nginx/html

CMD sed -i -e "s|SERVER_URL_PLACEHOLDER|${ATLAS_URI}|g" /usr/share/nginx/html/main.*.js && nginx -g 'daemon off;'

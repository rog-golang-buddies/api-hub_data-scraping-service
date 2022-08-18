# Data Scraping Service
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/rog-golang-buddies/api-hub_data-scraping-service/main.svg)](https://results.pre-commit.ci/latest/github/rog-golang-buddies/api-hub_data-scraping-service/main)

## Description
Service asynchronously process user request to add new Open API.
In other words this service processes content of Open API file, transforms it to the ASD (API Specification Document) model and sends next to the storage and update service.

### Main functions (To Do)
1. Listen to queue events (links to open API yaml/json files)
2. Check link availability
3. Retrieve file content
4. Validate content
5. Parse content into an ASD model
6. Put ASD model with metadata to the storage and update service queue

### Starting service
The easiest way to start an application is to do it with docker.
If you have docker you just need to run a command from the project root
`docker-compose -f ./docker/docker-compose-dev.yml up -d --build`.
And `docker-compose -f ./docker/docker-compose-dev.yml down` to stop.
You can observe queues, and send and retrieve messages from queues via the web interface available by address http://localhost:15672.
login/password = guest/guest.

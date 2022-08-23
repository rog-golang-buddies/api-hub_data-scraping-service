# Data Scraping Service

[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/rog-golang-buddies/api-hub_data-scraping-service/main.svg)](https://results.pre-commit.ci/latest/github/rog-golang-buddies/api-hub_data-scraping-service/main)

## Description

Service asynchronously process user request to add new Open API.
In other words, this service processes the content of the Open API file and transforms it into the ASD (API
Specification Document) model and sends it next to the storage and update service.

### Starting service

The easiest way to start an application is to do it with docker.
If you have docker you just need to run a command from the project root
`docker-compose -f ./docker/docker-compose-dev.yml up -d --build`.
And `docker-compose -f ./docker/docker-compose-dev.yml down` to stop.
You can observe queues, and send and retrieve messages from queues via the web interface available by
the address http://localhost:15672.
login/password = guest/guest.

### MVP version

1. Listen for the events with the static links to the open API specification files.
2. Download & parse openapi specification into a common API specification document(ASD) (view for the UI part).
3. Send notification to the API gateway if required (depends on the flag; look 'How it works' section)
4. Post ASD to the result queue.

#### Communication model

Consume requests with the file urls and notification flag
Default listen queue name: data-scraping-asd
Request:

```json5
{
    "file_url": "https://developer.atlassian.com/cloud/trello/swagger.v3.json",
    "is_notify_user": true
}
```

If "is_notify_user" is true then this service must post notifications to the separate queue. A notification contains one
field with an error model. If an error happens it will contain an error otherwise nil.
Default notification queue name: gateway-scrape-notifications
Example:

```json5
{
    "error": {
        "cause": "file exceed the limit: 5242880",
        "message": "error while processing url"
    }
}
```

If the parsing process has been completed correctly then the result will be posted to the result queue and delivered to
the 'storage and update service'
Default result queue name: storage-update-asd
The model is too big, so I don't give its description here - see the code for details.

#### How to check functionality manually using the RabbitMQ management page

1. Start service as mentioned in the 'Start service' section
2. Go to http://localhost:15672 and login as guest/guest
3. Go to the Queue tab.
4. Check that data-scraping-asd queue has been already presented here
5. Expand 'Add a new queue' section under the 'Overview' and add 2 queues: 'gateway-scrape-notifications' and
   'storage-update-asd'
6. Go into the data-scraping-asd queue and expand the 'Publish message' section under the charts
7. Add request body and publish a message
8. You can check service logs with `docker logs dss`, return to the Queues tab and check result messages in the queues
   using the "Get messages" section

### Known current limitations (TO DO)

1. Supported only swagger 3.0 version.
2. Ignore field constraints (max length and etc.)

### Main functions

1. Listen to queue events (links to open API yaml/json files)
2. Check link availability
3. Retrieve file content (there is a limit of file size - by default it's 5 Mb)
4. Validate content
5. Parse content into an ASD model
6. Put ASD model with metadata to the storage and update service queue

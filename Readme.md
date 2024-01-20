<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<a name="readme-top"></a>

<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">PowerLog Explorer</h3>

  <p align="center">
    A complete full stack log explorer
  </p>
</div>

<!-- TABLE OF CONTENTS -->

## Table of Contents

  <ol>
    <li>
      About The Project
    </li>
    <li>
     Getting Started
    </li>
    <li>Usage</li>
    <li>Features</li>
    <li>System Design</li>
    <li>Some QnA</li>
    <li>Contributing</li>
    <li>License</li>
    <li>Contact</li>
  </ol>

<!-- ABOUT THE PROJECT -->

## About The Project

![Capture-2023-11-19-004922](https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/9f661760-c4b9-4c30-8e0d-8b8c7cdc0320)

This project is created in order checkout logs in a more organised manner. We have all been in a position where there are too many logs and we are not able to filter them out. This project tries to solve this issue. It has both a frontend and a backend which is highly scalable and can handle millions of requests. More of the systen design in the later section.

### Built With

This project is built with

**Frontend**

1. ReactJS
2. TypeScript
3. Vite
4. TailwindCSS

**Backend**

1. Golang (GIN Framework)

**Databases**

1. MongoDB
2. Redis

<!-- GETTING STARTED -->

## Getting Started

Getting the project running is our first task and there are 2 ways to do so. One is manual installation and the other is dockerised installation (recommended).

### Prerequisites

The following things are required to get started

1. Docker [( link )](https://docs.docker.com/engine/install/)
2. NodeJS [( link )](https://nodejs.org/en/download)
3. Yarn. ( If not installed then run `npm i -g yarn` )
4. Golang [( link )](https://go.dev/doc/install) ( Not Required in Dockerised Installation )
5. MongoDB [( link )](https://www.mongodb.com/docs/manual/administration/install-community/) ( Not Required in Dockerised Installation )
6. Redis Stack Server [( link )](https://redis.io/docs/install/install-stack/) ( Not Required in Dockerised Installation )

### Setting Up Backend [ Dockerised ] [**Recommended**]

This will be actually very simple

1. Clone the repo

```bash
git clone https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK.git
```

2. `cd november-2023-hiring-Shubhrajyoti-Dey-FrosTiK`
3. `sudo docker-compose up`

And thats it !!!

It will take time to setup as it will download every dependency and all but thats it. You dont need to worry about anything. Not even `env`.

**Note**

In this dockerised setup it doesnot consider that you have `MongoDB` or `Redis` running locally. If you have then also it will work as both are running on different ports.

Also note that initially the `backend` or the `job` service may not start but just restart it once every other service is running. (This happens as we are all dependent on MongoDB and Redis to start up before we start the servers)

These are the `PORT` mappings I have done for you

```.env
REDIS      8888
MONGODB    9999
Backend    3000
Job        No Port
```

## Setting up Backend [Manual Setup]

First clone the repo

```bash
git clone https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK.git
```

Now lets understand what is going on.

These are the dependencies:

1. MongoDB Server running
2. Redis Stack Server running

And we need the port numbers for both so check now what ports are they running

The default ports are as follows:

```.env
MongoDB    27017
Redis      6379
```

Now there are 2 servers in the backend to run.

1. `Backend` (`/backend`)
2. `Job` (`/job`)

Both are imdependent jobs and none require the other to be on to startup.

**Running the Backend:**

1. `cd backend`
2. Now set 2 `env` variables with name `ATLAS_URI` and `REDIS_HOST`. I generally run these commands to set it up.

```.env
export ATLAS_URI=mongodb://localhost:27017
export REDIS_HOST=localhost:6379
```

3. Now lets install the dependencies

```bash
go mod install
```

4. Now lets run the server

```
go run .
```

5. The `backend` server should start at PORT `3000`
   <img width="1344" alt="Screenshot 2023-11-19 at 1 34 33 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/8c64d7e0-009e-4870-b363-033d365de525">

**Running the Job**

Go to the root of the project and then follow

1. Lets get in the directory first

```
cd job
```

2. Set the env. Same as of the `backend`. (It should be exactly the same)

```.env
export ATLAS_URI=mongodb://localhost:27017
export REDIS_HOST=localhost:6379
```

3. Now lets install the dependencies

```bash
go mod install
```

4. Start the runner

```
go run .
```

5. Job should be running now

<img width="1344" alt="Screenshot 2023-11-19 at 1 35 29 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/0e3864a6-b345-443a-a8e4-fac47246fba0">

## Setting up Frontend

First go to the root of the project then follow the steps

1. Go into the directory

```
cd frontend
```

2. Install the dependencies. Do note that you should use `yarn`

```
yarn install
```

3. Set up the env variables. Make a `.env` inside the `/frontend` folder (our active directory) and paste exactly this. **Note:** The backend should be running by now.

```.env
VITE_BACKEND=localhost:3000
```

4. Start the server

```
yarn dev
```

5. The server should start at PORT `3333`

## Port Mappings

```.env
Backend    3000
Frontend   3333
```

<!-- USAGE EXAMPLES -->

## Usage

The project is actually very simple to use.

**To insert logs**

```bash
curl -XPOST -H "Content-type: application/json" -d '{
  "level": "error",
  "message": "Failed to Redis",
  "resourceId": "server-1234",
  "timestamp": "2023-09-15T08:00:00Z",
  "traceId": "abc-xyz-123",
  "spanId": "span-456",
  "commit": "5e5342f",
  "metadata": {
    "parentResourceId": "server-0987"
  }
}' 'http://localhost:3000'
```

<img width="1118" alt="Screenshot 2023-11-19 at 1 49 30 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/dc184290-0919-423c-8597-1aa61541aea6">

**Other Endpoints**

```
GET    /        Gets the latest [X] amount of logs
GET    /search  Returns the latest [X] logs with search filters
```

**`GET Latest 10`**

```bash
curl -XGET -H 'page-size: 10' -H 'page-number: 0' -H 'Cache-Control: no-cache' 'http://localhost:3000'
```

<img width="1119" alt="Screenshot 2023-11-19 at 1 59 08 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/74521478-84d7-453c-8e7a-eb809111330a">

**`GET Search Listing`**

There are several options here and these are `query params` which are accepted here

```.json
{
  "level": "",
  "levelRegex": "",
  "message": "",
  "messageRegex": "",
  "resourceId": "",
  "resourceIdRegex": "",
  "timestamp": "",
  "timestampRegex": "",
  "traceId": "",
  "traceIdRegex": "",
  "spanId": "",
  "spanIdRegex": "",
  "commit": "",
  "commitRegex": "",
  "parentResourceId": "",
  "parentResourceIdRegex": "",
  "timeStart": "",
  "timeEnd": "",
  "fullTextSearch": "",
  "pageNumber": 0,
  "pageSize": 60
}
```

**Note**

Here all the above params are `optional` and the type of the params are depicted and should not be changed. Also not that any time parameter needs to be in `ISOString` format eg `2029-11-12T11:45:26.371Z`

Also note that the filters follow a `&` relationship i.e if you apply 2 filters only the data which satisfies bothe the `filters` will be returned

This is a sample curl request

```bash
curl -XGET 'http://localhost:3000?fullTextSearch="db"&level="error"&timeStart="2020-11-21T18:30:00.000Z"'
```

<img width="1112" alt="Screenshot 2023-11-19 at 2 05 38 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/b12f1e2e-0bdb-4d1b-9d6c-575f25d82761">

**Frontend**

Open `http://localhost:3333` in your browser.

![Capture-2023-11-19-004922](https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/5f169dd0-73ea-4317-94fa-56f194970741)

There are 2 tabs.

1. Real Time Logs: This is a list of real time logs. Any filters applied on it will also be real time.
2. Full Text Search: This is not a real time log. You can apply full text search over here but the data will not be updated real time.

Filtering can be done by 2 ways:

1. Server side filtering: This will filter all the logs available and get you the output. Just press on the filter button on the right side of the webpage and the modal will open to set your filters. Here in this modal you can also choose to search via `regex` by clicking on the toggle.
   
<p align="center">
<img height="600" alt="Screenshot 2023-11-19 at 2 05 38 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/bf8ea28e-f831-4d64-b889-740a718c58de">
</p>

2. Client side filtering: This will filter only the data which has been already fetched from the server. This is very fast and can be used to filter data quicky from the available logs.

Click on the 3 dots in the side of the column

<p align="center">
<img width="274" alt="Screenshot 2023-11-19 at 2 14 00 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/38c570f0-43ba-4c7a-b6f7-ca383c72176c">
</p>

Now click on filter

<p align="center">
<img alt="Screenshot 2023-11-19 at 2 14 00 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/39e7461c-3520-4f75-a5df-441d56e45b27">
</p>

Now the column filters will appear and you can apply filters.

<p align="center">
<img alt="Screenshot 2023-11-19 at 2 14 00 AM" src="https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/9657c898-2700-4d9b-b123-5f928fc8b825">
</p>

<!-- Features -->

## Featires

- [x] Log Ingestor
  - [x] Mechanism to ingest logs in the provided format.
  - [x] Ensure scalability to handle high volumes of logs efficiently
  - [x] Mitigate potential bottlenecks such as I/O operations, database write speeds, etc.
  - [x] Logs are ingested via an HTTP server, which runs on port `3000` by default.
- [x] Query Interface (WEB UI)
  - [x] Include filters based on
    - [x] level
    - [x] message
    - [x] resourceId
    - [x] timestamp
    - [x] traceId
    - [x] spanId
    - [x] commit
    - [x] metadata.parentResourceId
  - [x] Efficient and quick search results.
- [x] Extra Features
  - [x] Implement search within specific date ranges.
  - [x] Log count filter to reduce DB load and increase filter flexibility
  - [x] Utilize regular expressions for search.
  - [x] Allow combining multiple filters.
  - [x] Provide real-time log ingestion and searching capabilities.
  - [ ] role-based access to the query interface. [ NOT IMPLEMENTED ]
  - [x] Both Client + Server side filtering
  - [x] Advanced options for client side sorting + column manipulation
  - [x] Advanced caching used for more optimal performance

## System Design

This section will talk about the decisions taken and why have these decisions taken.

### What is **NOT** done

By looking at the problem statement for the first time a very simple architecture comes in mind which is just a client server REST architecture like this

![Untitled-2023-11-19-0242](https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/bc0bf6e2-803b-4995-8f1e-a7f9cc7202df)

These are the flaws of the architecture if we look on a higher level:

1. If there are millions of request for log entry our DB will be a bottleneck and we would have massive costing.
2. As our DB is clogged we would also have worse response times with data entry / fetch.
3. Now suppose our frontend takes the data from the backend. But if there is any entry our data becomes stale and we again need to fetch it. One thing which could be done here is making a poller but thats not optimal.
4. Now suppose you have millions of clients who want the logs with different queries. So you will do million queries which will be again very costly.

### What is done

We have changed the architecture significantly to solve the above mentioned problems. The current architecture looks like this

![Untitled-2023-11-19-0242](https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/53766ec5-f9a1-431b-857a-3cb5859c7aff)

Lets dive a bit deep in this

1. Backend takes all the logs submitted and appends it to a `RedisQueue`. As `Redis` is a `in-memory` store the operations are super fast.
2. Now there is different service called `job` whose job is to empty the `RedisQueue` by pushing the data to `MongoDB` in batch. This minizes the `DB` traffic significantly as millions of round trips and for loops are saved here. This will also decrease the costing of the solution as it significantly decreases the `DB` calls.
3. So now the data ingestion is handled at scale but we also need to show the data real time. So `WebSocket` is chosen here as a medium of communication between frontend and backend to minimize `expensive polling`.
4. The `job` also pushes a messege in `RedisPubSub` for `backend` server to consume.
5. One thing to keep in mind that we also need to support real time filtering of data. So we also need to keep track of which client has which filters applied. So `backend` server mantains this `mapping` of `filter` accross connections.
6. `MongoDB` has been indexed for better performance.
7. Whenever the `backend` server starts it spawns a `goRoutine` whose task is to check the `RedisPubSub` and send updated filtered query output to each connected client. It takes the updated connection list by the `pointer reference` to a global variable in `backend`.
8. Whenver the `backend` receives a `websocket` connection another `goRoutine` is spawn whole only task is to interact with the client and update the `filters` stored according to what the client wants. This allows the `backend` to send `real-time` filtered data.
9. The `backend` also uses a **`custom wrapper on MongoDB`** named **`Mongik`** writen by me iteself ( [link](https://github.com/FrosTiK-SD/mongik) ) which reduces DB calls by caching data and managing cache invalidation and aggregate pipelines all by itself. Thus the performance is improved and the `DB` calls are also reduced significantly. `Mongik` works like this.

![Mongik](https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK/assets/75834423/37ac8940-1f5c-48f1-af46-d22d98a48b93)

Mongik also invalidates cache efficiently to ensure that stale data is not served (Not shown in the diagram)

## Some QnA

**Q. In this architecture also there will be millions of call in the DB when updating the client ?**

The answer is **no**. When there is an update in the DB and there are millions of client `subscribed`, the `goRoutine 1` (described earlier) combines all the `filters` of the connected clients and make a **single DB call** and sends the result to the client which improves performance and reduces DB calls.

**Q. What if the database is down ?**

No problem. As `job` is using `RedisQueue` to push to `DB` it will retry when `DB` is up.

**Q. What if backend is down ?**

If the backend is down then also `job` will continue working to push the remaining `logs` to the `DB` as both are `independent` services.

This type of an architecture is also very scalable and fault tollerant as each component of the architecture can be scaled according to the needs which also makes the architecture cost effective and flexible.

In this way all the pain points are addressed which makes the architecture so scalable.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the `MIT License`. See `LICENSE.txt` for more information.

## Contact

Shubhrajyoti Dey - toshubhrajyotidey@gmail.com

Project Link: https://github.com/dyte-submissions/november-2023-hiring-Shubhrajyoti-Dey-FrosTiK.git

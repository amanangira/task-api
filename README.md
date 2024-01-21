## Introduction
A simple CRUD API set for managing user tasks.

### Libraries being used 
1. Router - [Chi](https://github.com/go-chi/chi)
2. PostgreSQL Client - [sqlx](https://github.com/jmoiron/sqlx)
3. CLI Utility - [Makefile](https://www.gnu.org/software/make/)
4. Database Migrations - [DBMate](https://github.com/amacneil/dbmate), do `dbmate help` to see available options.


### Local setup 
1. Copy `.env.example` to `.env`.
2. Make sure the value of `POSTGRES_MOUNT_PATH` is a valid directory, update value or create directory accordingly. 
3. Install [Docker desktop](https://docs.docker.com/engine/install/).
4. In your working directory run `docker-compose up`, use `-d` flag to detach it from your shell session. This should get your PostgreSQL server running. Check with `docker-compose ps`.
5. Install [make](https://www.gnu.org/software/make/#download).
6. Install DBMate [DBMate](https://github.com/amacneil/dbmate?tab=readme-ov-file#installation).
7. Run migrations, `make migrate`.
8. Run development server, `make run`.


### Important environment variables 
A sample environment variable file is part of the repository `.env.example`. You simply do `cp .env.example .env` to create your own `.env` with defaults. 

1. Development related environment vars.
   - `SERVER_PORT` - to configure the default port on which the dev server would run.
   - `IS_DEBUG` - enables debug logging such as API paths and number of middlewares attached to it. Can be extended to 
   set log level.
2. Database variables - variables prefixed with `POSTGRES` are used by the application and docker-compose to configure
    PostgreSQL service with the desired connection details. 
3. DBMate variables - we are using a library for managing our DB migrations called [DBMate](https://github.com/amacneil/dbmate). You can check about 
    its environments variables on the library page.  

### APIs
#### Endpoints
- POST /api/task has 
- GET /api/task has 
- PATCH /api/task/{taskID} has 
- DELETE /api/task/{taskID} has 
- GET /api/task/{taskID} has 
- GET /health has 

#### Sample Body 
```json
{
    "id": "",
    "title" : "sample",
    "description" : "new description",
    "priority" : "p0",
    "created_at" : "2001-03-24T16:21:21.269Z",
    "updated_at": "2001-03-24T16:21:21.269Z",
    "due_at": "2001-03-24T16:21:21.269Z"
}
```

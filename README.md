### `Momentum Server`

<b>`LIVE`</b>

https://services.momentumchurch.dev/

<b>`CMS (Typescript + React)`</b>

https://services.momentumchurch.dev/dashboard

----

Momentum Server is a Fiber API for the Momentum mobile apps and CMS. It
provides User, Meals, Sermons, Payments, Transactions, Messaging, Notifications,
Events, and Volunteer Services functionality.

The backend uses PostgreSQL with `pgx`/`sqlx` and Goose migrations, Stripe for
payments, and Redis for caching. Every `/api/v1` route requires bearer-token
authentication. API errors use a consistent JSON response with `code` and
`message` fields.

-----

- #### `Go` Dependencies
  - PostgreSQL via `pgx`/`sqlx`
  - Goose for database migrations
  - Redis for caching
  - Fiber
  - Stripe
  - Firebase

- #### `Typescript` Dependencies
  - React
  - TanStack Query
  - TailwindCSS
  - React Router
  - Flowbite React
  - Firebase Authentication

The React CMS is exposed by the backend, and React Router is used for in-app
routing. The backend serves each CMS route and React Router maps the route to
its corresponding page.

The CMS compiler uses TypeScript 7. `@typescript/typescript6` is kept as the
compatibility API for tools that have not added TypeScript 7 support yet.

The CMS follows the same dependency boundaries as the backend:

- `domain/` contains entities and contracts.
- `application/` contains feature services and use cases.
- `infrastructure/` contains HTTP, authentication, and query adapters.
- `interfaces/` contains the React presentation layer.
- `main.tsx` is the CMS composition root, analogous to `main.go`.

## Architecture

### Backend

```text
main.go
├── domain/                         Entities and repository contracts
├── application/                   Use cases, services, configuration, cache
└── infrastructure/
    ├── http/                      Fiber handlers and middleware
    ├── integration/               Stripe, SermonCloud, Planning Center
    ├── messaging/                 Email and Firebase messaging
    └── persistence/               PostgreSQL, Redis, repositories, migrations
```

### CMS

```text
cms/src/
├── domain/                        Entities and repository contracts
├── application/                   Feature services and use cases
├── infrastructure/               HTTP, Firebase, and query adapters
├── interfaces/                    Router, providers, pages, and components
└── main.tsx                       Composition root
```

### Architecture at a glance

Solid arrows show request and runtime flow. Dashed arrows show dependency
direction, implementation, and composition-root wiring.

```mermaid
flowchart TB
    mobile["Momentum mobile apps"]
    browser["CMS browser"]

    subgraph product["Momentum Server"]
        presentation["Interfaces<br/>Fiber handlers + React presentation"]
        application["Application<br/>Use cases + feature services"]
        domain["Domain<br/>Entities + contracts"]
        infrastructure["Infrastructure<br/>Persistence + integrations + adapters"]
        composition["Composition roots<br/>main.go + main.tsx"]
    end

    postgres[("PostgreSQL")]
    redis[("Redis")]
    integrations["Stripe<br/>Firebase<br/>SermonCloud<br/>Planning Center"]

    mobile --> presentation
    browser --> presentation
    presentation --> application
    application --> domain
    infrastructure -. "implements domain contracts" .-> domain
    composition -. "wires dependencies" .-> presentation
    composition -.-> application
    composition -.-> infrastructure
    infrastructure --> postgres
    infrastructure --> redis
    infrastructure --> integrations
```

----

### Database (Neon Postgres)

The application uses PostgreSQL connection settings from `DATABASE_URL`.

- Set `DATABASE_URL` to your PostgreSQL connection URL.
- Goose migrations run automatically during application startup.
- Migration files live in `infrastructure/persistence/db/migrations`.

### Application configuration

Create a `.env` file in the repository root for local development:

```dotenv
ENVIRONMENT=development
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/momentum
REDIS_URL=redis://localhost:6379
API_KEY=replace-with-a-local-development-key
```

Required:

- `DATABASE_URL`
- `REDIS_URL`
- `API_KEY`

Optional:

- `FIREBASE_PROJECT_ID`
- `SERMONCLOUD_URL`
- `STRIPE_SECRET_KEY`
- `STRIPE_PUBLISHABLE_KEY`
- `STRIPE_VERSION`
- `PLANNING_CENTER_URL`
- `PLANNING_CENTER_TOKEN_URL`
- `EMAIL_APP_PASSWORD`

`PORT` defaults to `8080`. Never commit real credentials. The browser-facing
API key is client configuration and is not a secret; backend authorization
remains authoritative.

### Developer commands

The root `Makefile` provides the common local workflows:

```bash
make help
make run
make build
make test
make test-postgres
make swagger
make openapi
make openapi-check
make fmt
make check
make docker-build
```

The Go handler annotations remain the source of API documentation. `make
openapi` regenerates the Swagger 2.0 files used by the existing Swagger UI and
converts the same generated contract into the validated OpenAPI 3 document at
[`docs/openapi.json`](./docs/openapi.json). The OpenAPI document is generated,
not manually edited, and is the future source for client tooling such as CMS
transport type generation.

Build the CMS before starting the Go server:

```bash
cd cms
yarn install --frozen-lockfile
yarn lint
yarn build
cd ..
make run
```

The local dashboard is served at:

http://localhost:8080/dashboard

For CMS development, `yarn dev` watches the frontend build output. The Go
server continues to serve the generated files from `cms/dist`.

### PostgreSQL integration tests

The repository integration tests can run against an ephemeral PostgreSQL
container:

```bash
make test-postgres
```

The script starts PostgreSQL on `127.0.0.1:55432`, runs the Goose migrations and
repository integration tests, then removes the container and its temporary
volume.

----

### [API Documentation](https://services.momentumchurch.dev/swagger/index.html)

The API provides users and the CMS with access to create, edit, retrieve, and
upload content.

Each `/api/v1` route requires a bearer token.

-----

### [Users](https://services.momentumchurch.dev/swagger/index.html#/Users)

Users are allowed to create accounts, retrieve, update, and delete their
information.

------

#### [Get User By Id](https://services.momentumchurch.dev/swagger/index.html#/Users/get_api_v1_users__userId_)

This endpoint allows users to get their own information.

<img width="1430" alt="Screenshot of the get user endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/20bc554d-2b7e-4272-b85f-bb75a121ae15">

------

#### [Post a User](https://services.momentumchurch.dev/swagger/index.html#/Users/post_api_v1_users)

This endpoint allows users to post user information.

<img width="1430" alt="Screenshot of the post user endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/49171553-ae04-4afa-aba3-2aa97b280cf2">

------

#### [Update a User](https://services.momentumchurch.dev/swagger/index.html#/Users/put_api_v1_users)

This endpoint allows users to update their own information.

<img width="1430" alt="Screenshot of the update user endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/44e4ce84-b8dd-4dd0-b0b0-39341a69e868">

------

#### [Delete a User](https://services.momentumchurch.dev/swagger/index.html#/Users/delete_api_v1_users__userId_)

This endpoint allows users to delete their own information.

<img width="1430" alt="Screenshot of the delete user endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ce047e45-0c38-48af-b986-c6010651468b">

-----

### [Transactions](https://services.momentumchurch.dev/swagger/index.html#/Transactions)

Users are allowed to create transactions after making a payment, then view and
delete their transactions.

------

#### [Post a Transaction](https://services.momentumchurch.dev/swagger/index.html#/Transactions/post_api_v1_transactions)

This endpoint allows users to post a transaction.

<img width="1430" alt="Screenshot of the post transaction endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/609455c4-2ffb-4781-95b9-ecb1806e2588">

------

#### [Delete a Transaction](https://services.momentumchurch.dev/swagger/index.html#/Transactions/delete_api_v1_transactions__transactionId_)

This endpoint allows users to delete a transaction.

<img width="1430" alt="Screenshot of the delete transaction endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/cb2ec928-77c5-406e-9e8d-36890c210aa8">

------

#### [Get User Transactions](https://services.momentumchurch.dev/swagger/index.html#/Transactions/get_api_v1_transactions__userId_)

This endpoint allows users to get their own transactions.

<img width="1430" alt="Screenshot of the get transactions endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/1b967c37-7d87-48cf-83e3-08651316951d">

-----

### [Meals](https://services.momentumchurch.dev/swagger/index.html#/Meals)

Users are allowed to organize meals. A meal contains the receiver's contact,
food preferences, allergies, dates to send the meals, and volunteers to bring
the meals.

-----

#### [Get All Meals](https://services.momentumchurch.dev/swagger/index.html#/Meals/get_api_v1_meals)

This endpoint allows users to fetch all meals available.

<img width="1430" alt="Screenshot of the get meals endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/1899cd3f-0813-44c1-8801-b7a50358e846">

-----

#### [Post a Meal](https://services.momentumchurch.dev/swagger/index.html#/Meals/post_api_v1_meals)

This endpoint allows users to post a meal.

<img width="1430" alt="Screenshot of the post meal endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ed153b05-7368-4d46-b805-56490fb31d3f">

------

#### [Post a Meal Participant](https://services.momentumchurch.dev/swagger/index.html#/Meals/post_api_v1_meals_meal)

This endpoint allows users to volunteer as a participant in a meal.

<img width="1430" alt="Screenshot of the post meal participant endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/fb3a12b6-2446-4408-a9ab-a9be5badf3c0">

------

### [Sermons](https://services.momentumchurch.dev/swagger/index.html#/Sermons)

Users are allowed to view sermons. They contain preacher and video playback
metadata.

------

#### [Get Sermons By Page #](https://services.momentumchurch.dev/swagger/index.html#/Sermons/get_api_v1_sermons)

This endpoint allows users to get sermons.

<img width="1430" alt="Screenshot of the get sermons endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/15c9b5bb-aba3-4c1d-b27c-7f09bfd0b582">

-----

### [Payments](https://services.momentumchurch.dev/swagger/index.html#/Payments)

Users are allowed to post payment information to initiate Stripe checkout.

------

#### [Post a Payment](https://services.momentumchurch.dev/swagger/index.html#/Payments/post_api_v1_payments)

This endpoint allows users to post payments.

<img width="1430" alt="Screenshot of the post payment endpoint" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/f4b32fb6-4bd0-47d4-8128-4df7fe4d46d2">

-----

### [Messages](https://services.momentumchurch.dev/swagger/index.html#/Messages)

The CMS uses these endpoints to manage published and staged messages:

```text
GET  /api/v1/messages/admin
GET  /api/v1/messages/unpublished
POST /api/v1/messages
PUT  /api/v1/messages
```

The CMS routes are:

```text
/dashboard
/messages
/notifications
/payments
/users
```

## Production build

The multi-stage [Dockerfile](./Dockerfile) installs the CMS dependencies,
builds the React assets, copies `cms/dist` into the Go image, and compiles the
single production binary.

```bash
make docker-build
```

The generated CMS `dist` directory is intentionally ignored by Git and created
during the Docker build.

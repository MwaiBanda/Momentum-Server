# Momentum Server

Backend and content-management platform for the Momentum Church mobile apps and
website.

[Live API](https://services.momentumchurch.dev/) ·
[CMS dashboard](https://services.momentumchurch.dev/dashboard) ·
[Swagger documentation](https://services.momentumchurch.dev/swagger/index.html)

## Overview

Momentum Server is a Go API backed by PostgreSQL and Redis. It supports the
Momentum product surface across:

- User accounts
- Sermons and events
- Messages and passages
- Meals and volunteer services
- Payments and transactions
- Push notifications

The repository also contains the React CMS used to manage messages,
notifications, and payment data. The CMS is compiled into static assets and
served by the Go application, so the project can be deployed as one container.

## Why this project

This project is an exploration of a maintainable full-stack architecture for a
small product team:

- Direct SQL and Goose migrations instead of an ORM abstraction.
- Explicit dependency boundaries between business logic and integrations.
- A backend composition root that makes runtime dependencies visible.
- A React CMS structured with the same domain, application, infrastructure, and
  interface boundaries as the backend.
- Automated validation through unit tests, integration tests, linting, Swagger
  generation, and production Docker builds.

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

`main.go` is the composition root. It loads configuration, opens dependencies,
runs migrations, constructs repositories and services, and registers HTTP
handlers and routes.

### CMS

```text
cms/src/
├── domain/                        Entities and repository contracts
├── application/                   Feature services and use cases
├── infrastructure/               HTTP, Firebase, and query adapters
├── interfaces/                    Router, providers, pages, and components
└── main.tsx                       Composition root
```

The CMS presentation layer does not call Axios or Firebase directly. Those
details are hidden behind application-facing contracts and infrastructure
adapters.

### System flow

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
    infrastructure -. "implements ports and adapters" .-> application
    composition -. "wires dependencies" .-> presentation
    infrastructure --> postgres
    infrastructure --> redis
    infrastructure --> integrations
```

## Technology

### Backend

- Go 1.26
- Fiber
- PostgreSQL
- `pgx` and `sqlx`
- Goose migrations
- Redis
- Stripe
- Firebase Cloud Messaging
- Swagger

### CMS

- React 19
- TypeScript 7
- TanStack Query 5
- React Router
- Tailwind CSS
- Flowbite React
- Firebase Authentication
- Vite

## Local development

### Prerequisites

- Go 1.26+
- Node.js 24+
- Yarn 1.22+
- PostgreSQL
- Redis
- Docker for integration tests and production image builds

### Configuration

Create a `.env` file in the repository root for local development:

```dotenv
ENVIRONMENT=development
PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/momentum
REDIS_URL=redis://localhost:6379
API_KEY=replace-with-a-local-development-key
```

Optional integrations:

```dotenv
FIREBASE_PROJECT_ID=
SERMONCLOUD_URL=
STRIPE_SECRET_KEY=
STRIPE_PUBLISHABLE_KEY=
STRIPE_VERSION=
PLANNING_CENTER_URL=
PLANNING_CENTER_TOKEN_URL=
EMAIL_APP_PASSWORD=
```

Never commit real credentials. The browser-facing API key is a client
configuration value, not a secret; backend authorization remains authoritative.

### Build and run

Build the CMS first so its assets are available to the Go server:

```bash
cd cms
yarn install --frozen-lockfile
yarn build
cd ..
make run
```

The local dashboard is served at
`http://localhost:8080/dashboard`.

For CMS development, `yarn dev` watches the frontend build output. The Go
server continues to serve the generated files from `cms/dist`.

## Useful commands

Run these from the repository root:

```bash
make help            # Show available commands
make run             # Start the API and dashboard
make build           # Build the Go binary
make test            # Run the Go test suite
make test-postgres   # Run repository integration tests with ephemeral Postgres
make swagger         # Regenerate Swagger documentation
make fmt             # Format Go source files
make check           # Check Git diff whitespace
make docker-build    # Build the production image
```

CMS checks:

```bash
cd cms
yarn lint
yarn build
```

## Testing

The test suite includes:

- Application service tests.
- Repository tests.
- Fiber handler and API contract tests.
- Redis cache tests.
- PostgreSQL integration tests using Docker.
- CMS TypeScript, ESLint, and accessibility checks.

The PostgreSQL integration workflow starts a temporary database on
`127.0.0.1:55432`, applies the Goose migrations, runs repository tests, and
removes the temporary container and volume.

## API surface

All `/api/v1` routes require bearer-token authentication.

| Area | Examples |
| --- | --- |
| Users | Create, retrieve, update, and delete users |
| Messages | Published, unpublished, passages, and notes |
| Transactions | Create, list, and delete transactions |
| Payments | Start Stripe checkout |
| Meals | Create meals and assign volunteers |
| Sermons | List sermon content |
| Events | List Planning Center events |
| Notifications | Send Firebase notifications |

### CMS routes

| Path | Purpose |
| --- | --- |
| `/dashboard` | CMS home |
| `/messages` | Manage published and staged messages |
| `/notifications` | Send user notifications |
| `/payments` | View transactions |
| `/users` | User area |

### API routes

All routes below are mounted under `/api/v1` and require bearer-token
authentication.

| Resource | Routes |
| --- | --- |
| Users | `POST /users`, `GET /users/:userId`, `PUT /users`, `DELETE /users/:userId` |
| Meals | `POST /meals`, `GET /meals`, `PUT /meals`, `DELETE /meals/:mealId` |
| Meal volunteers | `POST /meals/meal`, `PUT /meals/meal` |
| Payments | `POST /payments` |
| Transactions | `GET /transactions`, `GET /transactions/:userId`, `POST /transactions`, `DELETE /transactions/:transactionId` |
| Sermons | `GET /sermons` |
| Volunteer services | `GET /services`, `POST /services`, `GET /services/:type`, `PUT /services/day` |
| Messages | `POST /messages`, `PUT /messages`, `GET /messages/admin`, `GET /messages/unpublished`, `GET /messages/:userId`, `DELETE /messages/:messageId` |
| Message notes | `POST /messages/notes`, `PUT /messages/notes`, `DELETE /messages/notes/:noteId` |
| Notifications | `POST /notifications` |
| Events | `GET /events` |

The CMS message workflow uses:

```text
GET  /api/v1/messages/admin
GET  /api/v1/messages/unpublished
POST /api/v1/messages
PUT  /api/v1/messages
```

Interactive API documentation is available through
[Swagger UI](https://services.momentumchurch.dev/swagger/index.html).

## Production build

The multi-stage [Dockerfile](./Dockerfile) performs the complete production
build:

1. Installs CMS dependencies with the locked Yarn version.
2. Builds the React CMS.
3. Copies `cms/dist` into the Go image.
4. Compiles the Go server.
5. Runs the single production binary.

Build it locally with:

```bash
make docker-build
```

## Project status

The project is actively maintained as a portfolio and learning project. The
current focus is keeping the architecture explicit while improving the
reliability and polish of the CMS workflows.

### `Momentum Server`
<b>`LIVE`</b>
https://services.momentumchurch.dev/

<b>`CMS (TypeScript + React)`</b>
https://services.momentumchurch.dev/dashboard

----
Momentum Server is a Fiber API for the Momentum mobile apps and CMS. It provides user,
meal, sermon, payment, transaction, messaging, notification, event, and volunteer-service
functionality.

The backend uses PostgreSQL with `pgx`/`sqlx` and Goose migrations, Stripe for payments,
and Redis for caching. Every `/api/v1` route requires bearer-token authentication.
API errors use a consistent JSON response with `code` and `message` fields.

-----
- #### `Go` Dependencies
  - PostgreSQL via `pgx`/`sqlx`
  - Goose for database migrations
  - Redis for Caching
  - Fiber
- #### `CMS` Dependencies
    - React
    - React Query
    - TailwindCSS
    - React Router
    
The React CMS is built from `cms/` and its compiled assets are served by the backend.
React Router handles the CMS application routes.

----
### Database (Neon Postgres)

The application uses PostgreSQL connection settings from `DATABASE_URL`.

- Set `DATABASE_URL` to your PostgreSQL connection URL.

Goose migrations run automatically during application startup. Migration files live in
`infrastructure/persistence/db/migrations`.

### Developer commands

The root `Makefile` provides the common local workflows:

```bash
make help
make run
make test
make test-postgres
make swagger
make fmt
make check
```

### Application configuration

Configuration is loaded once at startup through `application/config`.

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

`PORT` defaults to `8080`. Development environments may load values from `.env`; production environments should provide configuration through the process environment.

### PostgreSQL integration tests

The repository integration tests can run against an ephemeral PostgreSQL container:

```bash
make test-postgres
```

The script starts PostgreSQL on `127.0.0.1:55432`, runs the Goose migrations and repository integration tests, then removes the container and its temporary volume.

----
### [API Documentation](https://services.momentumchurch.dev/swagger/index.html)
The interactive Swagger documentation describes the available API endpoints.

-----
### [Users](https://services.momentumchurch.dev/swagger/index.html#/Users)
Users are allowed to create accounts, retrieve, update, and delete their information. 

------

#### [Get User By Id](https://services.momentumchurch.dev/swagger/index.html#/Users/get_api_v1_users__userId_)
This endpoint allows users to get their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 42 15 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/20bc554d-2b7e-4272-b85f-bb75a121ae15">

------

#### [Post a User](https://services.momentumchurch.dev/swagger/index.html#/Users/post_api_v1_users)
This endpoint allows users to post user information
<img width="1430" alt="Screenshot 2023-08-19 at 8 42 26 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/49171553-ae04-4afa-aba3-2aa97b280cf2">

------

#### [Update a User](https://services.momentumchurch.dev/swagger/index.html#/Users/put_api_v1_users)
This endpoint allows users to update their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 40 15 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/44e4ce84-b8dd-4dd0-b0b0-39341a69e868">

------

#### [Delete a User](https://services.momentumchurch.dev/swagger/index.html#/Users/delete_api_v1_users__userId_)
This endpoint allows users to delete their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 41 10 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ce047e45-0c38-48af-b986-c6010651468b">

-----
### [Transactions](https://services.momentumchurch.dev/swagger/index.html#/Transactions)
Users can create transactions after making a payment, then view and delete their transactions.

------

#### [Post a Transaction](https://services.momentumchurch.dev/swagger/index.html#/Transactions/post_api_v1_transactions)
This endpoint allows users to post a transaction
<img width="1430" alt="Screenshot 2023-08-19 at 9 50 05 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/609455c4-2ffb-4781-95b9-ecb1806e2588">

------

#### [Delete a Transaction](https://services.momentumchurch.dev/swagger/index.html#/Transactions/delete_api_v1_transactions__transactionId_)
This endpoint allows users to delete a transaction
<img width="1430" alt="Screenshot 2023-08-19 at 9 50 29 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/cb2ec928-77c5-406e-9e8d-36890c210aa8">

------

#### [Get User Transactions](https://services.momentumchurch.dev/swagger/index.html#/Transactions/get_api_v1_transactions__userId_)
This endpoint allows users to get their own transactions
<img width="1430" alt="Screenshot 2023-08-19 at 9 50 49 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/1b967c37-7d87-48cf-83e3-08651316951d">

-----

### [Meals](https://services.momentumchurch.dev/swagger/index.html#/Meals)
Users are allowed to organize meals, a meal contains the receiver's contact, food preferences/allergies, dates to send the meals, and volunteers to bring meals  

-----

#### [Get All Meals](https://services.momentumchurch.dev/swagger/index.html#/Meals/get_api_v1_meals)
This endpoint allows users to fetch all meals available
<img width="1430" alt="Screenshot 2023-08-19 at 3 15 44 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/1899cd3f-0813-44c1-8801-b7a50358e846">

-----

#### [Post a Meal](https://services.momentumchurch.dev/swagger/index.html#/Meals/post_api_v1_meals)
This endpoint allows users to post a meal
<img width="1430" alt="Screenshot 2023-08-19 at 7 53 48 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ed153b05-7368-4d46-b805-5647f3e704bc">

------

#### [Post a Meal Participant](https://services.momentumchurch.dev/swagger/index.html#/Meals/post_api_v1_meals_meal)
This endpoint allows users to volunteer as a participant in a meal
<img width="1430" alt="Screenshot 2023-08-19 at 8 14 19 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/fb3a12b6-2446-4408-a9ab-a9be5badc3f0">

------


### [Sermons](https://services.momentumchurch.dev/swagger/index.html#/Sermons)
Users are allowed to view sermons, they contain preacher and video playback metadata 

------

#### [Get Sermons By Page #](https://services.momentumchurch.dev/swagger/index.html#/Sermons/get_api_v1_sermons)
This endpoint allows users to get sermons
<img width="1430" alt="Screenshot 2023-08-19 at 10 59 58 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/15c9b5bb-aba3-4c1d-b27c-7f09bfd0b582">



------

### [Payments](https://services.momentumchurch.dev/swagger/index.html#/Payments)
Users can submit payment information to initiate Stripe checkout.

------

#### [Post a Payment](https://services.momentumchurch.dev/swagger/index.html#/Payments/post_api_v1_payments)
This endpoint allows users to post payments
<img width="1430" alt="Screenshot 2023-08-19 at 10 54 21 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/f4b32fb6-4bd0-47d4-8128-4df7fe4d46d2">

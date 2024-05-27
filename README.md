### `Momentum Server`
<b>`LIVE`</b>
https://services.momentumchurch.dev/

<b>`CMS(Typescript + React)`</b>
https://services.momentumchurch.dev/dashboard

----
Backend built with Fiber for Android &amp; iOS apps providing User, Meals, Semons, Payments & Transactions CRUD functionality, 
interacts with Stripe API to facilitate payments, Prisma as ORM to connect to a Postgres instance, and provides CRUD functionality.
With, Redis for caching. Each v1 API route is protected with token authentication for each request sent a bearer token is required.

-----
- #### `Go` Dependencies
  - Prisma ORM connects to Postgres DB
  - Redis for Caching
  - Fiber
- #### `Typescript` Dependencies
    - React
    - React Query
    - TailwindCSS
    - React Router
    
React frontend is exposed in the backend, and react-router is used for in-app routing. Additionally, the backend does expose each in-app route
and react-router handles mapping the route to its corresponding page

----
### [API Documentation](https://services.momentumchurch.dev/)
The API provides users to access, edit and upload content

-----
### [Users](https://services.momentumchurch.dev/index.html#/Users)
Users are allowed to create accounts, retrieve, update, and delete their information. 

------

#### [Get User By Id](https://services.momentumchurch.dev/index.html#/Users/get_api_v1_users__userId_)
This endpoint allows users to get their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 42 15 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/20bc554d-2b7e-4272-b85f-bb75a121ae15">

------

#### [Post a User](https://services.momentumchurch.dev/index.html#/Users/post_api_v1_users)
This endpoint allows users to post user information
<img width="1430" alt="Screenshot 2023-08-19 at 8 42 26 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/49171553-ae04-4afa-aba3-2aa97b280cf2">

------

#### [Update a User](https://services.momentumchurch.dev/index.html#/Users/put_api_v1_users)
This endpoint allows users to update their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 40 15 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/44e4ce84-b8dd-4dd0-b0b0-39341a69e868">

------

#### [Delete a User](https://services.momentumchurch.dev/index.html#/Users/delete_api_v1_users__userId_)
This endpoint allows users to delete their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 41 10 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ce047e45-0c38-48af-b986-c6010651468b">

-----
### [Transactions](https://services.momentumchurch.dev/index.html#/Transactions)
Users are allowed to create transactions after making a payment, they view and delete transac 

------

#### [Post a Transaction](https://services.momentumchurch.dev/index.html#/Transactions/post_api_v1_transactions)
This endpoint allows users to post a transaction
<img width="1430" alt="Screenshot 2023-08-19 at 9 50 05 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/609455c4-2ffb-4781-95b9-ecb1806e2588">

------

#### [Delete a Transaction](https://services.momentumchurch.dev/index.html#/Transactions/delete_api_v1_transactions__transactionId_)
This endpoint allows users to delete a transaction
<img width="1430" alt="Screenshot 2023-08-19 at 9 50 29 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/cb2ec928-77c5-406e-9e8d-36890c210aa8">

------

#### [Get User Transactions](https://services.momentumchurch.dev/index.html#/Transactions/get_api_v1_transactions__userId_)
This endpoint allows users to get their own transactions
<img width="1430" alt="Screenshot 2023-08-19 at 9 50 49 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/1b967c37-7d87-48cf-83e3-08651316951d">

-----

### [Meals](https://services.momentumchurch.dev/index.html#/Meals)
Users are allowed to organize meals, a meal contains the receiver's contact, food preferences/allergies, dates to send the meals, and volunteers to bring meals  

-----

#### [Get All Meals](https://services.momentumchurch.dev/index.html#/Meals/get_api_v1_meals)
This endpoint allows users to fetch all meals available
<img width="1430" alt="Screenshot 2023-08-19 at 3 15 44 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/1899cd3f-0813-44c1-8801-b7a50358e846">

-----

#### [Post a Meal](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals)
This endpoint allows users to post a meal
<img width="1430" alt="Screenshot 2023-08-19 at 7 53 48 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ed153b05-7368-4d46-b805-5647f3e704bc">

------

#### [Post a Meal Participant](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals_meal)
This endpoint allows users to volunteer as a participant in a meal
<img width="1430" alt="Screenshot 2023-08-19 at 8 14 19 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/fb3a12b6-2446-4408-a9ab-a9be5badc3f0">

------


### [Sermons](https://services.momentumchurch.dev/index.html#/Sermons)
Users are allowed to view sermons, they contain preacher and video playback metadata 

------

#### [Get Sermons By Page #](https://services.momentumchurch.dev/index.html#/Sermons/get_api_v1_sermons)
This endpoint allows users to get sermons
<img width="1430" alt="Screenshot 2023-08-19 at 10 59 58 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/15c9b5bb-aba3-4c1d-b27c-7f09bfd0b582">



------

### [Payments](https://services.momentumchurch.dev/index.html#/Payments)
Users are allowed to post payment information to initiate stripe checkoot

------

#### [Post a Payment](https://services.momentumchurch.dev/index.html#/Payments/post_api_v1_payments)
This endpoint allows users to post payments
<img width="1430" alt="Screenshot 2023-08-19 at 10 54 21 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/f4b32fb6-4bd0-47d4-8128-4df7fe4d46d2">


### Momentum Server
----
Backend built with Fiber for Android &amp; iOS apps providing User, Meals, Semons, Payments & Transactions CRUD functionality, 
interacts with Stripe API to facilitate payments, Prisma as ORM to connect to a Postgres instance, and provides CRUD functionality.
With, Redis for caching

-----
- Postgres DB
- Prisma ORM
- Redis for Caching
- Fiber 
----
### [API Documentation](https://services.momentumchurch.dev/)
The API provides users to access, edit and upload content

-----
### [Users](https://services.momentumchurch.dev/index.html#/Meals)
Users are allowed to create accounts, retrieve, update, and delete their information. 

------

#### [Get User By Id](https://services.momentumchurch.dev/index.html#/Meals/get_api_v1_meals)
This endpoint allows users to get their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 42 15 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/20bc554d-2b7e-4272-b85f-bb75a121ae15">

------

#### [Post a User](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals)
This endpoint allows users to post user information
<img width="1430" alt="Screenshot 2023-08-19 at 8 42 26 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/49171553-ae04-4afa-aba3-2aa97b280cf2">

------

#### [Update a User](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals)
This endpoint allows users to update their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 40 15 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/44e4ce84-b8dd-4dd0-b0b0-39341a69e868">

------

#### [Delete a User](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals_meal)
This endpoint allows users to delete their own information
<img width="1430" alt="Screenshot 2023-08-19 at 8 41 10 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ce047e45-0c38-48af-b986-c6010651468b">

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

-----

#### [Post a Meal Participant](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals_meal)
This endpoint allows users to volunteer as a participant to a meal
<img width="1430" alt="Screenshot 2023-08-19 at 8 14 19 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/fb3a12b6-2446-4408-a9ab-a9be5badc3f0">

-----


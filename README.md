## Momentum-Server
Backend built with Fiber for Android &amp; iOS apps providing User, Meals, Semons, Payments & Transactions CRUD functionality, 
interacts with Stripe API to facilitate payments, Prisma as ORM to connect to a Postgres instance, and provides CRUD functionality.
With, Redis for caching

### Stack 
- Postgres DB
- Prisma ORM
- Redis for Caching
- Fiber 
----
### [API Documentation](https://services.momentumchurch.dev/)
The API provides users to access, edit and upload content
### [Meals](https://services.momentumchurch.dev/index.html#/Meals)
Users are allowed to organize meals, a meal contains the receiver's contact, food preferences/allergies, dates to send the meals, and volunteers to bring meals  
#### [Get All Meals](https://services.momentumchurch.dev/index.html#/Meals/get_api_v1_meals)
This endpoint allows users to fetch all meals available
<img width="1430" alt="Screenshot 2023-08-19 at 3 15 44 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/1899cd3f-0813-44c1-8801-b7a50358e846">
#### [Post a Meal](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals)
This endpoint allows users to post a meal
<img width="1430" alt="Screenshot 2023-08-19 at 7 53 48 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/ed153b05-7368-4d46-b805-5647f3e704bc">
#### [Post a Meal Participant](https://services.momentumchurch.dev/index.html#/Meals/post_api_v1_meals_meal)
This endpoint allows users to volunteer as a participant to a meal
<img width="1430" alt="Screenshot 2023-08-19 at 8 14 19 PM" src="https://github.com/MwaiBanda/Momentum-Server/assets/49708426/fb3a12b6-2446-4408-a9ab-a9be5badc3f0">

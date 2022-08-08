# Property Finder Golang Bootcamp Final Project

## Introduction
&nbsp;&nbsp;&nbsp;&nbsp;This project is a basket service which will be a Go-based REST API. In this project users can add, drop, update, remove products in basket.Existing items in basket will be available for purchase by customers (Payment mechanism not included project). 
Records of paid orders are kept and discounts are provided for subsequent orders. 
The tax rates of the products in the Project are three types as 1%, 8% and 18%. 
There are a total of 10 products as representative in the project. 
These have been added to the project as SQL databases. 
User interface is not included in the project. 
`Mozilla Framework` was used while developing the project. Also `MySQL` are used in project.
In addition, Golang for business logic and `Postman` for unit testing and REST API tests were used.

### Assumptions
1. All purchased products were purchased within the same month.
2. Monthly limit, shopping series discount and discounted products are not fixed and may change over time.


---

## Functions
* **Create User**
  * You need to create a membership to be able to make any transactions in the market. Membership process is kept simple in the project. Only username is required to register.
* **List All User**
  * Show all users and users information in the database
* **Set Active User**
  * After the membership is created, you need to activate your membership in order to be able to take action. This function can be thought of as a login function. This process was also kept simple in the project.
* **List Product**
  * List all products in market
* **Add Product**
  * Add product to basket and total price are changed according to products which is in basket
* **Drop Product**
  * Decreases the number of items in the basket by one.
* **Delete Product**
  * It completely removes the product in the basket.
* **Show Order**
  * Shows all items in the cart, the total amount of the order, the tax amount, the discount amount and the amount due
* **Show All Past Orders**
  * By entering the user ID, it displays all the orders made by the user together with the contents of the basket and the receipt.

---
## Business Logic
* **Multiple Product Discount**
  * If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.
* **Complete Monthly Limit Discount**
  * If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off. The monthly limit has been determined as 500 units in the project.
* **Shopping Streak Discount**
  * Every fourth order whose total is more than given amount which is determined 100 unit in the project may have discount  depending on products. Products whose VAT is %1 don’t have any discount
    but products whose VAT (Value Added Tax) is %8 and %18 have discount of %10 and %15
    respectively
* **Result**
  * Only one discount can be applied at a time. Only the highest discount should be applied.

## Conclusion
&nbsp;&nbsp;&nbsp;&nbsp;While developing the project, attention was paid to coupling and cohesion. Attempted to apply clean code and clean architecture principles.
Implemented well-known REST API patterns. Created their own errors for error handling in project.
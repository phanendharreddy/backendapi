## Authorization and Authentication service in Golang

## The problem 
The problem at hand is to create a backend API service in Golang that handles authorization and authentication for a web app. The web app is a simple platform where users in an organization can sign in and list all other users in their organization. Login is performed by supplying a username, password combination, with all passwords hashed when stored in a database for security purposes. The user should be logged in with a JWT token that expires after one hour, and they should be able to receive a new access token using a 'Refresh token' with a validity of 24 hours. The user should also be able to log out. There are admin privileges assigned to a few users, giving them the ability to add new user accounts or delete existing user accounts from their organization. Non-admin users should be able to see other user accounts but shouldn't be able to add/delete any user accounts. Additionally, any user shouldn't be able to view/add/delete user accounts into any other organization. The API should follow REST API conventions, and the API should cover functionalities like user login, user logout, admin user adding a new user account, admin user deleting an existing user account from their organization, and listing all users in their organization. Unit tests should be added for each API endpoint.

## The solution 
By using Gin framework from Golang Language. I've developed this APIs for Login, Logout, CreateUser, GetUsers and DeleteUser functions. By invoking this APIs in the end user application can be implemented a optimized application. 

## Goals 
With the help of API developed in this project. It'll be easier to implement a application by using this project.

## Implementation design 
Lay out the plan and how to execute it. It should match the achieved goals stated in your proposal. List all known dependencies.

## Risks & Mitigation 
List risks that may prevent you from successfully completing your project and your mitigation for each risk. Include other commitments you have during the program (job, internship, etc).

## About me 

I've recently graduated from computer science and engineering stream. My computational skills will help this project to complete effectively. I've worked with academic projects and prosonnel projects as well and then this is the first opportunity with open source organizations. My extensive background in computer science, combined with my passion for open source projects, gives me a unique perspective that I believe would be valuable to your organization.

## Contact information 

Full Name : Phanendhar Reddy Kusuma
Email Id : phanendharreddykusuma@gmail.com
Mobile Number : 7729026081
Languages Known : English, Telugu and Hindi.
GitHub : phanendharreddy
LinkedIn : phanendharreddy
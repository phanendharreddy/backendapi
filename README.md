## Authorization and Authentication service in Golang

## The problem 
The problem at hand is to create a backend API service in Golang that handles authorization and authentication for a web app. The web app is a simple platform where users in an organization can sign in and list all other users in their organization. Login is performed by supplying a username, password combination, with all passwords hashed when stored in a database for security purposes. The user should be logged in with a JWT token that expires after one hour, and they should be able to receive a new access token using a 'Refresh token' with a validity of 24 hours. The user should also be able to log out. There are admin privileges assigned to a few users, giving them the ability to add new user accounts or delete existing user accounts from their organization. Non-admin users should be able to see other user accounts but shouldn't be able to add/delete any user accounts. Additionally, any user shouldn't be able to view/add/delete user accounts into any other organization. The API should follow REST API conventions, and the API should cover functionalities like user login, user logout, admin user adding a new user account, admin user deleting an existing user account from their organization, and listing all users in their organization.

## Goals 
API Endpoints
The API follows REST API conventions and covers the following functionalities:

    POST /users/login: User Login
    POST /users/logout: User Logout
    POST /users/createuser: A new User account(by providing the First_name, Last_name, Password, Email, Phone and User_type)
    DELETE /users/:id: Admin User deletes an existing User account from their organization
    GET /users: List all Users in their organization

## Requirements

To run this API, you will need the following:

    Golang framework (version >= 1.16)
    MongoDB (version == 6.0.5)

## Implementation design 
To build the project, you first need to define the different parts of the project, such as the API endpoints and the data models. Once you have a clear idea of what the project will look like, you can start building it using the Gin web framework. Gin makes building web applications easier by providing pre-built tools for handling things like HTTP requests and routing. To keep user data secure, you will also use a few special tools, like JWT and Bcrypt. These tools will help to generate secure passwords and encrypt user data, so it can't be easily accessed by unauthorized users. After building and testing your code, you can then deploy your project to a web server so that other people can access it online. With your project live, you'll be able to see how people use it and make any necessary adjustments to improve it over time.

## Instructions to setup

1. First, install the latest version of Golang and MongoDB.

2. Then clone this repository using the command:

    git clone https://github.com/HousewareHQ/houseware---backend-engineering-octernship-phanendharreddy

3. Navigate to the cloned repository directory by using the command:

    cd houseware---backend-engineering-octernship-phanendharreddy

4. Install the dependencies required for the project by using the command: 

    go mod download

5. After installing the dependencies, run the project using the command: 

    go run main.go

6. The project should now be running and accessible at http://localhost:9000.


## Timeline
I'm glad to say that this assignment was completed in less than a week. The first few days were spent exploring and familiarizing myself with Golang and its libraries, as well as testing out different database options. By mid-week, I had settled on MongoDB as the database for the project and began implementing the login and logout functionalities, followed by creating and deleting users. By the end of the week, I had completed the main functionalities of the assignment.

## Risks and mitigation
In the process of completing this assignment, I first ensured that I had all the necessary requirements, including the latest version of Golang and MongoDB, as well as Postman for testing. I referred to the Golang documentation to familiarize myself with the methods and objects required to write the code, and I also identified the libraries and dependencies that would be necessary for the project.

One of the main challenges I faced was connecting to the database using PostGreSQL, as there were issues with the URL and some of the methods were deprecated. To mitigate this risk, I researched other database options and found that MongoDB would be a better fit for this assignment. I was able to quickly incorporate the necessary MongoDB methods into the code, which allowed the Login() and Logout() POST methods to work as intended.

Once logged in, users are given a token that allows them to perform certain operations, such as getusers, createuser, and deleteuser. The database is closed after a user logs out, which restricts access to these operations. I initially thought that the GetUsers() function was restricted to admin users only, but upon further review of the problem statement, I realized this was not the case. I removed the admin requirement in the code, and users can now call the GetUsers() function by passing their token.

The DeleteUser() operation is restricted to admin users only, meaning regular users cannot delete other users. This functionality is working correctly. However, I encountered a challenge with the CreateUser() function, which was supposed to be restricted to admin users only. Although I attempted to incorporate this restriction into the code, it did not work as expected. Therefore, currently, all users can perform the CreateUser() operation.

## About me 
I've recently graduated from computer science and engineering stream. My computational skills will help this project to complete effectively. I put my knowledge to the test on several projects. <br>
My first academic project was - “Cyberbullying Detection with Deep Learning.” We created this project keeping in mind that people have started spending more time on social networking sites where cyberbullying has become a societal issue that must be tackled utilising machine learning approaches. To identify cyberbullying in social media content, we propose employing a Char-CNN (Character-level Convolutional Neural Network) model. As a part of the project team, I was in charge of creating Python code for characteristics. We used the Spyder IDE from the anaconda package to write the code. I assisted with project documentation also.  <br>
My second academic project was - “[Extension of Lexicon Algorithm for Sarcasm Detection](https://github.com/phanendharreddy/sarcasamDetection).” In this project, we extended the lexicon algorithm to develop more efficient systems for sarcasm detection. We used Jupyter Notebook because it is ideal for data exploration and sentiment analysis. I played a key role in developing the existing algorithms, and I also managed the project's repository with Git and GitHub, which are version control systems.<br>
 For personal projects, my first project was - “[Weather Web application](https://github.com/phanendharreddy/weatherApp).” Here, I’ve created a weather web application that pulls data from the Forecast API and displays the weather information. I wrote the structure in HTML and used JavaScript to connect it to the Forecast API. <br>
My second personal project was - “[Private Browser](https://github.com/phanendharreddy/ownBrowser).” Using Python and PyQt5, I created a private web browser that enables users to surf the web without leaving a record of their behaviour on the device. Tabbed browsing, back/forward navigation, page refreshing, and URL loading are some of its features. Moreover, it renders online pages using the QtWebEngine module and disables local storage to stop websites from collecting user data. I use this programme for very secure financial functions. Also, this programme is used by my personal family and a handful of my acquaintances. <br>
I'm thrilled to hear about Houseware's vision to empower the next generation of knowledge workers by putting the data warehouse in their hands, in the language they speak. As a computer science student with a passion for open source projects and extensive experience in backend development, I believe that Houseware's mission is perfectly aligned with my interests and skill set. I'm excited about the opportunity to work as a Backend engineer with the Houseware team on multiple customer-facing projects, especially with the focus on technical architecture and backend engineering. I'm confident that I can own the backend and infrastructure stack end-to-end, understand the business use cases, map it to the best-in-class engineering systems, and maintain a great developer experience. Now, I'm familiarized in Golang and Python and have prior experience building backend systems, along with hands-on experience with AWS/GCP. I'm very interested in this GitHub Octernship Program and believe that working with Houseware would be an incredible opportunity to gain exposure to open source projects and contribute my knowledge towards an organization with a vision I am passionate about. The monthly stipend of $600 USD would also provide valuable financial support during the Octernship period. Overall, I am excited about the possibility of joining the Houseware team and contributing to the development of a forward-looking product that empowers the next generation of knowledge workers.
## Contact information 

Full Name           : Phanendhar Reddy Kusuma <br>
E-mail              : phanendharreddykusuma@gmail.com <br>
Mobile Number       : 7729026081 <br> 
I can speak  : English, Telugu and Hindi. <br>
[GitHub](https://www.github.com/phanendharreddy/) <br>
[LinkedIn](https://www.linkedin.com/in/phanendharreddy/) <br>
[Twitter](https://twitter.com/phanendharr_ddy)
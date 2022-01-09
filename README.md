# KPI Dashboard Application
Go application featuring authentication, jwt authorization, RESTful APIs.

Related project's React based frontend application is accesible via URL below:

[github.com/EnesToraman/React-KPI-Dashboard](https://github.com/EnesToraman/React-KPI-Dashboard)

### Project Overview

- Dashboard application for use of people who would evaluate KPIs of an airport operations. 
- After logging in, a jwt token is created to make requests with. HttpOnly cookies are used for carrying and storing the token.
- People who use this application fall under two roles: MANAGER and STAFF. Managers are able to view more graphs due to role authorization.
- While signing up, any user is signed up as STAFF. In order to make any changes on role, database administrator needs to be contacted.
- Graphs include KPIs like Revenue per Day, Average Ticket Price per Day etc. 

### Database Schema:

![Database Schema](/sql-script/database-schema.png)

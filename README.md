# JSON CRUD API in Go

## Routes
| Route | Method | Description|
| ------ | ------ | ------ |
| API_URL/ | GET | List all the posts registered |
| API_URL/post/:id | GET | List a specific post by it's ID |
| API_URL/post | POST | Saved a new posts in the database |
| API_URL/post/:id | PATCH | Updated a specific posts by it's ID |
| API_URL/post/:id | DELETE | "Delete" a post |


## Tech
Technologies used for the construction of this api are::

- [GO](https://go.dev/) - Programming language.
- [Gorm](https://gorm.io/) - ORM to facilitate database connection.
- [Gin](https://gin-gonic.com/) - Gin is a web framework written in Go.
- [Postgres](https://www.postgresql.org/) - Open source object-oriented relational database management system.
# Recipes API
Recipes API with basic CRUD operations developed in Go. Deployed in Heroku: https://recipes-api-go.herokuapp.com/

The application uses PostgreSQL as the database, and Golang libraries mux and gorm, to handle object-relational mpaping and the http router respectively.

## Models
Located in the `models` directory are the models for the entities that the application uses.
### Recipe
```
{
    "ID":           int
    "CreatedAt":    time
    "UpdatedAt":    time
    "DeletedAt":    time
    "name":         string
    "description":  string
    "image":        string
    "requiredTime": int
    "difficulty":   int
    "servings":     int
    "steps":        string
    "ingredients": []Ingredient
}
```
The attributes `ID`, `CreatedAt`, `UpdatedAt`, and `DeletedAt` are automatically generated when a recipe is inserted to the DB. The primary key for this model is the `ID`.

### Ingredient
```
{
    "CreatedAt":    time
    "UpdatedAt":    time
    "name":         string
    "type":         string
}
```
The attributes `CreatedAt`, `UpdatedAt`, and `DeletedAt` are automatically generated when a recipe is inserted to the DB. The primary key for this model is the `name`.

## Routes
### Ingredients
|  Verb | Route  |
|---|---|
|  GET |  `/api/ingredients` |
|  GET |  `/api/ingredient/:name` |
|  POST | `/api/ingredient` |
|  PUT |  `/api/ingredient` |
|  DELETE |  `/api/ingredient/:name` |
Both `POST` amd `PUT` operations require an ingredient object be sent as the body of the request.

### Recipes
|  Verb | Route  |
|---|---|
|  GET |  `/api/recipes` |
|  GET |  `/api/recipe/:id` |
|  POST | `/api/recipe` |
|  PUT |  `/api/recipe` |
|  DELETE |  `/api/recipe/:id` |
Both `POST` amd `PUT` operations require a recipe object be sent as the body of the request.

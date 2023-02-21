# Item-catalog
## Purpose
This application is one in muliple that combined completes a auction service. The Item-catalog handles the CRUD operations of any items added to the auction service, and exposes this functionality through rest api.

## TechStack
The basic layout for this application is as follows:
    
    mongodb -> go-chi-restapi -> golang-apigateway (TODO) -> angular UI (TODO)


## Development
This application tries to its best ability to follow a test first based development. That means that you should first create a basic test that can confirm that you new functionality works as intenden. See Test section.

There main components that complete this application is:


1. Every new entity feature that needs to be added should first and foremost created in models folder. Example, I want to create a feature regarding books. Then i need to create a book entity in the models folder. The entitys in the models folder serves two purposes. 1. It must contain the entity structure. 2. It should only handle the basic operations for that entity such as fetching, adding, updating towards a db or other third party services. 

2. Router folder. This folder should only contain the entity routes and naming should be based on the entity that you are creating. Example: I want to create a book enity, then then router-file would be book-routes.

3. Every router file should have a controller in controllers folder that handles any business logic needed and call the underlying repository/entity-model. The router should call the controller.
## Testing

To run the tests simply run:
    
    go test -v







# MedPortal Project

## Technology Stack

- **Backend:** Golang with Gin framework  
- **Database:** PostgreSQL (using Neon provider)  
- **Frontend:** React with TypeScript  

## Design

The project follows a **Repository Design Pattern** with clear separation of concerns across layers:

- **Repository:** Responsible for interacting with the database  
- **Handler:** Manages API requests and responses  
- **Service:** Contains business logic and ties the repository and handlers together  

### Authentication

- Uses **JWT tokens** that store user information, crucial for role determination  
- Integrated with **AuthMiddleware** to allow route access based on specific user roles or shared access  

## Testing

*(Add testing details here if available)*

## Documentation

- API documentation is available via **Swagger**, providing an interactive overview of all routes  
- Additional project details are documented in this README file  

## Video Link

*(Add video link here)*


# MedPortal Project
## Deployment Link

*https://diligent-perception-production.up.railway.app/*

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

- API documentation is available at [API Documentation](https://solar-flare-912061.postman.co/workspace/My-Workspace~4c0d0c05-bc59-4f02-876f-ff184cd7c804/collection/29570289-8c567590-464f-49dc-bce7-7832f5f246fb?action=share&source=copy-link&creator=29570289) via Postman 
- Additional project details are documented in this README file  

## Video Link

*(Add video link here)*



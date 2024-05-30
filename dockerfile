# Use the official MongoDB image from Docker Hub
FROM mongo:latest

# Set environment variables
ENV MONGO_INITDB_ROOT_USERNAME=admin
ENV MONGO_INITDB_ROOT_PASSWORD=password

# Expose the default MongoDB port
EXPOSE 27017

# Define the command to run MongoDB
CMD ["mongod"]


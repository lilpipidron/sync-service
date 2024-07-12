# sync-service

# User Guide: Managing Clients and Algorithm Status

## Endpoints

### 1. Add Client

- **URL:** `/addClient`
- **Method:** POST
- **Description:** This endpoint allows adding a new client to the system with the specified details.
- **Request Body:**

```yaml
{
  "client_name": "Client Name",
  "version": 1,
  "image": "client-image",
  "cpu": "1", "memory": "1Gi",
  "need_restart": false
}
```

**Response Codes:**

- `201 Created`: Client successfully added.
- `400 Bad Request`: Invalid request body format.
- `500 Internal Server Error`: Failed to add client.

### 2. Update Client

- **URL:** `/updateClient`
- **Method:** PUT
- **Description:** This endpoint enables updating an existing client with the specified details.
- **Request Body:**

```yaml
{
  "id": 1,
  "client_name": "Updated Client Name",
  "version": 2,
  "image": "updated-image",
  "cpu": "2",
  "memory": "2Gi",
  "need_restart": true
}
```

**Response Codes:**
- `200 OK`: Client successfully updated.
- `400 Bad Request`: Invalid request body format.
- `500 Internal Server Error`: Failed to update client.

### 3. Delete Client

- **URL:** `/deleteClient/{id}`
- **Method:** DELETE
- **Description:** This endpoint deletes a client identified by `{id}`.
- **Response Codes:**
    - `204 No Content`: Client successfully deleted.
    - `404 Not Found`: Client with the specified ID not found.
    - `500 Internal Server Error`: Failed to delete client.

### 4. Update Algorithm Status

- **URL:** `/updateAlgorithmStatus`
- **Method:** PUT
- **Description:** This endpoint updates the status of an algorithm based on the provided request.
- **Request Body:**

```yaml
{
  "id": 1,
  "client_id": 1,
  "vwap": true,
  "twap": false,
  "hft": true
}
```
**Response Codes:**
- `200 OK`: Algorithm status successfully updated.
- `400 Bad Request`: Invalid request body format.
- `404 Not Found`: Algorithm status with the specified ID not found.
- `500 Internal Server Error`: Failed to update algorithm status.
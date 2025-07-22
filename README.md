# RaribleAPI Test Task

**Author:** [Daniil Herasymenko](https://github.com/DanHerasymenko)

**Default swagger link:** [http://localhost:8081/swagger/index.html#/](http://localhost:8081/swagger/index.html#/)

---

## Technologies Used

- Golang 1.24.0
- Web Framework: Gin
- Swagger: swaggo/gin-swagger
- Docker
- External API: [Rarible API](https://api.rarible.org)

---

## Run with Docker

1. Clone the repository.
2. Create a `.env` file in the root directory using the template below.
3. Put your Rarible API key in `RARIBLE_API_KEY` in the `.env` file.
4. Build and run the application:

```bash
docker build -t rarible-api .
docker run --name rarible-api -p 8081:8081 rarible-api
```

---

## Example `.env` File

```env
# Application
APP_PORT=:8081

# Rarible API
RARIBLE_API_URL=https://api.rarible.org/v0.1
RARIBLE_API_KEY=your_key_here
```

---

## Implemented Endpoints

| Method | Endpoint                                 | Description                                 |
|--------|------------------------------------------|---------------------------------------------|
| GET    | `/api/rarible/ownerships/{id}`           | Get NFT ownership by ID                     |
| POST   | `/api/rarible/traits/rarity`             | Get NFT traits rarity                       |

---

## Example of GET request parameter to get list by NFT id

**Request:**
```
Ownership ID: ETHEREUM:0x60e4d786628fea6478f785a6d7e704777c86a7c6
```

---

## Example of POST request body to get NFT rarity

**Body:**
```json
{
  "collectionId": "ETHEREUM:0x60e4d786628fea6478f785a6d7e704777c86a7c6",
  "properties": [
    {
      "key": "Hat",
      "value": "Halo"
    },
    {
      "key": "Background",
      "value": "Blue"
    },
    {
      "key": "Eyes",
      "value": "Laser"
    }
  ]
}
```
---
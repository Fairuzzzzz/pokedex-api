# Pokedex API
A RESTful API service for managing pokemon teams using data from [PokeAPI](https://pokeapi.co/).

## Features
- User authentication (signup & login)
- Pokemon search
- Team management (create, delete, get details)
- Pokemon management in teams (add, remove, list)

## Installation & Setup

### Steps
1. Clone the repository
```bash
git clone https://github.com/yourusername/pokedex-api.git
cd pokedex-api
```

2. Install dependencies
```bash
go mod download
```

3. Configure database
- Create PostgreSQL database
- Update database configuration in `internal/configs/config.yaml`

4. Run the application
```bash
go run cmd/main.go
```

The server will start on port 9999 by default.

## API Endpoints

### Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/memberships/sign-up` | Register new user |
| POST | `/memberships/login` | Login user |

### Pokemon
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/pokemon/:name` | Search pokemon by name |

### Team Management
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/team/create-team` | Create new pokemon team |
| GET | `/team/list-team` | Get list of teams |
| GET | `/team/get-team` | Get team details |
| POST | `/team/delete-team` | Delete a team |

### Pokemon in Team
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/pokemon-team/list/:teamID` | List pokemon in a team |
| POST | `/pokemon-team/add` | Add pokemon to team |
| POST | `/pokemon-team/remove` | Remove pokemon from team |

## Usage Examples

### Register New User
```bash
curl -X POST http://localhost:9999/memberships/sign-up \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "username": "username",
    "password": "password123"
}'
```

### Login
```bash
curl -X POST http://localhost:9999/memberships/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "password123"
}'
```

### Search Pokemon
```bash
curl -X GET http://localhost:9999/pokemon/pikachu \
-H "Authorization: <access_token>"
```

### Create Team
```bash
curl -X POST http://localhost:9999/team/create-team \
-H "Authorization: <access_token>" \
-H "Content-Type: application/json" \
-d '{
    "teamName": "My Team"
}'
```

### Add Pokemon to Team
```bash
curl -X POST http://localhost:9999/pokemon-team/add \
-H "Authorization: <access_token>" \
-H "Content-Type: application/json" \
-d '{
    "teamID": 1,
    "pokemonName": "pikachu"
}'
```

## Notes
- Maximum 6 pokemon per team
- Token expires in 10 minutes
- All endpoints except signup/login require authentication token

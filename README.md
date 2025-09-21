# Skyliner

A full-stack flight booking platform with Travelers, Agents, and Admins. Built with SvelteKit, Go, PostgreSQL, and Stripe.

## Features

- ✈️ **Flight Search**: One-way, round-trip, and multi-city flight searches
- 💺 **Live Seat Maps**: Real-time seat availability and selection
- 🧾 **Booking Management**: Create, manage, and track flight bookings
- 💳 **Secure Payments**: Stripe Checkout and Billing Portal integration
- 🔔 **Real-time Updates**: Live price changes and seat availability
- 👩‍💼 **Agent Console**: Tools for travel agents and administrators
- 🌙 **Dark Mode**: Toggle between light and dark themes
- 📱 **Responsive Design**: Works seamlessly on desktop and mobile
- 🔐 **Secure Authentication**: JWT-based authentication with role management

## Tech Stack

### Backend
- **Go 1.22+** with Gin framework
- **TypeScript** for type safety
- **PostgreSQL** database with GORM ORM
- **Redis** for caching and sessions
- **JWT** authentication with refresh tokens
- **WebSockets** for real-time features
- **Stripe** for payment processing
- **Jest** for testing

### Frontend
- **SvelteKit** with TypeScript
- **Vite** for fast development
- **Tailwind CSS** for styling
- **Svelte Router** for navigation
- **@tanstack/svelte-query** for data fetching
- **Svelte stores** for state management
- **Recharts** for data visualization
- **Vitest** for testing

### Prerequisites
- Go 1.22+
- Node.js 18+
- Docker and Docker Compose
- PostgreSQL 17+ (if not using Docker)
- Redis 7+ (if not using Docker)
- Git

### 1. Clone the repository
```bash
git clone https://github.com/ttiimmothy/skyliner
cd skyliner
```

### 2. For development

#### 1. Start the database
```bash
docker compose up -d postgres redis
```

#### 2. Set up the backend
```bash
cd server
go mod download
cp env.example .env
# Edit .env with your database URL and JWT secret
make migrate
make seed
make dev
```

#### 3. Set up the frontend
```bash
cd frontend
npm install
cp env.example .env
# Edit .env with your API URL
npm run dev -- --port 5193
```

#### 4. Access the application
- Frontend: http://localhost:5193
- Backend API: http://localhost:8080/api
- Database: localhost:5432
- Redis: localhost:6379

## Demo Credentials

The application comes with seeded demo data. You can log in with:
- **Email**: traveler@example.com (or agent@example.com, admin@example.com)
- **Password**: password123

## Project Structure

```
skyliner/
├── frontend/                 # SvelteKit application
│   ├── Dockerfile
│   ├── package.json
│   ├── src/
│   │   ├── components/       # UI components
│   │   ├── routes/           # Page components
│   │   ├── lib/              # API client and stores
│   │   └── app.css           # Global styles
│   ├── vite.config.ts
│   └── tailwind.config.js
├── server/                   # Go API server and database layer
│   ├── Dockerfile
│   ├── go.mod
│   ├── Makefile
│   └── internal/
│       ├── cmd/server/       # Application entry point
│       ├── http/             # HTTP handlers and middleware
│       ├── db/               # Database models and migrations
│       ├── auth/             # Authentication logic
│       ├── payments/         # Stripe integration
│       ├── ws/               # WebSocket hub
│       └── tests/            # Backend tests
├── docker-compose.yml        # Development environment
├── test-setup.sh            # Setup verification script
└── README.md
```

## API Endpoints

### Authentication
- `POST /api/v1/auth/signup` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout
- `POST /api/v1/auth/refresh` - Refresh access token

### Search
- `GET /api/v1/airports` - Get airports
- `GET /api/v1/airlines` - Get airlines
- `POST /api/v1/search` - Search flights
- `GET /api/v1/flights/:id/seatmap` - Get seat map

### Bookings
- `POST /api/v1/bookings` - Create booking
- `GET /api/v1/bookings/:id` - Get booking
- `POST /api/v1/bookings/:id/issue` - Issue booking
- `POST /api/v1/bookings/:id/cancel` - Cancel booking

### Payments
- `POST /api/v1/payments/checkout-session` - Create checkout session
- `POST /api/v1/payments/billing-portal` - Create billing portal
- `POST /webhooks/stripe` - Stripe webhook

### Admin/Agent
- `GET /api/v1/admin/bookings` - Get all bookings
- `POST /api/v1/admin/bookings/:id/waive` - Waive booking
- `POST /api/v1/admin/reprice` - Reprice bookings

## Development

### Running Tests
```bash
# Backend tests
cd server
make test

# Frontend tests
cd frontend
npm test

# All tests
npm test
```

### Linting
```bash
# Backend linting
cd server
make lint

# Frontend linting
cd frontend
npm run lint

# All linting
npm run lint
```

### Database Management
```bash
# Run migrations
cd server
make migrate

# Seed database
make seed

# Reset database (if needed)
make reset
```

## Docker Deployment

### Build and run with Docker Compose
```bash
# Build all services
docker compose build

# Start all services
docker compose up -d

# View logs
docker compose logs -f
```

### Individual Docker builds
```bash
# Backend
cd server
docker build -t skyliner-backend .

# Frontend
cd frontend
docker build -t skyliner-frontend .
```

## Environment Variables

### Backend (.env)
```env
DATABASE_URL="postgresql://postgres:postgres@localhost:5432/flightdb?sslmode=disable"
REDIS_URL="redis://localhost:6379"
JWT_ACCESS_TTL="15m"
JWT_REFRESH_TTL="168h"
JWT_SECRET="your-super-secret-jwt-key-here"
GOOGLE_CLIENT_ID=""
STRIPE_SECRET_KEY=""
STRIPE_WEBHOOK_SECRET=""
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api/v1
VITE_STRIPE_PUBLISHABLE_KEY=pk_test_xxx
```

## WebSocket Events

- `priceTick:{searchHash}` - Price updates
- `seatUpdate:{flightId}` - Seat availability changes
- `bookingStatus:{bookingId}` - Booking status updates

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with modern web technologies
- Inspired by popular flight booking platforms
- Designed for simplicity and ease of use

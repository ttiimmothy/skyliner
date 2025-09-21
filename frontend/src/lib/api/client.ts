import { browser } from '$app/environment';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

class ApiClient {
  private baseURL: string;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseURL}${endpoint}`;
    
    // Get auth token from localStorage if available
    const token = browser ? localStorage.getItem('accessToken') : null;
    
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      ...(options.headers as Record<string, string>),
    };

    if (token) {
      headers.Authorization = `Bearer ${token}`;
    }

    const response = await fetch(url, {
      ...options,
      headers,
    });

    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: 'Unknown error' }));
      throw new Error(error.error || `HTTP ${response.status}`);
    }

    return response.json();
  }

  // Auth endpoints
  async signup(data: {
    email: string;
    password: string;
    first_name: string;
    last_name: string;
  }) {
    return this.request('/auth/signup', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async login(data: { email: string; password: string }) {
    return this.request('/auth/login', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async logout() {
    return this.request('/auth/logout', {
      method: 'POST',
    });
  }

  // Search endpoints
  async getAirports() {
    return this.request('/airports');
  }

  async getAirlines() {
    return this.request('/airlines');
  }

  async searchFlights(data: {
    trip_type: string;
    legs: Array<{
      origin: string;
      destination: string;
      date: string;
    }>;
    passengers: number;
    cabin?: string;
    flexibility?: number;
  }) {
    return this.request('/search', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getSeatMap(flightId: number) {
    return this.request(`/flights/${flightId}/seatmap`);
  }

  // Booking endpoints
  async createBooking(data: {
    segments: Array<{
      flight_id: number;
      fare_id: number;
      seat_number?: string;
    }>;
    passengers: Array<{
      first_name: string;
      last_name: string;
      email?: string;
      phone?: string;
      date_of_birth?: string;
      passport?: string;
      ssr?: string;
    }>;
    seats?: Array<{
      segment_id: number;
      seat_id: number;
    }>;
    extras?: Array<{
      type: string;
      price: number;
    }>;
  }) {
    return this.request('/bookings', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getBooking(bookingId: number) {
    return this.request(`/bookings/${bookingId}`);
  }

  async issueBooking(bookingId: number) {
    return this.request(`/bookings/${bookingId}/issue`, {
      method: 'POST',
    });
  }

  async cancelBooking(bookingId: number) {
    return this.request(`/bookings/${bookingId}/cancel`, {
      method: 'POST',
    });
  }

  // Payment endpoints
  async createCheckoutSession(data: { booking_id: number }) {
    return this.request('/payments/checkout-session', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async createBillingPortal(data: { booking_id: number }) {
    return this.request('/payments/billing-portal', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }
}

export const apiClient = new ApiClient(API_BASE_URL);

import { writable } from 'svelte/store';

export interface FlightSegment {
  id: number;
  flight_id: number;
  fare_id: number;
  seat_number?: string;
  flight: {
    id: number;
    number: string;
    airline: {
      id: number;
      code: string;
      name: string;
    };
    origin: {
      id: number;
      code: string;
      name: string;
      city: string;
      country: string;
    };
    destination: {
      id: number;
      code: string;
      name: string;
      city: string;
      country: string;
    };
    departure_time: string;
    arrival_time: string;
    duration: number;
    stops: number;
  };
  fare: {
    id: number;
    class: string;
    fare_type: string;
    base_price: number;
    currency: string;
    available: number;
  };
}

export interface CartItem {
  segments: FlightSegment[];
  passengers: number;
  totalPrice: number;
  currency: string;
}

export interface CartState {
  items: CartItem[];
  totalItems: number;
  totalPrice: number;
  currency: string;
}

const initialState: CartState = {
  items: [],
  totalItems: 0,
  totalPrice: 0,
  currency: 'USD'
};

export const cartStore = writable<CartState>(initialState);

export function addToCart(item: CartItem) {
  cartStore.update(state => {
    const newItems = [...state.items, item];
    const totalItems = newItems.reduce((sum, item) => sum + item.passengers, 0);
    const totalPrice = newItems.reduce((sum, item) => sum + item.totalPrice, 0);
    
    return {
      ...state,
      items: newItems,
      totalItems,
      totalPrice
    };
  });
}

export function removeFromCart(index: number) {
  cartStore.update(state => {
    const newItems = state.items.filter((_, i) => i !== index);
    const totalItems = newItems.reduce((sum, item) => sum + item.passengers, 0);
    const totalPrice = newItems.reduce((sum, item) => sum + item.totalPrice, 0);
    
    return {
      ...state,
      items: newItems,
      totalItems,
      totalPrice
    };
  });
}

export function clearCart() {
  cartStore.set(initialState);
}

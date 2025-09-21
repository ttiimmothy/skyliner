import { writable } from 'svelte/store';

export interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  role: 'traveler' | 'agent' | 'admin';
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface AuthState {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
}

const initialState: AuthState = {
  user: null,
  accessToken: null,
  refreshToken: null,
  isAuthenticated: false,
  isLoading: true
};

export const authStore = writable<AuthState>(initialState);

export function setUser(user: User, accessToken: string, refreshToken: string) {
  authStore.update(state => ({
    ...state,
    user,
    accessToken,
    refreshToken,
    isAuthenticated: true,
    isLoading: false
  }));
  
  // Store tokens in localStorage
  localStorage.setItem('accessToken', accessToken);
  localStorage.setItem('refreshToken', refreshToken);
}

export function clearUser() {
  authStore.update(state => ({
    ...state,
    user: null,
    accessToken: null,
    refreshToken: null,
    isAuthenticated: false,
    isLoading: false
  }));
  
  // Clear tokens from localStorage
  localStorage.removeItem('accessToken');
  localStorage.removeItem('refreshToken');
}

export function setLoading(loading: boolean) {
  authStore.update(state => ({
    ...state,
    isLoading: loading
  }));
}

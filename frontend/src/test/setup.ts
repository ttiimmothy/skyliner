import '@testing-library/jest-dom';
import { vi, expect } from 'vitest';

// Mock SvelteKit stores
vi.mock('$app/stores', () => ({
	page: {
		subscribe: vi.fn(() => () => {}),
		url: new URL('http://localhost:3000')
	},
	navigating: {
		subscribe: vi.fn(() => () => {})
	}
}));

// Mock SvelteKit navigation
vi.mock('$app/navigation', () => ({
	goto: vi.fn(),
	beforeNavigate: vi.fn(),
	afterNavigate: vi.fn()
}));

// Mock SvelteKit environment
vi.mock('$app/environment', () => ({
	browser: true
}));

// Mock environment
Object.defineProperty(import.meta, 'env', {
	value: {
		VITE_API_URL: 'http://localhost:8080/api/v1'
	}
});

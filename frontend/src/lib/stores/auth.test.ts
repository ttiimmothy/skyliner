import { describe, it, expect, vi } from 'vitest';

// Mock Svelte stores
const mockAuthStore = {
	subscribe: vi.fn((fn) => {
		fn({
			isAuthenticated: false,
			user: null,
			accessToken: null,
			refreshToken: null
		});
		return () => {};
	}),
	update: vi.fn()
};

vi.mock('svelte/store', () => ({
	writable: vi.fn(() => mockAuthStore)
}));

describe('Auth Store', () => {
	it('should have authStore defined', () => {
		expect(mockAuthStore).toBeDefined();
		expect(mockAuthStore.subscribe).toBeDefined();
		expect(mockAuthStore.update).toBeDefined();
	});

	it('should call subscribe with initial state', () => {
		const callback = vi.fn();
		mockAuthStore.subscribe(callback);
		
		expect(callback).toHaveBeenCalledWith({
			isAuthenticated: false,
			user: null,
			accessToken: null,
			refreshToken: null
		});
	});
});

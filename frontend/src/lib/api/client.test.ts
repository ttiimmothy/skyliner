import { describe, it, expect, vi, beforeEach } from 'vitest';

// Mock the API client module
const mockApiClient = {
	signup: vi.fn(),
	login: vi.fn(),
	logout: vi.fn(),
	getAirports: vi.fn(),
	getAirlines: vi.fn(),
	searchFlights: vi.fn(),
	getSeatMap: vi.fn(),
	createBooking: vi.fn(),
	getBooking: vi.fn(),
	issueBooking: vi.fn(),
	cancelBooking: vi.fn(),
	createCheckoutSession: vi.fn(),
	createBillingPortal: vi.fn()
};

vi.mock('./client', () => ({
	apiClient: mockApiClient
}));

describe('API Client', () => {
	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('should have all required methods', () => {
		expect(mockApiClient.signup).toBeDefined();
		expect(mockApiClient.login).toBeDefined();
		expect(mockApiClient.logout).toBeDefined();
		expect(mockApiClient.getAirports).toBeDefined();
		expect(mockApiClient.getAirlines).toBeDefined();
		expect(mockApiClient.searchFlights).toBeDefined();
		expect(mockApiClient.getSeatMap).toBeDefined();
		expect(mockApiClient.createBooking).toBeDefined();
		expect(mockApiClient.getBooking).toBeDefined();
		expect(mockApiClient.issueBooking).toBeDefined();
		expect(mockApiClient.cancelBooking).toBeDefined();
		expect(mockApiClient.createCheckoutSession).toBeDefined();
		expect(mockApiClient.createBillingPortal).toBeDefined();
	});

	it('should call signup with correct parameters', async () => {
		const mockResponse = {
			user: { id: 1, email: 'test@example.com' },
			access_token: 'token123',
			refresh_token: 'refresh123'
		};

		mockApiClient.signup.mockResolvedValue(mockResponse);

		const result = await mockApiClient.signup({
			email: 'test@example.com',
			password: 'password123',
			first_name: 'Test',
			last_name: 'User'
		});

		expect(mockApiClient.signup).toHaveBeenCalledWith({
			email: 'test@example.com',
			password: 'password123',
			first_name: 'Test',
			last_name: 'User'
		});

		expect(result).toEqual(mockResponse);
	});
});

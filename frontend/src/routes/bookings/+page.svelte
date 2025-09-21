<script lang="ts">
  import { onMount } from 'svelte';
  import { authStore } from '$lib/stores/auth';
  import { apiClient } from '$lib/api/client';

  let bookings = $state<any[]>([]);
  let isLoading = $state(true);
  let error = $state<string | null>(null);
  let selectedStatus = $state('all');

  // Mock booking data for demonstration
  const mockBookings = [
    {
      id: 'BK001',
      status: 'confirmed',
      bookingDate: '2024-01-15',
      totalAmount: 599.98,
      passengers: 2,
      flights: [
        {
          number: 'AA100',
          airline: 'American Airlines',
          origin: { code: 'JFK', name: 'John F. Kennedy International Airport', city: 'New York' },
          destination: { code: 'LAX', name: 'Los Angeles International Airport', city: 'Los Angeles' },
          departure: '2024-02-15T08:30:00Z',
          arrival: '2024-02-15T14:00:00Z',
          duration: '5h 30m',
          class: 'Economy'
        }
      ]
    },
    {
      id: 'BK002',
      status: 'pending',
      bookingDate: '2024-01-20',
      totalAmount: 899.99,
      passengers: 1,
      flights: [
        {
          number: 'BA200',
          airline: 'British Airways',
          origin: { code: 'JFK', name: 'John F. Kennedy International Airport', city: 'New York' },
          destination: { code: 'LHR', name: 'London Heathrow Airport', city: 'London' },
          departure: '2024-03-10T22:15:00Z',
          arrival: '2024-03-11T10:15:00Z',
          duration: '7h 0m',
          class: 'Business'
        }
      ]
    },
    {
      id: 'BK003',
      status: 'cancelled',
      bookingDate: '2024-01-10',
      totalAmount: 299.99,
      passengers: 1,
      flights: [
        {
          number: 'AF300',
          airline: 'Air France',
          origin: { code: 'LHR', name: 'London Heathrow Airport', city: 'London' },
          destination: { code: 'CDG', name: 'Charles de Gaulle Airport', city: 'Paris' },
          departure: '2024-01-25T14:30:00Z',
          arrival: '2024-01-25T16:45:00Z',
          duration: '1h 15m',
          class: 'Economy'
        }
      ]
    }
  ];

  onMount(async () => {
    try {
      // In a real app, this would fetch from the API
      // const response = await apiClient.getBookings();
      // bookings = response.bookings || [];
      
      // For now, use mock data
      bookings = mockBookings;
    } catch (err) {
      error = 'Failed to load bookings';
      console.error('Error loading bookings:', err);
    } finally {
      isLoading = false;
    }
  });

  function getStatusColor(status: string) {
    switch (status) {
      case 'confirmed': return 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200';
      case 'pending': return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-200';
      case 'cancelled': return 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200';
      default: return 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200';
    }
  }

  function formatDate(dateString: string) {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    });
  }

  function formatDateTime(dateString: string) {
    return new Date(dateString).toLocaleString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  }

  function filteredBookings() {
    if (selectedStatus === 'all') return bookings;
    return bookings.filter(booking => booking.status === selectedStatus);
  }

  function cancelBooking(bookingId: string) {
    if (confirm('Are you sure you want to cancel this booking?')) {
      // In a real app, this would call the API
      console.log('Cancelling booking:', bookingId);
      // Update the booking status
      bookings = bookings.map(booking => 
        booking.id === bookingId ? { ...booking, status: 'cancelled' } : booking
      );
    }
  }
</script>

<svelte:head>
  <title>My Bookings - Skyliner</title>
  <meta name="description" content="View and manage your flight bookings on Skyliner" />
</svelte:head>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <div class="mb-8">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white">My Bookings</h1>
    <p class="mt-2 text-gray-600 dark:text-gray-400">Manage your flight reservations and view booking details</p>
  </div>

  <!-- Status Filter -->
  <div class="mb-6">
    <div class="flex space-x-4">
      <button
        onclick={() => selectedStatus = 'all'}
        class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 {selectedStatus === 'all' ? 'bg-primary-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300 dark:bg-gray-700 dark:text-gray-300 dark:hover:bg-gray-600'}"
      >
        All ({bookings.length})
      </button>
      <button
        onclick={() => selectedStatus = 'confirmed'}
        class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 {selectedStatus === 'confirmed' ? 'bg-primary-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300 dark:bg-gray-700 dark:text-gray-300 dark:hover:bg-gray-600'}"
      >
        Confirmed ({bookings.filter(b => b.status === 'confirmed').length})
      </button>
      <button
        onclick={() => selectedStatus = 'pending'}
        class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 {selectedStatus === 'pending' ? 'bg-primary-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300 dark:bg-gray-700 dark:text-gray-300 dark:hover:bg-gray-600'}"
      >
        Pending ({bookings.filter(b => b.status === 'pending').length})
      </button>
      <button
        onclick={() => selectedStatus = 'cancelled'}
        class="px-4 py-2 rounded-lg text-sm font-medium transition-colors duration-200 {selectedStatus === 'cancelled' ? 'bg-primary-600 text-white' : 'bg-gray-200 text-gray-700 hover:bg-gray-300 dark:bg-gray-700 dark:text-gray-300 dark:hover:bg-gray-600'}"
      >
        Cancelled ({bookings.filter(b => b.status === 'cancelled').length})
      </button>
    </div>
  </div>

  <!-- Loading State -->
  {#if isLoading}
    <div class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
    </div>
  {/if}

  <!-- Error State -->
  {#if error}
    <div class="bg-red-50 dark:bg-red-900 border border-red-200 dark:border-red-700 rounded-lg p-4 mb-6">
      <div class="flex">
        <svg class="h-5 w-5 text-red-400" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"></path>
        </svg>
        <div class="ml-3">
          <h3 class="text-sm font-medium text-red-800 dark:text-red-200">Error</h3>
          <p class="mt-1 text-sm text-red-700 dark:text-red-300">{error}</p>
        </div>
      </div>
    </div>
  {/if}

  <!-- Bookings List -->
  {#if !isLoading && !error}
    {#if filteredBookings().length === 0}
      <div class="text-center py-12">
        <svg class="mx-auto h-12 w-12 text-gray-400 dark:text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.29-1.009-5.824-2.57M15 6.75a3 3 0 11-6 0 3 3 0 016 0z"></path>
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900 dark:text-white">No bookings found</h3>
        <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">
          {selectedStatus === 'all' ? 'You haven\'t made any bookings yet.' : `No ${selectedStatus} bookings found.`}
        </p>
        <div class="mt-6">
          <a href="/" class="btn btn-primary">Search Flights</a>
        </div>
      </div>
    {:else}
      <div class="space-y-6">
        {#each filteredBookings() as booking}
          <div class="card">
            <div class="flex items-start justify-between mb-4">
              <div>
                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Booking #{booking.id}</h3>
                <p class="text-sm text-gray-600 dark:text-gray-400">Booked on {formatDate(booking.bookingDate)}</p>
              </div>
              <div class="flex items-center space-x-3">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium {getStatusColor(booking.status)}">
                  {booking.status.charAt(0).toUpperCase() + booking.status.slice(1)}
                </span>
                {#if booking.status === 'confirmed' || booking.status === 'pending'}
                  <button
                    onclick={() => cancelBooking(booking.id)}
                    class="text-red-600 hover:text-red-500 text-sm font-medium"
                  >
                    Cancel
                  </button>
                {/if}
              </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-4">
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Total Amount</p>
                <p class="text-lg font-semibold text-gray-900 dark:text-white">${booking.totalAmount.toFixed(2)}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Passengers</p>
                <p class="text-lg font-semibold text-gray-900 dark:text-white">{booking.passengers}</p>
              </div>
              <div>
                <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Flights</p>
                <p class="text-lg font-semibold text-gray-900 dark:text-white">{booking.flights.length}</p>
              </div>
            </div>

            <!-- Flight Details -->
            <div class="space-y-4">
              {#each booking.flights as flight}
                <div class="border border-gray-200 dark:border-gray-700 rounded-lg p-4">
                  <div class="flex items-center justify-between mb-3">
                    <div class="flex items-center space-x-3">
                      <div class="w-8 h-8 bg-primary-100 dark:bg-primary-900 rounded-full flex items-center justify-center">
                        <span class="text-xs font-bold text-primary-600 dark:text-primary-400">{flight.airline.charAt(0)}</span>
                      </div>
                      <div>
                        <p class="font-medium text-gray-900 dark:text-white">{flight.airline}</p>
                        <p class="text-sm text-gray-500 dark:text-gray-400">Flight {flight.number}</p>
                      </div>
                    </div>
                    <span class="inline-flex items-center px-2 py-1 rounded text-xs font-medium bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-200">
                      {flight.class}
                    </span>
                  </div>

                  <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                    <div>
                      <p class="text-sm font-medium text-gray-500 dark:text-gray-400">From</p>
                      <p class="font-semibold text-gray-900 dark:text-white">{flight.origin.code}</p>
                      <p class="text-sm text-gray-600 dark:text-gray-400">{flight.origin.city}</p>
                      <p class="text-sm text-gray-500 dark:text-gray-500">{formatDateTime(flight.departure)}</p>
                    </div>
                    <div class="flex items-center justify-center">
                      <div class="flex-1 border-t border-gray-300 dark:border-gray-600"></div>
                      <div class="px-3 text-sm text-gray-500 dark:text-gray-400">{flight.duration}</div>
                      <div class="flex-1 border-t border-gray-300 dark:border-gray-600"></div>
                    </div>
                    <div>
                      <p class="text-sm font-medium text-gray-500 dark:text-gray-400">To</p>
                      <p class="font-semibold text-gray-900 dark:text-white">{flight.destination.code}</p>
                      <p class="text-sm text-gray-600 dark:text-gray-400">{flight.destination.city}</p>
                      <p class="text-sm text-gray-500 dark:text-gray-500">{formatDateTime(flight.arrival)}</p>
                    </div>
                  </div>
                </div>
              {/each}
            </div>
          </div>
        {/each}
      </div>
    {/if}
  {/if}
</div>

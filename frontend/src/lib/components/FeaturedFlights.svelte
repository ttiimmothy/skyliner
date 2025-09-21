<script lang="ts">
  import { onMount } from 'svelte';
  import { apiClient } from '$lib/api/client';

  let featuredFlights: any[] = [];
  let isLoading = true;

  onMount(async () => {
    try {
      // Get some sample flights for featured section
      const response = await apiClient.searchFlights({
        trip_type: 'one-way',
        legs: [
          {
            origin: 'JFK',
            destination: 'LAX',
            date: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000).toISOString().split('T')[0]
          }
        ],
        passengers: 1,
        cabin: '',
        flexibility: 3
      }) as { flights: any[] };
      
      featuredFlights = response.flights?.slice(0, 3) || [];
    } catch (error) {
      console.error('Failed to load featured flights:', error);
    } finally {
      isLoading = false;
    }
  });

  function formatPrice(price: number, currency: string = 'USD') {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: currency
    }).format(price);
  }

  function formatDuration(minutes: number) {
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return `${hours}h ${mins}m`;
  }

  function formatTime(dateString: string) {
    return new Date(dateString).toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit'
    });
  }
</script>

<div class="text-center mb-12">
  <h2 class="text-3xl font-bold text-gray-900 dark:text-gray-200 mb-4">Featured Flights</h2>
  <p class="text-lg text-gray-600 dark:text-gray-400">Discover amazing destinations at great prices</p>
</div>

{#if isLoading}
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    {#each Array.from({length: 3}) as _}
      <div class="card animate-pulse">
        <div class="h-4 bg-gray-200 rounded w-3/4 mb-4"></div>
        <div class="h-3 bg-gray-200 rounded w-1/2 mb-2"></div>
        <div class="h-3 bg-gray-200 rounded w-2/3 mb-4"></div>
        <div class="h-8 bg-gray-200 rounded w-1/3"></div>
      </div>
    {/each}
  </div>
{:else if featuredFlights.length > 0}
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
    {#each featuredFlights as flight}
      <div class="card hover:shadow-lg transition-shadow duration-200">
        <div class="flex justify-between items-start mb-4">
          <div>
            <h3 class="text-lg font-semibold text-gray-900">
              {flight.origin.code} → {flight.destination.code}
            </h3>
            <p class="text-sm text-gray-600">
              {flight.airline.name} {flight.number}
            </p>
          </div>
          <div class="text-right">
            <div class="text-2xl font-bold text-primary-600">
              {formatPrice(flight.fares[0]?.base_price || 0)}
            </div>
            <div class="text-sm text-gray-500">
              {flight.fares[0]?.class} class
            </div>
          </div>
        </div>

        <div class="space-y-2 mb-4">
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">Departure</span>
            <span class="font-medium">
              {formatTime(flight.departure_time)}
            </span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">Arrival</span>
            <span class="font-medium">
              {formatTime(flight.arrival_time)}
            </span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-gray-600">Duration</span>
            <span class="font-medium">
              {formatDuration(flight.duration)}
            </span>
          </div>
          {#if flight.stops > 0}
            <div class="flex justify-between text-sm">
              <span class="text-gray-600">Stops</span>
              <span class="font-medium">
                {flight.stops} {flight.stops === 1 ? 'stop' : 'stops'}
              </span>
            </div>
          {/if}
        </div>

        <div class="flex space-x-2">
          <button class="btn btn-primary flex-1">
            Select Flight
          </button>
          <button class="btn btn-outline">
            Details
          </button>
        </div>
      </div>
    {/each}
  </div>
{:else}
  <div class="text-center py-12">
    <div class="text-gray-400 mb-4">
      <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
      </svg>
    </div>
    <h3 class="text-lg font-medium text-gray-900 dark:text-gray-300 mb-2">No flights available</h3>
    <p class="text-gray-600">Try adjusting your search criteria or dates.</p>
  </div>
{/if}

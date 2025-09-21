<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { apiClient } from '$lib/api/client';
  import FlightCard from '$lib/components/FlightCard.svelte';
  import FiltersPanel from '$lib/components/FiltersPanel.svelte';

  let searchResults: any[] = [];
  let isLoading = true;
  let error = '';
  let filters = {
    priceRange: [0, 2000],
    airlines: [] as string[],
    stops: [] as number[],
    departureTime: [] as string[],
    cabinClass: [] as string[]
  };

  onMount(async () => {
    // Get search parameters from URL
    const urlParams = new URLSearchParams($page.url.search);
    const searchData = {
      trip_type: urlParams.get('trip_type') || 'one-way',
      legs: JSON.parse(urlParams.get('legs') || '[]'),
      passengers: parseInt(urlParams.get('passengers') || '1'),
      cabin: urlParams.get('cabin') || '',
      flexibility: parseInt(urlParams.get('flexibility') || '0')
    };

    try {
      const response = await apiClient.searchFlights(searchData) as { flights: any[] };
      searchResults = response.flights || [];
    } catch (err: any) {
      error = err.message || 'Search failed';
    } finally {
      isLoading = false;
    }
  });

  function handleFilterChange(newFilters: any) {
    filters = { ...filters, ...newFilters };
    // Apply filters to searchResults
    // This would typically involve re-filtering the results
  }

  function formatPrice(price: number, currency: string = 'USD') {
    return new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: currency
    }).format(price);
  }
</script>

<svelte:head>
  <title>Search Results - Skyliner</title>
</svelte:head>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <div class="flex flex-col lg:flex-row gap-8">
    <!-- Filters Sidebar -->
    <div class="lg:w-1/4">
      <FiltersPanel bind:filters onFilterChange={handleFilterChange} />
    </div>

    <!-- Results -->
    <div class="lg:w-3/4">
      <!-- Header -->
      <div class="mb-6">
        <h1 class="text-2xl font-bold text-gray-900 dark:text-white mb-2">
          {searchResults.length} flights found
        </h1>
        <p class="text-gray-600 dark:text-gray-400">
          {#if searchResults.length > 0}
            Showing flights from {searchResults[0].origin.code} to {searchResults[0].destination.code}
          {/if}
        </p>
      </div>

      <!-- Loading State -->
      {#if isLoading}
        <div class="space-y-4">
          {#each Array.from({length: 3}) as _}
            <div class="card animate-pulse">
              <div class="flex justify-between items-start mb-4">
                <div class="space-y-2">
                  <div class="h-4 bg-gray-200 rounded w-32"></div>
                  <div class="h-3 bg-gray-200 rounded w-24"></div>
                </div>
                <div class="h-8 bg-gray-200 rounded w-20"></div>
              </div>
              <div class="space-y-2">
                <div class="h-3 bg-gray-200 rounded w-full"></div>
                <div class="h-3 bg-gray-200 rounded w-3/4"></div>
                <div class="h-3 bg-gray-200 rounded w-1/2"></div>
              </div>
            </div>
          {/each}
        </div>
      {:else if error}
        <!-- Error State -->
        <div class="text-center py-12">
          <div class="text-red-400 mb-4">
            <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">Search Error</h3>
          <p class="text-gray-600 dark:text-gray-400 mb-4">{error}</p>
          <a href="/" class="btn btn-primary">Try Again</a>
        </div>
      {:else if searchResults.length === 0}
        <!-- No Results -->
        <div class="text-center py-12">
          <div class="text-gray-400 mb-4">
            <svg class="w-16 h-16 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.29-1.009-5.824-2.709M15 6.75a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
          </div>
          <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-2">No flights found</h3>
          <p class="text-gray-600 dark:text-gray-400 mb-4">
            Try adjusting your search criteria or dates to find more options.
          </p>
          <a href="/" class="btn btn-primary">Modify Search</a>
        </div>
      {:else}
        <!-- Results -->
        <div class="space-y-4">
          {#each searchResults as flight}
            <FlightCard {flight} />
          {/each}
        </div>
      {/if}
    </div>
  </div>
</div>

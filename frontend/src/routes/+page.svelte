<script lang="ts">
  import { goto } from '$app/navigation';
  import SearchForm from '$lib/components/SearchForm.svelte';
  import FeaturedFlights from '$lib/components/FeaturedFlights.svelte';

  function handleSearch(searchData: any) {
    // Build URL parameters for the search
    const params = new URLSearchParams();
    params.set('trip_type', searchData.trip_type || 'one-way');
    params.set('legs', JSON.stringify(searchData.legs || []));
    params.set('passengers', (searchData.passengers || 1).toString());
    params.set('cabin', searchData.cabin || '');
    params.set('flexibility', (searchData.flexibility || 0).toString());

    // Navigate to results page with search parameters
    goto(`/results?${params.toString()}`);
  }
</script>

<svelte:head>
  <title>Skyliner - Find Your Perfect Flight</title>
</svelte:head>

<div class="bg-gradient-to-br from-primary-600 to-primary-800 text-black dark:text-white dark:from-primary-700 dark:to-primary-900">
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24">
    <div class="text-center">
      <h1 class="text-4xl md:text-6xl font-bold mb-6">
        Find Your Perfect Flight
      </h1>
      <p class="text-xl md:text-2xl mb-8 text-primary-100 dark:text-primary-200">
        Book flights with ease and discover amazing destinations
      </p>
    </div>
    
    <div class="max-w-4xl mx-auto">
      <SearchForm search={handleSearch} />
    </div>
  </div>
</div>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
  <FeaturedFlights />
</div>
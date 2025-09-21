<script lang="ts">
  import { apiClient } from '$lib/api/client';

  let { search } = $props();

  let tripType = $state('round-trip');
  let passengers = $state(1);
  let cabin = $state('');
  let flexibility = $state(0);
  let legs = $state([
    { origin: '', destination: '', date: '' },
    { origin: '', destination: '', date: '' }
  ]);

  let airports = $state<any[]>([]);
  let isLoading = $state(false);

  // Load airports on mount
  async function loadAirports() {
    try {
      const response = await apiClient.getAirports() as { airports: any[] };
      airports = response.airports || [];
    } catch (error) {
      console.error('Failed to load airports:', error);
    }
  }

  // Initialize
  loadAirports();

  function addLeg() {
    legs = [...legs, { origin: '', destination: '', date: '' }];
  }

  function removeLeg(index: number) {
    if (legs.length > 1) {
      legs = legs.filter((_, i) => i !== index);
    }
  }

  function updateLeg(index: number, field: string, value: string) {
    legs = legs.map((leg, i) => 
      i === index ? { ...leg, [field]: value } : leg
    );
  }

  async function handleSearch() {
    if (!legs[0].origin || !legs[0].destination || !legs[0].date) {
      alert('Please fill in all required fields');
      return;
    }

    isLoading = true;
    try {
      const searchData = {
        trip_type: tripType,
        legs: legs.filter(leg => leg.origin && leg.destination && leg.date),
        passengers,
        cabin: cabin || undefined,
        flexibility
      };

      // Call the search function passed as prop
      search?.(searchData);
    } catch (error) {
      console.error('Search failed:', error);
      alert('Search failed. Please try again.');
    } finally {
      isLoading = false;
    }
  }

  function getToday() {
    return new Date().toISOString().split('T')[0];
  }
</script>

<div class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6 border border-gray-200 dark:border-gray-700">
  <form onsubmit={handleSearch}>
    <!-- Trip Type -->
    <div class="mb-6">
      <div class="label">Trip Type</div>
      <div class="flex space-x-4">
        <label class="flex items-center text-gray-700 dark:text-gray-300">
          <input
            type="radio"
            bind:group={tripType}
            value="one-way"
            class="mr-2 text-primary-600 focus:ring-primary-500"
          />
          One-way
        </label>
        <label class="flex items-center text-gray-700 dark:text-gray-300">
          <input
            type="radio"
            bind:group={tripType}
            value="round-trip"
            class="mr-2 text-primary-600 focus:ring-primary-500"
          />
          Round-trip
        </label>
        <label class="flex items-center text-gray-700 dark:text-gray-300">
          <input
            type="radio"
            bind:group={tripType}
            value="multi-city"
            class="mr-2 text-primary-600 focus:ring-primary-500"
          />
          Multi-city
        </label>
      </div>
    </div>

    <!-- Flight Legs -->
    <div class="mb-6">
      <div class="label">Flight Details</div>
      {#each legs as leg, index}
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-4">
          <!-- Origin -->
          <div>
            <div class="label">From</div>
            <select
              bind:value={leg.origin}
              onchange={(e) => updateLeg(index, 'origin', (e.target as HTMLSelectElement).value)}
              class="input"
              required
            >
              <option value="">Select airport</option>
              {#each airports as airport}
                <option value={airport.code}>{airport.code} - {airport.name}</option>
              {/each}
            </select>
          </div>

          <!-- Destination -->
          <div>
            <div class="label">To</div>
            <select
              bind:value={leg.destination}
              onchange={(e) => updateLeg(index, 'destination', (e.target as HTMLSelectElement).value)}
              class="input"
              required
            >
              <option value="">Select airport</option>
              {#each airports as airport}
                <option value={airport.code}>{airport.code} - {airport.name}</option>
              {/each}
            </select>
          </div>

          <!-- Date -->
          <div>
            <div class="label">Date</div>
            <input
              type="date"
              bind:value={leg.date}
              onchange={(e) => updateLeg(index, 'date', (e.target as HTMLSelectElement).value)}
              min={getToday()}
              class="input"
              required
            />
          </div>

          <!-- Remove button -->
          <div class="flex items-end">
            {#if legs.length > 1}
              <button
                type="button"
                onclick={() => removeLeg(index)}
                class="btn btn-outline text-red-600 border-red-300 hover:bg-red-50"
              >
                Remove
              </button>
            {/if}
          </div>
        </div>
      {/each}

      {#if tripType === 'multi-city'}
        <button
          type="button"
          onclick={addLeg}
          class="btn btn-outline"
        >
          Add Flight
        </button>
      {/if}
    </div>

    <!-- Passengers and Options -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
      <!-- Passengers -->
      <div>
        <div class="label">Passengers</div>
        <select bind:value={passengers} class="input">
          {#each Array.from({length: 9}, (_, i) => i + 1) as num}
            <option value={num}>{num} {num === 1 ? 'Passenger' : 'Passengers'}</option>
          {/each}
        </select>
      </div>

      <!-- Cabin Class -->
      <div>
        <div class="label">Cabin Class</div>
        <select bind:value={cabin} class="input">
          <option value="">Any</option>
          <option value="economy">Economy</option>
          <option value="business">Business</option>
          <option value="first">First</option>
        </select>
      </div>

      <!-- Flexibility -->
      <div>
        <div class="label">Date Flexibility</div>
        <select bind:value={flexibility} class="input">
          <option value={0}>Exact dates</option>
          <option value={1}>±1 day</option>
          <option value={3}>±3 days</option>
          <option value={7}>±1 week</option>
        </select>
      </div>
    </div>

    <!-- Search Button -->
    <div class="text-center">
      <button
        type="submit"
        disabled={isLoading}
        class="btn btn-primary px-8 py-3 text-lg disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {#if isLoading}
          <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white inline" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          Searching...
        {:else}
          Search Flights
        {/if}
      </button>
    </div>
  </form>
</div>

<script lang="ts">
  let { filters = $bindable({
    priceRange: [0, 2000],
    airlines: [],
    stops: [],
    departureTime: [],
    cabinClass: []
  }), onFilterChange } = $props();

  let airlines = [
    { code: 'AA', name: 'American Airlines' },
    { code: 'DL', name: 'Delta Air Lines' },
    { code: 'UA', name: 'United Airlines' },
    { code: 'BA', name: 'British Airways' },
    { code: 'AF', name: 'Air France' },
    { code: 'JL', name: 'Japan Airlines' }
  ];

  let stopOptions = [
    { value: 0, label: 'Non-stop' },
    { value: 1, label: '1 stop' },
    { value: 2, label: '2+ stops' }
  ];

  let timeOptions = [
    { value: 'morning', label: 'Morning (6AM-12PM)' },
    { value: 'afternoon', label: 'Afternoon (12PM-6PM)' },
    { value: 'evening', label: 'Evening (6PM-12AM)' },
    { value: 'night', label: 'Night (12AM-6AM)' }
  ];

  let cabinOptions = [
    { value: 'economy', label: 'Economy' },
    { value: 'business', label: 'Business' },
    { value: 'first', label: 'First' }
  ];

  function updatePriceRange() {
    onFilterChange?.({ priceRange: filters.priceRange });
  }

  function toggleArrayFilter(array: any[], value: any) {
    const newArray = array.includes(value)
      ? array.filter(item => item !== value)
      : [...array, value];
    return newArray;
  }

  function toggleAirline(airline: string) {
    filters.airlines = toggleArrayFilter(filters.airlines, airline);
    onFilterChange?.({ airlines: filters.airlines });
  }

  function toggleStop(stop: number) {
    filters.stops = toggleArrayFilter(filters.stops, stop);
    onFilterChange?.({ stops: filters.stops });
  }

  function toggleTime(time: string) {
    filters.departureTime = toggleArrayFilter(filters.departureTime, time);
    onFilterChange?.({ departureTime: filters.departureTime });
  }

  function toggleCabin(cabin: string) {
    filters.cabinClass = toggleArrayFilter(filters.cabinClass, cabin);
    onFilterChange?.({ cabinClass: filters.cabinClass });
  }

  function clearFilters() {
    filters = {
      priceRange: [0, 2000],
      airlines: [],
      stops: [],
      departureTime: [],
      cabinClass: []
    };
    onFilterChange?.(filters);
  }
</script>

<div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
  <div class="flex items-center justify-between mb-6">
    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Filters</h3>
    <button
      onclick={clearFilters}
      class="text-sm text-primary-600 hover:text-primary-500 dark:text-primary-400 dark:hover:text-primary-300"
    >
      Clear all
    </button>
  </div>

  <div class="space-y-6">
    <!-- Price Range -->
    <div>
      <h4 class="text-sm font-medium text-gray-900 dark:text-gray-300 mb-3">Price Range</h4>
      <div class="space-y-2">
        <input
          type="range"
          min="0"
          max="2000"
          step="50"
          bind:value={filters.priceRange[1]}
          oninput={updatePriceRange}
          class="w-full h-2 bg-gray-200 dark:bg-gray-600 rounded-lg appearance-none cursor-pointer"
        />
        <div class="flex justify-between text-sm text-gray-600 dark:text-gray-400">
          <span>$0</span>
          <span>${filters.priceRange[1]}</span>
        </div>
      </div>
    </div>

    <!-- Airlines -->
    <div>
      <h4 class="text-sm font-medium text-gray-900 dark:text-gray-300 mb-3">Airlines</h4>
      <div class="space-y-2">
        {#each airlines as airline}
          <label class="flex items-center">
            <input
              type="checkbox"
              checked={filters.airlines.includes(airline.code)}
              onchange={() => toggleAirline(airline.code)}
              class="rounded border-gray-300 dark:border-gray-600 text-primary-600 focus:ring-primary-500 dark:bg-gray-700"
            />
            <span class="ml-2 text-sm text-gray-700 dark:text-gray-300">{airline.name}</span>
          </label>
        {/each}
      </div>
    </div>

    <!-- Stops -->
    <div>
      <h4 class="text-sm font-medium text-gray-900 dark:text-gray-300 mb-3">Stops</h4>
      <div class="space-y-2">
        {#each stopOptions as option}
          <label class="flex items-center">
            <input
              type="checkbox"
              checked={filters.stops.includes(option.value)}
              onchange={() => toggleStop(option.value)}
              class="rounded border-gray-300 dark:border-gray-600 text-primary-600 focus:ring-primary-500 dark:bg-gray-700"
            />
            <span class="ml-2 text-sm text-gray-700 dark:text-gray-300">{option.label}</span>
          </label>
        {/each}
      </div>
    </div>

    <!-- Departure Time -->
    <div>
      <h4 class="text-sm font-medium text-gray-900 dark:text-gray-300 mb-3">Departure Time</h4>
      <div class="space-y-2">
        {#each timeOptions as option}
          <label class="flex items-center">
            <input
              type="checkbox"
              checked={filters.departureTime.includes(option.value)}
              onchange={() => toggleTime(option.value)}
              class="rounded border-gray-300 dark:border-gray-600 text-primary-600 focus:ring-primary-500 dark:bg-gray-700"
            />
            <span class="ml-2 text-sm text-gray-700 dark:text-gray-300">{option.label}</span>
          </label>
        {/each}
      </div>
    </div>

    <!-- Cabin Class -->
    <div>
      <h4 class="text-sm font-medium text-gray-900 dark:text-gray-300 mb-3">Cabin Class</h4>
      <div class="space-y-2">
        {#each cabinOptions as option}
          <label class="flex items-center">
            <input
              type="checkbox"
              checked={filters.cabinClass.includes(option.value)}
              onchange={() => toggleCabin(option.value)}
              class="rounded border-gray-300 dark:border-gray-600 text-primary-600 focus:ring-primary-500 dark:bg-gray-700"
            />
            <span class="ml-2 text-sm text-gray-700 dark:text-gray-300">{option.label}</span>
          </label>
        {/each}
      </div>
    </div>
  </div>
</div>

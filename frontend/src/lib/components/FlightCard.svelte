<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { cartStore, addToCart } from '$lib/stores/cart';

  const dispatch = createEventDispatcher();

  export let flight: any;

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

  function selectFlight() {
    // Add to cart
    const cartItem = {
      segments: [{
        id: 0,
        flight_id: flight.id,
        fare_id: flight.fares[0].id,
        flight: flight,
        fare: flight.fares[0]
      }],
      passengers: 1,
      totalPrice: flight.fares[0].base_price,
      currency: flight.fares[0].currency
    };
    
    addToCart(cartItem);
    dispatch('select', { flight });
  }

  function viewDetails() {
    dispatch('details', { flight });
  }
</script>

<div class="card hover:shadow-lg transition-shadow duration-200">
  <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between">
    <!-- Flight Info -->
    <div class="flex-1">
      <div class="flex items-center justify-between mb-4">
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
            {formatPrice(flight.fares[0]?.base_price || 0, flight.fares[0]?.currency)}
          </div>
          <div class="text-sm text-gray-500">
            {flight.fares[0]?.class} class
          </div>
        </div>
      </div>

      <!-- Flight Details -->
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-4 text-sm">
        <div>
          <div class="text-gray-600">Departure</div>
          <div class="font-medium">
            {formatTime(flight.departure_time)}
          </div>
          <div class="text-gray-500 text-xs">
            {flight.origin.name}
          </div>
        </div>
        
        <div>
          <div class="text-gray-600">Arrival</div>
          <div class="font-medium">
            {formatTime(flight.arrival_time)}
          </div>
          <div class="text-gray-500 text-xs">
            {flight.destination.name}
          </div>
        </div>
        
        <div>
          <div class="text-gray-600">Duration</div>
          <div class="font-medium">
            {formatDuration(flight.duration)}
          </div>
        </div>
        
        <div>
          <div class="text-gray-600">Stops</div>
          <div class="font-medium">
            {flight.stops === 0 ? 'Non-stop' : `${flight.stops} ${flight.stops === 1 ? 'stop' : 'stops'}`}
          </div>
        </div>
      </div>

      <!-- Fare Options -->
      {#if flight.fares.length > 1}
        <div class="mt-4">
          <div class="text-sm text-gray-600 mb-2">Fare Options:</div>
          <div class="flex space-x-4">
            {#each flight.fares as fare}
              <div class="flex items-center space-x-2">
                <span class="text-sm font-medium">{fare.fare_type}</span>
                <span class="text-sm text-gray-600">
                  {formatPrice(fare.base_price, fare.currency)}
                </span>
              </div>
            {/each}
          </div>
        </div>
      {/if}
    </div>

    <!-- Actions -->
    <div class="mt-4 lg:mt-0 lg:ml-6 flex flex-col space-y-2">
      <button
        on:click={selectFlight}
        class="btn btn-primary whitespace-nowrap"
      >
        Select Flight
      </button>
      <button
        on:click={viewDetails}
        class="btn btn-outline whitespace-nowrap"
      >
        View Details
      </button>
    </div>
  </div>
</div>

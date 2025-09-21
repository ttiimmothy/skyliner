<script lang="ts">
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';

  $: status = $page.status;
  $: message = $page.error?.message || 'An unexpected error occurred';

  function goHome() {
    goto('/');
  }

  function goBack() {
    history.back();
  }
</script>

<svelte:head>
  <title>{status} - Skyliner</title>
</svelte:head>

<div class="min-h-screen bg-gray-50 dark:bg-gray-900 flex items-center justify-center px-4 sm:px-6 lg:px-8">
  <div class="max-w-md w-full space-y-8 text-center">
    <div>
      <div class="mx-auto h-24 w-24 text-primary-600 dark:text-primary-400">
        {#if status === 404}
          <svg fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.29-1.009-5.824-2.57M15 6.75a3 3 0 11-6 0 3 3 0 016 0z"></path>
          </svg>
        {:else}
          <svg fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"></path>
          </svg>
        {/if}
      </div>
      <h2 class="mt-6 text-3xl font-extrabold text-gray-900 dark:text-white">
        {#if status === 404}
          Page Not Found
        {:else}
          {status} Error
        {/if}
      </h2>
      <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
        {#if status === 404}
          Sorry, we couldn't find the page you're looking for.
        {:else}
          {message}
        {/if}
      </p>
    </div>

    <div class="space-y-4">
      <button
        onclick={goHome}
        class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 dark:focus:ring-offset-gray-900"
      >
        Go Home
      </button>
      
      <button
        onclick={goBack}
        class="group relative w-full flex justify-center py-2 px-4 border border-gray-300 dark:border-gray-600 text-sm font-medium rounded-md text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 dark:focus:ring-offset-gray-900"
      >
        Go Back
      </button>
    </div>

    {#if status === 404}
      <div class="mt-8">
        <p class="text-sm text-gray-500 dark:text-gray-400">
          Try searching for flights or check our popular destinations:
        </p>
        <div class="mt-4 flex flex-wrap justify-center gap-2">
          <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-primary-100 text-primary-800 dark:bg-primary-900 dark:text-primary-200">
            New York (JFK)
          </span>
          <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-primary-100 text-primary-800 dark:bg-primary-900 dark:text-primary-200">
            Los Angeles (LAX)
          </span>
          <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-primary-100 text-primary-800 dark:bg-primary-900 dark:text-primary-200">
            London (LHR)
          </span>
          <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-primary-100 text-primary-800 dark:bg-primary-900 dark:text-primary-200">
            Paris (CDG)
          </span>
        </div>
      </div>
    {/if}
  </div>
</div>

<script lang="ts">
  import { goto } from '$app/navigation';

  function goHome() {
    goto('/');
  }

  function goBack() {
    history.back();
  }
</script>

<svelte:head>
  <title>404 - Page Not Found | Skyliner</title>
  <meta name="description" content="The page you're looking for doesn't exist. Find your perfect flight with Skyliner." />
</svelte:head>

<div class="min-h-screen bg-gray-50 dark:bg-gray-900 flex items-center justify-center px-4 sm:px-6 lg:px-8">
  <div class="max-w-md w-full space-y-8 text-center">
    <div>
      <div class="mx-auto h-24 w-24 text-primary-600 dark:text-primary-400">
        <svg fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.29-1.009-5.824-2.57M15 6.75a3 3 0 11-6 0 3 3 0 016 0z"></path>
        </svg>
      </div>
      <h1 class="mt-6 text-4xl font-extrabold text-gray-900 dark:text-white">
        404
      </h1>
      <h2 class="mt-2 text-2xl font-bold text-gray-700 dark:text-gray-300">
        Page Not Found
      </h2>
      <p class="mt-4 text-lg text-gray-600 dark:text-gray-400">
        Sorry, we couldn't find the page you're looking for.
      </p>
      <p class="mt-2 text-sm text-gray-500 dark:text-gray-500">
        The page might have been moved, deleted, or you entered the wrong URL.
      </p>
    </div>

    <div class="space-y-4">
      <button
        onclick={goHome}
        class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 dark:focus:ring-offset-gray-900 transition-colors duration-200"
      >
        <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
        </svg>
        Go Home
      </button>
      
      <button
        onclick={goBack}
        class="group relative w-full flex justify-center py-3 px-4 border border-gray-300 dark:border-gray-600 text-sm font-medium rounded-md text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 dark:focus:ring-offset-gray-900 transition-colors duration-200"
      >
        <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
        </svg>
        Go Back
      </button>
    </div>

    <div class="mt-8">
      <p class="text-sm text-gray-500 dark:text-gray-400 mb-4">
        Try searching for flights or check our popular destinations:
      </p>
      <div class="flex flex-wrap justify-center gap-2">
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
        <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-primary-100 text-primary-800 dark:bg-primary-900 dark:text-primary-200">
          Tokyo (NRT)
        </span>
        <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium bg-primary-100 text-primary-800 dark:bg-primary-900 dark:text-primary-200">
          San Francisco (SFO)
        </span>
      </div>
    </div>

    <div class="mt-8 pt-6 border-t border-gray-200 dark:border-gray-700">
      <p class="text-xs text-gray-400 dark:text-gray-500">
        Need help? <a href="/contact" class="text-primary-600 dark:text-primary-400 hover:underline">Contact our support team</a>
      </p>
    </div>
  </div>
</div>

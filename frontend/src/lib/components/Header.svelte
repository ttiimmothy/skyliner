<script lang="ts">
  import { authStore } from '$lib/stores/auth';
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  let isMenuOpen = false;
  let isDarkMode = false;

  onMount(() => {
    // Check for saved theme preference
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme === 'dark') {
      isDarkMode = true;
      document.documentElement.classList.add('dark');
    }
  });

  function toggleMenu() {
    isMenuOpen = !isMenuOpen;
  }

  function toggleDarkMode() {
    isDarkMode = !isDarkMode;
    if (isDarkMode) {
      document.documentElement.classList.add('dark');
      localStorage.setItem('theme', 'dark');
    } else {
      document.documentElement.classList.remove('dark');
      localStorage.setItem('theme', 'light');
    }
  }

  function handleLogout() {
    authStore.update(state => ({
      ...state,
      user: null,
      accessToken: null,
      refreshToken: null,
      isAuthenticated: false
    }));
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    goto('/');
  }
</script>

<header class="bg-white shadow-sm border-b border-gray-200 dark:bg-gray-900 dark:border-gray-700">
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
    <div class="flex justify-between items-center h-16">
      <!-- Logo -->
      <div class="flex items-center">
        <a href="/" class="flex items-center space-x-2">
          <div class="w-8 h-8 bg-primary-600 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold text-lg">S</span>
          </div>
          <span class="text-xl font-bold text-gray-900 dark:text-white">Skyliner</span>
        </a>
      </div>

      <!-- Navigation -->
      <nav class="hidden md:flex space-x-8">
        <a href="/" class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-sm font-medium dark:text-gray-300 dark:hover:text-primary-400">
          Search Flights
        </a>
        <a href="/bookings" class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-sm font-medium dark:text-gray-300 dark:hover:text-primary-400">
          My Bookings
        </a>
        {#if $authStore.isAuthenticated && $authStore.user?.role === 'agent'}
          <a href="/agent" class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-sm font-medium dark:text-gray-300 dark:hover:text-primary-400">
            Agent Console
          </a>
        {/if}
        {#if $authStore.isAuthenticated && $authStore.user?.role === 'admin'}
          <a href="/admin" class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-sm font-medium dark:text-gray-300 dark:hover:text-primary-400">
            Admin
          </a>
        {/if}
      </nav>

      <!-- User menu -->
      <div class="flex items-center space-x-4">
        <!-- Dark mode toggle -->
        <button
          onclick={toggleDarkMode}
          class="p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 dark:text-gray-500 dark:hover:text-gray-400 dark:hover:bg-gray-800"
          aria-label="Toggle dark mode"
        >
          {#if isDarkMode}
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z" clip-rule="evenodd"></path>
            </svg>
          {:else}
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"></path>
            </svg>
          {/if}
        </button>

        {#if $authStore.isAuthenticated}
          <!-- User dropdown -->
          <div class="relative">
            <button
              onclick={toggleMenu}
              class="flex items-center space-x-2 text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
            >
              <div class="w-8 h-8 bg-primary-100 rounded-full flex items-center justify-center">
                <span class="text-primary-600 font-medium">
                  {$authStore.user?.first_name?.[0]}{$authStore.user?.last_name?.[0]}
                </span>
              </div>
              <span class="text-gray-700 dark:text-gray-300">{$authStore.user?.first_name}</span>
              <svg class="w-4 h-4 text-gray-400 dark:text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
              </svg>
            </button>

            {#if isMenuOpen}
              <div class="absolute right-0 mt-2 w-48 bg-white dark:bg-gray-800 rounded-md shadow-lg py-1 z-50 border border-gray-200 dark:border-gray-700">
                <a href="/profile" class="block px-4 py-2 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700">
                  Profile
                </a>
                <a href="/bookings" class="block px-4 py-2 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700">
                  My Bookings
                </a>
                <a href="/billing" class="block px-4 py-2 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700">
                  Billing
                </a>
                <button
                  onclick={handleLogout}
                  class="block w-full text-left px-4 py-2 text-sm text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700"
                >
                  Sign out
                </button>
              </div>
            {/if}
          </div>
        {:else}
          <!-- Login/Signup buttons -->
          <div class="flex items-center space-x-4">
            <a href="/login" class="text-gray-700 hover:text-primary-600 px-3 py-2 rounded-md text-sm font-medium dark:text-gray-300 dark:hover:text-primary-400">
              Sign in
            </a>
            <a href="/signup" class="btn btn-primary">
              Sign up
            </a>
          </div>
        {/if}

        <!-- Mobile menu button -->
        <button
          onclick={toggleMenu}
          class="md:hidden p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 dark:text-gray-500 dark:hover:text-gray-400 dark:hover:bg-gray-800"
          aria-label="Toggle mobile menu"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</header>

<!-- Mobile menu -->
{#if isMenuOpen}
  <div class="md:hidden">
    <div class="px-2 pt-2 pb-3 space-y-1 sm:px-3 bg-white dark:bg-gray-900 border-t border-gray-200 dark:border-gray-700">
      <a href="/" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400">
        Search Flights
      </a>
      <a href="/bookings" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400">
        My Bookings
      </a>
      {#if $authStore.isAuthenticated && $authStore.user?.role === 'agent'}
        <a href="/agent" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400">
          Agent Console
        </a>
      {/if}
      {#if $authStore.isAuthenticated && $authStore.user?.role === 'admin'}
        <a href="/admin" class="block px-3 py-2 rounded-md text-base font-medium text-gray-700 dark:text-gray-300 hover:text-primary-600 dark:hover:text-primary-400">
          Admin
        </a>
      {/if}
    </div>
  </div>
{/if}

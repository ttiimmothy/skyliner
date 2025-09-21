<script lang="ts">
  import { goto } from '$app/navigation';
  import { apiClient } from '$lib/api/client';

  let email = $state('');
  let isSubmitting = $state(false);
  let isSubmitted = $state(false);
  let error = $state<string | null>(null);

  async function handleSubmit() {
    if (!email) {
      error = 'Please enter your email address';
      return;
    }

    if (!isValidEmail(email)) {
      error = 'Please enter a valid email address';
      return;
    }

    isSubmitting = true;
    error = null;

    try {
      // In a real app, this would call the API
      // await apiClient.forgotPassword(email);
      
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      isSubmitted = true;
    } catch (err) {
      error = 'Failed to send reset email. Please try again.';
      console.error('Password reset error:', err);
    } finally {
      isSubmitting = false;
    }
  }

  function isValidEmail(email: string) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  }

  function goToLogin() {
    goto('/login');
  }
</script>

<svelte:head>
  <title>Forgot Password - Skyliner</title>
  <meta name="description" content="Reset your Skyliner account password" />
</svelte:head>

<div class="min-h-screen bg-gray-50 dark:bg-gray-900 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
  <div class="mx-auto w-full max-w-md">
    <div class="flex justify-center">
      <a href="/" class="flex items-center space-x-2">
        <div class="w-10 h-10 bg-primary-600 rounded-lg flex items-center justify-center">
          <span class="text-white font-bold text-xl">S</span>
        </div>
        <span class="text-2xl font-bold text-gray-900 dark:text-white">Skyliner</span>
      </a>
    </div>
    <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900 dark:text-white">
      Forgot your password?
    </h2>
    <p class="mt-2 text-center text-sm text-gray-600 dark:text-gray-400">
      No worries! Enter your email address and we'll send you a link to reset your password.
    </p>
  </div>

  <div class="mt-8 mx-auto w-full max-w-md">
    <div class="bg-white dark:bg-gray-800 py-8 px-4 shadow sm:rounded-lg sm:px-10">
      {#if isSubmitted}
        <!-- Success State -->
        <div class="text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-green-100 dark:bg-green-900">
            <svg class="h-6 w-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
          <h3 class="mt-4 text-lg font-medium text-gray-900 dark:text-white">Check your email</h3>
          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
            We've sent a password reset link to <strong>{email}</strong>
          </p>
          <p class="mt-2 text-sm text-gray-500 dark:text-gray-500">
            Didn't receive the email? Check your spam folder or try again.
          </p>
          
          <div class="mt-6 space-y-3">
            <button
              onclick={() => { isSubmitted = false; email = ''; }}
              class="btn btn-primary w-full"
            >
              Send another email
            </button>
            <button
              onclick={goToLogin}
              class="btn btn-outline w-full"
            >
              Back to login
            </button>
          </div>
        </div>
      {:else}
        <!-- Form State -->
        <form onsubmit={handleSubmit} class="space-y-6">
          <div>
            <label for="email" class="label">Email address</label>
            <input
              id="email"
              name="email"
              type="email"
              bind:value={email}
              required
              class="input {error ? 'border-red-300 dark:border-red-600' : ''}"
              placeholder="Enter your email address"
            />
            {#if error}
              <p class="mt-1 text-sm text-red-600 dark:text-red-400">{error}</p>
            {/if}
          </div>

          <div>
            <button
              type="submit"
              disabled={isSubmitting}
              class="btn btn-primary w-full disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {#if isSubmitting}
                <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white inline" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Sending reset link...
              {:else}
                Send reset link
              {/if}
            </button>
          </div>
        </form>

        <div class="mt-6">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300 dark:border-gray-600"></div>
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-2 bg-white dark:bg-gray-800 text-gray-500 dark:text-gray-400">Remember your password?</span>
            </div>
          </div>

          <div class="mt-6">
            <button
              onclick={goToLogin}
              class="btn btn-outline w-full"
            >
              Back to login
            </button>
          </div>
        </div>
      {/if}
    </div>

    <!-- Help Section -->
    <div class="mt-6">
      <div class="bg-blue-50 dark:bg-blue-900 border border-blue-200 dark:border-blue-700 rounded-lg p-4 text-black dark:text-gray-500">
        <div class="flex">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-blue-400" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path>
            </svg>
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-blue-800 dark:text-blue-200">Need help?</h3>
            <div class="mt-2 text-sm text-blue-700 dark:text-blue-300">
              <p>If you're having trouble resetting your password, you can:</p>
              <ul class="mt-1 list-disc list-inside space-y-1">
                <li>Check your spam/junk folder for the reset email</li>
                <li>Make sure you're using the correct email address</li>
                <li>Contact our support team for assistance</li>
              </ul>
            </div>
            <div class="mt-3">
              <a href="/contact" class="text-sm font-medium text-blue-800 dark:text-blue-200 hover:text-blue-600 dark:hover:text-blue-300">
                Contact Support →
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

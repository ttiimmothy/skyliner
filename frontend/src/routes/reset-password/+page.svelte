<script lang="ts">
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { apiClient } from '$lib/api/client';

  let password = $state('');
  let confirmPassword = $state('');
  let isSubmitting = $state(false);
  let isSuccess = $state(false);
  let error = $state<string | null>(null);
  let token = $state('');

  // Get token from URL parameters
  $effect(() => {
    if ($page.url.searchParams.has('token')) {
      token = $page.url.searchParams.get('token') || '';
    }
  });

  async function handleSubmit() {
    if (!password) {
      error = 'Please enter a new password';
      return;
    }

    if (password.length < 8) {
      error = 'Password must be at least 8 characters long';
      return;
    }

    if (password !== confirmPassword) {
      error = 'Passwords do not match';
      return;
    }

    if (!token) {
      error = 'Invalid reset token. Please request a new password reset.';
      return;
    }

    isSubmitting = true;
    error = null;

    try {
      // In a real app, this would call the API
      // await apiClient.resetPassword(token, password);
      
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      isSuccess = true;
    } catch (err) {
      error = 'Failed to reset password. The token may have expired. Please request a new password reset.';
      console.error('Password reset error:', err);
    } finally {
      isSubmitting = false;
    }
  }

  function goToLogin() {
    goto('/login');
  }

  function goToForgotPassword() {
    goto('/forgot-password');
  }
</script>

<svelte:head>
  <title>Reset Password - Skyliner</title>
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
      Reset your password
    </h2>
    <p class="mt-2 text-center text-sm text-gray-600 dark:text-gray-400">
      Enter your new password below
    </p>
  </div>

  <div class="mt-8 mx-auto w-full max-w-md">
    <div class="bg-white dark:bg-gray-800 py-8 px-4 shadow sm:rounded-lg sm:px-10">
      {#if isSuccess}
        <!-- Success State -->
        <div class="text-center">
          <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-green-100 dark:bg-green-900">
            <svg class="h-6 w-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
            </svg>
          </div>
          <h3 class="mt-4 text-lg font-medium text-gray-900 dark:text-white">Password reset successful!</h3>
          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
            Your password has been successfully updated. You can now log in with your new password.
          </p>
          
          <div class="mt-6">
            <button
              onclick={goToLogin}
              class="btn btn-primary w-full"
            >
              Continue to login
            </button>
          </div>
        </div>
      {:else}
        <!-- Form State -->
        <form onsubmit={handleSubmit} class="space-y-6">
          <div>
            <label for="password" class="label">New password</label>
            <input
              id="password"
              name="password"
              type="password"
              bind:value={password}
              required
              class="input {error && error.includes('password') ? 'border-red-300 dark:border-red-600' : ''}"
              placeholder="Enter your new password"
              minlength="8"
            />
            <p class="mt-1 text-xs text-gray-500 dark:text-gray-400">
              Password must be at least 8 characters long
            </p>
          </div>

          <div>
            <label for="confirmPassword" class="label">Confirm new password</label>
            <input
              id="confirmPassword"
              name="confirmPassword"
              type="password"
              bind:value={confirmPassword}
              required
              class="input {error && error.includes('match') ? 'border-red-300 dark:border-red-600' : ''}"
              placeholder="Confirm your new password"
            />
          </div>

          {#if error}
            <div class="bg-red-50 dark:bg-red-900 border border-red-200 dark:border-red-700 rounded-lg p-3">
              <div class="flex">
                <svg class="h-5 w-5 text-red-400" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"></path>
                </svg>
                <div class="ml-3">
                  <p class="text-sm text-red-800 dark:text-red-200">{error}</p>
                </div>
              </div>
            </div>
          {/if}

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
                Resetting password...
              {:else}
                Reset password
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
              <span class="px-2 bg-white dark:bg-gray-800 text-gray-500 dark:text-gray-400">Need help?</span>
            </div>
          </div>

          <div class="mt-6 text-center">
            <button
              onclick={goToForgotPassword}
              class="text-sm text-primary-600 dark:text-primary-400 hover:text-primary-500 dark:hover:text-primary-300"
            >
              Request a new reset link
            </button>
          </div>
        </div>
      {/if}
    </div>

    <!-- Security Notice -->
    <div class="mt-6">
      <div class="bg-yellow-50 dark:bg-blue-900 border border-yellow-200 dark:border-yellow-700 rounded-lg p-4 text:black dark:text-gray-500">
        <div class="flex">
          <div class="flex-shrink-0">
            <svg class="h-5 w-5 text-yellow-400" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
            </svg>
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-yellow-800 dark:text-yellow-200">Security notice</h3>
            <div class="mt-2 text-sm text-yellow-700 dark:text-yellow-300">
              <p>For your security, this password reset link will expire in 1 hour. If you need a new link, please request another password reset.</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

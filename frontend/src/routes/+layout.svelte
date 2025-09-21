<script lang="ts">
  import favicon from '$lib/assets/favicon.svg';
	import '../app.css';
  import { onMount } from 'svelte';
  import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
  import { authStore } from '$lib/stores/auth';
  import Header from '$lib/components/Header.svelte';
  import Footer from '$lib/components/Footer.svelte';

  let queryClient: QueryClient;

  onMount(() => {
    queryClient = new QueryClient({
      defaultOptions: {
        queries: {
          staleTime: 5 * 60 * 1000, // 5 minutes
          retry: 1
        }
      }
    });
  });
</script>

<svelte:head>
	<!-- <link rel="icon" href="/favicon.svg" type="image/svg+xml" />
	<link rel="icon" href="/favicon.ico" type="image/x-icon" /> -->
  <link rel="icon" href={favicon} />
  <title>Skyliner - Flight Booking Platform</title>
  <meta name="description" content="Book flights with ease on Skyliner" />
</svelte:head>

{#if queryClient}
  <QueryClientProvider client={queryClient}>
    <div class="min-h-screen bg-gray-50 dark:bg-gray-900 flex flex-col">
      <Header />
      <main class="flex-1">
        <slot />
      </main>
      <Footer />
    </div>
  </QueryClientProvider>
{/if}
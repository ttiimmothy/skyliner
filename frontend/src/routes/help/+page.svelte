<script lang="ts">
  let searchQuery = $state('');
  let selectedTopic = $state('all');

  const helpTopics = [
    {
      id: 'getting-started',
      title: 'Getting Started',
      icon: '🚀',
      description: 'Learn how to use Skyliner for the first time',
      articles: [
        {
          title: 'How to create an account',
          description: 'Step-by-step guide to setting up your Skyliner account',
          readTime: '2 min read'
        },
        {
          title: 'How to search for flights',
          description: 'Learn how to find the best flight deals',
          readTime: '3 min read'
        },
        {
          title: 'Understanding fare types',
          description: 'Basic economy vs standard vs flexible fares explained',
          readTime: '4 min read'
        },
        {
          title: 'Setting up notifications',
          description: 'Get alerts for price drops and flight updates',
          readTime: '2 min read'
        }
      ]
    },
    {
      id: 'booking',
      title: 'Booking Flights',
      icon: '✈️',
      description: 'Everything about booking and managing your flights',
      articles: [
        {
          title: 'How to book a flight',
          description: 'Complete guide to booking flights on Skyliner',
          readTime: '5 min read'
        },
        {
          title: 'Payment methods accepted',
          description: 'Learn about our secure payment options',
          readTime: '2 min read'
        },
        {
          title: 'Booking for multiple passengers',
          description: 'How to book flights for groups and families',
          readTime: '3 min read'
        },
        {
          title: 'Seat selection guide',
          description: 'How to choose your preferred seats',
          readTime: '3 min read'
        },
        {
          title: 'Adding special requests',
          description: 'Meal preferences, wheelchair assistance, and more',
          readTime: '2 min read'
        }
      ]
    },
    {
      id: 'managing',
      title: 'Managing Bookings',
      icon: '📋',
      description: 'Change, cancel, or modify your existing bookings',
      articles: [
        {
          title: 'How to view your bookings',
          description: 'Access and manage all your flight reservations',
          readTime: '2 min read'
        },
        {
          title: 'Changing flight dates',
          description: 'Step-by-step guide to modifying your travel dates',
          readTime: '4 min read'
        },
        {
          title: 'Cancelling a booking',
          description: 'How to cancel flights and understand refund policies',
          readTime: '3 min read'
        },
        {
          title: 'Adding extra baggage',
          description: 'How to purchase additional baggage allowance',
          readTime: '2 min read'
        },
        {
          title: 'Upgrading your seat',
          description: 'How to upgrade to premium economy or business class',
          readTime: '3 min read'
        }
      ]
    },
    {
      id: 'checkin',
      title: 'Check-in & Travel',
      icon: '🎫',
      description: 'Prepare for your journey with check-in and travel tips',
      articles: [
        {
          title: 'Online check-in guide',
          description: 'How to check in online and get your boarding pass',
          readTime: '3 min read'
        },
        {
          title: 'What to bring to the airport',
          description: 'Essential documents and items for your trip',
          readTime: '4 min read'
        },
        {
          title: 'Baggage policies explained',
          description: 'Understanding carry-on and checked baggage rules',
          readTime: '5 min read'
        },
        {
          title: 'Travel tips and advice',
          description: 'Expert tips for a smooth travel experience',
          readTime: '6 min read'
        },
        {
          title: 'What if I miss my flight?',
          description: 'Steps to take if you miss your scheduled departure',
          readTime: '3 min read'
        }
      ]
    },
    {
      id: 'account',
      title: 'Account & Profile',
      icon: '👤',
      description: 'Manage your account settings and preferences',
      articles: [
        {
          title: 'Updating your profile',
          description: 'How to change your personal information',
          readTime: '2 min read'
        },
        {
          title: 'Changing your password',
          description: 'Keep your account secure with a strong password',
          readTime: '2 min read'
        },
        {
          title: 'Managing notifications',
          description: 'Control which emails and alerts you receive',
          readTime: '2 min read'
        },
        {
          title: 'Linking frequent flyer accounts',
          description: 'Connect your airline loyalty programs',
          readTime: '3 min read'
        },
        {
          title: 'Deleting your account',
          description: 'How to permanently delete your Skyliner account',
          readTime: '2 min read'
        }
      ]
    },
    {
      id: 'technical',
      title: 'Technical Support',
      icon: '🔧',
      description: 'Troubleshooting and technical assistance',
      articles: [
        {
          title: 'Website not loading properly',
          description: 'Common solutions for website issues',
          readTime: '3 min read'
        },
        {
          title: 'Mobile app troubleshooting',
          description: 'Fix common issues with our mobile app',
          readTime: '4 min read'
        },
        {
          title: 'Email not received',
          description: 'What to do if you don\'t receive confirmation emails',
          readTime: '2 min read'
        },
        {
          title: 'Browser compatibility',
          description: 'Supported browsers and system requirements',
          readTime: '2 min read'
        },
        {
          title: 'Two-factor authentication',
          description: 'Set up and manage 2FA for enhanced security',
          readTime: '3 min read'
        }
      ]
    }
  ];

  function filteredTopics() {
    let filtered = helpTopics;
    
    if (selectedTopic !== 'all') {
      filtered = helpTopics.filter(topic => topic.id === selectedTopic);
    }
    
    if (searchQuery) {
      filtered = filtered.map(topic => ({
        ...topic,
        articles: topic.articles.filter(article => 
          article.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
          article.description.toLowerCase().includes(searchQuery.toLowerCase())
        )
      })).filter(topic => topic.articles.length > 0);
    }
    
    return filtered;
  }

  const topicOptions = [
    { value: 'all', label: 'All Topics' },
    { value: 'getting-started', label: 'Getting Started' },
    { value: 'booking', label: 'Booking Flights' },
    { value: 'managing', label: 'Managing Bookings' },
    { value: 'checkin', label: 'Check-in & Travel' },
    { value: 'account', label: 'Account & Profile' },
    { value: 'technical', label: 'Technical Support' }
  ];
</script>

<svelte:head>
  <title>Help Center - Skyliner</title>
  <meta name="description" content="Get help with booking flights, managing reservations, and using Skyliner's features" />
</svelte:head>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <div class="text-center mb-12">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-4">Help Center</h1>
    <p class="text-lg text-gray-600 dark:text-gray-400 mb-8">
      Find answers, guides, and support for all your travel needs
    </p>

    <!-- Search Bar -->
    <div class="max-w-2xl mx-auto">
      <div class="relative">
        <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
          <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
          </svg>
        </div>
        <input
          type="text"
          bind:value={searchQuery}
          placeholder="Search help articles..."
          class="input pl-12 text-lg"
        />
      </div>
    </div>
  </div>

  <!-- Quick Actions -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
    <a href="/faq" class="card hover:shadow-lg transition-shadow duration-200">
      <div class="flex items-center space-x-4">
        <div class="w-12 h-12 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white">FAQ</h3>
          <p class="text-gray-600 dark:text-gray-400">Quick answers to common questions</p>
        </div>
      </div>
    </a>

    <a href="/contact" class="card hover:shadow-lg transition-shadow duration-200">
      <div class="flex items-center space-x-4">
        <div class="w-12 h-12 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Contact Us</h3>
          <p class="text-gray-600 dark:text-gray-400">Get in touch with our support team</p>
        </div>
      </div>
    </a>

    <a href="/bookings" class="card hover:shadow-lg transition-shadow duration-200">
      <div class="flex items-center space-x-4">
        <div class="w-12 h-12 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center">
          <svg class="w-6 h-6 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-lg font-semibold text-gray-900 dark:text-white">My Bookings</h3>
          <p class="text-gray-600 dark:text-gray-400">Manage your flight reservations</p>
        </div>
      </div>
    </a>
  </div>

  <!-- Filter -->
  <div class="mb-8">
    <div class="flex flex-col sm:flex-row gap-4 items-center justify-between">
      <h2 class="text-xl font-semibold text-gray-900 dark:text-white">Help Articles</h2>
      <select bind:value={selectedTopic} class="input md:w-64">
        {#each topicOptions as option}
          <option value={option.value}>{option.label}</option>
        {/each}
      </select>
    </div>
  </div>

  <!-- Help Topics -->
  <div class="space-y-8">
    {#each filteredTopics() as topic}
      <div class="card">
        <div class="flex items-start space-x-4 mb-6">
          <div class="text-3xl">{topic.icon}</div>
          <div>
            <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-2">{topic.title}</h3>
            <p class="text-gray-600 dark:text-gray-400">{topic.description}</p>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          {#each topic.articles as article}
            <a href="/help/article" class="block p-4 border border-gray-200 dark:border-gray-700 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors duration-200">
              <h4 class="font-medium text-gray-900 dark:text-white mb-2">{article.title}</h4>
              <p class="text-sm text-gray-600 dark:text-gray-400 mb-2">{article.description}</p>
              <span class="text-xs text-gray-500 dark:text-gray-500">{article.readTime}</span>
            </a>
          {/each}
        </div>
      </div>
    {/each}
  </div>

  <!-- No Results -->
  {#if filteredTopics().length === 0}
    <div class="text-center py-12">
      <svg class="mx-auto h-12 w-12 text-gray-400 dark:text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.29-1.009-5.824-2.57M15 6.75a3 3 0 11-6 0 3 3 0 016 0z"></path>
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900 dark:text-white">No articles found</h3>
      <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">
        Try adjusting your search terms or browse all topics.
      </p>
    </div>
  {/if}

  <!-- Contact Support -->
  <div class="mt-12 bg-primary-50 dark:bg-gray-800 rounded-lg p-8">
    <div class="text-center">
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">Need more help?</h3>
      <p class="text-gray-600 dark:text-gray-400 mb-6">
        Our support team is available 24/7 to assist you with any questions or issues.
      </p>
      <div class="flex flex-col sm:flex-row gap-4 justify-center">
        <a href="/contact" class="btn btn-primary">
          Contact Support
        </a>
        <a href="tel:+15551234567" class="btn btn-outline">
          Call Us: (555) 123-4567
        </a>
      </div>
    </div>
  </div>
</div>

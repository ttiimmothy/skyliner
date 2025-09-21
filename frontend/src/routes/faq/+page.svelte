<script lang="ts">
  let searchQuery = $state('');
  let selectedCategory = $state('all');
  let openItems = $state<Set<number>>(new Set());

  const faqData = [
    {
      category: 'booking',
      title: 'Booking & Reservations',
      items: [
        {
          question: 'How do I book a flight?',
          answer: 'To book a flight, simply search for your desired destination and dates on our homepage, select your preferred flight, and follow the booking process. You can also use our mobile app for a seamless booking experience.'
        },
        {
          question: 'Can I book a flight for someone else?',
          answer: 'Yes, you can book flights for other people. During the booking process, you\'ll be asked to enter the passenger details. Make sure to use the correct information as it appears on their ID or passport.'
        },
        {
          question: 'How far in advance can I book a flight?',
          answer: 'You can typically book flights up to 11 months in advance, depending on the airline. Some airlines may allow bookings even further ahead. We recommend booking at least 2-3 weeks in advance for the best prices.'
        },
        {
          question: 'Can I hold a reservation before paying?',
          answer: 'Unfortunately, we cannot hold reservations without payment. All bookings must be completed with full payment at the time of reservation to secure your seat.'
        },
        {
          question: 'What payment methods do you accept?',
          answer: 'We accept all major credit cards (Visa, MasterCard, American Express), PayPal, and other secure payment methods. All transactions are encrypted and secure.'
        }
      ]
    },
    {
      category: 'cancellation',
      title: 'Cancellations & Refunds',
      items: [
        {
          question: 'Can I cancel my booking?',
          answer: 'Yes, you can cancel your booking through your account dashboard or by contacting our support team. However, cancellation policies vary by airline and fare type. Some tickets may be non-refundable.'
        },
        {
          question: 'How long does it take to process a refund?',
          answer: 'Refunds are typically processed within 7-14 business days after cancellation approval. The exact timeframe depends on the airline and your payment method. You\'ll receive an email confirmation once processed.'
        },
        {
          question: 'Are there cancellation fees?',
          answer: 'Cancellation fees depend on the airline and fare type. Basic economy tickets are usually non-refundable, while flexible fares may allow free cancellations. Check your booking details for specific terms.'
        },
        {
          question: 'Can I get a refund for a non-refundable ticket?',
          answer: 'Non-refundable tickets cannot be refunded, but you may be able to change your travel dates for a fee, depending on the airline\'s policy. Contact our support team for assistance.'
        }
      ]
    },
    {
      category: 'changes',
      title: 'Changes & Modifications',
      items: [
        {
          question: 'Can I change my flight date or time?',
          answer: 'Yes, you can modify your booking through your account dashboard or by contacting our support team. Changes may be subject to airline fees and fare differences.'
        },
        {
          question: 'How much does it cost to change a flight?',
          answer: 'Change fees vary by airline and fare type. Basic economy tickets usually have higher change fees, while flexible fares may allow free changes. You\'ll also need to pay any fare difference.'
        },
        {
          question: 'Can I change the passenger name on my booking?',
          answer: 'Name changes are generally not allowed on airline tickets due to security reasons. If you need to change a passenger, you may need to cancel and rebook. Contact our support team for assistance.'
        },
        {
          question: 'Can I add extra baggage after booking?',
          answer: 'Yes, you can add extra baggage through your account dashboard or by contacting the airline directly. Additional baggage fees will apply.'
        }
      ]
    },
    {
      category: 'checkin',
      title: 'Check-in & Boarding',
      items: [
        {
          question: 'When can I check in for my flight?',
          answer: 'Online check-in typically opens 24 hours before departure for domestic flights and 24-48 hours for international flights. You can check in through our website, mobile app, or directly with the airline.'
        },
        {
          question: 'What documents do I need for check-in?',
          answer: 'You\'ll need a valid government-issued photo ID (driver\'s license, passport) and your booking confirmation. For international flights, you\'ll also need a valid passport and any required visas.'
        },
        {
          question: 'Can I select my seat during check-in?',
          answer: 'Yes, you can usually select your seat during online check-in, subject to availability. Some airlines may charge for preferred seats or allow free seat selection based on your fare type.'
        },
        {
          question: 'What if I miss my flight?',
          answer: 'If you miss your flight, contact the airline immediately. Depending on the circumstances and fare type, you may be able to rebook on the next available flight for a fee.'
        }
      ]
    },
    {
      category: 'technical',
      title: 'Technical Support',
      items: [
        {
          question: 'I\'m having trouble logging into my account. What should I do?',
          answer: 'Try resetting your password using the "Forgot Password" link on the login page. If you continue to have issues, contact our support team with your email address and we\'ll help you regain access.'
        },
        {
          question: 'The website is not loading properly. What can I do?',
          answer: 'Try clearing your browser cache and cookies, or try using a different browser. If the problem persists, check your internet connection or contact our technical support team.'
        },
        {
          question: 'I didn\'t receive my booking confirmation email. What should I do?',
          answer: 'Check your spam/junk folder first. If you still don\'t see it, log into your account to view your booking details, or contact our support team to resend the confirmation.'
        },
        {
          question: 'Can I use the mobile app on multiple devices?',
          answer: 'Yes, you can use your account on multiple devices. Simply download the app and log in with your credentials. Your bookings and preferences will sync across all devices.'
        }
      ]
    }
  ];

  function toggleItem(index: number) {
    if (openItems.has(index)) {
      openItems.delete(index);
    } else {
      openItems.add(index);
    }
    openItems = new Set(openItems);
  }

  function filteredFAQs() {
    let filtered = faqData;
    
    if (selectedCategory !== 'all') {
      filtered = faqData.filter(category => category.category === selectedCategory);
    }
    
    if (searchQuery) {
      filtered = filtered.map(category => ({
        ...category,
        items: category.items.filter(item => 
          item.question.toLowerCase().includes(searchQuery.toLowerCase()) ||
          item.answer.toLowerCase().includes(searchQuery.toLowerCase())
        )
      })).filter(category => category.items.length > 0);
    }
    
    return filtered;
  }

  const categories = [
    { value: 'all', label: 'All Questions' },
    { value: 'booking', label: 'Booking & Reservations' },
    { value: 'cancellation', label: 'Cancellations & Refunds' },
    { value: 'changes', label: 'Changes & Modifications' },
    { value: 'checkin', label: 'Check-in & Boarding' },
    { value: 'technical', label: 'Technical Support' }
  ];
</script>

<svelte:head>
  <title>FAQ - Frequently Asked Questions | Skyliner</title>
  <meta name="description" content="Find answers to common questions about booking flights, cancellations, check-in, and more on Skyliner" />
</svelte:head>

<div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <div class="text-center mb-12">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-4">Frequently Asked Questions</h1>
    <p class="text-lg text-gray-600 dark:text-gray-400">
      Find answers to common questions about our flight booking service
    </p>
  </div>

  <!-- Search and Filter -->
  <div class="mb-8">
    <div class="flex flex-col md:flex-row gap-4 mb-6">
      <!-- Search -->
      <div class="flex-1">
        <div class="relative">
          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
          <input
            type="text"
            bind:value={searchQuery}
            placeholder="Search FAQs..."
            class="input pl-10"
          />
        </div>
      </div>

      <!-- Category Filter -->
      <div class="md:w-64">
        <select bind:value={selectedCategory} class="input">
          {#each categories as category}
            <option value={category.value}>{category.label}</option>
          {/each}
        </select>
      </div>
    </div>
  </div>

  <!-- FAQ Content -->
  <div class="space-y-8">
    {#each filteredFAQs() as category}
      <div class="card">
        <h2 class="text-xl font-semibold text-gray-900 dark:text-white mb-6">{category.title}</h2>
        
        <div class="space-y-4">
          {#each category.items as item, index}
            <div class="border border-gray-200 dark:border-gray-700 rounded-lg">
              <button
                onclick={() => toggleItem(index)}
                class="w-full px-6 py-4 text-left flex justify-between items-center hover:bg-gray-50 dark:hover:bg-gray-800 transition-colors duration-200"
              >
                <span class="font-medium text-gray-900 dark:text-white pr-4">{item.question}</span>
                <svg 
                  class="w-5 h-5 text-gray-500 dark:text-gray-400 transform transition-transform duration-200 {openItems.has(index) ? 'rotate-180' : ''}"
                  fill="none" 
                  stroke="currentColor" 
                  viewBox="0 0 24 24"
                >
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>
              
              {#if openItems.has(index)}
                <div class="px-6 pb-4">
                  <p class="text-gray-700 dark:text-gray-300 leading-relaxed">{item.answer}</p>
                </div>
              {/if}
            </div>
          {/each}
        </div>
      </div>
    {/each}
  </div>

  <!-- No Results -->
  {#if filteredFAQs().length === 0}
    <div class="text-center py-12">
      <svg class="mx-auto h-12 w-12 text-gray-400 dark:text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 12h6m-6-4h6m2 5.291A7.962 7.962 0 0112 15c-2.34 0-4.29-1.009-5.824-2.57M15 6.75a3 3 0 11-6 0 3 3 0 016 0z"></path>
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900 dark:text-white">No FAQs found</h3>
      <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">
        Try adjusting your search terms or browse all categories.
      </p>
    </div>
  {/if}

  <!-- Contact Support -->
  <div class="mt-12 bg-primary-50 dark:bg-gray-800 rounded-lg p-8">
    <div class="text-center">
      <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-2">Still have questions?</h3>
      <p class="text-gray-600 dark:text-gray-400 mb-6">
        Our support team is here to help you with any questions or concerns.
      </p>
      <div class="flex flex-col sm:flex-row gap-4 justify-center">
        <a href="/contact" class="btn btn-primary">
          Contact Support
        </a>
        <a href="/help" class="btn btn-outline">
          Help Center
        </a>
      </div>
    </div>
  </div>
</div>

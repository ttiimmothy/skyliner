<script lang="ts">
  let formData = $state({
    name: '',
    email: '',
    subject: '',
    message: ''
  });
  
  let isSubmitting = $state(false);
  let isSubmitted = $state(false);

  async function handleSubmit() {
    isSubmitting = true;
    
    // Simulate form submission
    await new Promise(resolve => setTimeout(resolve, 1000));
    
    isSubmitted = true;
    isSubmitting = false;
    
    // Reset form
    formData = {
      name: '',
      email: '',
      subject: '',
      message: ''
    };
  }
</script>

<svelte:head>
  <title>Contact Us - Skyliner</title>
  <meta name="description" content="Get in touch with Skyliner customer support for assistance with your flight bookings" />
</svelte:head>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
  <div class="max-w-3xl mx-auto">
    <div class="text-center mb-12">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-4">Contact Us</h1>
      <p class="text-lg text-gray-600 dark:text-gray-400">
        We're here to help! Get in touch with our customer support team for any questions or assistance.
      </p>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-12">
      <!-- Contact Information -->
      <div>
        <h2 class="text-2xl font-semibold text-gray-900 dark:text-white mb-6">Get in Touch</h2>
        
        <div class="space-y-6">
          <div class="flex items-start space-x-4">
            <div class="w-10 h-10 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center flex-shrink-0">
              <svg class="w-5 h-5 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">Email Support</h3>
              <p class="text-gray-600 dark:text-gray-400">support@skyliner.com</p>
              <p class="text-sm text-gray-500 dark:text-gray-500">We'll respond within 24 hours</p>
            </div>
          </div>

          <div class="flex items-start space-x-4">
            <div class="w-10 h-10 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center flex-shrink-0">
              <svg class="w-5 h-5 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">Phone Support</h3>
              <p class="text-gray-600 dark:text-gray-400">+1 (555) 123-4567</p>
              <p class="text-sm text-gray-500 dark:text-gray-500">Mon-Fri 9AM-6PM EST</p>
            </div>
          </div>

          <div class="flex items-start space-x-4">
            <div class="w-10 h-10 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center flex-shrink-0">
              <svg class="w-5 h-5 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">Office Address</h3>
              <p class="text-gray-600 dark:text-gray-400">
                Skyliner Inc.<br>
                123 Aviation Way<br>
                New York, NY 10001<br>
                United States
              </p>
            </div>
          </div>

          <div class="flex items-start space-x-4">
            <div class="w-10 h-10 bg-primary-100 dark:bg-primary-900 rounded-lg flex items-center justify-center flex-shrink-0">
              <svg class="w-5 h-5 text-primary-600 dark:text-primary-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-lg font-medium text-gray-900 dark:text-white">Live Chat</h3>
              <p class="text-gray-600 dark:text-gray-400">Available 24/7</p>
              <p class="text-sm text-gray-500 dark:text-gray-500">Get instant help from our support team</p>
            </div>
          </div>
        </div>

        <!-- FAQ Section -->
        <div class="mt-12">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white mb-4">Frequently Asked Questions</h3>
          <div class="space-y-4">
            <div class="border border-gray-200 dark:border-gray-700 rounded-lg p-4">
              <h4 class="font-medium text-gray-900 dark:text-white mb-2">How do I cancel my booking?</h4>
              <p class="text-sm text-gray-600 dark:text-gray-400">
                You can cancel your booking through your account dashboard or by contacting our support team. Cancellation policies vary by airline and fare type.
              </p>
            </div>
            <div class="border border-gray-200 dark:border-gray-700 rounded-lg p-4">
              <h4 class="font-medium text-gray-900 dark:text-white mb-2">What payment methods do you accept?</h4>
              <p class="text-sm text-gray-600 dark:text-gray-400">
                We accept all major credit cards, PayPal, and other secure payment methods. All transactions are encrypted and secure.
              </p>
            </div>
            <div class="border border-gray-200 dark:border-gray-700 rounded-lg p-4">
              <h4 class="font-medium text-gray-900 dark:text-white mb-2">How can I modify my booking?</h4>
              <p class="text-sm text-gray-600 dark:text-gray-400">
                You can modify your booking through your account dashboard or by contacting our support team. Changes may be subject to airline policies and fees.
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Contact Form -->
      <div>
        <h2 class="text-2xl font-semibold text-gray-900 dark:text-white mb-6">Send us a Message</h2>
        
        {#if isSubmitted}
          <div class="bg-green-50 dark:bg-green-900 border border-green-200 dark:border-green-700 rounded-lg p-6 mb-6">
            <div class="flex">
              <svg class="h-5 w-5 text-green-400" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
              </svg>
              <div class="ml-3">
                <h3 class="text-sm font-medium text-green-800 dark:text-green-200">Message sent successfully!</h3>
                <p class="mt-1 text-sm text-green-700 dark:text-green-300">We'll get back to you as soon as possible.</p>
              </div>
            </div>
          </div>
        {/if}

        <form onsubmit={handleSubmit} class="space-y-6">
          <div>
            <label for="name" class="label">Full Name</label>
            <input
              type="text"
              id="name"
              bind:value={formData.name}
              required
              class="input"
              placeholder="Enter your full name"
            />
          </div>

          <div>
            <label for="email" class="label">Email Address</label>
            <input
              type="email"
              id="email"
              bind:value={formData.email}
              required
              class="input"
              placeholder="Enter your email address"
            />
          </div>

          <div>
            <label for="subject" class="label">Subject</label>
            <select
              id="subject"
              bind:value={formData.subject}
              required
              class="input"
            >
              <option value="">Select a subject</option>
              <option value="booking">Booking Support</option>
              <option value="cancellation">Cancellation Request</option>
              <option value="refund">Refund Inquiry</option>
              <option value="technical">Technical Issue</option>
              <option value="general">General Inquiry</option>
              <option value="other">Other</option>
            </select>
          </div>

          <div>
            <label for="message" class="label">Message</label>
            <textarea
              id="message"
              bind:value={formData.message}
              required
              rows="6"
              class="input"
              placeholder="Please describe your inquiry in detail..."
            ></textarea>
          </div>

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
              Sending...
            {:else}
              Send Message
            {/if}
          </button>
        </form>
      </div>
    </div>
  </div>
</div>

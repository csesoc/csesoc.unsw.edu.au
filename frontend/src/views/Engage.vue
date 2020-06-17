<template>
  <v-app>
    <v-parallax
      height="300"
      src="https://images.unsplash.com/photo-1592251170558-01228a0a06f2?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1920&q=80"
    ></v-parallax>

    <v-tabs
      class="elevation-2"
      vertical
      dark
    >
      <v-tabs-slider></v-tabs-slider>

      <!-- Enquiry tab -->
      <v-tab>Enquiry</v-tab>
      <v-tab-item>
        <v-card flat tile>
          <EnquiryForm></EnquiryForm>
        </v-card>
      </v-tab-item>

      <!-- Feedback tab -->
      <v-tab>Feedback</v-tab>
      <v-tab-item>
        <v-card flat tile>
          <FeedbackForm></FeedbackForm>
        </v-card>
      </v-tab-item>
    
    </v-tabs>
  </v-app>
</template>

<script>
import EnquiryForm from '@/components/EnquiryForm';
import FeedbackForm from '@/components/FeedbackForm';

export default {
  data: () => ({

  }),
  components: {
    EnquiryForm,
    FeedbackForm
  }
};
</script>

<style scoped>
</style>

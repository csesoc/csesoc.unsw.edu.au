<template>
  <v-row justify="space-around">
    <v-col class="form-box">
      <v-form
        ref="form"
        v-model="valid"
      >
        <v-text-field
          v-model="name"
          label="Name"
        ></v-text-field>

        <v-text-field
          v-model="email"
          label="E-mail"
        ></v-text-field>

        <v-textarea
          v-model="body"
          :rules="bodyRules"
          label="Message"
          required
        ></v-textarea>

        <v-btn
          :disabled="!valid"
          color="blue accent-4"
          class="mr-4"
          @click="send"
        >
          Send
        </v-btn>
      </v-form>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'FeedbackForm',
  data: () => ({
    valid: true,
    name: '',
    email: '',
    body: '',
    bodyRules: [
      v => !!v || 'Message is required',
    ],
  }),

  methods: {
    send() {
      
    },
  },
}
</script>

<style scoped>
.form-box {
  padding-left: 10%;
  padding-right: 10%;
}
</style>
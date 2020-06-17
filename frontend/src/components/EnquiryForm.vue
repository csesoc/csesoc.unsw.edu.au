<template>
  <v-row justify="space-around">
    <v-col class="form-box">
      <v-form
        ref="form"
        v-model="valid"
      >
        <v-text-field
          v-model="name"
          :rules="nameRules"
          label="Name"
          required
        ></v-text-field>

        <v-text-field
          v-model="email"
          :rules="emailRules"
          label="E-mail"
          required
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
  name: 'ContactForm',
  data: () => ({
    valid: true,
    name: '',
    nameRules: [
      v => !!v || 'Name is required',
    ],
    email: '',
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
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
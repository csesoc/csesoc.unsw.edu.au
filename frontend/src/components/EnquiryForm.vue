<template>
  <v-row justify="space-around">
    <v-col class="form-box">
      <h2 v-if="this.messageSent">
        Message sent!
      </h2>

      <v-form
        v-else
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
import APIClient from '../utils/APIClient'
import {MAILING_URL} from '../utils/Constants'

export default {
  name: 'EnquiryForm',
  props: ['type'],
  data: () => ({
    valid: true,
    messageSent: false,
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
      APIClient.mailingAPI(MAILING_URL[this.type], this.name, this.email, this.body)
      .then((res) => {
        switch (res.status) {
          case 202:
            this.messageSent = true;
            console.log("Message sent: " + res);
            break;
          case 400:
            console.error("Invalid form: " + res);
            break;
          default:
            console.error("Failed to send message: " + res);
        }
      });
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
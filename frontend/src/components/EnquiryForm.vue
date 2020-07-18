<template>
  <v-row justify="space-around">
    <v-col class="form-box">
      <v-form
        ref="form"
        v-model="valid"
      >
        <label class="text-body-1 input-label">{{ this.type == "sponsorship" ? "Company Name" : "Name" }}</label>
        <v-text-field class="input" placeholder="Name" :rules="nameRules" v-model="name"></v-text-field>
        <label class="text-body-1 input-label"> Email </label>
        <v-text-field class="input" placeholder="John.smith@gmail.com" :rules="emailRules" v-model="email"></v-text-field>
        <label class="text-body-1 input-label"> Message </label>
        <v-textarea class="input" placeholder="Body" :rules="bodyRules" v-model="body"></v-textarea>
        <v-btn text style="margin-left:60%" :disabled="!valid" @click="send">Send</v-btn>
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
            this.body = "";
            this.email = "";
            this.name = "";
            console.log("Message sent: " + res);
            break;
          case 400:
            console.error("Invalid form: " + res);
            console.log(res);
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

.input {
  margin-left:15%;
  width:50%;
}

.input-label {
  padding-top:20px;
  float:left;
}
</style>
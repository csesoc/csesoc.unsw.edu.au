<!--
  MailingForm
  --
  This component contains an enquiry form which can be used to interact with mailing API's endpoints.
  --
  Props:
    - type: string - either "general", "sponsorship" or "feedback"
-->

<template>
  <v-row justify="space-around">
    <v-col class="form-box">
      <v-form ref="form" v-model="valid">
        <!-- Name -->
        <label class="text-body-1 input-label" :data-cy="`${this.type}-name-label`">
          {{ isType('sponsorship') ? "Company Name" : "Name" }}
          <span v-if="!isType('feedback')" class="required">*</span>
        </label>
        <v-text-field class="input" placeholder="John Smith" v-model="name" :data-cy="`${this.type}-name-field`"
          :rules="[ !isType('feedback') ? rules.required : true ]"></v-text-field>
        <!-- Email -->
        <label class="text-body-1 input-label" :data-cy="`${this.type}-email-label`"> Email
          <span v-if="!isType('feedback')" class="required">*</span>
        </label>
        <v-text-field class="input" placeholder="john.smith@email.com" v-model="email" :data-cy="`${this.type}-email-field`"
          :rules="[ rules.email, !isType('feedback') ? rules.required : true ]"></v-text-field>
        <!-- Message -->
        <label class="text-body-1 input-label" :data-cy="`${this.type}-message-label`"> Message
          <span class="required">*</span>
        </label>
        <v-textarea class="input" placeholder="Message" v-model="body" :data-cy="`${this.type}-message-field`"
          :rules="[ rules.required ]"></v-textarea>
        <!-- Send button -->
        <v-btn text style="margin-left:60%" :disabled="!valid" @click="send" :data-cy="`${this.type}-send-button`">Send</v-btn>
      </v-form>
    </v-col>
  </v-row>
</template>

<script>
import isEmail from 'validator/lib/isEmail';

import APIClient from '../utils/APIClient';
import { MAILING_URL } from '../utils/Constants';

export default {
  name: 'MailingForm',
  props: ['type'],
  data: () => ({
    valid: true,
    name: '',
    email: '',
    body: '',
    rules: {
      required: (value) => !!value || 'Required',
      email: (value) => value.length === 0 || isEmail(value) || 'Invalid e-mail',
    },
  }),
  methods: {
    isType(name) {
      return this.type === name;
    },
    send() {
      APIClient.mailingAPI(MAILING_URL[this.type], this.name, this.email, this.body)
        .then((res) => {
          switch (res.status) {
            case 202:
              this.$refs.form.reset();
              console.log(`Message sent: ${res}`);
              break;
            case 400:
              console.error(`Invalid form: ${res}`);
              console.log(res);
              break;
            default:
              console.error(`Failed to send message: ${res}`);
          }
        });
    },
  }
};
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

.required
{
  color: red;
}
</style>

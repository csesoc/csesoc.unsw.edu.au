 <template>
  <form id="login-form">
    <v-btn color="normal" text dark @click="dialog = true">Login</v-btn>
    <v-layout row justify-center>
      <v-dialog
        v-model="dialog"
        dark
        no-click-animation
        persistent
        max-width="600px"
      >
        <v-card>
          <v-form ref="form">
            <v-card-title>
              <span class="headline" colour="white">Login</span>
            </v-card-title>

            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex xs12>
                    <v-text-field
                      v-model="zid"
                      :rules="inputRules"
                      label="zID"
                      required
                    ></v-text-field>
                  </v-flex>

                  <v-flex xs12>
                    <v-text-field
                      :type="show ? 'text' : 'password'"
                      v-model="password"
                      @click:append="show = !show"
                      :rules="inputRules"
                      label="zPass"
                      required
                    ></v-text-field>
                  </v-flex>
                </v-layout>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                style="min-width:100px;min-height:50px;"
                text
                @click="dialog = false"
                >Close</v-btn
              >
              <v-btn
                color="blue darken-1"
                style="min-width:100px;min-height:50px;"
                text
                @click="login"
                >Submit</v-btn
              >
            </v-card-actions>
          </v-form>
        </v-card>
      </v-dialog>
    </v-layout>
  </form>
</template>

<script>

export default {
  data() {
    return {
      zid: '',
      password: '',
      dialog: false,
      show: false,
      inputRules: [
        v => v.trim() !== '' || 'You cannot leave this field empty',
      ],
      success: false,
    };
  },
  methods: {
    login() {
      console.log(`Hi ${this.zid}!`);
    },
  },
};
</script>

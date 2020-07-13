<template>
  <div>
    <!-- make header a seperate component! -->
    <header id="showcase">
      <v-img max-width="80vw" max-height="30vh" contain src="@/assets/csesocwhiteblue.png" />
    </header>
    <v-container class="margin" fluid>
      <h1 class="border text-h1 font-weight-bold" style="padding:25px;">Sponsors</h1>
      <h2 class="text-h4">Principal <br> Sponsors</h2>
      <v-container class="border" fluid>
          <div v-for="(sponsor, index) in tierOne" :key="sponsor.id" :style="marginStyle(index)">
              <div class="logo-margin">
                <img class="large-logo" :src="sponsor.logo" v-on:click="onClickModal(sponsor)"/>
              </div>
              <div v-if="(index + 1) % largeLogoFilter === 0 || index === tierOne.length - 1" style="clear:both;"></div>
          </div>
      </v-container>
      <div style="clear:both;"></div>
      <h2 class="text-h4">Major <br> Sponsors</h2>
      <v-container class="border" fluid>
          <div v-for="(sponsor, index) in tierTwo" :key="sponsor.id" :style="marginStyle(index)">
              <div class="logo-margin">
                <img class="mid-logo" :src="sponsor.logo" v-on:click="onClickModal(sponsor)"/>
              </div>
              <div v-if="(index + 1) % midLogoFilter === 0 || index === tierTwo.length - 1" style="clear:both;"></div>
          </div>
      </v-container>
      <div style="clear:both;"></div>
      <h2 class="text-h4">Affiliate <br> Sponsors</h2>
      <v-container class="border" fluid>
          <div v-for="(sponsor, index) in tierThree" :key="sponsor.id" :style="marginStyle(index)">
              <div class="logo-margin">
                  <img class="small-logo" :src="sponsor.logo" v-on:click="onClickModal(sponsor)"/>
              </div>
              <div v-if="(index + 1) % smallLogoFilter === 0 || index === tierThree.length - 1" style="clear:both;"></div>
          </div>
      </v-container>
      <div style="clear:both;"></div>
    </v-container>
    <SponsorModal v-bind:dialog="dialog" v-bind:title="currentSponsor.name" v-bind:body="currentSponsor.detail"> </SponsorModal>
    
    <h1 class="text-center text-h1 font-weight-bold">Become a <br> Sponsor</h1>
    <v-container fluid style="margin-left:15%">
      <label class="text-body-1 input-label"> Company Name </label>
      <v-text-field class="input" placeholder="Company Name" v-bind:value="newCompanyName" v-model="newCompanyName"></v-text-field>
      <label class="text-body-1 input-label" style="margin-right:150px;"> Email </label>
      <v-text-field class="input" placeholder="John.smith@gmail.com" v-bind:value="newCompanyName" v-model="newEmail"></v-text-field>
      <label class="text-body-1 input-label" style="margin-right:150px;"> Message </label>
      <v-textarea class="input" placeholder="Body" v-bind:value="newCompanyName" v-model="newMessage"></v-textarea>
      <v-btn text style="margin-left:62%" @click="onClickSend()">Send</v-btn>
    </v-container>
  </div>
</template>
  
<script type="text/javascript">
import APIClient  from '../utils/APIClient'
import SponsorModal from '@/components/SponsorModal'
import {MAILING_URL} from '../utils/Constants'
import Vue from 'vue'

export default {
  data: () => ({
    currentSponsor: {},
    sponsors: [],
    dialog: false,
    newCompanyName: "",
    newEmail: "",
    newMessage: "",

    //Constants
    largeLogoFilter: 3,
    midLogoFilter: 4,
    smallLogoFilter: 5
  }),
  components: {
    SponsorModal
  },
  computed: {
    // functions to determine sizing category of sponsor based on their value
    tierOne() {
      return this.sponsors.filter(x => x.tier == 2);
    },
    tierTwo() {
      return this.sponsors.filter(x => x.tier == 1);
    },
    tierThree() {
      return this.sponsors.filter(x => x.tier == 0);
    }
  },
  mounted() {
      APIClient.fetchSponsors()
      .then((responseJson) => {
        this.sponsors = responseJson;
      });
  },
  methods: {
    marginStyle: function(index) {
      var style = {};
      
      var limit = parseInt((index) / 3);
      if (limit % 2 == 0) {
        style["margin-left"] = "10%";
      } else {
        style["margin-left"] = "15%";
      }
      
      return style;
    },
    onClickModal: function (sponsor) {
      this.currentSponsor = sponsor;
      this.dialog = true;
    },
    onClickSend: function () {
      APIClient.mailingAPI(MAILING_URL['sponsorship'], this.newCompanyName, this.newEmail, this.newMessage)
      .then((res) => {
        switch (res.status) {
          case 202:
            this.newCompanyName = "";
            this.newEmail = "";
            this.newMessage = "";
            console.log("Message sent: " + res);
            break;
          case 400:
            console.error("Invalid form: " + res);
            break;
          default:
            console.error("Failed to send message: " + res);
        }
      });
    }
  }
};
</script>

<style scoped>
h1 {
  padding-top:50px;
  padding-bottom:50px;
}

h2 {
  padding-top:65px;
  padding-bottom:65px;
  margin-left:-2%;
  max-width:150px
}

#showcase {
  align-items: center;
  background-blend-mode: darken;
  background-image: url("../assets/black_lozenge_@2X.png");
  background-position: center;
  background-repeat: repeat;
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 100vw;
  justify-content: center;
  text-align: center;
}

#showcase img {
  max-height: 30vh;
  max-width: 30%;
}

.input {
  margin-left:15%;
  width:50%;
}

.input-label {
  padding:20px;
  float:left;
}

.border {
  border-left: 1px groove black;
  padding-bottom: 2%;
}

.margin {
  margin-left: 5%;
}

.logo-margin {
  float:left;
  padding:5%;
}

.large-logo {
  max-width:250px;
  max-height:205px;
}

.mid-logo {
  max-width:200px;
  max-height:205px;
}

.small-logo {
  max-width:150px;
  max-height:205px;
}
</style>
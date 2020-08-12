<!--
  Sponsor
  --
  This view corresponds to the Sponsor page of the website.
  This page consists of:
    - sponsors
    - become a sponsor form
-->

<template>
  <div>
    <header id="showcase">
      <v-img max-width="80vw" max-height="30vh" contain src="@/assets/csesocwhiteblue.png" />
    </header>
    <v-container class="margin" fluid>
      <h1 class="border text-h1 font-weight-bold" style="padding:25px;">Sponsors</h1>
      <h2 class="text-h4">Principal <br> Sponsors</h2>
      <v-container class="border" fluid>
          <div v-for="(sponsor, index) in tierOne" :key="sponsor.id" :style="marginStyle(index, largeLogoFilter)">
              <div class="logo-margin">
                <img class="large-logo" :src="`data:image/png;base64,${sponsor.logo}`" v-on:click="onClickModal(sponsor)"/>
              </div>
              <div v-if="(index + 1) % largeLogoFilter === 0 || index === tierOne.length - 1" style="clear:both;"></div>
          </div>
      </v-container>
      <div style="clear:both;"></div>
      <h2 class="text-h4">Major <br> Sponsors</h2>
      <v-container class="border" fluid>
          <div v-for="(sponsor, index) in tierTwo" :key="sponsor.id" :style="marginStyle(index, midLogoFilter)">
              <div class="logo-margin">
                <img class="mid-logo" :src="`data:image/png;base64,${sponsor.logo}`" v-on:click="onClickModal(sponsor)"/>
              </div>
              <div v-if="(index + 1) % midLogoFilter === 0 || index === tierTwo.length - 1" style="clear:both;"></div>
          </div>
      </v-container>
      <div style="clear:both;"></div>
      <h2 class="text-h4">Affiliate <br> Sponsors</h2>
      <v-container class="border" fluid>
          <div v-for="(sponsor, index) in tierThree" :key="sponsor.id" :style="marginStyle(index, smallLogoFilter)">
              <div class="logo-margin">
                  <img class="small-logo" :src="`data:image/png;base64,${sponsor.logo}`" v-on:click="onClickModal(sponsor)"/>
              </div>
              <div v-if="(index + 1) % smallLogoFilter === 0 || index === tierThree.length - 1" style="clear:both;"></div>
          </div>
      </v-container>
      <div style="clear:both;"></div>
    </v-container>
    <SponsorModal v-model="dialog" v-bind:title="currentSponsor.name" v-bind:body="currentSponsor.detail"> </SponsorModal>

    <h1 class="text-center text-h1 font-weight-bold">Become a <br> Sponsor</h1>
    <v-card flat tile style="margin-left:15%">
      <EnquiryForm type="sponsorship"></EnquiryForm>
    </v-card>
  </div>
</template>

<script type="text/javascript">
import SponsorModal from '@/components/SponsorModal';
import EnquiryForm from '@/components/EnquiryForm';
import APIClient from '../utils/APIClient';

export default {
  name: 'Sponsor',
  data: () => ({
    currentSponsor: {},
    sponsors: [],
    dialog: false,

    // Constants
    largeLogoFilter: 3,
    midLogoFilter: 4,
    smallLogoFilter: 5
  }),
  components: {
    SponsorModal,
    EnquiryForm
  },
  computed: {
    // functions to determine sizing category of sponsor based on their value
    tierOne() {
      return this.sponsors.filter((x) => x.tier === 2);
    },
    tierTwo() {
      return this.sponsors.filter((x) => x.tier === 1);
    },
    tierThree() {
      return this.sponsors.filter((x) => x.tier === 0);
    }
  },
  mounted() {
    APIClient.fetchSponsors()
      .then((responseJson) => {
        this.sponsors = responseJson;
      });
  },
  methods: {
    marginStyle(index, limit) {
      const style = {};

      const row = parseInt((index) / limit, 10);
      if (row % 2 === 0) {
        style['margin-left'] = '10%';
      } else {
        style['margin-left'] = '15%';
      }

      return style;
    },
    onClickModal(sponsor) {
      this.currentSponsor = sponsor;
      this.dialog = true;
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
  max-width:225px;
  max-height:205px;
}

.mid-logo {
  max-width:175px;
  max-height:205px;
}

.small-logo {
  max-width:125px;
  max-height:205px;
}
</style>

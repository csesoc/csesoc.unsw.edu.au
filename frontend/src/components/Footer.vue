<!--
  Footer
  --
  This component contains footer elements of the website, visible from all pages.
-->

<template>
  <v-footer dark padless id="tiers" >
    <v-card style="padding:75px;width:100%">

      <v-card-text class="white--text" style="float:left;width:25%">
        <v-container class="fill-height" fluid style="float:left;max-height: 200px;max-width:300px;">
          <v-img src="@/assets/csesocwhiteblue.png" />
          <br/>B03 CSE Building K17, 
          <br/> UNSW
          <br/>
        </v-container>
      </v-card-text>
      <v-card-text class="white--text" style="float:left;width:20%">
        <a href="https://www.facebook.com/csesoc">About</a>
        <br/>
        <a href="https://www.facebook.com/groups/csesoc">Events</a>
        <br/>
        <a href="https://www.linkedin.com/company/csesoc/about/">CSE Resources</a>
        <br/>
        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">Sponsorship</a>
        <br/>
        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">Engage</a>
      </v-card-text>
      <v-card-text class="white--text" style="float:left;width:20%">
        <a href="https://www.facebook.com/csesoc">Facebook Page</a>
        <br/>
        <a href="https://www.facebook.com/groups/csesoc">Instagram</a>
        <br/>
        <a href="https://www.linkedin.com/company/csesoc/about/">LinkedIn</a>
        <br/>
        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">Twitter</a>
        <br/>
        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">Youtube</a>
      </v-card-text>
<<<<<<< HEAD

      <v-divider></v-divider>

      <v-card-text class="white-text text-center">
        &copy; {{ new Date().getFullYear() }} &mdash;
        <strong>CSESoc UNSW</strong>
=======
      <v-card-text class="white--text" style="float:left;width:25%">
        <a href="https://www.facebook.com/csesoc">Notangle</a>
        <br/>
        <a href="https://www.facebook.com/groups/csesoc">Compclub</a>
        <br/>
        <a href="https://www.linkedin.com/company/csesoc/about/">Job Board</a>
        <br/>
        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">CSE UNSW</a>
        <br/>
        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">CSESoc Media</a>
>>>>>>> inital framework for footer
      </v-card-text>
    </v-card>
  </v-footer>
</template>

<script type="text/javascript">
import APIClient from '../utils/APIClient';

export default {
  name: 'Footer',
  data: () => ({
    // sponsors have name, logo, 'tier' value, and link
    sponsors: []
  }),
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
  }
};
</script>

<style scoped>
<<<<<<< HEAD
#tiers {
  margin-top: 0px;
}

#tiers-sponsors {
=======
#tiers__sponsors {
>>>>>>> inital framework for footer
  background: #4a4a4a;
  padding-top: 5vw;
}

.white-text {
  font-size: 16px;
  line-height: 28px;
}

.white-text a {
  color: white;
}

.white-text a:hover {
  color: #427bff;
}

.tiers-grid {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  margin-left: 1vw;
  margin-right: 1vw;
}

.tiers-box {
  margin: 20px;
}
<<<<<<< HEAD

.tiers-contact {
  margin-bottom: 1vw;
}
=======
>>>>>>> inital framework for footer
</style>

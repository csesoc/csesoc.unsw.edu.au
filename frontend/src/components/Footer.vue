<!--
  Footer
  --
  This component contains footer elements of the website, visible from all pages.
-->

<template>
  <v-footer dark padless id="tiers">
    <v-card id="tiers__sponsors" width="100vw" class="white--text text-center">
      <v-card-text class="white--text">
        <v-container class="fill-height" fluid style="max-height: 200px; max-width:300px">
          <v-img src="@/assets/csesocwhiteblue.png" />
          <br />B03 CSE Building K17, UNSW
          <br />
        </v-container>
      </v-card-text>

      <h2>CONTACT US</h2>
      <v-card-text class="white-text text-center tiers-contact">
        <a href="https://www.facebook.com/csesoc">Facebook Page</a>
        |
        <a href="https://www.facebook.com/groups/csesoc">Facebook Group</a>
        |
        <a href="https://www.linkedin.com/company/csesoc/about/">LinkedIn</a>

        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">csesoc@csesoc.org.au</a>
      </v-card-text>

      <v-divider></v-divider>

      <v-card-text class="white-text text-center">
        &copy; {{ new Date().getFullYear() }} &mdash;
        <strong>CSESoc UNSW</strong>
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
#tiers {
  margin-top: 0px;
}

#tiers-sponsors {
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

.tiers-contact {
  margin-bottom: 1vw;
}
</style>

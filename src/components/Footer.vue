<template>
  <v-footer dark padless id="tiers">
    <v-card id="tiers__sponsors" width="100vw" class="white--text text-center">
      <v-card-text class="white--text">
        <section class="tiers__grid">
          <a v-for="sponsor in tierOne" class="tiers__box" :href="sponsor.link" :key="sponsor.id">
            <img :src="sponsor.logo" style=" max-width:300px;max-height:105px;" />
          </a>
        </section>
        <section class="tiers__grid">
          <a v-for="sponsor in tierTwo" class="tiers__box" :href="sponsor.link" :key="sponsor.id">
            <img :src="sponsor.logo" style="max-width:200px;max-height:75px" />
          </a>
        </section>
        <section class="tiers__grid">
          <a v-for="sponsor in tierThree" class="tiers__box" :href="sponsor.link" :key="sponsor.id">
            <img :src="sponsor.logo" style="max-width:100px;max-height:50px" />
          </a>
        </section>
      </v-card-text>

      <h2>CONTACT US</h2>
      <v-card-text class="white--text text-center tiers__contact">
        <a href="https://www.facebook.com/csesoc">Facebook Page</a>
        |
        <a href="https://www.facebook.com/groups/csesoc">Facebook Group</a>
        |
        <a href="https://www.linkedin.com/company/csesoc/about/">LinkedIn</a>
        <br />B03 CSE Building K17, UNSW
        <br />
        <a id="tiers__link" href="mailto:csesoc@csesoc.org.au">csesoc@csesoc.org.au</a>
      </v-card-text>

      <v-divider></v-divider>

      <v-card-text class="white--text text-center">
        &copy; {{ new Date().getFullYear() }} &mdash;
        <strong>CSESoc UNSW</strong>
      </v-card-text>
    </v-card>
  </v-footer>
</template>


<style scoped>
#tiers {
  margin-top: 0px;
}

#tiers__sponsors {
  background: #4a4a4a;
  padding-top: 5vw;
}

.white--text {
  font-size: 16px;
  line-height: 28px;
}

.white--text a {
  color: white;
}

.white--text a:hover {
  color: #427bff;
}

.tiers__grid {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  margin-left: 1vw;
  margin-right: 1vw;
}

.tiers__box {
  margin: 20px;
}

.tiers__contact {
  margin-bottom: 1vw;
}
</style>

<script type="text/javascript">
export default {
  name: 'Footer',
  data: () => ({
    // sponsors have name, logo, 'tier' value, and link
    sponsors: []
  }),
  computed: {
    // functions to determine sizing category of sponsor based on their value
    tierOne() {
      return this.sponsors.filter(x => x.tier >= 1000);
    },
    tierTwo() {
      return this.sponsors.filter(x => x.tier >= 100 && x.tier < 1000);
    },
    tierThree() {
      return this.sponsors.filter(x => x.tier >= 10 && x.tier < 100);
    }
  },
  mounted() {
    fetch(
      '/api/sponsors/?token=null'
    )
      .then(r => r.json())
      .then((responseJson) => {
        this.sponsors = responseJson;
      });
  }
};
</script>

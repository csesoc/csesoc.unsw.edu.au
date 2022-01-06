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
    <div class="down-button" @click="onClickScroll" @scroll.passive="handleScroll" ref="downButton">
      <img src="@/assets/downbutton.png" />
    </div>
    <v-container class="margin" fluid>
      <h1 class="border" font-weight-bold style="padding:25px;">Sponsors</h1>
      <h2 class="text-h4">
        Principal <br />
        Sponsors
      </h2>
      <v-container class="border" fluid>
        <div
          v-for="(sponsor, index) in tierOne"
          :key="sponsor.id"
          :style="marginStyle(index, largeLogoFilter)"
        >
          <div class="logo-margin">
            <img
              class="large-logo logo"
              :src="`data:image/png;base64,${sponsor.logo}`"
              v-on:click="onClickModal(sponsor)"
            />
          </div>
          <div
            v-if="(index + 1) % largeLogoFilter === 0 || index === tierOne.length - 1"
            style="clear:both;"
          ></div>
        </div>
      </v-container>
      <div style="clear:both;"></div>
      <h2 class="text-h4">
        Major <br />
        Sponsors
      </h2>
      <v-container class="border" fluid>
        <div v-for="(sponsor, index) in tierTwo" :key="sponsor.id" :style="marginStyle(index, midLogoFilter)">
          <div class="logo-margin">
            <img
              class="mid-logo logo"
              :src="`data:image/png;base64,${sponsor.logo}`"
              v-on:click="onClickModal(sponsor)"
            />
          </div>
          <div
            v-if="(index + 1) % midLogoFilter === 0 || index === tierTwo.length - 1"
            style="clear:both;"
          ></div>
        </div>
      </v-container>
      <div style="clear:both;"></div>
      <h2 class="text-h4">
        Affiliate <br />
        Sponsors
      </h2>
      <v-container class="border" fluid>
        <div
          v-for="(sponsor, index) in tierThree"
          :key="sponsor.id"
          :style="marginStyle(index, smallLogoFilter)"
        >
          <div class="logo-margin">
            <img
              class="small-logo logo"
              :src="`data:image/png;base64,${sponsor.logo}`"
              v-on:click="onClickModal(sponsor)"
            />
          </div>
          <div
            v-if="(index + 1) % smallLogoFilter === 0 || index === tierThree.length - 1"
            style="clear:both;"
          ></div>
        </div>
      </v-container>
      <div style="clear:both;"></div>
    </v-container>
    <SponsorModal v-model="dialog" v-bind:title="currentSponsor.name" v-bind:body="currentSponsor.detail">
    </SponsorModal>
  </div>
</template>

<script type="text/javascript">
import SponsorModal from '@/components/SponsorModal';
import APIClient from '../utils/APIClient';

export default {
  name: 'Sponsor',
  data: () => ({
    currentSponsor: {},
    sponsors: [],
    dialog: false,
    scrollY: 0,

    // Constants
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
    APIClient.fetchSponsors().then((responseJson) => {
      this.sponsors = responseJson;
    });
    window.addEventListener('scroll', this.handleScroll, true);
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll, true);
  },
  methods: {
    marginStyle(index, limit) {
      const style = {};
      const row = parseInt(index / limit, 10);
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
    },
    calculateNextTitle() {
      const arr = Array.from(document.getElementsByClassName('text-h4'));
      // find the next element if there is such a thing
      for (let i = 0; i < arr.length; i += 1) {
        const element = arr[i];
        if (element.offsetTop - this.scrollY > 0) {
          return element;
        }
      }
      return null;
    },
    // on click wil take you to the next heading
    onClickScroll() {
      try {
        const next = this.calculateNextTitle(window.scrollY);
        if (next === null) {
          // scroll to top
          document.getElementById('showcase').scrollIntoView({ behavior: 'smooth' });
        } else {
          next.scrollIntoView({ behavior: 'smooth' });
        }
      } catch (err) {
        /* console.log(err); */
      }
    },
    // check if current windowTop beyond the next in line view
    // if it is beyond it, pop it from list of views
    // since we want the button to bring us down not back up again
    handleScroll(e) {
      this.scrollY = e.target.scrollTop;
      try {
        if (this.calculateNextTitle() == null) {
          // TODO flip button
          this.$refs.downButton.className = 'down-button down-button-rotate';
        } else {
          this.$refs.downButton.className = 'down-button';
        }
      } catch (err) {
        /* console.log(err); */
      }
    }
  }
};
</script>

<style scoped>
h1 {
  padding-top: 50px;
  padding-bottom: 50px;
}

h2 {
  padding-top: 65px;
  padding-bottom: 65px;
  margin-left: -2%;
  max-width: 150px;
}

#showcase {
  align-items: center;
  background-blend-mode: darken;
  background-image: url('../assets/black_lozenge_@2X.png');
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
  padding-right: 10vw;
  display: flex;
  flex-wrap: wrap;
}

.margin {
  margin-left: 5%;
  margin-bottom: 15%;
}

.logo-margin {
  float: left;
  padding: 5%;
}

.logo {
  transition: transform 0.2s;
}

.logo:hover {
  cursor: pointer;
  transform: scale(1.05);
}

.large-logo {
  max-width: 225px;
  max-height: 205px;
}

.mid-logo {
  max-width: 175px;
  max-height: 205px;
}

.small-logo {
  max-width: 125px;
  max-height: 205px;
}

.down-button {
  position: fixed;
  left: 50vw;
  top: 90vh;
  height: 50px;
  width: 50px;
  border-radius: 50%;
}

.down-button > img {
  object-fit: cover;
  width: 100%;
  filter: opacity(60%);
}

.down-button > img:hover {
  cursor: pointer;
  transform: scale(1.1);
}
.down-button-rotate {
  transform: rotate(180deg);
}

@media only screen and (max-width: 300px) {
  h2 {
    color: white;
  }
}
</style>

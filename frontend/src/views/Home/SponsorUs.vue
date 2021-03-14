<!--
  Sponsor Us
  --
  This view is for the sponsor us information on the bottom of the landing
  page.
-->

<template>
  <div id="sponsor" class="content">
    <v-container>
    <!-- "support csesoc" title -->
    <HeaderTitle title="SUPPORT CSESOC" />
    <!-- row for trophy image and text adjacent -->
    <v-row no-gutters v-for="info in supportInfo" :key="info.title">
      <v-col cols="4" align="center">
        <div class="icon-text-aligner">
          <img :src="info.logo">
        </div>
      </v-col>
      <!-- h4 and p text adjacent to trophy image -->
      <v-col cols="8">
        <h4>{{info.title}}</h4>
        <p>{{info.desc}}</p>
      </v-col>
    </v-row>
    <!-- "sponsor us" button -->
    <v-col cols="12" align="right">
      <RouterLink to="/sponsors" class="link" style="text-decoration: none;">
      <button class="btn--sponsorUs">
        Sponsor Us
      </button>
      </RouterLink>
    </v-col>
    </v-container>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'SponsorUs',
  data: () => ({
    // all information comes with a logo, title and description
    supportInfo: [
      {
        logo: require('../../assets/trophy.svg'),
        title: 'Reach Australia\'s Best Computing Graduates',
        desc: 'Many of our past members have gone on to work at our previous sponsor companies including Google, Facebook, Jane Street and Commonwealth Bank.'
      },
      {
        logo: require('../../assets/share.svg'),
        title: 'Interact With Our Huge Active Community',
        desc: 'We have an extremely active community of over 3000 CSE students, achieved through our offering of career, social and educational events.'
      },
    ]
  }),
  components: {
    HeaderTitle,
  }
};
</script>

<style lang="scss" scoped>
#sponsor {
  background-image: linear-gradient(transparent, transparent, $dark-color-1), url(../../assets/landing_page/sponsor_us.png);
  background-size: cover;
  background-attachment: local;
}

.col {
  margin-bottom: $space-xs;
}

// "Sponsor Us" button
.btn--sponsorUs {
  @extend .btn--lg;
}

.btn--sponsorUs:hover {
  transition: 0.4s;
  // when on hover, change background from transparent to $brand-color with light opacity
  background: rgba(102,255,255, 0.2);
}
</style>

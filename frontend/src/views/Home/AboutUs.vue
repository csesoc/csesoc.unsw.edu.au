<!--
  Mission
  --
  Mission statement of CSESOC
-->

<template>
  <div id="aboutus" class="content">
    <v-container>
      <HeaderTitle :title="'ABOUT US'" />
      <v-row no-gutters justify="end">
        <v-col xs="12" sm="6">
          <img data-cy="mission-img" contain :src="mission.image" class="image" />
        </v-col>
        <v-col xs="12" sm="6" align-self="center">
          <p data-cy="mission-statement" class="text--md">{{ mission.statement }}</p>
          <v-btn to="/about" block outlined color="white" icon class="btn--sm">
            SEE MORE
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'AboutUs',
  components: {
    HeaderTitle
  },
  data: () => ({
    mission: {
      image: require('../../assets/landing_page/mission.png'),
      statement:
        'We are one of the biggest and most active societies at UNSW, catering to over 3500 CSE students spanning across degrees in Computer Science, Software Engineering, Bioinformatics and Computer Engineering.'
    }
  })
};
</script>
<style lang="scss" scoped>
.image {
  max-width: 100%;
  max-height: 100%;
}
</style>

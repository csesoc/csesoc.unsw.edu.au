<!--
  Home
  --
  This view corresponds to the landing page of the website.
  This page consists of:
    - csesocs mission
    - upcoming events
    - slideshow
    - resources
    - join us
    - sponsor us
-->

<template>
  <div id="home">
    <PrimaryNavbar />
    <!-- Showcase -->
    <Showcase />
    <!-- Mission -->
    <AboutUs />
    <!-- Slideshow -->
    <Slideshow />
    <!-- Community Links -->
    <CommunityLinks />
    <!--Events-->
    <EventDisplay :events="eventItems" :updated="lastEventUpdate" />
    <!-- Student Resources -->
    <StudentResources />
    <!-- Sponsor -->
    <SponsorUs />
  </div>
</template>

<script>
import Showcase from '@/views/Home/Showcase';
import AboutUs from '@/views/Home/AboutUs';
import CommunityLinks from '@/views/Home/CommunityLink';
import EventDisplay from '@/views/Home/EventDisplay';
import Slideshow from '@/views/Home/Slideshow';
import StudentResources from '@/views/Home/StudentResources';
import SponsorUs from '@/views/Home/SponsorUs';
import APIClient from '../../utils/APIClient';
import PrimaryNavbar from '../../components/PrimaryNavbar';

export default {
  name: 'Home',
  data: () => ({
    drawer: false,
    eventItems: [],
    loadTime: new Date().getTime(),
    lastEventUpdate: 0
  }),
  components: {
    Showcase,
    AboutUs,
    CommunityLinks,
    EventDisplay,
    Slideshow,
    StudentResources,
    SponsorUs,
    PrimaryNavbar,
  },
  mounted() {
    APIClient.eventsAPI()
      .then((responseJson) => {
        this.eventItems = responseJson.events;
        this.lastEventUpdate = this.loadTime - responseJson.updated * 1000;
      });
  },
  methods: {
    scrollto(anchor) {
      const element = this.$refs[anchor];
      const top = element.offsetTop;
      window.scrollTo({ top, behavior: 'smooth' });
    }
  }
};
</script>

<style scoped lang="scss">
#home {
  background-color: $dark-color-1;
  box-sizing: border-box;
  color: $light-color;
}
</style>

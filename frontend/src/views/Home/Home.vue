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
  <div id="home" >
    <!-- Showcase -->
    <Showcase class="padding"/>
    <!-- Mission -->
    <Mission  class="padding"/>
    <!--Events-->
    <EventDisplay :events="eventItems" :updated="lastEventUpdate"  class="padding"/>
    <!-- Slideshow -->
    <Slideshow  class="padding"/>
    <!-- Student Resources -->
    <StudentResources  class="padding"/>
    <!-- Community Links -->
    <CommunityLinks  class="padding"/>
    <!-- Sponsor -->
    <SponsorUs class="padding"/>
  </div>
</template>

<script>
import Showcase from '@/views/Home/Showcase';
import Mission from '@/views/Home/Mission';
import CommunityLinks from '@/views/Home/CommunityLink';
import EventDisplay from '@/views/Home/EventDisplay';
import Slideshow from '@/views/Home/Slideshow';
import StudentResources from '@/views/Home/StudentResources';
import SponsorUs from '@/views/Home/SponsorUs';
import APIClient from '../../utils/APIClient';

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
    Mission,
    CommunityLinks,
    EventDisplay,
    Slideshow,
    StudentResources,
    SponsorUs,
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

.padding{
  padding: 0.5em;
}

</style>

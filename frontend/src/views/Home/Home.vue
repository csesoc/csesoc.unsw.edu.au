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
    <!-- Showcase -->
    <Showcase />
    <!-- Mission -->
    <Mission />
    <!-- Events -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle :title="'upcoming events'" />
      <v-row>
          <EventGrid :events="eventItems" v-if="time < 60 * 1440 * 60000"></EventGrid>
          <!-- <div
            class="fb-page"
            data-href="https://www.facebook.com/csesoc"
            data-tabs="events"
            data-small-header="false"
            data-adapt-container-width="true"
            data-hide-cover="false"
            data-show-facepile="false"
          >
            <blockquote cite="https://www.facebook.com/csesoc" class="fb-xfbml-parse-ignore">
              <a href="https://www.facebook.com/csesoc">Loading...</a>
            </blockquote>
          </div> -->
      </v-row>
      <v-row>
        <v-spacer></v-spacer>
        <a href = "https://www.facebook.com/csesoc/events" class="fb-event-link">See more events on our Facebook page! ⭝</a>
      </v-row>
    </v-container>
    <!-- Slideshow -->
    <div class="blue-cutout">
      <InfiniteSlideshow duration="32s" direction="reverse">
        <img height="250px" src="@/assets/banner-1.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-2.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-3.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-4.jpg" style="padding-left: 20px; padding-right: 20px;">
      </InfiniteSlideshow>
      <InfiniteSlideshow duration="35s" direction="reverse">
        <img height="250px" src="@/assets/banner-1.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-2.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-3.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-4.jpg" style="padding-left: 20px; padding-right: 20px;">
      </InfiniteSlideshow>
      <InfiniteSlideshow duration="30s" direction="reverse">
        <img height="250px" src="@/assets/banner-1.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-2.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-3.jpg" style="padding-left: 20px; padding-right: 20px;">
        <img height="250px" src="@/assets/banner-4.jpg" style="padding-left: 20px; padding-right: 20px;">
      </InfiniteSlideshow>
    </div>
    <!-- Resources -->
    <v-container>
      <HeaderTitle :title="'resources'" />
      <Preview :items="resourceItems"/>
    </v-container>

    <!-- Join Us -->
    <v-container ref="joinus">
      <CommunityLinks ></CommunityLinks>
    </v-container>

    <!-- Sponsor -->
    <SponsorUs />
  </div>
</template>

<script>
import Showcase from '@/views/Home/Showcase';
import Mission from '@/views/Home/Mission';
import SponsorUs from '@/views/Home/SponsorUs';
import HeaderTitle from '@/components/HeaderTitle';
import EventGrid from '@/components/EventGrid';
import InfiniteSlideshow from '@/components/InfiniteSlideshow';
import Preview from '@/components/Preview';
import CommunityLinks from '@/components/CommunityLink';
import APIClient from '../../utils/APIClient';

export default {
  name: 'Home',
  data: () => ({
    drawer: false,
    resourceItems: [],
    announceApiUri:
      'https://gist.githack.com/gawdn/79b9df83f2fd267a3287d13b9badce48/raw/7bfb85a4cb799712229bed5ea02234e773eb42d4/populated_list.json',
    announceItems: [],
    mediaApiUri:
      'https://gist.githack.com/gawdn/a590d5be689e3ffbee15c213928e3b4b/raw/bed82e02b6a4d01196a4390e1a8b12ccf5b377fa/media0a.json',
    mediaItems: [],
    items: [],
    eventItems: [],
    time: new Date().getTime()
  }),
  components: {
    Showcase,
    Mission,
    SponsorUs,
    HeaderTitle,
    EventGrid,
    InfiniteSlideshow,
    Preview,
    CommunityLinks
  },
  mounted() {
    fetch(this.mediaApiUri)
      .then((r) => r.json())
      .then((responseJson) => {
        this.mediaItems = responseJson;
      });

    APIClient.resourcesAPI('/preview')
      .then((responseJson) => {
        this.resourceItems = responseJson;
      })
      .catch((error) => {
        // fix this
        console.log(error);
      });

    APIClient.eventsAPI()
      .then((responseJson) => {
        this.eventItems = responseJson.events.slice(0, 3);
        this.time -= responseJson.updated * 1000;
      });
    fetch(this.announceApiUri)
      .then((r) => r.json())
      .then((responseJson) => {
        this.announceItems = responseJson;
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

.fb-event-link {
  text-decoration: none;
  font-weight: bold;
  color: rgb(54, 119, 243);
}

.blue-cutout {
  background: rgb(18, 76, 219);
  background: linear-gradient(
  125deg,
  rgba(18, 76, 219, 1) 0%,
  rgba(50, 112, 255, 1) 50%,
  rgba(30, 104, 255, 1) 100%
  );
  margin: 50px 0px;
  padding: 30px 0px;
}
</style>

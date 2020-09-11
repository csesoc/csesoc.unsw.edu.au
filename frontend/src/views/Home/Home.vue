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
    <!--Events-->
    <EventDisplay :events="eventItems" :updated="lastEventUpdate" />
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

    <!-- Student Resources -->
    <v-container class="resource--styles">
      <h1>STUDENT RESOURCES</h1>

      <v-row no-gutters class="row">
        <v-col xs="12" sm="8" class="pa-0">
          <a style="color: inherit" href="">
            <div class="box big">
              <h2>Job Board</h2>
              <h3>A place where CSESoc students can look for relevant job opportunities.</h3>
              <v-img src="@/assets/resource-job-board.png" contain class="image" />
            </div>
          </a>
        </v-col>
        <v-col xs="12" sm="4" class="pb-0">
          <a style="color: inherit" href="">
            <div class="box small rounded-lg">
              <h2>First Year Guide</h2>
              <v-img src="@/assets/resource-first-year-guide.png" contain class="image" />
            </div>
          </a>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col xs="12" sm="4" class="pa-0">
          <a style="color: inherit" href="https://media.csesoc.org.au/">
            <div class="box small rounded-lg">
              <h2>CSESoc Media</h2>
              <v-img src="@/assets/resource-media.png" class="image" />
            </div>
          </a>
        </v-col>
        <v-col xs="12" sm="8" class="pb-0" href="">
          <a style="color: inherit" href="">
            <div class="box big rounded-lg">
              <h2>Notangles</h2>
              <h3>Trimester timetabling tool - no more timetable tangles!</h3>
              <v-img src="@/assets/resource-notangles.png" class="image" />
            </div>
          </a>
        </v-col>
      </v-row>
    </v-container>

    <!-- Community Links -->
    <CommunityLinks />

    <!-- Sponsor -->
    <SponsorUs />
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';
import Showcase from '@/views/Home/Showcase';
import Mission from '@/views/Home/Mission';
import CommunityLinks from '@/views/Home/CommunityLink';
import SponsorUs from '@/views/Home/SponsorUs';
import EventDisplay from '@/views/Home/EventDisplay';
import InfiniteSlideshow from '@/components/InfiniteSlideshow';
import Preview from '@/components/Preview';
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
    loadTime: new Date().getTime(),
    lastEventUpdate: 0
  }),
  components: {
    Showcase,
    Mission,
    SponsorUs,
    CommunityLinks,
    HeaderTitle,
    EventDisplay,
    InfiniteSlideshow,
    Preview,
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
        this.eventItems = responseJson.events;
        this.lastEventUpdate = this.loadTime - responseJson.updated * 1000;
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

// Student resources
.resource--styles {
  color: $light-color;

  .row {
    margin-bottom: $space-xxs;
    text-decoration: none;

    a {
      color: inherit;
      text-decoration: none;

      .box {
        border-radius: 10px;
        display: flex;
        flex-direction: column;
        height: 100%;
        overflow: hidden;

        &.big{
          @include linearGradient($primary-color, $secondary-color-2);
        }
        &.small{
          @include linearGradient($primary-color, $secondary-color-1);
        }

        h2 {
          padding-top: $space-xs;
          padding-bottom: $space-xxs;
          margin: 0;
          padding-left: 35px;
          padding-right: 35px;
        }

        h3 {
          padding-left: 35px;
          padding-right: 35px;
        }

        .image {
          border-bottom-left-radius: 10px;
          border-bottom-right-radius: 10px;
          transition: transform .2s;
        }
      }

      .box:hover > .image {
        transform: scale(1.1);
      }
    }
  }
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

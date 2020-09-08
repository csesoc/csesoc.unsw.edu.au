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
    <div id="showcase">
      <div class="showcase-image">
        <h1 class="showcase--h1">
          We inspire the programmers of the future.
        </h1>

        <button
          class="btn--joinUs"
          data-cy=joinus-button
          @click="scrollto('joinus')">
            Join Us
        </button>
      </div>
    </div>
    <!-- Mission -->
    <div class="mission">
      <h1 class="mission--h1"> CSESOCS MISSION </h1>
      <p class=mission--p>
        "To empower every person and every organization on the planet to achieve more.
        We strive to create local opportunity, growth, and impact in every country around the world. "
      </p>
    </div>
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
    <v-container ref="joinus" class="joinUs">
      <h1 class="joinUs--h1" data-cy=communityLinks-title>BE PART OF THE COMMUNITY</h1>
      <h4 class="joinUs--h4" data-cy=communityLinks-bodyText>
        Find our stall at O-week, or just turn up to an event and get to know us!
      </h4>
      <CommunityLinks></CommunityLinks>
    </v-container>

    <!-- Sponsor -->
    <div class=sponsor >
      <v-container>
        <!-- "support csesoc" title -->
        <h1 class=sponsor--h1>SUPPORT CSESOC</h1>
        <!-- row for trophy image and text adjacent -->
        <v-row no-gutters>
          <!-- trophy image -->
          <v-col cols="4" align="center" class="pa-0">
            <img src="@/assets/trophy.svg" class="sponsor--img">
          </v-col>
          <!-- h4 and p text adjacent to trophy image -->
          <v-col cols="8" class="pa-0">
            <h4 class=sponsor--h4>Reach Australia's Best Computing Graduates</h4>
            <p class=sponsor--p> Many of our past members have gone on to work at our previous
              sponsor companies including Google, Facebook, Jane St and Commonwealth Bank. </p>
          </v-col>
        </v-row>
        <!-- row for network/share image and text adjacent -->
        <v-row no-gutters class="pa-0">
          <!-- network image -->
          <v-col cols="4" align="center" justify="center" class="pa-0">
            <img src="@/assets/share.svg" class="sponsor--img">
          </v-col>
          <!-- h4 and p text adjacent to network image -->
          <v-col cols="8" class="pa-0">
            <h4 class=sponsor--h4>Interact With Our Huge Active Community</h4>
            <p class=sponsor--p> We have an extremely active community of ~3000 CSE students,
              achieved through our offering of career, social and educational events. </p>
          </v-col>
        <!-- "sponsor us" button -->
        </v-row>
        <v-col cols="12" align="right">
          <RouterLink to="/sponsors" class="link" style="text-decoration: none;">
            <button class="btn--sponsorUs">
              Sponsor Us
            </button>
          </RouterLink>
        </v-col>
      </v-container>
    </div>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';
import EventGrid from '@/components/EventGrid';
import InfiniteSlideshow from '@/components/InfiniteSlideshow';
import CommunityLinks from '@/components/CommunityLink';
import APIClient from '../utils/APIClient';

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
    HeaderTitle,
    EventGrid,
    InfiniteSlideshow,
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
}

// Showcase
#showcase {
  height: 100vh;
  width: 100vw;
  text-align: left;
  margin-bottom: $space-sm;
}

.showcase--h1 {
  padding-left: 7%;
  @extend h1;
  color: $light-color;
  padding-right: 50vw;
  line-height: 95px;
  padding-top: $space-xl;
}

.btn--joinUs {
  @extend .btn--lg;
  border: 2px solid #ffffff;
  left: 80vw;
}

.btn--joinUs:hover {
  transition: 0.4s;
  background: rgba(102,255,255, 0.2);
}

.showcase-image {
  background-image: linear-gradient(transparent, transparent, $dark-color-1), url(../assets/landingPageHeaderBackground.png);
  height: 100vh;
  background-position: 75% 50%;
  background-size: 100%;
  background-size: cover;
}

// Mission
.mission {
  height: 120vh;
  width: 100vw;
  float: left;
  background-size: 100%;
  background-size: cover;
  position: relative;
  margin-bottom: $space-xl;
}

.mission--h1 {
  padding-left: 7%;
  @extend h1;
  color: $light-color;
  padding-top: $space-md;
  position: relative;
}

.mission--p {
  @extend h3;
  color: $light-color;
  padding-left: 67%;
  padding-right: 7%;
  padding-top: $space-md;
  position: relative;
}

// css for background image only
.mission::before {
  content: "";
  background-image: linear-gradient($dark-color-1, transparent, $dark-color-1), url(../assets/mission.jpg);
  background-size: cover;
  position: absolute;
  width: 60%;
  top: 0px;
  right: 0px;
  bottom: 0px;
  left: 0px;
}

.joinUs {
  padding-bottom: $space-xl;
}

.joinUs--h1 {
  color: $light-color;
  margin-bottom: $space-xs;
}

.joinUs--h4 {
  color: $light-color;
  margin-bottom: $space-md
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
          padding-left: 20px;
          padding-right: 20px;
        }

        h3 {
          padding-left: 20px;
          padding-right: 20px;
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

// SUPPORT CSESOC SECTION
.sponsor {
  margin-bottom: $space-xl;
  background-image: url(../assets/supportusbackground.png);
  background-size: cover;
  background-position: 20% 50%;
}

.sponsor--h1 {
  color: $light-color;
  margin-bottom:  $space-md;
}

.sponsor--h4 {
  color: $light-color;
  margin-bottom: $space-xxs;
}

.sponsor--p {
  color: $light-color;
  margin-bottom: $space-xs;
}

// for icon images
.sponsor--img {
  height: 100px;
  margin-top: $space-xxxs;
}

// "Sponsor Us" button
.btn--sponsorUs {
  @extend .btn--lg;
  border: 2px solid #ffffff;
}

.btn--sponsorUs:hover {
  transition: 0.4s;
  // when on hover, change background from transparent to $brand-color with light opacity
  background: rgba(102,255,255, 0.2);
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

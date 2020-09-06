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
    <div class=sponsor >
      <h1 class=sponsor--h1>SUPPORT CSESOC</h1>
      <div class=sponsor--block1>
        <img src="@/assets/share.svg" class="sponsor--img">
        <div class=sponsor--text>
          <h4 class=sponsor--h4>Reach Australia's Best Computing Graduates</h4>
          <p class=sponsor--p> Many of our past members have gone on to work at our previous
            sponsor companies including Google, Facebook, Jane St and Commonwealth Bank. </p>
        </div>
      </div>
      <div class=sponsor--block2>
        <img src="@/assets/trophy.svg" class="sponsor--img">
        <div class=sponsor--text>
          <h4 class=sponsor--h4>Interact With Our Huge Active Community</h4>
          <p class=sponsor--p> We have an extremely active community of ~3000 CSE students,
            achieved through our offering of career, social and educational events. </p>
        </div>
      </div>
      <RouterLink to="/sponsors" class="link" >
        <button class="btn--sponsorUs">
          Sponsor Us
        </button>
      </RouterLink>
    </div>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';
import EventGrid from '@/components/EventGrid';
import InfiniteSlideshow from '@/components/InfiniteSlideshow';
import Preview from '@/components/Preview';
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

/*
.mission {
  float: left;
  background-size: 100%;
  background-size: cover;
  position: relative;
  margin-bottom: $space-xl;
}*/
// Sponsor
.sponsor {
  margin-bottom: $space-xl;
}

.sponsor--h1 {
  margin-left: 7%;
  @extend h1;
  color: $light-color;
  margin-bottom:  $space-md;
}

.sponsor--h4 {
  //padding-left: 33%;
  @extend h4;
  color: $light-color;
  margin-bottom: $space-xxs;
}

.sponsor--p {
  @extend p;
  color: $light-color;
}

.sponsor--text {
  //width: 70%;
  opacity: 1;
  float: left;
  width: 70%;
}

.sponsor--img {
  height: 100px;
  float: left;
  width: 30%;
}

.sponsor--block1 {
  padding-bottom: $space-md;
}

.sponsor--block2 {
  padding-bottom: $space-md;
}

.btn--sponsorUs {
  @extend .btn--lg;
  border: 2px solid #ffffff;
  left: 75.5vw;
}

.btn--sponsorUs:hover {
  transition: 0.4s;
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

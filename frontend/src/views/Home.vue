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
    <!-- make header a seperate component! -->
    <header id="showcase">
      <!-- CSESoc Image Logo (to replace h1 below)

      <v-img
        v-if="$vuetify.breakpoint.mdAndUp"
        max-width="35vw"
        max-height="20vh"
        contain
        src="@/assets/csesocwhiteblue.png"
      />
      <v-img v-else max-width="80vw" max-height="30vh" contain src="@/assets/csesocwhiteblue.png" />
      -->

      <h1>
        CSESoc, Lorem ipsum dolor sit amet, consetetur
      </h1>

      <!--change this to scroll to Join Us section when created-->
      <a
        @click="scrollto('joinus')"
        target="_blank"
        v-ripple
        class="button"
      >Join Us</a>
    </header>

    <!-- CSESocs Mission -->
    <v-container id="mission">
      <HeaderTitle :title="'csesocs mission'" />
      <p>
        "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod temp or invidunt ut labore et dolore "
      </p>
    </v-container>

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

    <v-container>
      <HeaderTitle :title="'resources'" />
      <Preview :items="resourceItems"/>
    </v-container>

    <v-container>
      <CommunityLinks></CommunityLinks>
    </v-container>

    <!-- Support CSESoc -->
    <v-container id=sponsor >
      <HeaderTitle :title="'support csesoc'" />
      <p>
        Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor
        invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam
        et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est
        </p>
        <RouterLink to="/sponsors" class="link" >
          <a
            style="float:right"
            target="_blank"
            v-ripple
            class="button">
            Sponsor Us
          </a>
        </RouterLink>
    </v-container>

  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';
import EventGrid from '@/components/EventGrid';
import InfiniteSlideshow from '@/components/InfiniteSlideshow';
import Preview from '@/components/Preview';
import APIClient from '../utils/APIClient';
import CommunityLinks from '@/components/CommunityLink';

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

<style scoped>
#showcase {
  align-items: left;
  background-blend-mode: darken;
  background-image: url("../assets/black_lozenge_@2X.png");
  background-position: center;
  background-repeat: repeat;
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 100vw;
  justify-content: center;
  padding: 7% ;
  text-align: left;
}

#showcase h1 {
  color: #fff;
  font-size: 80px;
  width: 80%;
  font-weight: bolder;
  line-height: 95px;
}

#showcase img {
  max-height: 30vh;
  max-width: 30%;
}

#showcase p {
  font-size: 20px;
}

#showcase .button {
  text-align: center;
  background: rgb(54, 119, 243);
  border-radius: 0px;
  color: #fff;
  font-size: 40px;
  font-weight: bold;
  margin-top: 45px;
  padding: 15px 20px;
  text-decoration: none;
  width: 250px;
}

#showcase .button:hover {
  transition: 0.4s;
  background: rgb(54, 119, 243);
  color: #fff;
}

#mission p {
  font-size: 40px;
  align-items: center;
  text-align: center;
  padding: 5vh 11vw;
}

#sponsor {
  align-items: right;
  padding-bottom: 150px;
}

#sponsor p {
  font-size: 30px;
  padding: 0 0 25px 0;
}

#sponsor .button {
  text-align: center;
  background: rgb(54, 119, 243);
  border-radius: 0px;
  color: #fff;
  font-size: 30px;
  font-weight: bold;
  padding: 15px 20px;
  text-decoration: none;
  width: 250px;
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

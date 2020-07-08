
<template>
  <div id="home">
    <!-- make header a seperate component! -->
    <header id="showcase">
      <v-img
        v-if="$vuetify.breakpoint.mdAndUp"
        max-width="35vw"
        max-height="20vh"
        contain
        src="@/assets/csesocwhiteblue.png"
      />
      <v-img v-else max-width="80vw" max-height="30vh" contain src="@/assets/csesocwhiteblue.png" />
      <a
        href="https://www.arc.unsw.edu.au/clubs"
        target="_blank"
        v-ripple
        class="button"
      >Join on spArc</a>
      <br />
      <v-btn text icon color="white" @click="scrollto('content-start')">
        <v-icon>mdi-chevron-down</v-icon>
      </v-btn>
    </header>

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
        <!-- <v-col sm="12" md="5" lg="8">
          <HeaderTitle :title="'latest'" />
          <ListComponent :items="announceItems" />
        </v-col> -->
      </v-row>
      <v-row>
        <v-spacer></v-spacer>
        <a href = "https://www.facebook.com/csesoc/events" class="fb-event-link">See more events on our Facebook page! ⭝</a>
      </v-row>
    </v-container>

    <Slider :items="mediaItems" :title="'media'" class="my-10" />

    <v-container>
      <HeaderTitle :title="'resources'" />
      <NavGrid :items="resourcesItems"></NavGrid>
    </v-container>
  </div>
</template>

<script>
import NavGrid from '@/components/NavGrid';
import ListComponent from '@/components/ListComponent';
import Slider from '@/components/Slider';
import HeaderTitle from '@/components/HeaderTitle';
import EventGrid from '@/components/EventGrid';
import APIClient  from '../utils/APIClient';

export default {
  data: () => ({
    drawer: false,
    resourcesApiUrl:
      'https://gist.githack.com/gawdn/6fb68af4e994dd72e50fb360d299cbb6/raw/6fa351ba05f90ce0906c4c7accdf8c712f28f60d/resources0b.json',
    resourcesItems: [],
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
    NavGrid,
    ListComponent,
    Slider,
    HeaderTitle,
    EventGrid
  },

  mounted() {
    fetch(this.mediaApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.mediaItems = responseJson;
      });

    fetch(this.resourcesApiUrl)
      .then(r => r.json())
      .then((responseJson) => {
        this.resourcesItems = responseJson;
      });
    APIClient.eventsAPI()
      .then((responseJson) => {
        this.eventItems = responseJson.events.slice(0,3);
        this.time -= responseJson.updated * 1000;
      });
    fetch(this.announceApiUri)
      .then(r => r.json())
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
  align-items: center;
  background-blend-mode: darken;
  background-image: url("../assets/black_lozenge_@2X.png");
  background-position: center;
  background-repeat: repeat;
  display: flex;
  flex-direction: column;
  height: 100vh;
  width: 100vw;
  justify-content: center;
  padding: 0 20px;
  text-align: center;
}
#showcase img {
  max-height: 30vh;
  max-width: 30%;
}
#showcase p {
  font-size: 20px;
}
#showcase .button {
  background: rgb(54, 119, 243);
  border-radius: 10px;
  color: #fff;
  font-size: 18px;
  margin-top: 20px;
  padding: 10px 20px;
  text-decoration: none;
  width: 250px;
}
#showcase .button:hover {
  transition: 0.4s;
  background: rgb(54, 119, 243);
  color: #fff;
}

.fb-event-link {
  text-decoration: none;
  font-weight: bold;
  color: rgb(54, 119, 243);
}
</style>

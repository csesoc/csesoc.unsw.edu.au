
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
        @click="scrollto('content-start')"
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
      <Preview :items="resourceItems"/>
    </v-container>
  </div>
</template>

<script>
import ListComponent from '@/components/ListComponent';
import Slider from '@/components/Slider';
import HeaderTitle from '@/components/HeaderTitle';
import EventGrid from '@/components/EventGrid';
import Preview from '@/components/Preview';
import APIClient from '../utils/APIClient';

export default {
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
    ListComponent,
    Slider,
    HeaderTitle,
    EventGrid,
    Preview,
  },

  mounted() {
    fetch(this.mediaApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.mediaItems = responseJson;
      });

    APIClient.resourcesAPI('/preview')
      .then((responseJson) => {
        this.resourceItems = responseJson;
      })
      .catch((error) => {
        // fix this
        console.log(error)
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

#mission {
  padding: 5% 0
}

#mission p {
  font-size: 40px;
  align-items: center;
  text-align: center;
  padding: 5vh 11vw;
}

.fb-event-link {
  text-decoration: none;
  font-weight: bold;
  color: rgb(54, 119, 243);
}
</style>

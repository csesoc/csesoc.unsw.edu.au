
<template>
  <div id="home">
    <!-- make header a seperate component! -->
    <header id="showcase">
      <v-img max-width="35vw" max-height="20vh" contain="true" src="@/assets/csesocwhiteblue.png" />
      <a href="#" v-ripple class="button">Join on spArc</a>
      <br />
      <v-btn text icon color="white" @click="scrollto('content-start')">
        <v-icon>mdi-chevron-down</v-icon>
      </v-btn>
    </header>

    <v-container ref="content-start" style="padding: 20px 10px 10px 10px">
      <v-row>
        <v-col sm="12" md="6" lg="3">
          <HeaderTitle :title="'events'" />
          <v-lazy>
            <div
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
            </div>
          </v-lazy>
        </v-col>
        <v-col sm="12" md="6" lg="9">
          <HeaderTitle :title="'announcements'" />
          <ListComponent :listItems="listItems" />
        </v-col>
      </v-row>
    </v-container>

    <!-- media slider - could be showcase of recent posts?? -->
    <!-- how to pass in OPTIONS -->
    <Slider :items="items" :title="'media'" class="my-10" />

    <v-container style="padding: 20px 10px 10px 10px">
      <HeaderTitle :title="'resources'" />
      <NavGrid :gridItems="gridItems"></NavGrid>
    </v-container>
  </div>
</template>

<script>
import NavGrid from '@/components/NavGridSquare';
import ListComponent from '@/components/ListComponent';
import Slider from '@/components/Slider';
import HeaderTitle from '@/components/HeaderTitle';

export default {
  data: () => ({
    drawer: false,
    gridApiUri:
      'https://gistcdn.githack.com/esyw/f83b10232854534b64e475473406dfe6/raw/263a737400bc4e4e642a28cb9da9851ef76e3546/help.json',
    gridItems: [],
    listApiUri:
      'https://gist.githack.com/gawdn/79b9df83f2fd267a3287d13b9badce48/raw/7bfb85a4cb799712229bed5ea02234e773eb42d4/populated_list.json',
    listItems: [],
    items: []
  }),

  components: {
    NavGrid,
    ListComponent,
    Slider,
    HeaderTitle
  },

  mounted() {
    fetch(this.gridApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.gridItems = responseJson;
        this.items = responseJson;
      });

    fetch(this.listApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.listItems = responseJson;
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
  background: rgb(32, 62, 207);
  color: #fff;
}
</style>

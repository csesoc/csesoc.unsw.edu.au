<template>
  <v-app id="inspire">
    <Sidebar :drawer="drawer" />
    <v-app-bar app dark>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"> </v-app-bar-nav-icon>
      <v-container
        class="fill-height"
        fluid
        style="max-height: 64px; max-width:100px"
      >
        <v-img src="./assets/csesoclogobluewhite.png" />
      </v-container>
      <div class="flex-grow-1"></div>

      <LoginForm align="right" />
    </v-app-bar>
    <v-content>
      <v-container fluid>
        <BasePost />
        <TextPost />
        <h2>// UPCOMING EVENTS</h2>
        <EventsGrid />
      </v-container>
    </v-content>
    <v-footer dark app>
      <span class="white--text">&copy; 2019</span>
    </v-footer>
  </v-app>
</template>

<script>
import TextPost from './components/TextPost.vue';
import Sidebar from './components/Sidebar.vue';
import LoginForm from './components/LoginForm.vue';
import BasePost from './components/BasePost.vue';
import EventsGrid from './components/EventsGrid.vue';
// import SearchPage from './components/SearchPage.vue';

export default {
  props: {
    source: String,
  },
  components: {
    TextPost,
    Sidebar,
    LoginForm,
    BasePost,
    EventsGrid,
    // SearchPage,
  },
  data: () => ({
    drawer: false,
  }),
  created() {
    this.fetchText();
  },
  methods: {
    fetchText() {
      const xmlhttp = new XMLHttpRequest();
      xmlhttp.onreadystatechange = () => {
        if (this.readyState === 4 && this.status === 200) {
          const myObj = this.responseText;
          console.log(myObj);
        }
      };
      xmlhttp.open('GET', './components/test.txt', true);
      xmlhttp.send();
    },
  },
};
</script>


<style scoped>
img {
  max-height: 100%;
  max-width: 100%;
}
</style>

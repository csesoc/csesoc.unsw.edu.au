<template>
  <v-app id="main-app">
    <!-- Navigation bar/app bar goes here -->
    <v-content>
      <div>
        <Sidebar :drawer="drawer" />
        <v-app-bar app dark width="100vw">
          <v-app-bar-nav-icon class="ma-2" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
          <div class="flex-grow-1"></div>
          <router-link to="/">
            <v-container class="fill-height" fluid style="max-height: 64px; max-width:100px">
              <v-img src="@/assets/csesocwhiteblue.png" />
            </v-container>
          </router-link>
          <div class="flex-grow-1"></div>
          <LoginForm align="right" />
        </v-app-bar>

        <router-view style="overflow-x: hidden"></router-view>
        <Footer />
      </div>
    </v-content>
  </v-app>
</template>

<script>
import Footer from '@/components/Footer';
import Sidebar from '@/components/Sidebar.vue';
import LoginForm from '@/components/LoginForm';

export default {
  name: 'App',
  data: () => ({
    drawer: false,
    gridApiUri:
      'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    gridItems: [],
    listApiUri:
      'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    listItems: []
  }),

  components: {
    Footer,
    Sidebar,
    LoginForm
  },

  mounted() {
    fetch(this.gridApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.gridItems = responseJson;
      });
  }
};
</script>

<style>
/** Must keep html and body seperate to allow scrolling */
* {
  margin: 0;
  padding: 0;
  font-family: "Quicksand", sans-serif;
}

html {
  width: 100vw;
  height: 100vh;
  overflow-x: hidden;
}

body {
  width: 100vw;
  height: 100vh;
}
</style>

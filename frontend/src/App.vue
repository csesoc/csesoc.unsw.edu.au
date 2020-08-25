<template>
  <v-app id="main-app">
    <!-- Navigation bar/app bar goes here -->
    <v-app-bar app dark width="100vw">
      <!-- Logo -->
      <RouterLink to="/" v-on:click.native="showMenu = false">
        <v-container class="fill-height" style="max-height: 64px; max-width:100px">
          <v-img src="@/assets/csesocwhiteblue.png" />
        </v-container>
      </RouterLink>

      <div class="flex-grow-1"></div>

      <!-- Menu button -->
      <v-app-bar-nav-icon class="ma-2" @click.stop="showMenu = !showMenu"></v-app-bar-nav-icon>
    </v-app-bar>
    <v-main >
      <Menu v-if="showMenu" @shown="onMenuCollapse" />
      <RouterView style="overflow-x: hidden" />
      <Footer />
    </v-main>
  </v-app>
</template>

<script>
import Footer from '@/components/Footer';
import Menu from '@/components/Menu';

export default {
  name: 'App',
  data: () => ({
    showMenu: false
  }),
  components: {
    Footer,
    Menu,
  },
  methods: {
    onMenuCollapse(val) {
      this.showMenu = val;
    }
  }
};
</script>

<style lang="scss">
@import 'assets/font/stylesheet.css';
/** Must keep html and body seperate to allow scrolling */
* {
  margin: 0;
  padding: 0;
  font-family: 'Core Sans C', Arial, sans-serif;
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

<template>
  <v-app id="main-app">
    <!-- Navigation bar/app bar goes here -->
    <Navbar />
    <v-main class="pa-0">
      <Menu v-if="showMenu" @shown="onMenuCollapse" />
      <RouterView style="overflow-x: hidden" />
      <Footer />
    </v-main>
  </v-app>
</template>

<script>
import Footer from '@/components/Footer';
import Menu from '@/components/Menu';
import Navbar from '@/components/Navbar';

export default {
  name: 'App',
  data: () => ({
    showMenu: false
  }),
  components: {
    Footer,
    Menu,
    Navbar,
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

html, body {
  width: 100vw;
  height: 100vh;
  overflow-x: hidden;
  box-sizing: border-box;
}

.blend{
  mix-blend-mode: exclusion;
}
</style>

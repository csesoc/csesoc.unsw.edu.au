<template>
  <v-app id="main-app">
    <!-- Navigation bar/app bar goes here -->
    <v-content>
      <div>
        <v-app-bar app dark width="100vw">
          <!-- Logo -->
          <router-link to="/">
            <v-container class="fill-height" style="max-height: 64px; max-width:100px">
              <v-img src="@/assets/csesocwhiteblue.png" />
            </v-container>
          </router-link>

          <div class="flex-grow-1"></div>

          <!-- Menu button -->
          <v-app-bar-nav-icon class="ma-2" @click.stop="showMenu = !showMenu"></v-app-bar-nav-icon>
        </v-app-bar>

        <main>
          <Menu v-if="showMenu" @shown="onMenuCollapse" />
          <RouterView style="overflow-x: hidden" />
        </main>
        
        <Footer />
      </div>
    </v-content>
  </v-app>
</template>

<script>
import Footer from '@/components/Footer';
import Menu from '@/components/Menu.vue';
import LoginForm from '@/components/LoginForm';

export default {
  name: 'App',
  data: () => ({
    showMenu: false
  }),
  components: {
    Footer,
    Menu,
    LoginForm
  },
  methods: {
    onMenuCollapse(val) {
      this.showMenu = val;
    }
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

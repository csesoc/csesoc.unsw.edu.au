<template>
  <v-app id="main-app">
    <!-- Navigation bar/app bar goes here -->
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>

<script>

export default {
  name: 'App',
  components: {},
  data: () => ({}),
};
</script>

<style scoped>
</style>

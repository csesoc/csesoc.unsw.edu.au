<template>
  <v-container>
    <h1>You shall be redirected soon</h1>
    <!-- <v-row>
      <v-col class="pa-4" cols='6'>
        Some text
      </v-col>
      <v-col class="pa-4" cols='6'>
        Some other text
      </v-col>
    </v-row>-->

    <div class="btns">
      <button class="btn--sm btn--primary">Primary</button>
      <button class="btn--sm btn--secondary">Secondary</button>
      <button class="btn--sm btn--accent">Accent</button>
    </div>
  </v-container>
</template>
<script>
export default {
  name: 'Resources',
  props: ['route'],
  beforeRouteEnter(to, from, next) {
    document.location.href = 'https://media.csesoc.org.au/tag/csesoc/';
    next('/');
  }
};
</script>
<style lang="scss" scoped>
.col {
  background-color: $primary-color;
}
</style>

<template>
  <v-container class="nav__grid" fluid>
    <v-row class="align-left justify-left">
      <v-col sm="12" md="4" lg="4" v-for="item in items" :key="item.id">
        <v-card class="ma-3 grid__square" :href="item.link ? item.link : '#0'" target="_blank">
          <v-img
            :src="item.src"
            class="grid__img"
            gradient="to bottom, rgba(255,255,255,0.1) 0%, rgba(41,41,41,0.6) 50%, rgba(24,24,24,0.8) 100%"
          >
            <v-card-title
              class="justify-center fill-height align-center white--text font-weight-medium text-truncate"
              v-text="item.title"
            ></v-card-title>
          </v-img>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>


<style scoped>
.grid_square,
.grid__img {
  height: 28vh;
  background: rgb(8, 72, 255);
  background: linear-gradient(
    120deg,
    rgba(8, 72, 255, 1) 0%,
    rgba(0, 98, 214, 1) 50%,
    rgba(65, 115, 255, 1) 100%
  );
}
</style>

<script type="text/javascript">
export default {
  name: 'NavGrid',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['items']
};
</script>

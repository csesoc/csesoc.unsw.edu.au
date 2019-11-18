<template>
  <v-container class="nav__grid" fluid>
    <v-row align="left" justify="left">
      <v-col sm="6" md="4" lg="4" v-for="item in gridItems" :href="item.link" :key="item.id">
        <v-card class="ma-3 grid__square">
          <v-img
            :src="item.src"
            class="grid__img white--text"
            gradient="to bottom, rgba(0,0,0,.3), rgba(0,0,0,.8)"
          >
            <v-card-title class="justify-center fill-height align-center" v-text="item.title"></v-card-title>
          </v-img>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>


<style scoped>
.grid_square,
.grid__img {
  max-height: 28vh;
}
</style>

<script type="text/javascript">
export default {
  name: 'NavGrid',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['gridItems']
};
</script>

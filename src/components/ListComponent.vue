<template>
  <v-container v-if="listItems.length > 0">
    <v-row class="list--item mt-md-6 mb-md-6" v-for="item in listItems" :key="item.id">
      <v-col sm="12" md="4">
        <v-img :src="item.image" contain="true" class="list--img" v-if="item.image" />
        <div class="img--wrapper__default" v-else>
          <v-img
            class="list--img list--img__default"
            src="@/assets/csesocwhiteblue.png"
            contain="true"
          />
        </div>
      </v-col>
      <v-col>
        <div>
          <h2 class="listItemTitle" v-if="item.title">{{ item.title }}</h2>
          <h3 class="listItemSubtitle" v-if="item.subtitle">{{ item.subtitle }}</h3>
          <p>{{ item.content }}</p>
          <div class="text-center">
            <router-link :to="item.link" v-if="item.link">
              <v-btn rounded color="primary" dark style="float: left;">More -></v-btn>
            </router-link>
          </div>
        </div>
      </v-col>
    </v-row>
  </v-container>
  <v-container v-else>
    <p>No posts yet...stay tuned! 😎</p>
  </v-container>
</template>


<style scoped>
.list--item {
  border-top: 1px solid rgba(25, 25, 25, 0.1);
  max-width: 90vw;
}

.list--img {
  max-height: 40vh;
}

.img--wrapper__default {
  padding: 10% 10% 10% 10% !important;
  background-color: rgb(25, 25, 25, 0.9);
}
</style>


<script>
export default {
  name: 'ListComponent',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['listItems']
};
</script>

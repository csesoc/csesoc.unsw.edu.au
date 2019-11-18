<template>
  <!-- Insert post here-->
    <v-container class="pa-md-1 mx-lg-auto">
      <v-row v-for="item in listItems">
        <v-col class="pa-2" v-if="item.imageItem">
            <!-- use v-img -->
          <img v-if="item.image.length != 0" :src="item.image" />
        </v-col>
        <v-col>
          <div class="pa-md-6">
            <h2 class="listItemTitle">
              {{ item.title }}
            </h2>
            <h3 class="listItemSubtitle">{{ item.subtitle }}</h3>
            <p>
              {{ item.content }}
            </p>
            <div class="text-center">
              <router-link :to="item.link"
                ><v-btn rounded color="primary" dark style="float: left;"
                  >More -></v-btn
                ></router-link
              >
            </div>
          </div>
        </v-col>
      </v-row>
    </v-container>
</template>


<style scoped>
@import url("https://fonts.googleapis.com/css?family=Quicksand&display=swap");

h1 {
  font-size: 48px;
  font-weight: normal;
}

body {
  font-family: "Quicksand", sans-serif;
  text-align: left;
}

p {
  font-size: 14px;
  font-weight: 300;
  -webkit-font-smoothing: subpixel-antialiased;
}

p span {
  font-size: 16px;
  font-style: italic;
  display: inline-block;
}

.underline {
  border-bottom: 1px solid grey;
}

img {
  max-width: 100%;
  max-height: 100%;
}
</style>


<script>

    export default {
      name: 'ListComponent',
      // Must be passed from parent object
      // items have title, image url (src), and link
      props: ['listItems'],

    };
</script>

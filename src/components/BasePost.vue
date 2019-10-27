<template>
  <div>
    <v-container class="pa-md-10 mx-lg-auto">
      <!-- maybe a loop to print out path to home-->
      <div class="links">
        <span class="underline">home / post / </span>
        <span class="underline">{{ parentData.nav }}</span>
      </div>
      <v-flex xs12>
        <h1>#!/{{ parentData.header }}</h1>
        <h3>{{ parentData.posttime }}</h3>
        <h2>//{{ parentData.subtitle }}</h2>
      </v-flex>
    </v-container>
    <!-- Insert post here-->
    <div class="postText">
      <v-container class="pa-md-10 mx-lg-auto">
        <p>{{ parentData.text }}</p>
      </v-container>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BasePost',
  props: {
    parentData: Object,
  },
  data() {
    const title = 'Header';
    const subtitle = 'Subtitle';
    return {
      nav: title.toLowerCase(),
      header: title.toUpperCase(),
      subtitle: subtitle.toUpperCase(),
      posttime: '14th Oct 2019 9:19pm',
      // new lines need to be manually inserted so we need to parse them from the text
      text: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit',
    };
  },
};
</script>

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
</style>

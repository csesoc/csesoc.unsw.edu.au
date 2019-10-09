<template>
  <!-- Insert post here-->
  <div class="gridView">
    <v-container class="pa-md-10 mx-lg-auto">
      <v-row>
        <v-col>
          <img src="../assets/hackathonbanner.jpg" />
        </v-col>
        <v-col>
          <div class="pa-md-6">
            <h2 class="eventName">
              {{ event.name }}
            </h2>
            <h3 class="eventTime">{{ event.date }}</h3>
            <h3>{{ event.time }}</h3>
            <p>
              {{ event.info }}
            </p>
            <div class="text-center">
              <v-btn rounded color="primary" dark style="float: right;"
                >More -></v-btn
              >
            </div>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  name: 'EventsGrid',
  el: '#nav',
  data() {
    const title = 'Header';
    return {
      event: {
        name: 'Annual Hackathon'.toUpperCase(),
        date: '1st January 2020',
        time: '0:00 - 23:59',
        info: 'Do you love solving problems? Want to make a product that can potentially change the world?\n\n\
        Here’s your opportunity to do so with your team of up to six people at CSESoc’s annual hackathon!  You and your team will be given 24 hours to come up with a solution for the problem presented. ',
        image: '../assets/hackathonbanner.jpg',
      }
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

img {
  max-width: 100%;
  max-height: 100%;
}
</style>

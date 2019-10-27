<template>
  <div>
    <header id="showcase">
      <img src="../assets/csesoclogobluewhite.png" />
      <a href="#" class="button"> Join on spArc </a>
      <br />
      <v-btn text icon color="white">
        <v-icon>mdi-chevron-down</v-icon>
      </v-btn>
    </header>
    <ListComponent />
    <EventsGrid />
  </div>
</template>

<script>
// @ is an alias to /src
import ListComponent from '@/components/ListComponent.vue';
import EventsGrid from '@/components/EventsGrid.vue';

export default {
  name: 'Home',
  components: {
    ListComponent,
    EventsGrid,
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css?family=Quicksand&display=swap");

* {
  margin: 0;
  padding: 0;
}
#showcase {
  background-image: url("https://backgroundcheckall.com/wp-content/uploads/2017/12/black-tech-background-12.jpg");
  background-size: cover;
  background-position: center;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  padding: 0 20px;
}

#showcase img {
  max-width: 30%;
  max-height: 30vh;
}
#showcase h1 {
  font-size: 50px;
  line-height: 1.2;
}

#showcase p {
  font-size: 20px;
}

#showcase .button {
  font-family: "Quicksand", sans-serif;
  font-size: 18px;
  text-decoration: none;
  color: #fff;
  background: #3a76f8;
  padding: 10px 20px;
  border-radius: 10px;
  margin-top: 20px;
  width: 250px;
}

#showcase .button:hover {
  background: #3a76f8;
  color: #fff;
}
</style>

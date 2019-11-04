<div id="fb-root"></div>
<script async defer crossorigin="anonymous" src="https://connect.facebook.net/en_US/sdk.js#xfbml=1&version=v5.0"></script>

<template>
  <div>

      <Sidebar :drawer="drawer" />




<!-- transparent + no logo until scroll point ??-->
    <v-app-bar app dark>
      <v-app-bar-nav-icon class="ma-2" @click.stop="drawer = !drawer"> </v-app-bar-nav-icon>

      <router-link to="/"
        ><v-container
          class="fill-height"
          fluid
          style="max-height: 64px; max-width:100px"
        >
          <!-- <v-img class="ma-8" src="https://github.com/csesoc/csesoc.unsw.edu.au/blob/frontendCombined/src/assets/csesoclogobluewhite.png?raw=true" />  -->
      </v-container>
      </router-link>
       <div class="flex-grow-1"></div>
      <LoginForm align="right"/>
    </v-app-bar>


      <header id="showcase">
        <img src="https://github.com/csesoc/csesoc.unsw.edu.au/blob/frontendCombined/src/assets/csesoclogobluewhite.png?raw=true" />
        <a href="#" v-ripple class="button"> Join on spArc </a>
        <br />
        <v-btn text icon color="white" href="#content-start">
          <v-icon>mdi-chevron-down</v-icon>
        </v-btn>
      </header>
      <ListComponent class="pa-8 ma-8"/>
      <Slider :items="items" class="my-10"/>
    <NavGrid id='content-start' :gridItems="gridItems"></NavGrid>
    <Footer/>
  </div>
</template>

<script>
import NavGrid from '@/components/NavGrid';
import ListComponent from '@/components/ListComponent';
import Footer from '@/components/Footer';
import Sidebar from '@/components/Sidebar';
import Slider from '@/components/Slider';
import LoginForm from '@/components/LoginForm';


export default {
  data: () => ({
    drawer: false,
    gridApiUri: 'https://gistcdn.githack.com/esyw/f83b10232854534b64e475473406dfe6/raw/263a737400bc4e4e642a28cb9da9851ef76e3546/help.json',
    gridItems: [],
    listApiUri: 'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    listItems: [],
    items: [],
  }),

  components: {
    NavGrid,
    ListComponent,
    Footer,
    Sidebar,
    Slider,
    LoginForm,
  },

  mounted() {
    fetch(this.gridApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.gridItems = responseJson;
        this.items=responseJson;
      });
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
  background: rgb(54, 119,243);
  padding: 10px 20px;
  border-radius: 10px;
  margin-top: 20px;
  width: 250px;
}
#showcase .button:hover {
    transition:0.4s;
  background: rgb(32,62,207);
  color: #fff;
}
</style>

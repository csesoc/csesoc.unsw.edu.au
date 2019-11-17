
<template>
  <div id="home">


<!-- make header a seperate component! -->
      <header id="showcase">
        <img src="https://github.com/csesoc/csesoc.unsw.edu.au/blob/frontendCombined/src/assets/csesoclogobluewhite.png?raw=true" />
        <a href="#" v-ripple class="button"> Join on spArc </a>
        <br />
        <v-btn text icon color="white" href="#content-start">
          <v-icon>mdi-chevron-down</v-icon>
        </v-btn>
      </header>

      <v-container class="ma-12 mt-20">
      <v-row no-gutters>
        <v-col>
            <div class="fb-page ml-12" data-href="https://www.facebook.com/csesoc"
            data-tabs="timeline, events, messages" data-width="450" data-height="750" data-small-header="true"
            data-adapt-container-width="true" data-hide-cover="true" data-show-facepile="true">
            <blockquote cite="https://www.facebook.com/csesoc" class="fb-xfbml-parse-ignore">
            <a href="https://www.facebook.com/csesoc">CSESoc UNSW</a></blockquote></div>
        </v-col>
        <v-col>
            <h1> #!/UPCOMING EVENTS </h1>
            <ListComponent :listItems="listItems" />
        </v-col>
      </v-row>
      </v-container>




      <!-- media slider - could be showcase of recent posts?? -->
      <Slider :items="items" class="my-10"/>

      <Events/>
      <h1 class="ml-12 mt-12"> #!/RESOURCES </h1>
      <div class="square">
    <NavGrid id='content-start' :gridItems="gridItems"></NavGrid>
    </div>
  </div>
</template>

<script>
import NavGrid from '@/components/NavGridSquare';
import ListComponent from '@/components/ListComponent';
import Footer from '@/components/Footer';
import Sidebar from '@/components/Sidebar';
import Slider from '@/components/Slider';
import LoginForm from '@/components/LoginForm';
import Events from '@/components/Events'


export default {
  data: () => ({
    drawer: false,
    gridApiUri: 'https://gistcdn.githack.com/esyw/f83b10232854534b64e475473406dfe6/raw/263a737400bc4e4e642a28cb9da9851ef76e3546/help.json',
    gridItems: [],
    listApiUri: 'https://gistcdn.githack.com/esyw/d3801d0bc2b3cefb7fe704a328bb22e8/raw/0010238c3dc744514faf2c859110db5eb6cf9cbe/list-test.json',
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
    Events,
  },

  mounted() {
    fetch(this.gridApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.gridItems = responseJson;
        this.items=responseJson;
      }
    );

    fetch(this.listApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.listItems = responseJson;
      }
    );
  },

};
</script>

<style scoped>
    @import url("https://fonts.googleapis.com/css?family=Quicksand&display=swap");
* {
  margin: 0;
  padding: 0;
}
#home {
    font-family: "Quicksand", sans-serif;

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

.square {
    padding: 10px 300px 50px 300px;
}
</style>

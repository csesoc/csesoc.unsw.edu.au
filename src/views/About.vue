<template>
  <div>
      <Sidebar :drawer="drawer" />





      <v-app-bar app dark>
        <v-app-bar-nav-icon class="ma-2" @click.stop="drawer = !drawer"> </v-app-bar-nav-icon>
        <div class="flex-grow-1"></div>

        <router-link to="/"
          ><v-container
            class="fill-height"
            fluid
            style="max-height: 64px; max-width:100px"
          >
            <v-img class="ma-8" src="https://github.com/csesoc/csesoc.unsw.edu.au/blob/frontendCombined/src/assets/csesoclogobluewhite.png?raw=true" />
        </v-container>
        </router-link>
         <div class="flex-grow-1"></div>
        <LoginForm align="right"/>
      </v-app-bar>

<div class="content">


    <h1 class="mt-12"> #!/ABOUT </h1>
    <p>
        CSESoc is the principal representative body for computing students on campus. We are one of the biggest and most active societies at UNSW, catering to over 3500 CSE students spanning across degrees in Computer Science, Software Engineering, Bioinformatics and Computer Engineering. We are here to fulfil the needs of computing students and also promote computing in all its forms through weekly social and technical events throughout the year.

We are a society for the students, by the students. As a CSE student, here's how we can help you:

We organise and run weekly events, including Trivia, Movie, Poker, and Boardgame Nights, LAN Parties, Workshops, Code Jams, Tech Talks, and our famous Free Weekly BBQ.
We run a highly successful First Year Camp, offering new CSE students (both undergraduate and postgraduate) a weekend of Trivia, Dance Parties, Scavenger Hunts and general frivolity with the chance to meet and mingle with other newcomers.
The society is entirely run by CSE student volunteers. All CSE Students are automatically members, and there are plenty of opportunities to get involved with running events and activities.
We cater to new and returning students alike, with a mix of events designed to have something for everyone. Events are open to all members and nearly all are free to attend.
We're as passionate about computing as you are; we provide technical events to teach you new skills, as well as a fortnightly magazine to keep you updated on what's new and interesting within the university, school and industry.
Most of all, we're here to help you settle in and have a great time at UNSW!
</p>
<NavGrid id='content-start' :gridItems="gridItems"></NavGrid>
</div>
    <Footer/>
  </div>
</template>

<script>
import NavGrid from '@/components/NavGrid';
import ListComponent from '@/components/ListComponent';
import Footer from '@/components/Footer';
import Sidebar from '@/components/Sidebar.vue';
import LoginForm from '@/components/LoginForm';


export default {
  data: () => ({
    drawer: false,
    gridApiUri: 'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    gridItems: [],
    listApiUri: 'https://gistcdn.githack.com/gawdn/464b5ed74404481f7296fb24f9f28243/raw/c9f63e5117a1406db9af5266c8cfd448161bbfec/test_grid.json',
    listItems: [],
  }),

  components: {
    NavGrid,
    ListComponent,
    Footer,
    Sidebar,
    LoginForm
  },

  mounted() {
    fetch(this.gridApiUri)
      .then(r => r.json())
      .then((responseJson) => {
        this.gridItems = responseJson;
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

.content {
    padding: 2vw 15vw 5vw 15vw;
}
</style>

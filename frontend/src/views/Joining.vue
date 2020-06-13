<template>
  <v-app>
    <v-parallax
      height="300"
      src="https://images.unsplash.com/photo-1483577910943-59decea62c35?ixlib=rb-1.2.1&auto=format&fit=crop&w=3889&q=80"
    ></v-parallax>
    <v-expansion-panels flat accordion id="show" style="padding: 20px 20%">
      <v-expansion-panel>
        <v-expansion-panel-header class="title py-3">What is CSESOC?</v-expansion-panel-header>
        <v-expansion-panel-content>
          CSESoc is the principal representative body for computing students on campus. We are one of the biggest and most active societies at UNSW, catering to approximately 1600 CSE students. We are here to fulfill the social needs of computing students and also promote computing in all its forms through weekly social and technical events throughout the year.
          <br />
          <br />We are a society for the students, by the students. As a CSE student, here's how we can help you:
          <br />We organise and run weekly events, including Trivia, Movie, Poker, and Boardgame Nights, LAN Parties, Workshops, Code Jams, Tech Talks, and our famous Free Weekly BBQ.
          <br />
          <br />- We run a highly successful First Year Camp, offering new CSE students (both undergraduate and postgraduate) a weekend of Trivia, Dance Parties, Scavenger Hunts and general frivolity with the chance to meet and mingle with other newcomers.
          <br />- The society is entirely run by CSE student volunteers. All CSE Students are automatically members, and there are plenty of opportunities to get involved with running events and activities.
          <br />- We cater to new and returning students alike, with a mix of events designed to have something for everyone. Events are open to all members and nearly all are free to attend.
          <br />- We're as passionate about computing as you are; we provide technical events to teach you new skills, as well as a fortnightly magazine to keep you updated on what's new and interesting within the university, school and industry.
          <br />
          <br />Most of all, we're here to help you settle in and have a great time at UNSW!
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="title py-3">Not a Arc member?</v-expansion-panel-header>
        <v-expansion-panel-content>
          You need to sign up at Arc in order to join any society on campus.
          <br />Joining Arc (for free!) means you can unlock all the best bits of student life. Set yourself up for success at UNSW by joining Arc online now, then come to visit us on campus for your awesome Arc Membership pack!
          <br />
          <br />
          <br />CLUBS: gain access to over 300 Clubs & Societies
          <br />SPORT: play your way with Sport Clubs, Nationals & Social Sport
          <br />EVENTS: find free food and fun every day on campus
          <br />VOLUNTEERING: lend a hand or an hour to enhance
          <br />WELLNESS: discover self-care and become your best self
          <br />HELP: get support when you need it with Legal & Advocacy
          <br />ART & DESIGN: access resources, meet fellow creatives and take your work to the next level
          <br />HEAPS MORE: bike servicing, free stationery, cheap trips, discounts on campus are just the beginning
          <br />
          <br />Click
          <a href="https://arclimited.formstack.com/forms/arc_membership_signup">HERE</a>
          to join now!
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="title py-3">Not a CSE student?</v-expansion-panel-header>
        <v-expansion-panel-content>
          If you are not enrolled in CSE program, you need to
          <a
            href="https://docs.google.com/forms/d/e/1FAIpQLSfDrhmyDz6F3Q98EEouoUYENwOzCWG1tEes_wJKl8xRzR84gg/viewform"
          >sign up</a> as an associate member with the payment of 10$ per semester.
          <br />
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-card class="mx-auto" style="padding: 50px">
      <v-container fluid>
        <v-row dense>
          <v-col v-for="card in cards" :key="card.title" :cols="card.flex">
            <v-card>
              <v-img
                :src="card.src"
                class="white--text align-end"
                gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
                height="200px"
              >
                <v-card-title v-text="card.title"></v-card-title>
              </v-img>

              <v-card-actions>
                <v-spacer></v-spacer>

                <v-btn icon>
                  <v-icon>mdi-heart</v-icon>
                </v-btn>

                <v-btn icon>
                  <v-icon>mdi-bookmark</v-icon>
                </v-btn>

                <v-btn icon>
                  <v-icon>mdi-share-variant</v-icon>
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    cards: [
        {
          title: "CSESoc at UNSW O-Week!",
          src: "https://scontent-syd2-1.xx.fbcdn.net/v/t1.0-9/84027645_2941067712592848_9084286345676849152_o.jpg?_nc_cat=100&_nc_sid=b386c4&_nc_ohc=MKJ-STTDT0QAX_UX1Bp&_nc_ht=scontent-syd2-1.xx&oh=5cec341c7754c26c3e75fe4f09720b69&oe=5F07FB7C ",
          flex: 4     
        },
        {
          title: "CSESoc Scavenger Hunt!",
          src: "https://scontent-syd2-1.xx.fbcdn.net/v/t1.0-9/s960x960/100995672_3310483902317892_7839912922775027712_o.jpg?_nc_cat=108&_nc_sid=b386c4&_nc_ohc=aYkRm1P5zvkAX_NNU6A&_nc_ht=scontent-syd2-1.xx&_nc_tp=7&oh=d4aaaf8c745f38acf1e8533e5a89ace9&oe=5F0A0545",
          flex: 4
        },
        {
          title: "CSESoc Hoodie Preorders!",
          src: "https://scontent-syd2-1.xx.fbcdn.net/v/t1.0-9/97431643_3260560320643584_6121191148355584000_o.jpg?_nc_cat=103&_nc_sid=b386c4&_nc_ohc=Keipx6N9K8EAX_zlnPw&_nc_ht=scontent-syd2-1.xx&oh=0836a183f2f0774de5e022ebc3e2549f&oe=5F0A9711 ",
          flex: 4
        }
      ]
    }),
  components: []
};
</script>

<style scoped>
</style>

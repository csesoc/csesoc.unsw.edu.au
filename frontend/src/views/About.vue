<!--
  About
  --
  This view corresponds to the About page of the website.
  This page consists of:
    - about
    - executives & directors
    - subcommittees
    - history
-->

<template>
  <v-app>
    <SecondaryNavbar />
    <v-parallax height="300" src="../assets/fade.png"></v-parallax>
    <!-- About -->
    <v-container ref="content-start" style="padding: 0px 30px 10px 30px">
      <HeaderTitle title="About" padding="50px 0px 0px 0px"></HeaderTitle>
      <p>
        <b>
          CSESoc is the official representative body of computing students at UNSW. We are one of the largest
          and most active societies at UNSW, and the largest computing society in the southern hemisphere.
          CSESoc comprises ~10k members spanning across degrees in Computer Science, Software
          Engineering, Bioinformatics and Computer Engineering. We are here to fulfil the social, personal and
          professional needs of CSE students, and promote computing through a variety of forms.
        </b>
      </p>
      <p><b>We are a society for the students, by the students. Here’s an overview of what we do;</b></p>
      <ul class="about-list">
        <li>
          Run weekly social and educational events, including trivia, movie, boardgames nights, LAN parties,
          workshops, coding competitions, tech talks, and our famous free weekly BBQ.
        </li>
        <li>Create original media content, including Podcasts, articles, YouTube videos, and live streams</li>
        <li>
          Run a highly successful First Year Camp and Peer Mentoring program, offering new CSE students (both
          undergraduate and postgraduate) a chance to meet and mingle with other newcomers
        </li>
        <li>
          Engage students with industry sponsors and representatives to develop their professional capacity
          and curiosity
        </li>
        <li>
          Develop our own open-source projects for students to get learn new skills and develop tools for our
          community
        </li>
        <li>
          Facilitate an online community of ~6k Discord users, ~6.7k Facebook followers, ~1.2k YouTube subs, and
          ~1.6k Instagram followers
        </li>
      </ul>
    </v-container>
    <v-container ref="content-start" style="padding: 0px 30px 10px 30px">
      <HeaderTitle title="2022 Statistics" padding="50px 0px 0px 0px"></HeaderTitle>
      <p>
        <b>
          <ul class="stats-list">
            <li>🥳 100+ events (35+ careers related revents) 🥳</li>
            <li>💸 41 sponsors 💸</li>
            <li>💬 480 000 discord messages 💬</li>
            <li>✨ 200+ volunteers ✨</li>
            <li>📼 Tiktok: 27.2k views, 2k+ likes 📼</li>
            <li>📼 Podcast: 10k downloads, 600+ average viewers 📼</li>
            <li>🚸 500+ high school students reached 🚸</li>
          </ul>
        </b>
      </p>
    </v-container>
  </v-app>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';
import SecondaryNavbar from '../components/SecondaryNavbar';

const fs = require('@/assets/execs_directs.json');

export default {
  name: 'About',
  data: () => ({
    execsDirects: fs,
  }),
  components: {
    HeaderTitle,
    SecondaryNavbar,
  },
};
</script>

<style scoped>
td {
  padding-top: 5px;
  padding-bottom: 5px;
}

th {
  padding-right: 10px;
}

.about-list > li {
  padding-bottom: 1vh;
}

.stats-list {
  list-style-type: none;
  padding: 0;
  margin: 0;
}
</style>

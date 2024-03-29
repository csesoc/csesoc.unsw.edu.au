<!--
  Engage
  --
  This view corresponds to the Engage page of the website.
  This page consists of:
    - joining
    - social links
    - faq
    - enquiry and feedback forms
-->

<template>
  <v-app>
    <v-parallax height="300" src="../assets/fade.png"></v-parallax>

    <!-- Joining -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="Joining"></HeaderTitle>
      <p>
        <b>
          Everybody is welcome to join CSESoc! To get started as a UNSW student, all you have to do is
          <a href="https://docs.google.com/forms/d/e/1FAIpQLSfDrhmyDz6F3Q98EEouoUYENwOzCWG1tEes_wJKl8xRzR84gg/viewform"
            >sign up with spARC</a
          >! <a href="https://member.arc.unsw.edu.au/s/clubdetail?clubid=0016F0000371VxAQAU"> join us!</a></b
        >
      </p>
      <p>
        <b
          >Even if you aren't a member, our Discord, social media and resources are mostly freely available to everybody. We look
          forward to hearing from you!</b
        >
      </p>
      <v-expansion-panels flat accordion id="show">
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Joining UNSW spARC</v-expansion-panel-header>
          <v-expansion-panel-content>
            You need to sign up at Arc in order to join any society on campus.
            <br />Joining Arc (for free!) means you can unlock all the best bits of student life.
            Set yourself up for success at UNSW by joining Arc online now, then come to visit us
            on campus for your awesome Arc Membership pack!
            <br /><br />
            <br />CLUBS: gain access to over 300 Clubs & Societies
            <br />SPORT: play your way with Sport Clubs, Nationals & Social Sport
            <br />EVENTS: find free food and fun every day on campus
            <br />VOLUNTEERING: lend a hand or an hour to enhance
            <br />WELLNESS: discover self-care and become your best self
            <br />HELP: get support when you need it with Legal & Advocacy
            <br />ART & DESIGN: access resources, meet fellow creatives and take your work to the next level
            <br />HEAPS MORE: bike servicing, free stationery, cheap trips, discounts on campus are just the beginning
            <br />
            <br />Click <a href="https://arclimited.formstack.com/forms/arc_membership_signup">HERE</a> to join now!
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">I'm not a CSE or UNSW student, can I join?</v-expansion-panel-header>
          <v-expansion-panel-content>
            If you are not enrolled in a CSE program or course, you need to
            <a href="https://docs.google.com/forms/d/e/1FAIpQLSfDrhmyDz6F3Q98EEouoUYENwOzCWG1tEes_wJKl8xRzR84gg/viewform"
              >sign up</a
            >
            as an associate member with a payment of $10 per semester.
            <br />
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-container>

    <!-- Social media links -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="Social links"></HeaderTitle>
      <NavGrid :items="socialLinks"></NavGrid>
    </v-container>

    <!-- FAQ -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="FAQ"></HeaderTitle>
      <v-expansion-panels flat accordion id="show" style="padding: 20px">
        <v-expansion-panel v-for="faq in faqLinks" :key="faq.question">
          <v-expansion-panel-header class="title py-3">{{ faq.question }}</v-expansion-panel-header>
          <v-expansion-panel-content>{{ faq.answer }}</v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-container>

    <!-- Forms -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <h2>Still have some questions?</h2>
      <br />

      <v-btn-toggle v-model="activeForm" mandatory>
        <v-btn value="general" data-cy="general-form-selector">
          Enquiry
        </v-btn>
        <v-btn value="feedback" data-cy="feedback-form-selector">
          Feedback
        </v-btn>
      </v-btn-toggle>

      <v-card flat tile>
        <MailingForm :type="this.activeForm"></MailingForm>
      </v-card>
    </v-container>
  </v-app>
</template>

<script>
import NavGrid from '@/components/NavGrid';
import HeaderTitle from '@/components/HeaderTitle';
import MailingForm from '@/components/MailingForm';

import APIClient from '../utils/APIClient';

export default {
  name: 'Engage',
  data: () => ({
    activeForm: 'general',
    socialLinks: [],
    faqLinks: []
  }),
  components: {
    NavGrid,
    HeaderTitle,
    MailingForm
  },

  mounted() {
    APIClient.socialsAPI().then((responseJson) => {
      this.socialLinks = responseJson;
    });

    APIClient.faqsAPI().then((responseJson) => {
      this.faqLinks = responseJson;
    });
  }
};
</script>

<style scoped></style>

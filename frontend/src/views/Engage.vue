<template>
  <v-app>
    <v-parallax
      height="300"
      src="https://images.unsplash.com/photo-1592251170558-01228a0a06f2?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1920&q=80"
    ></v-parallax>

    <!-- Social media links -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="Social links"></HeaderTitle>
    </v-container>

    <!-- FAQ -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="FAQ"></HeaderTitle>
      <v-expansion-panels flat accordion id="show" style="padding: 20px 20%">
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">What is CSESoc?</v-expansion-panel-header>
          <v-expansion-panel-content>
            CSESoc is the principal representative body for computing students on campus. We are one of the biggest and most active societies at UNSW, catering to approximately 1600 CSE students. We are here to fulfill the social needs of computing students and also promote computing in all its forms through weekly social and technical events throughout the year.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">How can we help CSE students?</v-expansion-panel-header>
          <v-expansion-panel-content>
            We are a society for the students, by the students. As a CSE student, we can help you through the weekly events we run and organise; these include: trivia, movie, poker, and boardgame nights, LAN Parties, workshops, code jams, tech talks, and our famous FREE weekly BBQ. Additionally:
            <br />- We run a highly successful First Year Camp, offering new CSE students (both undergraduate and postgraduate) a weekend of Trivia, Dance Parties, Scavenger Hunts and general frivolity with the chance to meet and mingle with other newcomers.
            <br />- The society is entirely run by CSE student volunteers. All CSE Students are automatically members, and there are plenty of opportunities to get involved with running events and activities.
            <br />- We cater to new and returning students alike, with a mix of events designed to have something for everyone. Events are open to all members and nearly all are free to attend.
            <br />- We're as passionate about computing as you are; we provide technical events to teach you new skills, as well as a fortnightly magazine to keep you updated on what's new and interesting within the university, school and industry.
            <br />
            <br />Most of all, we're here to help you settle in and have a great time at UNSW!
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">I'm not an Arc member, can I join?</v-expansion-panel-header>
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
          <v-expansion-panel-header class="title py-3">I'm not a CSE student, can I join?</v-expansion-panel-header>
          <v-expansion-panel-content>
            If you are not enrolled in CSE program, you need to
            <a
              href="https://docs.google.com/forms/d/e/1FAIpQLSfDrhmyDz6F3Q98EEouoUYENwOzCWG1tEes_wJKl8xRzR84gg/viewform"
            >sign up</a> as an associate member with the payment of 10$ per semester.
            <br />
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-container>



    <!-- Forms -->
    <v-tabs
      class="elevation-2"
      vertical
      dark
    >
      <v-tabs-slider></v-tabs-slider>

      <!-- Enquiry tab -->
      <v-tab>Enquiry</v-tab>
      <v-tab-item>
        <v-card flat tile>
          <EnquiryForm></EnquiryForm>
        </v-card>
      </v-tab-item>

      <!-- Feedback tab -->
      <v-tab>Feedback</v-tab>
      <v-tab-item>
        <v-card flat tile>
          <FeedbackForm></FeedbackForm>
        </v-card>
      </v-tab-item>
    
    </v-tabs>
  </v-app>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';
import EnquiryForm from '@/components/EnquiryForm';
import FeedbackForm from '@/components/FeedbackForm';

export default {
  data: () => ({

  }),
  components: {
    HeaderTitle,
    EnquiryForm,
    FeedbackForm
  }
};
</script>

<style scoped>
</style>

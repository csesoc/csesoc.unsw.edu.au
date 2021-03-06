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

    <!-- About -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="About"></HeaderTitle>
      <p>
        CSESoc is the principal representative body for computing students on campus. We are one of the biggest and most
        active societies at UNSW, catering to over 3500 CSE students spanning across degrees in
        Computer Science, Software Engineering, Bioinformatics and Computer Engineering.
        We are here to fulfil the needs of computing students and also promote computing in all
        its forms through weekly social and technical events throughout the year.
      </p>
      <p>
        We are a society for the students, by the students. As a CSE student, here's how we can help you:
        We organise and run weekly events, including Trivia, Movie, Poker, and Boardgame Nights, LAN Parties,
        Workshops, Code Jams, Tech Talks, and our famous Free Weekly BBQ.
        We run a highly successful First Year Camp, offering new CSE students (both undergraduate and postgraduate)
        a weekend of Trivia, Dance Parties, Scavenger Hunts and general frivolity with the chance to meet and mingle
        with other newcomers.
        The society is entirely run by CSE student volunteers. All CSE Students are automatically members, and
        there are plenty of opportunities to get involved with running events and activities.
        We cater to new and returning students alike, with a mix of events designed to have something for everyone.
        Events are open to all members and nearly all are free to attend.
        We're as passionate about computing as you are; we provide technical events to teach you new skills,
        as well as a fortnightly magazine to keep you updated on what's new and interesting within the university,
        school and industry.
      </p>
      <p>
        Most of all, we're here to help you settle in and have a great time at UNSW!
      </p>
    </v-container>

    <!-- Executives & Directors -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="Executives & Directors"></HeaderTitle>
      <p>
        The Executives and Directors are responsible for organising the society to make sure that things get done. Execs
        are elected annually by CSE students at the end of the preceding year and Directors are selected by Execs.
      </p>
      <p>
        The CSESoc Executive and Director teams for the current and past years are:
      </p>

      <v-tabs class="elevation-2" vertical dark>
        <v-tabs-slider></v-tabs-slider>

        <v-tab v-for="team in this.execsDirects" :key="team.year">
          {{team.year}}
        </v-tab>
        <v-tab-item v-for="team in this.execsDirects" :key="team.year">
          <v-card flat tile>

            <v-row no-gutters>
              <v-col v-for="(tierValue, tierName) in team.data" :key="tierName" md="4" style="margin: 0px 50px">
                <table>
                  <thead>
                    <tr><th colspan="2">
                      <h3>{{tierName}}</h3>
                    </th></tr>
                  </thead>
                  <tbody>
                    <tr v-for="(roleValue, roleName) in tierValue" :key="roleName">
                      <th>{{roleName}}</th>
                      <td>
                        <tr v-for="name in roleValue" :key="name">
                          {{name}}
                        </tr>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </v-col>
            </v-row>

          </v-card>
        </v-tab-item>
      </v-tabs>
    </v-container>

    <!-- Subcommittees -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="Subcommittees"></HeaderTitle>
      <p>
        The CSESoc Subcommittees are an integral part of our society and are responsible for directly interacting with the
        students and sponsors through events, publications, this website, etc. Each subcommittee is run by the portfolio
        directors, and those directors are selected by the executive and is part of the director board.
        The CSESoc Executive and Directors hope you have enjoyed our events and activities - so much so that you wish to
        get involved!
        If you are interested in volunteering in a specific area or are thinking about running for an executive position
        in the future, then joining a subcom is the best way to get involved, meet other people and try something new!
      </p>
      <v-expansion-panels flat accordion id="show">
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Careers</v-expansion-panel-header>
          <v-expansion-panel-content>
            One of the core goals of CSESoc is to provide students with a large variety of opportunities to build and
            diversify their careers and professional portfolios. Our members are recognised around the globe as world
            class computing students and many of them go on to work at an extremely wide range of organisations in an
            assortment of different roles.
            <br /><br />
            The careers team aims to imbue our members with the skills, experience, and connections needed to forge a
            successful career in the technology industry. We organise and host events throughout semester designed to
            achieve this goal - such as tech talks and workshops, networking nights, resume and interview workshops,
            hackathons, coding competitions and so on. We also help manage sponsor relations, as sponsors play a key role
            in everything we do. This subcommittee helps its members develop both personally and professionally as they
            correspond with our points of contact in the world's biggest tech companies, such as Google, Microsoft, Amazon,
            Facebook, Atlassian and Canva.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">CompClub</v-expansion-panel-header>
          <v-expansion-panel-content>
            CompClub is the high school outreach program of the Computer Science and Engineering Society (CSESoc) at UNSW
            Sydney. CompClub exists to introduce high school students to the basics of Computer Science and to encourage
            students to study CSE. CompClub runs standalone workshops throughout the year at UNSW and participating high
            schools as well as biannual programs at UNSW known as the Summer and Winter Schools.
            <br /><br />
            If you are a teacher or educator who is interested in contacting us, please email us at compclub@csesoc.org.au.
            If you are a high school student, keep an eye out for our Winter Workshops.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Creative</v-expansion-panel-header>
          <v-expansion-panel-content>
            The Creative Team is responsible for expressing our creative freedoms whilst
            integrating what we create fully with CSESoc branding and vision. We believe in
            combining the logical and analytical side of CSE with creativity and artistic
            expression. Not only this, we provide an avenue for subcom to learn something
            new, come up with creative CSE projects and develop your creative minds.
            <br /><br />
            The Events Team works with other portfolios to deliver event-based design
            demands such as banners on a request basis, focused on developing a professional
            relationship with not only CSESoc, but more importantly our sponsored
            companies!
            <br /><br />
            The Video Team handles the development of videos and editing for
            publishing/advertising. This year we want to explore new territories of motion
            graphics and animations to develop videos such as animated video banners,
            informational videos and promotional videos for events.
            <br /><br />
            The Merchandise Team is responsible for coming up with innovative ideas for our
            upcoming merchandise - such as stickers, hoodies and team shirts.
            <br /><br />
            The wildcard is an all-rounder role which covers all aspects of creative expression.
            This includes, but not limited to, organising subcom upskilling workshops,
            intersociety design related collaboration, general creative design and logo design.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Marketing</v-expansion-panel-header>
          <v-expansion-panel-content>
            The Marketing team is responsible for the effective promotion and marketing of
            CSESoc and its events. CSESoc is a premier society at UNSW so we aim to ensure
            that our image is continually strengthened while delivering valuable events and
            opportunities to CSE students.
            <br /><br />
            The Social Media team will be in charge of the creation and promotion of events
            and opportunities provided by the CSESoc society, sometimes including content
            creation. We also look at reviewing current and creating new strategies to
            maximise our reach. We also collaborate with other portfolios to ensure maximum
            exposure and success of events.
            <br /><br />
            The Photography/Video Team is a great place for those who are creatively inclined,
            the role will include taking photos and videos at social events, creating and
            managing media, and creating promotional content.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Media</v-expansion-panel-header>
          <v-expansion-panel-content>
            The Media Team is responsible for cementing the CSESoc’s brand and student voice, managing content that
            enriches the academic, social and cultural experience of CSE students all while acting as a bridge for
            students ideas, opinions and advice to the wider society. Some of the type of content media is responsible for
            producing include but not limited to are podcasts, videos, articles and the First Year Guide.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Socials</v-expansion-panel-header>
          <v-expansion-panel-content>
            The Socials Team plans, organises and executes all of CSESoc’s social events -
            including the weekly BBQs, pub crawl, ice skating, cardboard night, and lots more.
            In planning these events, we do our best to not only create opportunities for
            students across all years to form friendships, but to also create spaces where
            these bonds can continue to cultivate throughout their entire university careers.
            The best part about being part of this team is how rewarding it feels to watch as
            everything comes together and as everyone enjoys themselves within an
            environment that we created! In addition, you will also be joining a family of about
            50 people, and we will be going on trips and doing bonding activities together.
            Being part of the Socials Team is also a good way to develop leadership, teamwork
            and communication skills, and is something employers love to see on a resume.
            <br /><br />
            This year, we hope to expand our reach by engaging with our members more,
            doing our best to instigate friendships between strangers, and collaborating with
            our societies (and potentially other universities). We are looking for people from
            all years who are committed, passionate and excited about meeting new people
            and fostering these bonds!
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Projects</v-expansion-panel-header>
          <v-expansion-panel-content>
            CSESoc is a society for computing students, so it's only fair we do some computing.
            The Projects Team is a place to create student-led projects with the goal of building
            transferable technical and soft skills so our members can thrive in university, work and
            anything else you want to do.
            <br /><br />
            These projects are normally built for university students, which means when these
            projects are completed, they will be used by yourself and your peers! It is a great way to
            get involved in CSESoc while branching out and trying a technology or language that you
            are interested in.
            <br /><br />
            We encourage everyone to join the Projects team, regardless of experience level. Even if you don't feel like
            you can contribute technically, we want you for your ideas and your enthusiasm. Projects can be a great chance
            to learn some of the gaps in your knowledge that lectures might have left out, and to offer your first taste
            of designing and building things for the real world.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Student Network</v-expansion-panel-header>
          <v-expansion-panel-content>
            The Stunet Team works as part of the Socials portfolio to better the student experience!
            We aim to help incoming CSE students to integrate into life in CSE at UNSW,
            facilitate students connecting with each other, combat the feeling of "imposters
            syndrome" felt by new CSE students, elevate well-being of CSE students in general.
            We run workshops, mentoring programs, study sessions, and help facilitate CSE
            Events.
            <br /><br />
            Stunet works closely and Socials - and our role is quite diverse! Additionally, this year
            we are hoping to expand our program by improving cohesion and increase
            interaction between the year levels and undergraduate-postgraduate.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Workshops</v-expansion-panel-header>
          <v-expansion-panel-content>
            The Workshops team organises student-run educational events throughout the year to help CSE students (and anyone else,
            really) make use of both new, and widely used technologies. Additionally, Workshops promotes the use of Linux through
            installfests by helping students set up and solve problems with Linux.
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-container>

    <!-- History -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="History"></HeaderTitle>
      <p>
        CSESoc was formed in October 2006 from the old CompSoc and SESoc societies. CompSoc helped represent the interests
        of students studying Computer Engineering, Computer Science and postgraduate courses, while SESoc was the
        representative body for Software Engineering students. Both societies provided technical and social support to
        their members. In the best interests of everyone, the societies merged to provide a better experience to all CSE
        students.
      </p>
      <p>
        CSESoc now represents students enrolled in Computer Science, Computer Engineering, Software Engineering,
        Bioinformatics Engineering, or a post‐graduate program administered by CSE (research or coursework).
      </p>
      <p>
        Even today CSESoc continues to be an integral part of the student experience. Many students make the most of their
        time at university by joining a working group in first year to get a taste of the society. If you are enthusiastic
        and interested you can nominate yourself or be nominated for a position in the Exec at the end of the year.
      </p>
      <p>
        Being part of a society is a great way to meet new people and gain extra skills that employers look for in the
        industry.
      </p>
    </v-container>
  </v-app>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

const fs = require('@/assets/execs_directs.json');

export default {
  name: 'About',
  data: () => ({
    execsDirects: fs
  }),
  components: {
    HeaderTitle
  }
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
</style>

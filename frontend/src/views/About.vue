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
      <HeaderTitle title="Executives_&_Directors"></HeaderTitle>
      <p>
        The Executives and Directors are responsible for organising the society to make sure that things get done. Execs 
        are elected annually by CSE students at the end of the preceding year and Directors are selected by Execs.
      </p>
      <p>
        The CSESoc Executive and Director team for 2020 is:
      </p>

      <table class="table table-bordered table-striped">
        <thead>
          <tr><th colspan="2">
            <h3>Executives</h3>
          </th><th colspan="2">
            <h3>Directors</h3>
          </th></tr>
        </thead>
        <tbody>
          <tr>
            <th width="20%">Co-Presidents</th>
            <td width="30%">Shane Kadish<br>Tammy Zhong</td>
            <th width="20%">Careers</th>
            <td>Evangeline Endacott<br>Low Khye Ean<br>Nicholas Duller<br>Yan Zhai</td>
          </tr>
          <tr>
            <th>Secretary</th>
            <td>Sam Push</td>
            <th>CompClub</th>
            <td>Livia Wijayanti<br>Timothy Ryan</td>
          </tr>
          <tr>
            <th>Treasurer</th>
            <td>Teresa Feng</td>
            <th>Creative</th>
            <td>Elizabeth Zhong<br>Jia Min Guo</td>
          </tr>
          <tr>
            <th>Arc Delegate</th>
            <td>Jeremy Chiu</td>
            <th>Marketing</th>
            <td>Isaac Joshi<br>Jessica Feng</td>
          </tr>
          <tr>
            <th>Grievance Officer</th>
            <td>Tom Kunc</td>
            <th>Media</th>
            <td>Aditi Chandra<br>Clarence Shijun Feng</td>
          </tr>
          <tr>
            <th>&nbsp;</th>
            <td>&nbsp;</td>
            <th>Socials</th>
            <td>Frances Lee<br>Van-Roy Trinh</td>
          </tr>
          <tr>
            <th>&nbsp;</th>
            <td>&nbsp;</td>
            <th>Software Projects</th>
            <td>Gordon Zhong<br>Leesa Kristina Dang</td>
          </tr>
          <tr>
            <th>&nbsp;</th>
            <td>&nbsp;</td>
            <th>Student Network</th>
            <td>Sachin Krishnamoorthy<br>Shrey Somaiya</td>
          </tr>
          <tr>
            <th>&nbsp;</th>
            <td>&nbsp;</td>
            <th>Workshops</th>
            <td>Michael Gribben</td>
          </tr>
        </tbody>
      </table>


    </v-container>

    <!-- Subcommittees -->
    <v-container ref="content-start" style="padding: 20px 30px 10px 30px">
      <HeaderTitle title="Subcommittees"></HeaderTitle>
      <p>
        The CSESoc Subcommittees are an integral part of CSESoc and are responsible for directly interacting with the 
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
            achieve this goal - such as sponsor talks and workshops, networking nights, resume and interview workshops, 
            hackathons, and so on. We also help manage sponsor relations, as sponsors play a key role in everything we do. 
            This year, we're also planning on launching a number of exciting professional development initiatives, which, 
            combined with our events, look to form a strong platform in maximising opportunities for all of our 
            constituents.
            <br /><br />
            If that sounds like an amazing and meaningful role to play, then watch out for the next CSESoc recruitment drive - we're looking for passionate and dedicated people to get involved in the Careers team!
            <br /><br />
            Contact Email: careers@csesoc.org.au
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">CompClub</v-expansion-panel-header>
          <v-expansion-panel-content>
            CompClub is the high school outreach program of the Computer Science and Engineering Society (CSESoc) at UNSW 
            Sydney. CompClub exists to introduce high school students to the basics of Computer Science and to encourage 
            students to study CSE. CompClub runs ​standalone workshops throughout the year at UNSW and participating high 
            schools as well as biannual programs at UNSW known as the Summer and Winter Schools.
            <br /><br />
            If you are a teacher or educator who is interested in contacting us, please email us at compclub@csesoc.org.au.
            If you are a high school student, have a look at our Winter Workshops.
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Creative</v-expansion-panel-header>
          <v-expansion-panel-content>
            
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Marketing</v-expansion-panel-header>
          <v-expansion-panel-content>
            CSESoc's Marketing team is all about promoting the countless sponsored career events and the amazing social 
            events that happen throughout uni life. The marketing team allows you to help market CSESoc to students and 
            our sponsors. From designing merchandise, to creating a Facebook event for one of our fantastic social events, 
            tech talks or workshops, to taking photos and videos to capture these moments in an everlasting medium. Come 
            along, join in - all levels of experience welcome! To get involved with Marketing, lookout for the next 
            recruitment drive.
            <br /><br />
            Contact Email: marketing@csesoc.org.au
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
            CSESoc's Social team is in charge of making university life unforgettable, fun and to forge a strong community 
            within the society. We are responsible for the weekly barbeques, the classic Cardboard Night and Trivia 
            Nights, as well as the renown CSESoc Cruise! We are looking for YOU to help provide new ideas and improve on 
            our events. If you are keen on expanding your personal bubble and are looking to develop your skills, whilst 
            making a difference to student life of CSESoc members, then look no further.
            <br /><br />
            Contact Email: socials@csesoc.org.au
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Software Projects</v-expansion-panel-header>
          <v-expansion-panel-content>
            CSESoc is a society for computing students, so it's only fair we do some computing. The CSESoc Projects team 
            works on projects like the society's website, and other student led initiatives. We also work with the 
            Workshop team to run workshops, to build skills that can be applied in ours' as well as projects.
            We want to see the Projects team grow, and give members the skills their project ideas to reality, both inside 
            and outside the society.
            <br /><br />
            We encourage everyone to join the Projects team, regardless of experience level. Even if you don't feel like 
            you can contribute technically, we want you for your ideas and your enthusiasm. Projects can be a great chance 
            to learn some of the gaps in your knowledge that lectures might have left out, and to offer your first taste 
            of designing and building things for the real world.
            <br /><br />
            Contact Email: software.projects@csesoc.org.au
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Student Network</v-expansion-panel-header>
          <v-expansion-panel-content>
            
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header class="title py-3">Workshops</v-expansion-panel-header>
          <v-expansion-panel-content>
            The Workshops team organises student-run educational events throughout the year to help CSE students (and 
            anyone else, really) make use of both new, and widely used technologies. Additionally, Workshops promotes the 
            use of Linux through installfests by helping students set up and solve problems with Linux. This enables CSE 
            students to go above and beyond their degree, and we want your help! Either as someone with a lot of 
            experience to spread knowledge, or someone with a little experience to translate from 1337speak to n00bspeak, 
            or someone with no experience to provide a roadmap of what CSE students want to know.
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-container>



  </v-app>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

export default {
  data: () => ({

  }),
  components: {
    HeaderTitle
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
</style>

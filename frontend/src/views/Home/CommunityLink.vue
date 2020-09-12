<!--
  Community Links
  --
  This view contains the links to community social channels.
-->

<template>
<!-- Join Us Icons -->
  <div id="community" class="content">
    <v-container data-cy="communityLinks">
      <v-row no-gutters>
        <HeaderTitle
          title="BE PART OF THE COMMUNITY"
          subtitle="Find our stall at O-week, or just turn up to an event and get to know us!"
        />
      </v-row>
      <v-row class="pa-0" data-cy="communityLinks-images">
        <!-- Display community links -->
        <v-col cols="4" align="center" class="pa-0" v-for="channel in communityLinks" :key="channel.title">
          <a :href="channel.link">
            <img :src="channel.logo" class="logo zoom" data-cy="community-links-image">
          </a>
          <h2>{{channel.title}}</h2>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'CommunityLinks',
  data: () => ({
    communityLinks: [
      // community links have a link, logo and title
      {
        link: 'https://csesoc-community.slack.com/',
        logo: require('../../assets/slack_logo.png'),
        title: 'Slack Community',
      },
      {
        link: 'https://bitly.com/CSESocCourseChats',
        logo: require('../../assets/facebook_logo.png'),
        title: 'Facebook Group',
      },
      {
        link: 'https://forms.office.com/Pages/ResponsePage.aspx?id=pM_2PxXn20i44Qhnufn7o6ecLZTBorREjnXuTY-JfmBUMEpOMFBDTU1UWkhBWllWRTNPOVJFMUNCRi4u',
        logo: require('../../assets/discord_logo.png'),
        title: 'Discord Server',
      }
    ]
  }),
  components: {
    HeaderTitle
  }
};
</script>

<style scoped lang="scss">
.logo {
  width: 150px;
  height: auto;
  margin-bottom: $space-xxs;
  transition: transform .2s;
}

.logo:hover {
  transform: scale(1.1);
}
</style>

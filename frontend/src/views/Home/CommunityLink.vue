<!--
  Community Links
  --
  This view contains the links to community social channels.
-->

<template>
<!-- Join Us Icons -->
  <div id="community" class="content">
    <v-container data-cy="communityLinks">
      <!-- <v-row no-gutters>
        <HeaderTitle
          title="GET IN TOUCH"
        />
      </v-row> -->
      <v-row>
        <HeaderTitle
          title="GET IN TOUCH"
        />
      </v-row>
      <v-layout row wrap justify-space-between data-cy="communityLinks-images">
        <v-col cols="12" xs="12" sm="4" md="4" v-for="channel in communityLinks" :key="channel.title" >
          <v-btn block outlined color="white" target="_blank" :href="channel.link">
            <!-- <img :src="channel.logo" class="logo" data-cy="community-links-image"> -->
            <v-icon class="icons">{{`mdi-${channel.logo}`}}</v-icon>
            <div class="link-text">{{channel.title}}</div>
          </v-btn>
        </v-col>
        <!-- <v-flex md3> -->
          <!-- <v-btn outline block class="primary"> 1</v-btn> -->
          <!-- <a target="_blank" :href="channel.link">
              <img :src="channel.logo" class="logo" data-cy="community-links-image">
            </a>
            <h4 class="hidden-sm-and-down">{{channel.title}}</h4> -->
        <!-- </v-flex>
        <v-flex md3>
          <v-btn outline block class="primary"> 1</v-btn>
        </v-flex> -->
      </v-layout>
      <div class="d-flex justify-space-between icon-container">
        <div v-for="channel in communityLinks2" :key="channel.logo">
          <v-btn :href="channel.link" icon color="white light-2">
            <v-icon>{{`mdi-${channel.logo}`}}</v-icon>
          </v-btn>
        </div>
      </div>
      <!-- <v-layout>
        <v-icon></v-icon>
      </v-layout> -->
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
        // https://forms.office.com/Pages/ResponsePage.aspx?id=pM_2PxXn20i44Qhnufn7o6ecLZTBorREjnXuTY-JfmBUMEpOMFBDTU1UWkhBWllWRTNPOVJFMUNCRi4u
        link: 'https://discord.com/invite/b5ynAtP',
        logo: 'discord',
        title: 'Discord',
      },
      {
        link: 'https://www.facebook.com/csesoc/',
        logo: 'facebook',
        title: 'Facebook',
      },
      {
        link: 'https://www.facebook.com/groups/2509117190/?ref=pages_group_cta',
        logo: 'account-group',
        title: 'Facebook Group',
      }
    ],

    communityLinks2: [
      {
        link: 'https://twitter.com/csesoc?lang=en',
        logo: 'twitter',
      },
      {
        link: 'https://www.youtube.com/channel/UC1JHpRrf9j5IKluzXhprUJg',
        logo: 'youtube-subscription',
      },
      {
        link: 'https://www.instagram.com/csesoc_unsw/?hl=en',
        logo: 'instagram',
      },
      {
        link: 'https://www.tiktok.com/@csesoc?lang=en',
        logo: 'music-note-eighth',
      },
      {
        link: 'https://www.facebook.com/csesoc/photos/a.279434445422868/4432507913448813/',
        logo: 'wechat',
      },
      {
        link: 'https://www.linkedin.com/company/csesoc/',
        logo: 'linkedin',
      },

    ]
  }),
  components: {
    HeaderTitle,
  }
};
</script>

<style scoped lang="scss">
.logo {
  width: 30px;
  height: auto;
 // margin-bottom: $space-xxs;
  transition: transform .2s;

  @media only screen and (max-width: $breakpoint-medium) {
    width: 30px;
  }
}

.logo:hover {
  transform: scale(1.1);
}
.content {
  @media screen and (max-width: 960px) {
    padding-left: 20px;
    padding-right: 20px;
  }
}
.v-btn {
  min-height: 70px;
}

.v-icon {
  margin-right: 10px;
  @media screen and (min-width: 960px) {
    font-size: 40px;
  }
  @media screen and (min-width: 845px) {
    font-size: 30px;
  }
  @media screen and (min-width: 788px){
    font-size: 25px;
  }
  @media screen and (min-width: 700px){
    font-size: 20px;
  }
  @media screen and (min-width: 600px){
    font-size: 18px;
  }
  @media screen and (max-width: 600px){
    font-size: 30px;
  }
}

.link-text {
  font-weight: bold;
  @media screen and (min-width: 1264px) {
    font-size: 30px;
  }
  @media screen and (min-width: 845px) {
    font-size: 20px;
  }
  @media screen and (min-width: 788px){
    font-size: 18px;
  }
  @media screen and (min-width: 700px) {
    font-size: 14px;
  }
  @media screen and (min-width: 600px) {
    font-size: 11px;
  }
  @media screen and (max-width: 600px){
    font-size: 25px;
  }

  .icon-container {
    width: 500px;
    height: 10px;
  }
}

</style>

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
          <v-btn block outlined color="white" target="_blank" :href="channel.link" class="logo">
            <i :class="channel.logo"></i>
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
      <div class="d-flex justify-space-between icon-container hide-xs-and-down">
        <div v-for="channel in communityLinks2" :key="channel.logo">
          <v-btn :href="channel.link" color="white light-2" target="_blank" icon class="other-btn">
            <i :class="channel.logo"></i>
          </v-btn>
        </div>
      </div>
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
        logo: 'fab fa-discord',
        title: 'Discord',
      },
      {
        link: 'https://www.facebook.com/csesoc/',
        logo: 'fab fa-facebook',
        title: 'Facebook',
      },
      {
        link: 'https://www.facebook.com/groups/2509117190/?ref=pages_group_cta',
        logo: 'fas fa-users',
        title: 'Facebook Group',
      }
    ],

    communityLinks2: [
      {
        link: 'https://twitter.com/csesoc?lang=en',
        logo: 'fab fa-twitter',
      },
      {
        link: 'https://www.youtube.com/channel/UC1JHpRrf9j5IKluzXhprUJg',
        logo: 'fab fa-youtube',
      },
      {
        link: 'https://www.instagram.com/csesoc_unsw/?hl=en',
        logo: 'fab fa-instagram',
      },
      {
        link: 'https://www.tiktok.com/@csesoc?lang=en',
        logo: 'fab fa-tiktok',
      },
      {
        link: 'https://www.facebook.com/csesoc/photos/a.279434445422868/4432507913448813/',
        logo: 'fab fa-weixin',
      },
      {
        link: 'https://www.linkedin.com/company/csesoc/',
        logo: 'fab fa-linkedin-in',
      },

    ]
  }),
  components: {
    HeaderTitle,
  },

  computed: {
    size() {
      if (this.$vuetify.breakpoint.width <= 520) {
        return 30;
      }
      switch (this.$vuetify.breakpoint.name) {
        case 'xs': return 40;
        case 'sm': return 30;
        case 'md': return 40;
        case 'lg': return 50;
        case 'xl': return 50;
        default: return 10;
      }
    },
  }
};
</script>

<style scoped lang="scss">
.logo {
  width: 30px;
  margin-bottom: $space-xxs;
  transition: transform .2s;
  min-height: 70px;
  @media screen and (min-width: 0px) and (max-width: 599px) {
    min-height: 50px;
  }
}

.logo:hover {
  transform: scale(1.05);
}
.content {
  padding-top: 50px;
  @media screen and (max-width: 960px) {
    padding-left: 20px;
    padding-right: 20px;
  }
}

.logo i {
  margin-right: 30px;
  @media screen and (min-width: 1904px) {
    font-size: 45px;
  }
  @media screen and (min-width: 1264px) and (max-width: 1903px) {
    font-size: 40px;
    margin-right: 20px;
  }
  @media screen and (min-width: 871px) and (max-width: 1263px) {
    font-size: 30px;
    margin-right: 15px;
  }
  @media screen and (min-width: 727px) and (max-width: 944px) {
    font-size: 27px;
    margin-right: 10px;
  }
  @media screen and (min-width: 600px) and (max-width: 726px) {
    font-size: 25px;
    margin-right: 5px;
  }
  @media screen and (min-width: 0px) and (max-width: 599px) {
    font-size: 25px;
    margin-right: 5px;
  }
}

.link-text {
  font-weight: bold;
  @media screen and (min-width: 1904px) {
    font-size: 30px;
  }
  @media screen and (min-width: 1264px) and (max-width: 1903px) {
    font-size: 25px;
  }
  @media screen and (min-width: 945px) and (max-width: 1263px) {
    font-size: 20px;
  }
  @media screen and (min-width: 727px) and (max-width: 944px) {
    font-size: 15px;
  }
  @media screen and (min-width: 600px) and (max-width: 726px) {
    font-size: 11px;
  }
  @media screen and (min-width: 0px) and (max-width: 599px) {
    font-size: 18px;
  }
}

.icon-container {
  width: 70%;
  margin: 0 auto;
  //512
  @media (min-width: 512px) and (max-width: 599px) {
    width: 80%;
  }
  @media (min-width: 421px) and (max-width: 512px) {
    width: 90%;
  }
  @media (min-width: 0px) and (max-width: 420px) {
    width: 100%;
  }
}

.other-btn:hover {
  transform: scale(1.3);
}

.other-btn i {
  @media screen and (min-width: 1904px) {
    font-size: 45px;
  }
  @media screen and (min-width: 1264px) and (max-width: 1903px) {
    font-size: 40px;
  }
  @media screen and (min-width: 871px) and (max-width: 1263px) {
    font-size: 30px;
  }
  @media screen and (min-width: 727px) and (max-width: 944px) {
    font-size: 27px;
  }
  @media screen and (min-width: 600px) and (max-width: 726px) {
    font-size: 25px;
  }
  @media screen and (min-width: 0px) and (max-width: 599px) {
    font-size: 26px;
  }
}

</style>

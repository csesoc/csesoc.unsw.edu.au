<!--
  Footer
  --
  This component contains footer elements of the website, visible from all pages.
-->

<template>
  <v-footer color="#000033" padless>
    <v-container class="white--text" style="padding:60px 72px" fluid>
      <v-row no-gutters>
        <v-col cols="12" sm="5" order-sm="7">
          <div class="white-text socials">
            <h4 data-cy="footer-media-title">Social Media</h4>
            <v-row style="margin-top:20px" no-gutters>
              <v-col cols="6">
                <v-list-item class="anchor-items" v-for="url in mediaLeft" :key="url.name">
                  <v-list-item-content data-cy="footer-media-link">
                    <a :href="url.url">
                      {{ url.name }}
                    </a>
                  </v-list-item-content>
                </v-list-item>
              </v-col>
              <v-col cols="6">
                <v-list-item class="anchor-items" v-for="url in mediaRight" :key="url.name">
                  <v-list-item-content data-cy="footer-media-link">
                    <a :href="url.url">
                      {{ url.name }}
                    </a>
                  </v-list-item-content>
                </v-list-item>
              </v-col>
            </v-row>
          </div>
        </v-col>
        <v-col order-sm="5">
          <!-- Puts padding between logo and socials -->
        </v-col>
        <v-col
          cols="12"
          sm="5"
          order-sm="1"
          :class="{ 'mt-10': $vuetify.breakpoint.smAndDown, 'mt-0': $vuetify.breakpoint.mdAndUp }"
        >
          <div class="white-text" style="margin-bottom:20%" data-cy="footer-address">
            <a href="/">
              <img class="img-size" src="@/assets/csesoc-logo-white.svg" data-cy="footer-logo" />
            </a>
            <br />B03 CSE Building K17, UNSW
            <br />
            <a href="mailto:csesoc@csesoc.org.au"> csesoc@csesoc.org.au </a>
          </div>
          © {{new Date().getFullYear()}} — CSESoc UNSW
        </v-col>
      </v-row>
    </v-container>
  </v-footer>
</template>

<script type="text/javascript">
export default {
  name: 'Footer',
  data: () => ({
    // sponsors have name, logo, 'tier' value, and link
    csesocInternal: [
      { name: 'Home', url: '/' },
      { name: 'About', url: '/#/about' },
      { name: 'Resources', url: '/#/resources' },
      { name: 'Sponsorships', url: '/#/sponsors' },
      { name: 'Contact Us', url: '/#/engage' }
    ],
    csesocExternal: [
      { name: 'CSESoc Media', url: 'https://media.csesoc.org.au/' },
      { name: 'CSESoc Blog', url: 'https://blog.csesoc.org.au/' },
      { name: 'CSESoc Compclub', url: 'https://compclub.csesoc.unsw.edu.au/' },
      { name: 'CSE Faculty', url: 'https://www.engineering.unsw.edu.au/computer-science-engineering/' },
      { name: 'UNSW Handbook', url: 'https://www.handbook.unsw.edu.au/' }
    ],
    mediaLeft: [
      {
        name: 'Discord Community',
        url:
          'https://discord.gg/rwNkSjRzHs'
      },
      { name: 'Facebook Page', url: 'https://www.facebook.com/csesoc/' },
      { name: 'Facebook Group', url: 'https://www.facebook.com/csesoc/' },
      { name: 'LinkedIn', url: 'https://www.linkedin.com/company/csesoc/' }
    ],
    mediaRight: [
      { name: 'Instagram', url: 'https://www.instagram.com/csesoc_unsw/' },
      { name: 'YouTube', url: 'https://www.youtube.com/channel/UC1JHpRrf9j5IKluzXhprUJg' },
      { name: 'TikTok', url: 'https://www.tiktok.com/@csesoc?lang=en' }
    ]
  })
};
</script>

<style lang="scss" scoped>
.white-text {
  @extend p;
}
.white-text a {
  text-decoration: none;
  color: $light-color;
  padding-right: $space-xxs;
}
.white-text a:hover {
  font-weight: 900;
  transition-duration: 0.1s;
}
.nav-divider {
  padding-bottom: $space-xxs;
  border-bottom: 2px solid white;
}
.anchor-items {
  margin-left: -15px;
}
.img-size {
  max-width: 300px;
  max-height: 200px;
}
.socials {
  margin-top: 60px;
}
</style>

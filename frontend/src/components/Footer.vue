<!--
  Footer
  --
  This component contains footer elements of the website, visible from all pages.
-->

<template>
  <v-footer color="#000033" padless >
    <v-container class="white--text" style="padding:60px 72px" fluid>
      <v-row no-gutters>
        <v-col
          sm="6"
        >
          <div class="white-text" style="margin-bottom:20%" data-cy="footer-address">
          <a href="/">
            <img class="img-size" src="@/assets/csesoc-logo-white.svg" data-cy="footer-logo"/>
          </a>
          <br/>B03 CSE Building K17, UNSW
          <br/>
          <a href="mailto:csesoc@csesoc.org.au" > csesoc@csesoc.org.au </a>
          </div>
          © 2020 — CSESoc UNSW
        </v-col>
        <v-col
          sm="6"
        >
          <div class="white-text">
            <div class="nav-divider" data-cy = "footer-internal-link">
              <a v-for="url in csesocInternal" :key="url.name" :href="url.url">
              {{ url.name }}
              </a>
            </div>
            <v-row style="margin-top:25px" no-gutters>
              <v-col
                sm="6"
              >
                <h4 data-cy="footer-media-title">Social Media</h4>
                <v-list-item class="anchor-items" v-for="url in media" :key="url.name">
                  <v-list-item-content data-cy="footer-media-link">
                    <a :href="url.url">
                      {{ url.name }}
                    </a>
                  </v-list-item-content>
                </v-list-item>
              </v-col>
              <v-col
                sm="6"
              >
                <h4 data-cy="footer-resources-title">For your better future</h4>
                <v-list-item class="anchor-items" v-for="url in csesocExternal" :key="url.name">
                  <v-list-item-content data-cy=footer-resources-link>
                    <a :href="url.url">
                      {{ url.name }}
                    </a>
                  </v-list-item-content>
                </v-list-item>
              </v-col>
            </v-row>
          </div>
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
      { name: 'Engage', url: '/#/engage' }
    ],
    csesocExternal: [
      { name: 'CSESoc Media', url: 'https://media.csesoc.org.au/' },
      { name: 'CSESoc Blog', url: 'https://blog.csesoc.org.au/' },
      { name: 'CSESoc Compclub', url: 'https://compclub.csesoc.unsw.edu.au/' },
      { name: 'UNSW CSE', url: 'https://www.engineering.unsw.edu.au/computer-science-engineering/' },
      { name: 'UNSW Handbook', url: 'https://www.handbook.unsw.edu.au/' }
    ],
    media: [
      { name: 'Facebook Page', url: 'https://www.facebook.com/csesoc/' },
      { name: 'Facebook Group', url: 'https://www.facebook.com/csesoc/' },
      { name: 'Instagram', url: 'https://www.instagram.com/csesoc_unsw/' },
      { name: 'Discord Community', url: 'https://forms.office.com/Pages/ResponsePage.aspx?id=pM_2PxXn20i44Qhnufn7o6ecLZTBorREjnXuTY-JfmBUMEpOMFBDTU1UWkhBWllWRTNPOVJFMUNCRi4u' },
      { name: 'Slack Community', url: 'https://csesoc-community.slack.com/' },
      { name: 'LinkedIn', url: 'https://www.linkedin.com/company/csesoc/' },
      { name: 'TikTok', url: '#' },
      { name: 'YouTube', url: 'https://www.youtube.com/channel/UC1JHpRrf9j5IKluzXhprUJg' }
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
  padding-right:$space-xxs;
}

.white-text a:hover {
  font-weight: 900;
  transition-duration: 0.5s;
}

.nav-divider {
  padding-bottom: $space-xxs;
  border-bottom:2px solid white;
}

.anchor-items {
  margin-left:-15px;
}

.img-size {
  max-width:300px;
  max-height:200px;
}
</style>

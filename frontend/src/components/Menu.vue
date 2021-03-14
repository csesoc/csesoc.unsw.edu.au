<!--
  Menu
  --
  This component contains an overlay with links to all the pages of the website.
  The links have a grow-on-hover feature.
-->

<template>
  <div class="cover">
    <div class="frame">
      <!-- Header -->
      <v-container class="stack-elem">
        <v-row no-gutters>
          <!-- Logo -->
          <v-col cols="2">
            <RouterLink to="/">
              <v-img @click.stop="hide" class="logo-btn" width="140px" height="32px" src="@/assets/csesoc-logo-white.svg" />
            </RouterLink>
          </v-col>
          <!-- Spacer -->
          <v-spacer></v-spacer>
          <!-- Close -->
          <v-col cols="1">
            <v-img @click.stop="hide" class="close-btn" width="32px" height="32px" src="@/assets/close-icon.svg" />
          </v-col>
        </v-row>
      </v-container>

      <!-- Links -->
      <v-container class="body stack-elem">
        <v-row no-gutters style="height: 100%;">
          <!-- Page links -->
          <v-col xs="12" sm="6">
            <div class="page-links-container" @click.stop="hide">
              <RouterLink to="/about" class="link" data-cy="about-link">
                <h2>001 | About</h2>
              </RouterLink>
              <RouterLink to="/resources" class="link" data-cy="resources-link">
                <h2>010 | Resources</h2>
              </RouterLink>
              <RouterLink to="/sponsors" class="link" data-cy="sponsors-link">
                <h2>011 | Sponsors</h2>
              </RouterLink>
              <RouterLink to="/engage" class="link" data-cy="engage-link">
                <h2>100 | Engage</h2>
              </RouterLink>
            </div>
          </v-col>
          <!-- Social links -->
          <v-col xs="12" sm="6">
            <div class="social-links-container">
              <div class="push-down">
                <div v-for="link in socialLinks" :key="link.name" class="link" data-cy="menu-social-link">
                  <a target="_blank" :href="link.url" >
                    {{ link.name }}
                  </a>
                </div>
              </div>
            </div>
          </v-col>
        </v-row>
      </v-container>

      <!-- Footer -->
      <v-container class="stack-elem">
        <v-row no-gutters class="footer">
          <p style="display: inline;">
            <img height="14px" src="@/assets/moon-icon.png" style="padding-right: 5px;">
            We're sorry, only Dark Mode is available at CSESoc.
          </p>
        </v-row>
      </v-container>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Menu',
  data: () => ({
    socialLinks: [
      { name: 'Facebook', url: 'https://www.facebook.com/csesoc/' },
      { name: 'Instagram', url: 'https://www.instagram.com/csesoc_unsw/' },
      { name: 'Discord Community', url: 'https://bitly.com/CSESocCourseChats' },
      { name: 'Slack Community', url: 'https://csesoc-community.slack.com/' },
      { name: 'YouTube', url: 'https://www.youtube.com/channel/UC1JHpRrf9j5IKluzXhprUJg' },
      { name: 'LinkedIn', url: 'https://www.linkedin.com/company/csesoc/' },
      { name: 'TikTok', url: '#' },
    ]
  }),
  methods: {
    hide() {
      // Relaying status to parent
      this.$emit('shown', false);
    }
  }
};
</script>

<style lang="scss" scoped>
.cover {
  z-index: 5; /* display on top of nav-bar */
  display: block;
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: $dark-color-1;
}

.frame {
  display: flex;
  flex-direction: column;
  margin-top: $space-xxs;
  margin-right: 4.5em;
  margin-bottom: $space-xs;
  margin-left: 4.5em;
  height: 100%;
}

.stack-elem {
  padding: 0;
}

.split {
  display: flex;
  flex-direction: row;
  flex: 1;
  margin-top: $space-md;
}

.body {
  margin-top: $space-md;
  height: 100%;
}

.footer {
  margin-bottom: $space-xs;
  color: white;
}

.logo-btn {
  margin-left: 5px;
  margin-right: auto;
  cursor: pointer;
}

.close-btn {
  margin-left: auto;
  cursor: pointer;
}

.page-links-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;

  .link {
    text-decoration: none; /* Remove underline from links */
    color: white;
    cursor: pointer; /* To simulate a link */
    font-weight: normal;
    transition: all .2s ease-in-out;

    h2 {
      cursor: pointer; /* To simulate a link */
    }

    &:hover {
      font-weight: bolder;
      font-size: $text-lg;
    }
  }
}

.social-links-container {
  height: 100%;
  display: flex;
  flex-direction: row;
  align-items: flex-end;

  .push-down {
    flex-direction: column;
    align-items: flex-end;
    width: 100%;

    .link {
      text-align: right;
      color: white;
      cursor: pointer; /* To simulate a link */
      font-weight: normal;
      transition: all .2s ease-in-out;
      width: fit-content;
      margin-left: auto;

      a {
        @extend p; /* Copy <p> */
        text-decoration: none; /* Remove underline from links */
        color: white;
        cursor: pointer; /* To simulate a link */
      }

      &:hover {
        font-weight: bolder;
        font-size: $text-sm;
      }
    }
  }
}
</style>

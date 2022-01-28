<!--
  PrimaryNavbar
  --
  This component is static on the top of Home page and links to different parts of the screen.
  Also links to the Sponsorship page.
-->

<template>
  <v-app-bar app color="rgba(0,0,51,0.7)" elevation="0" width="100vw">
    <!-- Large screen display -->
    <v-container class="hidden-sm-and-down">
      <v-row class="d-flex justify-center">
        <v-col class="d-flex align-center justify-start">
          <RouterLink to="/" style="text-decoration: none;" v-on:click.native="showMenu = false">
            <v-img contain max-height="60px" max-width="100px" src="@/assets/csesocwhiteblue.png" />
          </RouterLink>
        </v-col>
        <v-col class="d-flex align-center justify-center">
          <a @click="goto('aboutus')" style="text-decoration: none;">
            <v-btn text color="white">ABOUT US</v-btn>
          </a>
        </v-col>
        <v-col class="d-flex align-center justify-center">
          <a @click="goto('community')" style="text-decoration: none;">
            <v-btn text color="white">GET IN TOUCH</v-btn>
          </a>
        </v-col>
        <v-col class="d-flex align-center justify-center">
          <a @click="goto('event-display')" style="text-decoration: none;">
            <v-btn text color="white">EVENTS</v-btn>
          </a>
        </v-col>
        <v-col class="d-flex align-center justify-center">
          <a @click="goto('student-resources')" style="text-decoration: none;">
            <v-btn text color="white">RESOURCES</v-btn>
          </a>
        </v-col>
        <v-col class="d-flex align-center justify-center">
          <v-btn text to="/sponsors" color="white">SPONSORSHIP</v-btn>
        </v-col>
      </v-row>
    </v-container>

    <!-- Smaller screen container -->
    <v-container class="hidden-md-and-up">
      <v-row class="d-flex justify-space-between">
        <v-col class="d-flex align-center">
          <RouterLink to="/" style="text-decoration: none;" v-on:click.native="showMenu = false">
            <v-img contain max-height="40px" max-width="100px" src="@/assets/csesocwhiteblue.png" />
          </RouterLink>
        </v-col>
        <v-col class="d-flex justify-end">
          <v-menu bottom left>
            <template v-slot:activator="{ on, attrs }">
              <v-btn dark icon v-bind="attrs" v-on="on">
                <v-icon>mdi-menu</v-icon>
              </v-btn>
            </template>
            <v-list>
              <v-list-item>
                <a @click="goto('aboutus')" style="text-decoration: none; color: black;">
                  About us
                </a>
              </v-list-item>
              <v-list-item>
                <a @click="goto('community')" style="text-decoration: none; color: black;">
                  Get in touch
                </a>
              </v-list-item>
              <v-list-item>
                <a @click="goto('event-display')" style="text-decoration: none; color: black;">
                  Events
                </a>
              </v-list-item>
              <v-list-item>
                <a @click="goto('student-resources')" style="text-decoration: none; color: black;">
                  Resources
                </a>
              </v-list-item>
              <v-list-item>
                <RouterLink style="text-decoration: none; color: black;" to="/sponsors">
                  Sponsorship
                </RouterLink>
              </v-list-item>
            </v-list>
          </v-menu>
        </v-col>
      </v-row>
    </v-container>
  </v-app-bar>
</template>

<script type="text/javascript">
export default {
  name: 'PrimaryNavbar',
  methods: {
    goto(refName) {
      const element = document.getElementById(refName);
      element.scrollIntoView({ block: 'start', behavior: 'smooth' });
    }
  }
};
</script>

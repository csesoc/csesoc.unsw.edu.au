<!--
  SecondaryNavbar
  --
  This component is static on the top of the About and Sponsorship pages. It links to
  each one and the home `/` page.
-->

<template>
  <v-app-bar app color="rgba(0,0,51,0.9)" elevation="0" width="100vw">
    <!-- Large screen display -->
    <v-container class="hidden-sm-and-down">
      <v-row class="d-flex">
        <v-col class="d-flex ps-10 align-center justify-start">
          <RouterLink to="/" style="text-decoration: none;" v-on:click.native="showMenu = false">
            <v-img contain max-height="60px" max-width="100px" src="@/assets/csesocwhiteblue.png" />
          </RouterLink>
        </v-col>
        <v-col class="d-flex align-center justify-end pr-10">
          <v-btn text to="/about" class="mr-2" color="white">ABOUT US</v-btn>
          <v-btn text to="/sponsors" class="ms-2" color="white">SPONSORSHIP</v-btn>
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
                <RouterLink to="/about" style="text-decoration: none; color: black;">
                  About us
                </RouterLink>
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
  name: 'SecondaryNavbar'
};
</script>

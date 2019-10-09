<template>
  <v-footer
    dark
    padless
  >
    <v-card
      class="flex"
      flat
      tile
      color="#3A76F8"
    >
  <!--
      <v-card-title  align="right">
          <strong>CONTACT US</strong>

        <v-spacer></v-spacer>

<br>
CSESoc
<br>
//  office location //
<br>
For enquiries about our society
<br>
Please email info(@)csesoc.org.au
<br>
For enquiries about sponsorship
<br>
Please email sponsorship(@)csesoc.org.au
<br>
      </v-card-title>
-->

      <v-card-actions class="grey darken-3 justify-left">
        © Copyright {{ new Date().getFullYear() }} — <strong>CSESoc</strong>
      </v-card-actions>
    </v-card>
  </v-footer>
</template>

<script>
export default {
  data: () => ({
    icons: [
      'fab fa-facebook',
      'fab fa-twitter',
      'fab fa-google-plus',
      'fab fa-linkedin',
      'fab fa-instagram',
    ],
  }),
};
</script>

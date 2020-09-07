<template>
<!-- Join Us Section -->
<!-- data-cy="communityLinks-images" -->
    <v-container data-cy="communityLinks" id="join">
      <v-row class="pa-0">
        <v-col cols="4" align="center" class="pa-0">
          <a href="https://www.facebook.com/groups/csesoc">
            <img src="@/assets/facebook_logo.png" class="zoom" data-cy="communityLinks-facebookImage">
          </a>
        </v-col>
        <v-col cols="4" align="center" class="pa-0">
          <a href="https://csesoc-community.slack.com/">
            <img src="@/assets/slack_logo.png" class="zoom" data-cy="communityLinks-slackImage">
          </a>
        </v-col>
        <v-col cols="4" align="center" class="pa-0">
          <a href="https://www.facebook.com/groups/csesoc">
            <img src="@/assets/discord_logo.png" class="zoom" data-cy="communityLinks-discordImage">
          </a>
        </v-col>
      </v-row>
    </v-container>
</template>

<script>
</script>

<style scoped lang="scss">

#join img {
  width: 150px;
  height: auto;
  //padding-left: 4%;
  //padding-right: 4%;
}

.zoom {
  transition: transform .2s;
}

.zoom:hover {
  transform: scale(1.1);
}
</style>

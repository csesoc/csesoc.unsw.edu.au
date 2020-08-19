<template>
<!-- Join Us Section -->
    <v-container id="join" ref="joinus">
      <h1>
        Join Us
      </h1>
      <p>
        Find our stall at O-week, or just turn up an event and get to know us!
      </p>
      <h2>
        Join our community online
      </h2>
      <div id="joinImg" >
        <a href="https://www.facebook.com/groups/csesoc">
          <img src="@/assets/facebook_logo.png" class="zoom">
        </a>
        <a href="https://csesoc-community.slack.com/">
          <img src="@/assets/slack_logo.png" class="zoom">
        </a>
        <!-- need to add in the csesoc discord link-->
        <a href="https://www.facebook.com/groups/csesoc">
          <img src="@/assets/discord_logo.png" class="zoom">
        </a>
      </div>
    </v-container>
</template>

<script>
</script>

<style scoped>

#join {
  align-items: center;
  text-align: center;
}

#join h1 {
  font-size: 80px;
}

#join h2 {
  font-size: 40px;
  padding-bottom: 5vh;
}

#join p {
  font-size: 35px;
  padding: 0 15vw 3vh;
}

#join img {
  width: 20%;
  height: auto;
  padding-left: 4%;
  padding-right: 4%;
}

#joinImg {
  padding-bottom: 5%;
}

.zoom {
  transition: transform .2s;
}

.zoom:hover {
  transform: scale(1.1);
}
</style>
